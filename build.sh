#!/bin/bash

BASE=$(pwd)
BASE=${BASE/mystock*/}"mystock"
echo "Base folder detected as :" $BASE
MYBIN=$BASE/bin/
mkdir -p $MYBIN

echo "Building linux version"
go build -o ${MYBIN}mystock $BASE/cmd/mystock/
chmod +x ${MYBIN}mystock

echo "Building windows version"
GOOS=windows GOARCH=amd64 go build -o ${MYBIN}mystock.exe $BASE/cmd/mystock/
chmod +x ${MYBIN}mystock

