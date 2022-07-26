package main

import (
	"github.com/Kwynto/agonist/cmd/desktop"
)

// The program is launched from main.go due to the peculiarities of "Fyne"
// and some peculiarities of the implementation of this program.
// Running through main.go avoids compilation problems.

func main() {
	desktop.Run()
}
