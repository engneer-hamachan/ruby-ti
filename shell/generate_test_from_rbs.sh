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
	"os/exec"
	"strings"
	"testing"
)

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
