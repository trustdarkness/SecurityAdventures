#!/bin/bash

function compile_frontend () {
  echo "compiling frontend..."

  mkdir -p src/scripts/js
  mkdir -p src/styles/css

  coffee --compile --output src/scripts/js/ src/scripts/coffee/
  less src/styles/less/* > src/styles/css/main.css

  echo "...finished"
}

compile_frontend
