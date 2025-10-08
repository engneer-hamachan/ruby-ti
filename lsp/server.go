package lsp

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"ti/base"
	"time"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var handler protocol.Handler

// ドキュメントの内容を保持する
var (
	documents = make(map[string]string)
	docMutex  sync.RWMutex
)

func NewServer() *server.Server {
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textDocumentCompletion,
		TextDocumentDidOpen:    textDocumentDidOpen,
		TextDocumentDidChange:  textDocumentDidChange,
		TextDocumentDidClose:   textDocumentDidClose,
	}

	server := server.NewServer(&handler, "ruby-ti", false)
	return server
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{".", "@"},
	}

	// ドキュメント同期の設定
	capabilities.TextDocumentSync = protocol.TextDocumentSyncKindIncremental

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

func analyzeContent(content string, filename string) error {
	// TSignaturesをクリア（tiコマンドの出力で完全に置き換える）
	base.TSignatures = nil

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

	// tiコマンドを実行（-aオプションで補完候補を取得）
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, tiPath, tmpFile.Name(), "-a")
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
		if strings.HasPrefix(line, "%") {
			line = strings.TrimPrefix(line, "%")
			parts := strings.SplitN(line, ":::", 2)
			methodName := parts[0]
			detail := ""
			if len(parts) == 2 {
				detail = parts[1]
			}

			if !methodSet[methodName] {
				methodSet[methodName] = true
				base.TSignatures = append(base.TSignatures, base.Sig{
					Contents: methodName,
					Detail:   detail,
				})
			}
		}
	}

	return nil
}

func textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	docMutex.Lock()
	documents[params.TextDocument.URI] = params.TextDocument.Text
	docMutex.Unlock()

	// ドキュメントを解析してTSignaturesを更新（エラーは無視）
	_ = analyzeContent(params.TextDocument.Text, string(params.TextDocument.URI))

	return nil
}

func textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	docMutex.Lock()
	uri := params.TextDocument.URI
	for _, changeAny := range params.ContentChanges {
		change, ok := changeAny.(protocol.TextDocumentContentChangeEvent)
		if !ok {
			continue
		}

		if change.Range == nil {
			// Full document update
			documents[uri] = change.Text
		} else {
			// Incremental update
			documents[uri] = applyChange(documents[uri], change)
		}
	}
	content := documents[uri]
	docMutex.Unlock()

	// 変更後のドキュメントを解析してTSignaturesを更新（エラーは無視）
	_ = analyzeContent(content, string(uri))

	return nil
}

func textDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	docMutex.Lock()
	defer docMutex.Unlock()
	delete(documents, params.TextDocument.URI)
	return nil
}

func applyChange(content string, change protocol.TextDocumentContentChangeEvent) string {
	if change.Range == nil {
		return change.Text
	}

	lines := strings.Split(content, "\n")
	start := change.Range.Start
	end := change.Range.End

	// 開始位置までの文字列
	before := ""
	for i := 0; i < int(start.Line); i++ {
		before += lines[i] + "\n"
	}
	if int(start.Line) < len(lines) {
		before += lines[start.Line][:start.Character]
	}

	// 終了位置以降の文字列
	after := ""
	if int(end.Line) < len(lines) {
		after = lines[end.Line][end.Character:]
		for i := int(end.Line) + 1; i < len(lines); i++ {
			after += "\n" + lines[i]
		}
	}

	return before + change.Text + after
}

func fuzzyMatch(pattern, text string) bool {
	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)

	patternIdx := 0
	for i := 0; i < len(text) && patternIdx < len(pattern); i++ {
		if text[i] == pattern[patternIdx] {
			patternIdx++
		}
	}
	return patternIdx == len(pattern)
}

func getCurrentWord(content string, line, character uint32) string {
	lines := strings.Split(content, "\n")
	if int(line) >= len(lines) {
		return ""
	}

	currentLine := lines[line]
	if int(character) > len(currentLine) {
		character = uint32(len(currentLine))
	}

	// カーソル位置から単語の開始位置を探す
	start := int(character)
	for start > 0 {
		ch := currentLine[start-1]
		if !isWordChar(ch) {
			break
		}
		start--
	}

	return currentLine[start:character]
}

func isWordChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		ch == '_' || ch == '@' || ch == '$'
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	var items []protocol.CompletionItem

	// ドキュメントの内容を取得
	docMutex.RLock()
	content, ok := documents[params.TextDocument.URI]
	docMutex.RUnlock()

	query := ""
	if ok {
		// カーソル位置から入力中の単語を取得
		query = getCurrentWord(content, params.Position.Line, params.Position.Character)
	}

	// TSignaturesから補完候補を生成（曖昧検索でフィルタリング）
	// queryが空の場合は何も返さない
	if query != "" {
		for _, sig := range base.TSignatures {
			if fuzzyMatch(query, sig.Contents) {
				items = append(items, protocol.CompletionItem{
					Label:  sig.Contents,
					Kind:   &[]protocol.CompletionItemKind{protocol.CompletionItemKindMethod}[0],
					Detail: &sig.Detail,
				})
			}
		}
	}

	return items, nil
}
