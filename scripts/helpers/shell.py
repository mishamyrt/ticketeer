import subprocess

def shell(command: str, cwd: str = None) -> str:
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