# isutil

utility commands for ISUCON

## Requirements

go 1.16 or later

## How to use

```
go install github.com/1m-yen-driven/isutil/cmd/...@latest
cd /path/to/go/project
structs ./... | peco | xargs -I@ tags -struct @ -key json ./...
```
