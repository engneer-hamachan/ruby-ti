#!/bin/bash

# Generate builtin JSON from RBS file
# Usage: ./shell/rbs_to_json.sh <rbs_file>

if [ $# -eq 0 ]; then
    echo "Usage: $0 <rbs_file>"
    exit 1
fi

RBS_FILE=$1

if [ ! -f "$RBS_FILE" ]; then
    echo "Error: File '$RBS_FILE' not found"
    exit 1
fi

ruby rb_tools/rbs_to_json.rb "$RBS_FILE" > ./builtin.json

echo "Generated builtin.json from $RBS_FILE"
