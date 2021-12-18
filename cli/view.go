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

func PrintAddOrder() {
	fmt.Println(`
#********* Auftrag erstellen ***********
# gib folgendes ein:
# "Datum", "Kunden Id", "Fahrer Id"
`)
}

func PrintUpdateOrder() {
	fmt.Println(`
#********* Order bearbeiten ***********
# gib folgendes ein wobei "ID" von
# einer bestehenden Order sein muss:
# "ID", "Datum", "Kunden Id", "Fahrer Id"
`)
}

func PrintDeleteOrder() {
	fmt.Println(`
#********* Auftrag löschen ***********
# gib die ID des zu löschenden Auftrags ein:
`)
}

func PrintDriverMenu() {
	fmt.Println(`
########################################
#************* Mitarbeiter *************
#********* Wähle eine Option ***********
# 1. Mitarbeiter erstellen
# 2. Mitarbeiter bearbeiten
# 3. Mitarbeiter löschen
#
# q. Zurück
#
##************* Mitarbeiter *************
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


#************* Mitarbeiter *************
`)
}

func PrintAddCustomer() {
	fmt.Println(`
#********* Kunde erstellen ***********
# gib folgendes ein:
# "ID", "Name", "Vorname", "Telefon Nr.", "Strasse", "PLZ", "Ort", "Land" 
`)
}

func PrintUpdateCustomer() {
	fmt.Println(`
#********* Kunde bearbeiten ***********
# gib folgendes ein wobei "ID" von
# einem bestehenden Kunden sein muss:
# "ID", "Name", "Vorname", "Telefon Nr.", "Strasse", "PLZ", "Ort", "Land" 
`)
}

func PrintDeleteCustomer() {
	fmt.Println(`
#********* Kunde löschen ***********
# gib die ID des zu löschenden Kunden ein:
`)
}

func PrintAddProduct() {
	fmt.Println(`
#********* Produkt erstellen ***********
# gib folgendes ein:
# "Beschreibung", "Kategorie Id", "Name"
`)
}

func PrintUpdateProduct() {
	fmt.Println(`
#********* Produkt bearbeiten ***********
# gib folgendes ein wobei "ID" von
# einem bestehenden Produkt sein muss:
# "ID", "Beschreibung", "Kategorie Id", "Name"
`)
}

func PrintDeleteProduct() {
	fmt.Println(`
#********* Produkt löschen ***********
# gib die ID des zu löschenden Produkts ein:
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
