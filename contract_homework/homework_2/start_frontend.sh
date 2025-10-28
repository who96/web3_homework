#!/usr/bin/env bash
set -euo pipefail

PORT=5173
DOC_ROOT="frontend"

if lsof -ti :"$PORT" >/dev/null 2>&1; then
  echo "Killing existing process on port $PORT..."
  lsof -ti :"$PORT" | xargs kill -9 >/dev/null 2>&1 || true
fi

echo "Serving $DOC_ROOT on http://localhost:$PORT"
python3 -m http.server -d "$DOC_ROOT" "$PORT"
