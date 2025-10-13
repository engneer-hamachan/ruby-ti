package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"ti/base"
	"time"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var handler protocol.Handler
var documentContents = make(map[string]string)
var responseSignatures = []base.Sig{}

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
		TextDocumentDefinition: textDocumentDefinition,
	}

	server := server.NewServer(&handler, "ruby-ti", false)
	return server
}

func initialize(
	ctx *glsp.Context,
	params *protocol.InitializeParams,
) (any, error) {

	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
			"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			".", "_",
		},
	}

	syncKind := protocol.TextDocumentSyncKindFull

	capabilities.TextDocumentSync =
		protocol.TextDocumentSyncOptions{
			OpenClose: &[]bool{true}[0],
			Change:    &syncKind,
			Save:      &protocol.SaveOptions{IncludeText: &[]bool{true}[0]},
		}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    "ruby-ti-lsp",
			Version: &[]string{"beta"}[0],
		},
	}, nil
}

func initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(ctx *glsp.Context) error {
	return nil
}

func setTrace(ctx *glsp.Context, params *protocol.SetTraceParams) error {
	return nil
}

func textDocumentDidOpen(
	ctx *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {

	documentContents[params.TextDocument.URI] = params.TextDocument.Text
	return nil
}

var changeEvent struct {
	Text string `json:"text"`
}

func textDocumentDidChange(
	ctx *glsp.Context,
	params *protocol.DidChangeTextDocumentParams,
) error {

	if len(params.ContentChanges) < 1 {
		return nil
	}

	change := params.ContentChanges[0]

	changeEventBytes, err := json.Marshal(change)
	if err != nil {
		return nil
	}

	if err := json.Unmarshal(changeEventBytes, &changeEvent); err == nil {
		documentContents[params.TextDocument.URI] = changeEvent.Text
	}

	return nil
}

func textDocumentDidSave(
	ctx *glsp.Context,
	params *protocol.DidSaveTextDocumentParams,
) error {

	content := params.Text
	documentContents[params.TextDocument.URI] = *content

	return nil
}

func textDocumentCompletion(
	ctx *glsp.Context,
	params *protocol.CompletionParams,
) (any, error) {

	var items []protocol.CompletionItem

	content, ok := documentContents[params.TextDocument.URI]
	if !ok {
		return nil, nil
	}

	analyzeContent(content, params.Position.Line)

	for _, sig := range responseSignatures {
		items =
			append(items, protocol.CompletionItem{
				Label:  sig.Contents,
				Detail: &sig.Detail,
			})
	}

	return items, nil
}

func textDocumentDefinition(
	ctx *glsp.Context,
	params *protocol.DefinitionParams,
) (any, error) {

	logger(fmt.Sprintf("Definition request at line=%d, char=%d", params.Position.Line, params.Position.Character))

	content, ok := documentContents[params.TextDocument.URI]
	if !ok {
		logger("Document content not found")
		return nil, nil
	}

	// カーソル位置のテキストから対象のメソッド名を抽出
	lines := strings.Split(content, "\n")
	if int(params.Position.Line) >= len(lines) {
		logger("Line out of range")
		return nil, nil
	}

	currentLine := lines[params.Position.Line]
	logger(fmt.Sprintf("Current line: %s", currentLine))

	methodName := extractMethodName(currentLine, int(params.Position.Character))
	logger(fmt.Sprintf("Extracted method name: %s", methodName))

	if methodName == "" {
		logger("Method name is empty")
		return nil, nil
	}

	// h.test 1 のときは h だけに、test 1 のときは test だけにする
	targetForPrefix := extractTargetForPrefix(currentLine, int(params.Position.Character))
	logger(fmt.Sprintf("Target for prefix: %s", targetForPrefix))

	// ドットが含まれているかチェック（レシーバがあるか）
	hasDot := strings.Contains(strings.Fields(currentLine)[0], ".")

	// メソッド名だけを残した行に置き換える
	modifiedLines := make([]string, len(lines))
	copy(modifiedLines, lines)
	modifiedLines[params.Position.Line] = targetForPrefix
	modifiedContent := strings.Join(modifiedLines, "\n")

	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp("", "ruby-ti-lsp-*.rb")
	if err != nil {
		return nil, nil
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(modifiedContent); err != nil {
		return nil, nil
	}
	tmpFile.Close()

	// ti {file} -a {row} で @prefix を取得
	prefixInfo := getPrefixInfo(tmpFile.Name(), int(params.Position.Line)+1)
	logger(fmt.Sprintf("Prefix info: %s", prefixInfo))

	if prefixInfo == "" {
		logger("Prefix info is empty")
		return nil, nil
	}

	// @{frame}:::{class} からクラス情報を抽出
	parts := strings.SplitN(strings.TrimPrefix(prefixInfo, "@"), ":::", 2)
	if len(parts) < 2 {
		logger("Failed to parse prefix info")
		return nil, nil
	}
	frame := parts[0]
	class := parts[1]
	logger(fmt.Sprintf("Frame: %s, Class: %s", frame, class))

	// ti {file} --define で全メソッド定義を取得
	definitions := getMethodDefinitions(tmpFile.Name())
	logger(fmt.Sprintf("Found %d definitions", len(definitions)))

	// ドットが入っていなくて、frameがunknownだったらtoplevelメソッド
	searchFrame := frame
	searchClass := class
	if !hasDot && frame == "unknown" {
		searchFrame = "unknown"
		searchClass = "unknown"
		logger("Toplevel method detected (no dot)")
	}
	logger(fmt.Sprintf("Searching with Frame: %s, Class: %s", searchFrame, searchClass))

	// マッチするメソッド定義を検索
	// フォーマット: %{frame}:::{class}:::{method}:::{filename}:::{row}
	for _, def := range definitions {
		defParts := strings.SplitN(def, ":::", 5)
		if len(defParts) < 5 {
			continue
		}

		defFrame := defParts[0]
		defClass := defParts[1]
		defMethod := defParts[2]
		defFilename := defParts[3]
		defRow := defParts[4]

		logger(fmt.Sprintf("Checking: frame=%s, class=%s, method=%s", defFrame, defClass, defMethod))

		if defFrame == searchFrame && defClass == searchClass && strings.HasPrefix(defMethod, methodName) {
			logger(fmt.Sprintf("Match found! File: %s, Row: %s", defFilename, defRow))

			// 定義位置を返す
			var row uint32
			fmt.Sscanf(defRow, "%d", &row)
			if row > 0 {
				row-- // 0-based indexing
			}

			// tmpファイルのパスの場合は、元のファイルのURIを使う
			targetURI := params.TextDocument.URI
			if !strings.Contains(defFilename, "ruby-ti-lsp-") {
				// 別ファイルの定義の場合はそのパスを使う
				targetURI = protocol.DocumentUri("file://" + defFilename)
			}

			location := protocol.Location{
				URI: targetURI,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      row,
						Character: 0,
					},
					End: protocol.Position{
						Line:      row,
						Character: 0,
					},
				},
			}

			return location, nil
		}
	}

	logger("No matching definition found")
	return nil, nil
}

