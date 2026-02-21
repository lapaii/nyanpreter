#!/usr/bin/env bash

# build-all.sh - compile nyassembler and run it on all .nyan programs in this directory
# Usage: simply execute from the test-programs directory (it builds all .nyan files)

set -euo pipefail

# locate the nyassembler directory relative to this script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
NYASSEMBLER_DIR="${SCRIPT_DIR}/../nyassembler"

# build the assembler
echo "Building nyassembler in ${NYASSEMBLER_DIR}..."
cd "${NYASSEMBLER_DIR}"
go build .

# path to assembler executable
NYASSEMBLER_EXEC="${NYASSEMBLER_DIR}/nyassembler"

# go back to test-programs directory
cd "${SCRIPT_DIR}"

echo "Processing .nyasm files in ${SCRIPT_DIR}..."

shopt -s nullglob
for src in *.nyasm; do
    out="${src%.nyasm}.nyobj"
    echo "  assembling $src -> $out"
    "${NYASSEMBLER_EXEC}" --input "$src" --output "$out"
done

echo "Done."
