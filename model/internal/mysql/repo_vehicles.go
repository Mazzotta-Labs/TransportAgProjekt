package mysql

import (
	"TransportAgProjekt/model/entity"
)

func (r *MySqlRepository) FindAllVehicle() []entity.Vehicle {
	var vehicles []entity.Vehicle
	result, err := db.Query("select * from vehicle")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var vehicle entity.Vehicle

		err := result.Scan(&vehicle.VehicleId, &vehicle.Brand, &vehicle.Number)
		if err != nil {
			panic(err.Error())
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles
}

func (r *MySqlRepository) AddVehicle(vehicle entity.Vehicle) {
	vehicleId := vehicle.VehicleId
	brand := vehicle.Brand
	//driverId := vehicle.DriverId
	number := vehicle.Number

	stmt, err := db.Prepare("insert into vehicle values (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(vehicleId, brand, number)
	if err != nil {
		panic(err.Error())
	}

}

func (r *MySqlRepository) UpdateVehicle(vehicle entity.Vehicle) {
	//TODO
}

func (r *MySqlRepository) DeleteVehicle(vehicle entity.Vehicle) {
	//TODO
}
