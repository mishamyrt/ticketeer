#!/bin/sh

if [ "$TICKETEER_VERBOSE" = "1" ] || [ "$TICKETEER_VERBOSE" = "true" ]; then
  set -x
fi

if [ "$TICKETEER" = "0" ]; then
  exit 0
fi

call_ticketeer() {
  binName="ticketeer"
  if test -n "$TICKETEER_BIN"; then
    "$TICKETEER_BIN" "$@"
  elif $binName -h >/dev/null 2>&1; then
    $binName "$@"
  else
    dir="$(git rev-parse --show-toplevel)"
    osArch=$(uname | tr '[:upper:]' '[:lower:]')
    cpuArch=$(uname -m | sed 's/aarch64/arm64/;s/x86_64/x64/')
    if test -f "$dir/node_modules/ticketeer-${osArch}-${cpuArch}/bin/$binName"
    then
      "$dir/node_modules/ticketeer-${osArch}-${cpuArch}/bin/$binName" "$@"
    elif test -f "$dir/node_modules/ticketeer/bin/index.js"
    then
      "$dir/node_modules/ticketeer/bin/index.js" "$@"
    elif yarn ticketeer -h >/dev/null 2>&1; then
      yarn ticketeer "$@"
    elif pnpm ticketeer -h >/dev/null 2>&1; then
      pnpm ticketeer "$@"
    else
      echo "ERROR: Can't find ticketeer in PATH."
      echo "Make sure ticketeer is available in your environment and re-try."
      echo "To skip these checks use --no-verify git argument or set TICKETEER=0 env variable."
      exit 1
    fi
  fi
}

call_ticketeer apply "$@"
