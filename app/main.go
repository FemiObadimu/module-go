package main

import (
	"fmt"

	"github.com/FemiObadimu/module-go/toolkitModule"
)

func main() {
	var randomStr toolkitModule.ToolkitModule
	s := randomStr.RandomString(10)
	fmt.Println(s)

}
