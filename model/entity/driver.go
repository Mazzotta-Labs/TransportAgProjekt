package entity

// Driver as type gerade mit Fahrzeug normalisiert da jeder Fahrer nur 1 Fahrzeug besitzen kann.
type Driver struct {
	DriverId  string
	Name      string
	Prename   string
	VehicleId string
	Brand     string
	Number    string
}
