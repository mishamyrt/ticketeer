#!/usr/bin/env python3
"""Publish ticketeer to package repositories"""
from abc import ABC, abstractmethod
from argparse import ArgumentParser, ArgumentTypeError
from pathlib import Path
import os
import sys
import json
import shutil
import re
from helpers.shell import shell, ShellError

PACKAGES_DIR = Path("packaging")
DIST_DIR = Path("dist")

class TicketeerPackaging(ABC):
    """Packaging interface"""

    @property
    @abstractmethod
    def name(self) -> str:
        """Name of the package"""

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
    def build(self):
        """Build package(s)"""

    @abstractmethod
    def publish(self):
        """Publish package(s)"""

class NPMPackaging(TicketeerPackaging):
    """NPM ticketeer packaging"""

    _token: str | None
    _packages: list[Path]
    _binary_sources = {
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

    def __init__(self, packages: list[Path], token=None):
        self._packages = packages
        self._token = token

    @property
    def name(self) -> str:
        return "npm"

    @staticmethod
    def from_dir(path: Path, token=None):
        """Create NPM packaging from directory"""
        manifests = path.glob("*/package.json")
        packages = [manifest.parent for manifest in manifests]
        return NPMPackaging(packages, token)

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
            if name not in self._binary_sources:
                if name == "ticketeer":
                    continue
                print(f"Warning: no binary for {name}")
                return
            src = dist_src / self._binary_sources[name]
            shutil.copy(src, package)

    def publish(self):
        node_token = os.environ.get("NODE_AUTH_TOKEN")
        os.environ["NODE_AUTH_TOKEN"] = self._token
        for package in self._packages:
            print(f"  - {package.name}")
            shell("npm publish", cwd=package)
        if node_token is not None:
            os.environ["NODE_AUTH_TOKEN"] = node_token

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

    def build(self):
        print("  nothing to build")

class GitPackaging(TicketeerPackaging):
    """Git ticketeer packaging"""
    _version_file = Path("internal/ticketeer/version.go")
    _version_re = r'const version = \"(.*)\"'

    @property
    def name(self) -> str:
        return "git"

    def __str__(self):
        return f"- git repository\n  - {Path.cwd()}"

    def set_version(self, version):
        shell(f"git tag v{version}")
        with self._version_file.open("r", encoding="utf-8") as file:
            content = file.read()
        content = re.sub(self._version_re, f'const version = "{version}"', content)
        with self._version_file.open("w", encoding="utf-8") as file:
            file.write(content)

    def copy_binaries(self, _: Path):
        pass

    def publish(self):
        print("  handled by Makefile")

    def build(self):
        print("  nothing to build")

class PyPiPackaging(TicketeerPackaging):
    """PyPI ticketeer packaging"""

    _token: str | None
    _packages: list[Path]
    _submodule_version_re = r'\"ticketeer_(.*)==.*;(.*)\"'
    _binary_sources: dict[str, list[tuple[str, str]]] = {
        "ticketeer_darwin": [
            ("ticketeer_darwin_amd64_v1/ticketeer", "ticketeer_amd64"),
            ("ticketeer_darwin_arm64_v8.0/ticketeer", "ticketeer_arm64"),
        ],
        "ticketeer_freebsd": [
            ("ticketeer_freebsd_amd64_v1/ticketeer", "ticketeer_amd64"),
            ("ticketeer_freebsd_arm64_v8.0/ticketeer", "ticketeer_arm64"),
        ],
        "ticketeer_linux": [
            ("ticketeer_linux_amd64_v1/ticketeer", "ticketeer_amd64"),
            ("ticketeer_linux_arm64_v8.0/ticketeer", "ticketeer_arm64"),
        ],
        "ticketeer_openbsd": [
            ("ticketeer_openbsd_amd64_v1/ticketeer", "ticketeer_amd64"),
            ("ticketeer_openbsd_arm64_v8.0/ticketeer", "ticketeer_arm64"),
        ],
        "ticketeer_windows": [
            ("ticketeer_windows_amd64_v1/ticketeer.exe", "ticketeer_amd64.exe"),
            ("ticketeer_windows_arm64_v8.0/ticketeer.exe", "ticketeer_arm64.exe"),
        ],
    }

    def __init__(self, packages: list[Path], token=None):
        self._packages = packages
        self._token = token

    @property
    def name(self) -> str:
        return "pypi"

    @staticmethod
    def from_dir(path: Path, token=None):
        """Create NPM packaging from directory"""
        manifests = path.glob("*/pyproject.toml")
        packages = [manifest.parent for manifest in manifests]
        return PyPiPackaging(packages, token)

    def __str__(self):
        packages = self._packages.copy()
        base_package = "ticketeer"
        for package in packages:
            if package.name == base_package:
                packages.remove(package)

        output = f"- pypi packages:\n  - {base_package}"
        for package in packages:
            output += f"\n  - {package.name}"
        return output

    def set_version(self, version: str):
        runner_package: Path
        for package in self._packages:
            with open(package / ".version", "w", encoding="utf-8") as file:
                file.write(version)
            if package.name == "ticketeer":
                runner_package = package
        runner_manifest = runner_package / "pyproject.toml"
        with open(runner_manifest, "r", encoding="utf-8") as file:
            pyproject = file.read()
        pyproject = re.sub(
            self._submodule_version_re,
            lambda match: f'"ticketeer_{match.group(1)}=={version};{match.group(2)}"',
            pyproject)
        with open(runner_manifest, "w", encoding="utf-8") as file:
            file.write(pyproject)

    def copy_binaries(self, dist_src: Path):
        readme_src = "README.md"
        for package in self._packages:
            readme_dst = package / readme_src
            shutil.copy(readme_src, readme_dst)
            name = package.name
            if name not in self._binary_sources or name == "ticketeer":
                continue
            for src, dst in self._binary_sources[name]:
                src = dist_src / src
                dst = package / package.name / dst
                shutil.copy(src, dst)

    def build(self):
        for package in self._packages:
            print(f"  - {package.name}")
            shell("python -m build", cwd=package)

    def publish(self):
        if self._token is None:
            raise RuntimeError("No token provided")
        for package in self._packages:
            print(f"  - {package.name}")
            shell((
                "python -m twine upload dist/* "
                "--repository pypi "
                f"--password {self._token} "
                "--non-interactive"
            ), cwd=package)

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

    pypi_token = os.environ.get("TICKETEER_PYPI_TOKEN")
    if pypi_token is None:
        print("TICKETEER_PYPI_TOKEN environment variable not set")
        sys.exit(1)

    npm_token = os.environ.get("TICKETEER_NPM_TOKEN")
    if npm_token is None:
        print("TICKETEER_NPM_TOKEN environment variable not set")
        sys.exit(1)

    packaging: list[TicketeerPackaging] = [
        NPMPackaging.from_dir(args.packages / "npm", token=npm_token),
        PyPiPackaging.from_dir(args.packages / "pypi", token=pypi_token),
        GitPackaging(),
    ]


    for pack in packaging:
        print(pack)

    print(f"Target version: {args.version}")

    print("Updating version...")
    for pack in packaging:
        pack.set_version(args.version)
    print("Copying binaries...")
    for pack in packaging:
        pack.copy_binaries(DIST_DIR)
    print("Building packages...")
    for pack in packaging:
        print("- " + pack.name)
        pack.build()
    print("Publishing...")
    for pack in packaging:
        print("- " + pack.name)
        pack.publish()
    print("Done")

if __name__ == "__main__":
    try:
        main()
    except ShellError as exc:
        print(f"Command {exc.cmd} failed")
        out = exc.stdout
        if len(exc.stderr.strip()) > 0:
            out += f"\n\n{exc.stderr}"
        print(out)
        sys.exit(1)
    except Exception as exc: # pylint: disable=broad-except
        print(exc)
        sys.exit(1)
