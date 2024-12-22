#!/usr/bin/env python3
"""Combine go coverage files"""
import sys
from argparse import ArgumentParser

def combine_coverage(files: list[str]) -> str:
    """Combine go coverage files"""
    content: str = ""
    common_mode: str | None = None
    for file_path in files:
        with open(file_path, "r", encoding="utf-8") as file:
            mode = file.readline().rstrip()
            file_content = file.read()
            if common_mode is None:
                common_mode = mode
            elif common_mode != mode:
                raise ValueError(f"Coverage files have different mode: {common_mode} != {mode}")
            content += file_content
    return f"{common_mode}\n{content}"

if __name__ == "__main__":
    parser = ArgumentParser(description="Combine Go coverage files")
    parser.add_argument('files', nargs='+', help='Go txt coverage files')
    parser.add_argument('-o', '--output', help='Output file', default="combined.cover.txt")
    args = parser.parse_args()

    if len(args.files) < 2:
        print("Error: not enough arguments")
        print("Usage: combine_coverage.py <file1> <file2> ...")
        sys.exit(1)

    combined = combine_coverage(args.files)
    with open(args.output, "w", encoding="utf-8") as f:
        f.write(combined)

    print(f"Combined {len(args.files)} files into {args.output}")