// extractMethodName extracts method name from line at cursor position
// Examples: "h.test 1" -> "test", "test 1" -> "test"
func extractMethodName(line string, col int) string {
	if col > len(line) {
		col = len(line)
	}

	// カーソル位置の単語の開始位置を探す
	start := col
	for start > 0 && isWordChar(line[start-1]) {
		start--
	}

	// カーソル位置の単語の終了位置を探す
	end := col
	for end < len(line) && isWordChar(line[end]) {
		end++
	}

	if start == end {
		return ""
	}

	return line[start:end]
}

// extractTargetForPrefix extracts target for -a option
// Examples: "h.test 1" -> "h", "test 1" -> "test"
func extractTargetForPrefix(line string, col int) string {
	if col > len(line) {
		col = len(line)
	}

	// カーソル位置の単語の開始位置を探す
	start := col
	for start > 0 && isWordChar(line[start-1]) {
		start--
	}

	// カーソル位置の単語の終了位置を探す
	end := col
	for end < len(line) && isWordChar(line[end]) {
		end++
	}

	if start == end {
		return ""
	}

	// ドットがあるかチェック（start より前を見る）
	dotPos := -1
	for i := start - 1; i >= 0; i-- {
		if line[i] == '.' {
			dotPos = i
			break
		} else if line[i] != ' ' && line[i] != '\t' {
			// ドット以外の文字があったら終了
			break
		}
	}

	if dotPos >= 0 {
		// h.test の場合、h を返す
		// ドットより前の単語を探す
		dotStart := dotPos
		for dotStart > 0 && isWordChar(line[dotStart-1]) {
			dotStart--
		}
		if dotStart < dotPos {
			return line[dotStart:dotPos]
		}
	}

	// ドットがない場合は、メソッド名だけを返す
	return line[start:end]
}

// isWordChar checks if a byte is part of a word (letter, digit, underscore)
func isWordChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_'
}

// getPrefixInfo gets @prefix info using ti -a
func getPrefixInfo(filename string, row int) string {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ti", filename, "-a", fmt.Sprintf("%d", row))
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "@") {
			return line
		}
	}

	return ""
}

// getMethodDefinitions gets all method definitions using ti --define
func getMethodDefinitions(filename string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ti", filename, "--define")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	var definitions []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "%") {
			definitions = append(definitions, strings.TrimPrefix(line, "%"))
		}
	}

	return definitions
}
