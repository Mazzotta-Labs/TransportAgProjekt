package main

import "TransportAgProjekt1/cli"

func main() {

	cli.ClearTerminal()
	cli.PrintMenue()
	for true {
		cli.ExecuteCommand()
	}

}
