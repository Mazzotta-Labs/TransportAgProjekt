package main

import "TransportAgProjekt/cli"

func main() {

	cli.ClearTerminal()
	cli.PrintMenue()
	for true {
		cli.ExecuteCommand()
	}

}
