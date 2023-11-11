package main

import (
	"fmt"
	"runtime"
)

func main() {

	if os := runtime.GOOS; os == "Linux" || os == "OS X. " {
		fmt.Println("Esto es Windows")
	} else {
		fmt.Println("Esto es Windwdos ", os)
	}

}
