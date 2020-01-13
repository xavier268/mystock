#!/bin/bash

BASE=$(pwd)
BASE=${BASE/mystock*/}"mystock"
echo "Base folder detected as :" $BASE
MYBIN=$BASE/bin/


mkdir $MYBIN
go build -o $MYBIN $BASE/... 
