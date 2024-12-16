#!/usr/bin/env python3
"""Update version packaging files"""
import sys
from pathlib import Path
import json
import shutil
import subprocess

NPM_PACKAGES_DIR = Path("packaging/npm")
DIST_DIR = Path("dist")



def _sh(command: str, cwd: str = None) -> str:
    """Run a shell command and return output"""
    process = subprocess.run(
        command,
        shell=True,
        check=True,
        capture_output=True,
        text=True,
        cwd=cwd
    )
    if process.returncode != 0:
        raise subprocess.CalledProcessError(
            process.returncode, command, process.stderr
        )
    return process.stdout


npm_binary_mapping = {
    "ticketeer_darwin_amd64_v1/ticketeer": "ticketeer-darwin-x64/ticketeer",
    "ticketeer_darwin_arm64_v8.0/ticketeer": "ticketeer-darwin-arm64/ticketeer",
    "ticketeer_freebsd_amd64_v1/ticketeer": "ticketeer-freebsd-x64/ticketeer",
    "ticketeer_freebsd_arm64_v8.0/ticketeer": "ticketeer-freebsd-arm64/ticketeer",
    "ticketeer_linux_amd64_v1/ticketeer": "ticketeer-linux-x64/ticketeer",
    "ticketeer_linux_arm64_v8.0/ticketeer": "ticketeer-linux-arm64/ticketeer",
    "ticketeer_openbsd_arm64_v8.0/ticketeer": "ticketeer-openbsd-arm64/ticketeer",
    "ticketeer_openbsd_amd64_v1/ticketeer": "ticketeer-openbsd-x64/ticketeer",
    "ticketeer_windows_amd64_v1/ticketeer.exe": "ticketeer-windows-x64/ticketeer.exe",
    "ticketeer_windows_arm64_v8.0/ticketeer.exe": "ticketeer-windows-arm64/ticketeer.exe",
}

def set_npm_version(content: str, version: str) -> str:
    """Update version in package.json"""
    package = json.loads(content)
    package["version"] = version
    if "optionalDependencies" in package:
        for key, _ in package["optionalDependencies"].items():
            package["optionalDependencies"][key] = version
    return json.dumps(package, indent=2)

def pack_npm(target: str) -> str:
    """Update version in all files"""
    # Set npm packages version
    packages = list(NPM_PACKAGES_DIR.glob("*/package.json"))
    for package in packages:
        print(f"Updating {package}")
        with open(package, "r", encoding="utf-8") as f:
            content = f.read()
        content = set_npm_version(content, target)
        with open(package, "w", encoding="utf-8") as f:
            f.write(content)
    # Copy binaries
    for key, value in npm_binary_mapping.items():
        print(f"Copying {key} to {value}")
        src = DIST_DIR / key
        dst = NPM_PACKAGES_DIR / value
        shutil.copy(src, dst)
    # Publish npm packages
    for package in packages:
        package_dir = package.parent
        print(f"Publishing {package_dir}")
        try:
            _sh("npm publish", cwd=package_dir)
        except subprocess.CalledProcessError as e:
            print(f"Failed to publish {package_dir}")
            print(e.stderr)
            continue
    print(f"Published {target}")
    return target

if __name__ == "__main__":
    pack_npm(sys.argv[1])
