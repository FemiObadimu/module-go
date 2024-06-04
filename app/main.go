package main

import (
	"fmt"

	toolkitmodule "github.com/FemiObadimu/module-go"
)

func main() {
	var randomStr toolkitmodule.ToolkitModule
	s := randomStr.RandomString(5)
	fmt.Println(s)

}
