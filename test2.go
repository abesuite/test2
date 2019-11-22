package main

import (
	"fmt"
	"github.com/abesuite/test1"
)
import "github.com/abesuite/test1/crypto"

func main() {
	test1.TestDepend()

	var i = 2 * crypto.TestRipemd160Size()
	fmt.Print(i)
}
