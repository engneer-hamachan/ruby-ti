package lsp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"ti/base"
	"time"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var handler protocol.Handler
var documentContents = make(map[string]string) // URI -> content のマップ

func NewServer() *server.Server {
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textDocumentCompletion,
		TextDocumentDidOpen:    textDocumentDidOpen,
		TextDocumentDidChange:  textDocumentDidChange,
		TextDocumentDidSave:    textDocumentDidSave,
	}

	server := server.NewServer(&handler, "ruby-ti", false)
	return server
}

func logger(log string) {
	log += "\n"
	f, err := os.OpenFile("./lsp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(log)
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{".", "@"},
	}

	// ドキュメント同期の設定（Full syncで全文を受け取る）
	syncKind := protocol.TextDocumentSyncKindFull
	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		OpenClose: &[]bool{true}[0],
		Change:    &syncKind,
		Save:      &protocol.SaveOptions{IncludeText: &[]bool{true}[0]},
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    "ruby-ti",
			Version: &[]string{"0.1.0"}[0],
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	return nil
}

func analyzeContent(content string, line uint32) error {
	// TSignaturesをクリア（tiコマンドの出力で完全に置き換える）
	base.TSignatures = nil

	// カーソル位置の行を取得
	lines := strings.Split(content, "\n")
	if int(line) < len(lines) {
		currentLine := lines[line]
		// 行が . で終わっている場合、. を削除して構文を完全にする
		trimmed := strings.TrimSpace(currentLine)
		if strings.HasSuffix(trimmed, ".") {
			// 末尾の . を削除
			lines[line] = strings.TrimSuffix(currentLine, ".")
			content = strings.Join(lines, "\n")
		}
	}

	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp("", "lsp-*.rb")
	if err != nil {
		logger(fmt.Sprintf("Error creating temp file: %v", err))
		return nil // エラーは無視してビルトインのみ使用
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// コンテンツを書き込み
	if _, err := tmpFile.WriteString(content); err != nil {
		logger(fmt.Sprintf("Error writing temp file: %v", err))
		return nil
	}
	tmpFile.Close()

	// tiコマンドのパスを取得（実行ファイルと同じディレクトリ）
	exePath, err := os.Executable()
	if err != nil {
		logger(fmt.Sprintf("Error getting executable path: %v", err))
		return nil
	}
	tiPath := filepath.Join(filepath.Dir(exePath), "ti")

	// tiコマンドを実行（-aオプションで補完候補を取得、行番号も渡す）
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	// 行番号は0ベースなので+1して1ベースに変換
	cmdLine := fmt.Sprintf("%s %s -a %d", tiPath, tmpFile.Name(), line+1)
	logger(fmt.Sprintf("Running command: %s", cmdLine))
	logger(fmt.Sprintf("File content:\n%s", content))

	cmd := exec.CommandContext(ctx, tiPath, tmpFile.Name(), "-a", fmt.Sprintf("%d", line+1))
	output, err := cmd.Output()
	if err != nil {
		// タイムアウトやエラーは無視
		logger(fmt.Sprintf("Error running ti command: %v", err))
		return nil
	}

	logger(fmt.Sprintf("ti output length: %d", len(output)))
	logger("ti output:")
	logger(string(output))
	logger("end ti output")

	// 出力を解析して補完候補を抽出（重複を除外するためにmapを使用）
	methodSet := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		// %で始まる行が補完候補（形式: %content:::detail）
		line, ok := strings.CutPrefix(line, "%")
		if !ok {
			continue
		}

		parts := strings.SplitN(line, ":::", 2)
		if len(parts) < 2 {
			continue
		}
		methodName := parts[0]
		detail := parts[1]

		if !methodSet[detail] {
			methodSet[detail] = true
			base.TSignatures = append(base.TSignatures, base.Sig{
				Contents: methodName,
				Detail:   detail,
			})
		}
	}

	logger(fmt.Sprintf("Found %d completion candidates", len(base.TSignatures)))

	return nil
}

func textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	// ドキュメント内容を保存
	documentContents[params.TextDocument.URI] = params.TextDocument.Text
	return nil
}

func textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	// ContentChangesからテキストを取得して解析
	if len(params.ContentChanges) > 0 {
		// Full syncモードなので最初の要素に全テキストが入っている
		change := params.ContentChanges[0]

		// JSONとして再度マーシャル・アンマーシャルして構造体に変換
		changeBytes, err := json.Marshal(change)
		if err == nil {
			var changeEvent struct {
				Text string `json:"text"`
			}
			if err := json.Unmarshal(changeBytes, &changeEvent); err == nil {
				text := changeEvent.Text
				logger("===Change===")
				logger(text)
				logger("==========")

				// ドキュメント内容を更新
				documentContents[params.TextDocument.URI] = text
			}
		}
	}

	return nil
}

func textDocumentDidSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	content := params.Text

	// ドキュメント内容を更新
	documentContents[params.TextDocument.URI] = *content

	return nil
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	var items []protocol.CompletionItem

	// キャッシュから最新のコンテンツを取得
	content, ok := documentContents[params.TextDocument.URI]
	if !ok {
		// キャッシュにない場合はファイルから読み込む
		filePath := strings.TrimPrefix(params.TextDocument.URI, "file://")
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return items, nil
		}
		content = string(fileContent)
	}

	logger("===Comp===")
	logger(content)
	logger("==========")

	// カーソル位置の行番号で解析実行
	// Position.Lineは0ベースなので、そのまま渡す（analyzeContent内で+1される）
	err := analyzeContent(content, params.Position.Line)
	if err != nil {
		logger(err.Error())
	}

	// TSignaturesから補完候補を生成（全件返す）
	for _, sig := range base.TSignatures {
		items = append(items, protocol.CompletionItem{
			Label:  sig.Contents,
			Kind:   &[]protocol.CompletionItemKind{protocol.CompletionItemKindMethod}[0],
			Detail: &sig.Detail,
		})
	}

	return items, nil
}
