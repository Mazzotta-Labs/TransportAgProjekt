package mysql

import "TransportAgProjekt/model/entity"

func (r *MySqlRepository) FindAllDriver() []entity.Driver {
	var drivers []entity.Driver
	result, err := db.Query("select * from Driver join Vehicle using (id)")
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
	driverId := driver.DriverId
	name := driver.Name
	prename := driver.Prename
	vehicleId := driver.VehicleId
	brand := driver.Brand
	number := driver.Number

	stmt, err := db.Prepare("insert into model values (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(driverId, name, prename, vehicleId, brand, number)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateDriver(driver entity.Driver) {
	//TODO
}

func (r *MySqlRepository) DeleteDriver(driver entity.Driver) {
	//TODO
}
