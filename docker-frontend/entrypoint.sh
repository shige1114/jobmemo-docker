#!/bin/bash
set -e

if [ ! -e "/frontend/package.json" ]; then
  echo 'nextjsを新規インストール'
  npm init -y
  npm install create-next-app
  npx create-next-app@latest jobapp --use-npm --typescript
  rm -rf jobapp/.gitignore
fi

if [ ! -d "/frontend/jobapp/node_modules" ]; then
  echo 'jobappの環境構築'
  npm install
fi

cd jobapp

# Then exec the container's main process (what's set as CMD in the Dockerfile).
exec "$@"