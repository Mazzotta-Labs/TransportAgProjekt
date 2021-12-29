package entity

// Driver as type gerade mit Fahrzeug normalisiert da jeder Fahrer nur 1 Fahrzeug besitzen kann.
type Driver struct {
	DriverId  int
	Name      string
	Prename   string
	VehicleId int
	Brand     string
	Number    string
}
