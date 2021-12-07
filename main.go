package main

import "mvc-booklibrary/cli"

func main() {

	cli.ClearTerminal()
	cli.PrintMenue()
	for true {
		cli.ExecuteCommand()
	}


}
