package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"ti/base"
	"time"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var handler protocol.Handler
var documentContents = make(map[string]string)

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

func initialize(
	ctx *glsp.Context,
	params *protocol.InitializeParams,
) (any, error) {

	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{".", "@"},
	}

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

func analyzeContent(content string, line uint32) error {
	base.TSignatures = nil

	content = removeTaiilDot(content, line)

	tmpFile, err := os.CreateTemp("", "ruby-ti-lsp-*.rb")
	if err != nil {
		return nil
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		return nil
	}

	tmpFile.Close()

	ctx, cancel :=
		context.WithTimeout(context.Background(), 1000*time.Millisecond)

	defer cancel()

	cmd :=
		exec.CommandContext(
			ctx,
			"ti",
			tmpFile.Name(),
			"-a",
			fmt.Sprintf("%d", line+1),
		)

	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	setTSignatures(output)

	return nil
}

func textDocumentDidOpen(
	ctx *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {

	documentContents[params.TextDocument.URI] = params.TextDocument.Text
	return nil
}

func textDocumentDidChange(
	ctx *glsp.Context,
	params *protocol.DidChangeTextDocumentParams,
) error {

	if len(params.ContentChanges) > 0 {
		change := params.ContentChanges[0]

		changeEventBytes, err := json.Marshal(change)
		if err == nil {
			var changeEvent struct {
				Text string `json:"text"`
			}

			if err := json.Unmarshal(changeEventBytes, &changeEvent); err == nil {
				documentContents[params.TextDocument.URI] = changeEvent.Text
			}
		}
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

	for _, sig := range base.TSignatures {
		items = append(items, protocol.CompletionItem{
			Label:  sig.Contents,
			Detail: &sig.Detail,
		})
	}

	return items, nil
}
