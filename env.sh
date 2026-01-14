#!/bin/bash
PROJECT_ROOT="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
export PROJECT_ROOT
export PATH="$PROJECT_ROOT/bin:$PATH"

if command -v python >/dev/null 2>&1; then
  export PROJECT_PYTHON="$(command -v python)"
elif command -v python3 >/dev/null 2>&1; then
  export PROJECT_PYTHON="$(command -v python3)"
else
  echo "[env.sh] ERROR: Could not find python/python3, please install or activate the environment first." >&2
  return 1
fi

echo "[env.sh] OK: PROJECT_ROOT=$PROJECT_ROOT"
echo "[env.sh] OK: PROJECT_PYTHON=$PROJECT_PYTHON"
echo "[env.sh] Now you use project manager: aspm -h"