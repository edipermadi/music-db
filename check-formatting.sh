#!/usr/bin/env sh

OUTPUT=$(gofmt -l .)
if [ -n "$OUTPUT" ]; then
      echo "ERROR: code not formatted, please run gofmt -w .";
      exit 1;
fi
