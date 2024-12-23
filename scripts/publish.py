#!/usr/bin/env python3
"""Publish ticketeer to package repositories"""
from abc import ABC, abstractmethod
from argparse import ArgumentParser, ArgumentTypeError
from pathlib import Path
import sys
import json
import shutil
import re
from helpers.shell import shell, ShellError

PACKAGES_DIR = Path("packaging")
DIST_DIR = Path("dist")

npm_binary = {
    "ticketeer-darwin-x64": "ticketeer_darwin_amd64_v1/ticketeer",
    "ticketeer-darwin-arm64": "ticketeer_darwin_arm64_v8.0/ticketeer",
    "ticketeer-freebsd-x64": "ticketeer_freebsd_amd64_v1/ticketeer",
    "ticketeer-freebsd-arm64": "ticketeer_freebsd_arm64_v8.0/ticketeer",
    "ticketeer-linux-x64": "ticketeer_linux_amd64_v1/ticketeer",
    "ticketeer-linux-arm64": "ticketeer_linux_arm64_v8.0/ticketeer",
    "ticketeer-openbsd-arm64": "ticketeer_openbsd_arm64_v8.0/ticketeer",
    "ticketeer-openbsd-x64": "ticketeer_openbsd_amd64_v1/ticketeer",
    "ticketeer-windows-x64": "ticketeer_windows_amd64_v1/ticketeer.exe",
    "ticketeer-windows-arm64": "ticketeer_windows_arm64_v8.0/ticketeer.exe",
}

class TicketeerPackaging(ABC):
    """Packaging interface"""

    @abstractmethod
    def __str__(self) -> str:
        """Describe package(s)"""

    @abstractmethod
    def set_version(self, version: str):
        """Update version in package(s) metadata file(s)"""

    @abstractmethod
    def copy_binaries(self, dist_src: Path):
        """Copy binaries into package(s) from dist directory"""

    @abstractmethod
    def publish(self):
        """Publish package(s)"""

class NPMPackaging(TicketeerPackaging):
    """NPM ticketeer packaging"""

    _packages: list[Path]

    def __init__(self, packages: list[Path]):
        self._packages = packages

    @staticmethod
    def from_dir(path: Path):
        """Create NPM packaging from directory"""
        manifests = path.glob("*/package.json")
        packages = [manifest.parent for manifest in manifests]
        return NPMPackaging(packages)

    def _set_package_version(self, package: Path, version: str):
        manifest_path = package / "package.json"
        with manifest_path.open("r", encoding="utf-8") as file:
            manifest = json.load(file)
        manifest["version"] = version
        if "optionalDependencies" in manifest:
            for key, _ in manifest["optionalDependencies"].items():
                manifest["optionalDependencies"][key] = version
        with manifest_path.open("w", encoding="utf-8") as file:
            json.dump(manifest, file, indent=2)

    def set_version(self, version: str):
        for package in self._packages:
            self._set_package_version(package, version)

    def copy_binaries(self, dist_src: Path):
        for package in self._packages:
            name = package.name
            if name not in npm_binary:
                if name == "ticketeer":
                    continue
                print(f"Warning: no binary for {name}")
                return
            src = dist_src / npm_binary[name]
            shutil.copy(src, package)

    def publish(self):
        for package in self._packages:
            print(f"Publishing {package.name}")
            try:
                shell("npm publish", cwd=package)
            except ShellError as exc:
                print(f"Failed to publish {package.name}")
                print(exc.stderr)
                continue

    def __str__(self):
        packages = self._packages.copy()
        base_package = "ticketeer"
        for package in packages:
            if package.name == base_package:
                packages.remove(package)

        output = f"- npm packages:\n  - {base_package}"
        for package in packages:
            output += f"\n  - {package.name}"
        return output

class GitPackaging(TicketeerPackaging):
    """Git ticketeer packaging"""

    def __str__(self):
        return f"- git repository\n  - {Path.cwd()}"

    def set_version(self, version):
        shell(f"git tag v{version}")

    def copy_binaries(self, _: Path):
        pass

    def publish(self):
        # Publishing is handled by Makefile
        pass


def semantic_version(s: str) -> str:
    """Semantic version validator"""
    semver_re = r"v?[0-9]+\.[0-9]+\.[0-9]+(-[a-z0-9\.]+)?"
    if not re.match(semver_re, s):
        raise ArgumentTypeError(f"not a valid semantic version: {s!r}")
    if s.startswith("v"):
        s = s[1:]
    return s

def main():
    """Script entry point"""
    parser = ArgumentParser(
        description="Publish ticketeer to package repositories.",
        usage="publish.py <version>",
    )
    parser.add_argument(
        "version",
        help="Version to publish. Should be a semantic version (e.g. 1.2.3, 1.2.3-rc.1)",
        type=semantic_version)
    parser.add_argument(
        "-d", "--dist",
        help=(
            "Directory with compiled binaries. "
            "The script expects the folder structure that goreleaser creates"),
        type=Path,
        metavar="",
        default=DIST_DIR)

    parser.add_argument(
        "-p", "--packages",
        help="Directory with ticketeer packages",
        type=Path,
        default=PACKAGES_DIR,
        metavar="")

    try:
        args = parser.parse_args()
    except ArgumentTypeError as exc:
        print(exc)
        parser.print_usage()
        sys.exit(1)

    packaging: list[TicketeerPackaging] = [
        NPMPackaging.from_dir(args.packages / "npm"),
        GitPackaging(),
    ]

    for pack in packaging:
        print(pack)

    print(f"Target version: {args.version}")

    print("Copying binaries...")
    for pack in packaging:
        pack.copy_binaries(DIST_DIR)
    print("Updating version...")
    for pack in packaging:
        pack.set_version(args.version)
    print("Publishing...")
    for pack in packaging:
        pack.publish()
    print("Done")

if __name__ == "__main__":
    main()
