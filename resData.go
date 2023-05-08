package main

import (
	"embed"
	"strings"
)

//go:embed data
var resFileSys embed.FS

func ReadResData(path string) (ret []byte, err error) {
	path = strings.ReplaceAll(path, "\\", "/")
	ret, err = resFileSys.ReadFile(path)

	return
}
