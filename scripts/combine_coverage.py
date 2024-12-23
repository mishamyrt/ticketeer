#!/usr/bin/env python3
"""Combine multiple go coverage files into a single one"""
import sys
from argparse import ArgumentParser, Action, ArgumentTypeError

def minimum_length(min_value: int):
    """Creates minimum length action"""
    class MinimumLength(Action):
        """Minimum length action"""
        def __call__(self, _, namespace, values, option_string=None):
            if len(values) < min_value:
                raise ArgumentTypeError(
                    f"Expected at least {min_value} arguments, "
                    f"got {len(values)}")
            setattr(namespace, self.dest, values)
    return MinimumLength

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

def main():
    """Script entry point"""
    parser = ArgumentParser(
        description="Combine Go coverage files",
        usage="combine_coverage.py <file1> <file2> ... [-o <output_file>]")
    parser.add_argument(
        'files',
        help='Go txt coverage files',
        nargs='+',
        action=minimum_length(2))
    parser.add_argument(
        '-o', '--output',
        help='Output file',
        default="combined.cover.out")

    try:
        args = parser.parse_args()
    except ArgumentTypeError as e:
        print(e)
        parser.print_usage()
        sys.exit(1)

    try:
        combined = combine_coverage(args.files)
    except ValueError as exc:
        print(exc)
        sys.exit(1)
    except FileNotFoundError as exc:
        print(f"Coverage file not found: {exc.filename}")
        sys.exit(1)

    try:
        with open(args.output, "w", encoding="utf-8") as f:
            f.write(combined)
    except OSError as exc:
        print(f"Failed to write to {args.output}: {exc}")
        sys.exit(1)

    print(f"Combined {len(args.files)} files into {args.output}")

if __name__ == "__main__":
    main()
