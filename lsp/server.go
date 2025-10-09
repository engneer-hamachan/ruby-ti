package lsp

import (
	"bufio"
	"context"
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
	}

	server := server.NewServer(&handler, "ruby-ti", false)
	return server
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

	// デバッグ用：コンテンツをtmp.rbに出力
	os.WriteFile("/tmp/tmp.rb", []byte(content), 0644)

	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp("", "lsp-*.rb")
	if err != nil {
		return nil // エラーは無視してビルトインのみ使用
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// コンテンツを書き込み
	if _, err := tmpFile.WriteString(content); err != nil {
		return nil
	}
	tmpFile.Close()

	// tiコマンドのパスを取得（実行ファイルと同じディレクトリ）
	exePath, err := os.Executable()
	if err != nil {
		return nil
	}
	tiPath := filepath.Join(filepath.Dir(exePath), "ti")

	// tiコマンドを実行（-aオプションで補完候補を取得、行番号も渡す）
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	// 行番号は0ベースなので+1して1ベースに変換
	cmd := exec.CommandContext(ctx, tiPath, tmpFile.Name(), "-a", fmt.Sprintf("%d", line+1))
	output, err := cmd.Output()
	if err != nil {
		// タイムアウトやエラーは無視
		return nil
	}

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

	return nil
}

func textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	// ドキュメント内容を保存
	documentContents[params.TextDocument.URI] = params.TextDocument.Text
	// ファイルオープン時は行番号0を指定
	_ = analyzeContent(params.TextDocument.Text, 0)
	return nil
}

func textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	// ContentChangesの最後の要素が最新の状態
	if len(params.ContentChanges) > 0 {
		lastChange := params.ContentChanges[len(params.ContentChanges)-1]

		// map[string]interface{} として扱う
		if changeMap, ok := lastChange.(map[string]interface{}); ok {
			if text, ok := changeMap["text"].(string); ok {
				documentContents[params.TextDocument.URI] = text
				os.WriteFile("/tmp/didchange.log", []byte(fmt.Sprintf("Updated: %d bytes\n", len(text))), 0644)
			}
		}
	}
	return nil
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	var items []protocol.CompletionItem

	// URIからファイルパスを取得してファイルを直接読み込む
	filePath := strings.TrimPrefix(params.TextDocument.URI, "file://")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return items, nil
	}
	content := string(fileContent)

	// カーソル位置の行番号で解析実行
	// Position.Lineは0ベースなので、そのまま渡す（analyzeContent内で+1される）
	_ = analyzeContent(content, params.Position.Line)

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
