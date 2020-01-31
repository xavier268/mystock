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
# Note that cgo and the related cross-comile options are needed !
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" CXX="x86_64-w64-mingw32-g++" go build -o ${MYBIN}mystock.exe $BASE/cmd/mystock/

