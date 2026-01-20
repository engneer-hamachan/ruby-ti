#!/bin/bash

# Script to generate test from sample.rb

# Build the project first
echo "Building project..."
go build .

# Check if build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

# Generate hash-based test filename
test_dir="test"
# Create a hash from the content of sample.rb and current timestamp to ensure uniqueness
content_hash=$(cat sample.rb | sha256sum | cut -c1-8)
test_name="${content_hash}"

# Check if this hash already exists and append counter if needed
counter=0
original_test_name="$test_name"
while [ -f "$test_dir/${test_name}.rb" ]; do
    ((counter++))
    test_name="${original_test_name}_${counter}"
done

# Run type checker on sample.rb to get expected output
echo "Running type checker on sample.rb..."
raw_output=$(./ti sample.rb "$@" 2>&1)
# Replace sample.rb with ./${test_name}.rb in the output for the test
output=$(echo "$raw_output" | sed "s|sample\.rb|\./${test_name}\.rb|g")
exit_code=$?
rb_file="$test_dir/${test_name}.rb"
go_test_file="$test_dir/${test_name}_test.go"

# Copy sample.rb to new test file
echo "Creating $rb_file..."
cp sample.rb "$rb_file"

# Build Go command arguments (comma-separated additional args)
go_args=""
for arg in "$@"; do
    # Escape quotes and backslashes in the argument
    escaped_arg=$(printf '%s' "$arg" | sed 's/\\/\\\\/g; s/"/\\"/g')
    go_args="$go_args, \"$escaped_arg\""
done

# Create Go test file
echo "Creating $go_test_file..."
cat > "$go_test_file" << EOF
package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test$(echo ${test_name:0:1} | tr '[:lower:]' '[:upper:]')${test_name:1}(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./${test_name}.rb"$go_args)

	output, _ := cmd.CombinedOutput()

	expectedOutput := \`$output\`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
EOF

echo "Generated test files:"
echo "  Ruby file: $rb_file"
echo "  Go test file: $go_test_file"
echo "  Expected output: $output"

# Test the generated test
echo "Running the generated test..."
cd test
capitalized_name=$(echo ${test_name:0:1} | tr '[:lower:]' '[:upper:]')${test_name:1}
go test -run "Test${capitalized_name}"
