#!/bin/bash

# Usage: sh shell/generate_test_from_rbs.sh path/to/file.rbs

if [ -z "$1" ]; then
  echo "Usage: sh shell/generate_test_from_rbs.sh path/to/file.rbs"
  exit 1
fi

RBS_FILE="$1"

if [ ! -f "$RBS_FILE" ]; then
  echo "Error: File '$RBS_FILE' not found"
  exit 1
fi

# Extract filename without extension
BASENAME=$(basename "$RBS_FILE" .rbs)

# Generate hash from file path for unique test name
HASH=$(echo -n "$RBS_FILE" | md5sum | cut -c1-8)

# Create test file in rbs_test directory
TEST_DIR="rbs_test"
mkdir -p "$TEST_DIR"

TEST_FILE="$TEST_DIR/${BASENAME}_${HASH}_test.go"

# Get expected output
EXPECTED_OUTPUT=$(sh ./shell/rbs_to_json_test.sh "$RBS_FILE" 2>/dev/null || echo "")

# Generate Go test file
cat > "$TEST_FILE" << GOEOF
package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// showDiff displays a diff by finding the first difference in the strings
func showDiff(expected, actual string) string {
	var diff strings.Builder
	diff.WriteString("\n=== DIFF ===\n")

	// Find first difference
	minLen := len(expected)
	if len(actual) < minLen {
		minLen = len(actual)
	}

	for i := 0; i < minLen; i++ {
		if expected[i] != actual[i] {
			start := i - 40
			if start < 0 {
				start = 0
			}
			end := i + 60
			expEnd := end
			actEnd := end
			if expEnd > len(expected) {
				expEnd = len(expected)
			}
			if actEnd > len(actual) {
				actEnd = len(actual)
			}

			diff.WriteString(fmt.Sprintf("First difference at position %d:\n", i))
			diff.WriteString(fmt.Sprintf("Expected: ...%s...\n", expected[start:expEnd]))
			diff.WriteString(fmt.Sprintf("Got:      ...%s...\n", actual[start:actEnd]))
			return diff.String()
		}
	}

	if len(expected) != len(actual) {
		diff.WriteString(fmt.Sprintf("Lengths differ: expected %d, got %d\n", len(expected), len(actual)))
		if len(expected) > minLen {
			diff.WriteString(fmt.Sprintf("Expected has extra: %s\n", expected[minLen:]))
		} else {
			diff.WriteString(fmt.Sprintf("Got has extra: %s\n", actual[minLen:]))
		}
	}

	return diff.String()
}

func Test_RBS_${BASENAME}_${HASH}(t *testing.T) {
	cmd := exec.Command("sh", "./shell/rbs_to_json_test.sh", "${RBS_FILE}")
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v, output: %s", err, string(output))
	}

	expectedOutput := \`${EXPECTED_OUTPUT}\`

	actualOutput := strings.TrimSpace(string(output))
	expectedOutputTrimmed := strings.TrimSpace(expectedOutput)

	if actualOutput != expectedOutputTrimmed {
		t.Errorf("Output mismatch:%s", showDiff(expectedOutputTrimmed, actualOutput))
	}
}
GOEOF

echo "Generated test file: $TEST_FILE"
