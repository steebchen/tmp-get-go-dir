package main

import (
	"fmt"
	"go/build"
)

func main() {
	ctx := build.Default
	pkg, err := ctx.Import("github.com/steebchen/tmp-get-go-dir", ".", build.FindOnly)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dir: %s\n", pkg.Dir)
}
