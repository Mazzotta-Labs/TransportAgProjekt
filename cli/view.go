package cli

import "fmt"

// view

func PrintMenue() {
	fmt.Println(`
##########################################
#******** Auftragsplanungs App ***********
#********* Wähle eine Option *************
# 1. Aufträge
# 2. Mitarbeiter
# 3. Kunden
# 4. Produkte
#
# q. Beenden
`)
}

func PrintOrderMenu() {
	fmt.Println(`
#########################################
#*************** Aufträge ***************
#********* Wähle eine Option ************
# 1. Auftrag erstellen
# 2. Auftrag bearbeiten
# 3. Auftrag löschen
#
# q. Zurück
`)
}

func PrintDriverMenu() {
	fmt.Println(`
########################################
#************* Mitarbeiter *************
#********* Wähle eine Option ***********
# 1. Auftrag erstellen
# 2. Auftrag bearbeiten
# 3. Auftrag löschen
#
# q. Zurück
`)
}

func PrintCustomerMenu() {
	fmt.Println(`
########################################
#*************** Kunden ****************
#********* Wähle eine Option ***********
# 1. Kunden erstellen
# 2. Kunden bearbeiten
# 3. Kunden löschen
#
# q. Zurück
`)
}

func PrintProductMenu() {
	fmt.Println(`
########################################
#************** Produkte ***************
#********* Wähle eine Option ***********
# 1. Produkte erstellen
# 2. Produkte bearbeiten
# 3. Produkte löschen
#
# q. Zurück
`)
}

func printGodBye() {
	fmt.Println("Good bye!")
}
