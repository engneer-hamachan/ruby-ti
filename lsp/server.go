package lsp

import (
	"bufio"
	"strings"
	"sync"
	"ti/base"
	"ti/context"
	"ti/eval"
	"ti/lexer"
	"ti/lexer/reader"
	"ti/parser"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var handler protocol.Handler

// ドキュメントの内容を保持する
var (
	documents       = make(map[string]string)
	docMutex        sync.RWMutex
	builtinSigs     []base.Sig // ビルトインのシグネチャを保存
	builtinSigsOnce sync.Once
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
	// 初回のみビルトインのシグネチャを保存
	builtinSigsOnce.Do(func() {
		builtinSigs = make([]base.Sig, len(base.TSignatures))
		copy(builtinSigs, base.TSignatures)
	})

	// TSignaturesをビルトインのみにリセット
	base.TSignatures = make([]base.Sig, len(builtinSigs))
	copy(base.TSignatures, builtinSigs)

	br := bufio.NewReader(strings.NewReader(content))

	for _, round := range context.GetRounds() {
		lr := reader.New(*br)
		l := lexer.New(lr)
		p := parser.New(l, filename)

		// cleanSimpleIdentifires
		for key, value := range base.TFrame {
			if value.IsIdentifierType() && key.Variable() == value.ToString() {
				delete(base.TFrame, key)
			}
		}

		ctx := context.NewContext("", "", round)
		evaluator := eval.Evaluator{}
		p.Errors = []error{}

		for {
			t, err := p.Read()
			if err != nil {
				return err
			}

			err = evaluator.Eval(&p, ctx, t)
			if err != nil {
				return err
			}

			if t == nil {
				break
			}
		}

		// 次のラウンドのために br をリセット
		br = bufio.NewReader(strings.NewReader(content))
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
