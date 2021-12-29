package mysql

import "TransportAgProjekt/model/entity"

func (r *MySqlRepository) FindAllDriver() []entity.Driver {
	var drivers []entity.Driver
	result, err := db.Query("select * from driver d join vehicle v using (id)")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var driver entity.Driver

		err := result.Scan(&driver.DriverId, &driver.Name, &driver.Prename, &driver.VehicleId, &driver.Brand, &driver.Number)
		if err != nil {
			panic(err.Error())
		}
		drivers = append(drivers, driver)
	}
	return drivers
}

func (r *MySqlRepository) AddDriver(driver entity.Driver) {
	name := driver.Name
	prename := driver.Prename
	vehicleId := driver.VehicleId

	stmt, err := db.Prepare("insert into driver (name, prename, vehicle_id) values (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(name, prename, vehicleId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateDriver(driver entity.Driver) {
	driverId := driver.DriverId
	name := driver.Name
	prename := driver.Prename
	vehicleId := driver.VehicleId

	stmt, err := db.Prepare("update `driver` set name = ?, prename = ?, vehicle_id = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(name, prename, vehicleId, driverId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) DeleteDriver(DriverId string) {
	stmt, err := db.Prepare("delete from driver where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(DriverId)
	if err != nil {
		panic(err.Error())
	}
}
