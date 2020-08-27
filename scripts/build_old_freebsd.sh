#!/bin/sh
DIR=/go/gmirror_exporter
cd $DIR
GOOS=freebsd GOARCH=amd64 make
cp $DIR/bin/gmirror_exporter /build/gmirror_exporter_freebsd_amd64
