package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Println("My variable is: ", x)
	x = 20
	fmt.Println("My variable is now: ", x)
	x = 4
	fmt.Println("My variable is now: ", x)
	y := 33 + x
	fmt.Println("My other variable is: ", y)
	x += 4
	fmt.Println("My first variable is now: ", x)
	y++
	fmt.Println("Adding one to my second var: ", y)
	y--
	fmt.Println("Subtracting one to my second var: ", y)

	z := "\nASsingned string here: \n"
	z += "keywords GO: \n"
	z += "break\ncase\nchan\nconst\ncontinue\ndefault\ndefer\nelse\nfallthrough\nfor\nfunc\ngo\ngoto\nif\nimport\ninterface\nmap\npackage\nrange\nreturn\nselect\nstruct\nswitch\ntype\nvar\n"

	fmt.Println(z)

	fmt.Println("Declaration forms: ")
	fmt.Println("1 - SHORT  =>   x := value")
	
	fmt.Println("2 - using 'var'  =>   var x = value")

	fmt.Println("IMPORTANT variable must be assigned on creation")

	fmt.Println("STRING " + "Operator ???")
	fmt.Println("Bool ", true)
	fmt.Print("Ending ")
	fmt.Print(" my test ")
	fmt.Print("just now")
	fmt.Println(".")
}
