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

func TestRbsToJsonJS(t *testing.T) {
	cmd := exec.Command("sh", "./shell/rbs_to_json_test.sh", "rbs/js.rbs")
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v, output: %s", err, string(output))
	}

	expectedOutput := `"{\n  \"frame\": \"Builtin\",\n  \"class\": \"JS\",\n  \"instance_methods\": [],\n  \"class_methods\": [\n    {\n      \"name\": \"global\",\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"document\",\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    }\n  ]\n}"
"{\n  \"frame\": \"JS\",\n  \"class\": \"Object\",\n  \"instance_methods\": [\n    {\n      \"name\": \"to_poro\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"Untyped\"\n        ]\n      }\n    },\n    {\n      \"name\": \"[]\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"Symbol\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"fetch\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"block_parameters\": [\n        \"::Object\"\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"Void\"\n        ]\n      }\n    },\n    {\n      \"name\": \"addEventListener\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"block_parameters\": [\n        \"String\"\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"Int\"\n        ]\n      }\n    },\n    {\n      \"name\": \"to_binary\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"String\"\n        ]\n      }\n    },\n    {\n      \"name\": \"querySelector\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"appendChild\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"JS::Object\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"removeChild\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"JS::Object\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"setAttribute\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        },\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"Void\"\n        ]\n      }\n    },\n    {\n      \"name\": \"removeAttribute\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"Void\"\n        ]\n      }\n    },\n    {\n      \"name\": \"replaceChild\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"JS::Object\"\n          ]\n        },\n        {\n          \"type\": [\n            \"JS::Object\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"createElement\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"createTextNode\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"className\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"classList\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"children\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"parentElement\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"tagName\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"innerHTML\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"innerHTML=\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"String\"\n        ]\n      }\n    },\n    {\n      \"name\": \"style\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"style=\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"String\"\n        ]\n      }\n    },\n    {\n      \"name\": \"textContent\",\n      \"arguments\": [],\n      \"return_type\": {\n        \"type\": [\n          \"JS::Object\"\n        ]\n      }\n    },\n    {\n      \"name\": \"textContent=\",\n      \"arguments\": [\n        {\n          \"type\": [\n            \"String\"\n          ]\n        }\n      ],\n      \"return_type\": {\n        \"type\": [\n          \"String\"\n        ]\n      }\n    }\n  ]\n}"`

	actualOutput := strings.TrimSpace(string(output))
	expectedOutputTrimmed := strings.TrimSpace(expectedOutput)

	if actualOutput != expectedOutputTrimmed {
		t.Errorf("Output mismatch:%s", showDiff(expectedOutputTrimmed, actualOutput))
	}
}
