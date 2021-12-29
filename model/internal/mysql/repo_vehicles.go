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
	brand := vehicle.Brand
	number := vehicle.Number

	stmt, err := db.Prepare("insert into vehicle values (?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(brand, number)
	if err != nil {
		panic(err.Error())
	}

}

func (r *MySqlRepository) UpdateVehicle(vehicle entity.Vehicle) {
	brand := vehicle.Brand
	number := vehicle.Number
	vehicleId := vehicle.VehicleId

	stmt, err := db.Prepare("update `vehicle` set brand = ?, number = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(brand, number, vehicleId, vehicleId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) DeleteVehicle(vehicleId string) {
	stmt, err := db.Prepare("delete from vehicle where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(vehicleId)
	if err != nil {
		panic(err.Error())
	}
}
