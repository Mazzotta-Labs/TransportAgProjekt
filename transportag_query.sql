CREATE DATABASE transportag;

USE transportag;

CREATE TABLE Address (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	street VARCHAR(100) NOT NULL,
	plz VARCHAR(100) NOT NULL,
	town VARCHAR(100) NOT NULL,
	country VARCHAR(100)
);

CREATE TABLE Customer (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	prename VARCHAR(100) NOT NULL,
	tel_nr VARCHAR(50),
    address_id INT NOT NULL,
    FOREIGN KEY(address_id) REFERENCES Address(id)
);

CREATE TABLE Vehicle (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	brand VARCHAR(100) NOT NULL,
	number VARCHAR(30) NOT NULL
);

CREATE TABLE Driver (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	prename VARCHAR(100) NOT NULL,
    vehicle_id INT NOT NULL,
    FOREIGN KEY(vehicle_id) REFERENCES Vehicle(id)
);

CREATE TABLE `Order` (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	order_date DATE NOT NULL,
    customer_id INT NOT NULL,
    driver_id INT NOT NULL,
    FOREIGN KEY(customer_id) REFERENCES Customer(id),
    FOREIGN KEY(driver_id) REFERENCES Driver(id)
);

CREATE TABLE Category (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    weight FLOAT NOT NULL,
	fragile BOOL DEFAULT(false),
	name VARCHAR(100) NOT NULL
);

CREATE TABLE Product (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(1000) NOT NULL,
	name VARCHAR(500) NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY(category_id) REFERENCES Category(id)
);

CREATE TABLE OrderToProduct(
	order_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY(order_id) REFERENCES `Order`(id)
    ON DELETE CASCADE,
    FOREIGN KEY(product_id) REFERENCES Product(id)
    ON UPDATE CASCADE,
  UNIQUE(order_id,product_id)
);


