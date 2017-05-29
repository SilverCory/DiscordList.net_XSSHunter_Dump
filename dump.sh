#!/bin/bash

CURRENT_DIR=$(pwd)
cd .. && 
go run "$CURRENT_DIR/dump.go" &&
cd "$CURRENT_DIR" &&
git add -A &&
git commit -m "Emails received."
git push
