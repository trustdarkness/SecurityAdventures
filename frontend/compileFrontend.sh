#!/bin/sh

function compile_frontend () {
  echo "compiling frontend..."

  mkdir -p src/scripts/js
  mkdir -p src/styles/css

  coffee --output src/scripts/js/ --compile src/scripts/coffee/
  lessc src/styles/less/* > src/styles/css/main.css

  echo "...finished"
}

compile_frontend
