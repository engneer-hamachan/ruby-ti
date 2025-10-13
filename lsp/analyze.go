package lsp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeContent(content string, line uint32) error {
	responseSignatures = nil

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

func findDefinition(content string, params *protocol.DefinitionParams) (any, error) {
	lines := strings.Split(content, "\n")
	if int(params.Position.Line) >= len(lines) {
		return nil, nil
	}

	currentLine := lines[params.Position.Line]
	methodName := extractMethodName(currentLine, int(params.Position.Character))
	if methodName == "" {
		return nil, nil
	}

	// h.test 1 のときは h だけに、test 1 のときは test だけにする
	targetForPrefix := extractTargetForPrefix(currentLine, int(params.Position.Character))

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
	if prefixInfo == "" {
		return nil, nil
	}

	// @{frame}:::{class} からクラス情報を抽出
	parts := strings.SplitN(strings.TrimPrefix(prefixInfo, "@"), ":::", 2)
	if len(parts) < 2 {
		return nil, nil
	}
	frame := parts[0]
	class := parts[1]

	// ti {file} --define で全メソッド定義を取得
	definitions := getMethodDefinitions(tmpFile.Name())

	// ドットが入っていなくて、frameがunknownだったらtoplevelメソッド
	searchFrame := frame
	searchClass := class
	if !hasDot && frame == "unknown" {
		searchFrame = "unknown"
		searchClass = "unknown"
	}

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

		if defFrame == searchFrame && defClass == searchClass && strings.HasPrefix(defMethod, methodName) {
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

	return nil, nil
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
