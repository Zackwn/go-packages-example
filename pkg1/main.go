package main

import (
	"fmt"

	"github.com/zackwn/packages-example/pkg2"
)

func main() {
	fmt.Println("Main func in pkg1")

	Another()
	pkg2.Pkg2()
}
