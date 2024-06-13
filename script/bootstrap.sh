#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=thriftbam
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}