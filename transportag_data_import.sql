
INSERT INTO `Category` (`weight`,`fragile`,`name`)
VALUES
  (52,"0","Vel Nisl Inc."),
  (19,"0","Metus Facilisis Lorem Limited"),
  (69,"0","Sociis Limited"),
  (2,"1","Mus Donec Dignissim Associates"),
  (21,"1","Pede Ultrices LLP");
  
INSERT INTO `Product` (`description`,`name`,`category_id`)
VALUES
  ("nibh lacinia orci,","Duncan",1),
  ("hendrerit","Mueller",5),
  ("vitae,","Leon",1),
  ("et, commodo at, libero. Morbi","Miller",2),
  ("tristique ac,","Eaton",4);
  
INSERT INTO `Vehicle` (`brand`,`number`)
VALUES
  ("Sed Congue LLC","68-723"),
  ("Posuere Corp.","LC7 4EQ"),
  ("Nec Quam LLP","831177"),
  ("Fringilla Mi Lacinia Corp.","27463"),
  ("Eget Industries","94-861");
  
INSERT INTO `Driver` (`name`,`prename`,`vehicle_id`)
VALUES
  ("Downs","Hollee",2),
  ("Rich","Gay",3),
  ("Richardson","Nelle",4),
  ("Sosa","Gavin",1),
  ("Larsen","Mallory",3);
  
INSERT INTO `Address` (`street`,`plz`,`town`,`country`)
VALUES
  ("963-8063 Nonummy Rd.","532512","Arequipa","Chile"),
  ("Ap #200-1244 Enim, St.","267715","Auburn","South Korea"),
  ("772-6450 Imperdiet, St.","6582","Çaldıran","Netherlands"),
  ("752-3574 Vitae Street","21024","Vandoeuvre-lès-Nancy","Belgium"),
  ("497-387 Eget Avenue","41205","Camporotondo di Fiastrone","New Zealand");

INSERT INTO `Customer` (`name`,`prename`,`tel_nr`,`address_id`)
VALUES
  ("Mooney","Abdul","(0137) 87688346",5),
  ("Beard","Plato","(0396) 91316760",3),
  ("Munoz","Shana","(07026) 7278683",3),
  ("Stewart","Quon","(041) 41863725",1),
  ("Patton","Risa","(043) 89285262",5);
  
INSERT INTO `Order` (`order_date`,`customer_id`,`driver_id`)
VALUES
  ("2022-11-24 19:52:11",5,1),
  ("2022-06-16 20:35:40",3,3),
  ("2021-04-19 15:10:36",3,4),
  ("2022-04-20 01:02:19",1,5),
  ("2022-10-10 22:36:10",5,3);

INSERT INTO `OrderToProduct` (`order_id`,`product_id`)
VALUES
  (5,1),
  (3,3),
  (3,4),
  (1,5),
  (5,3),  
  (2,2),
  (3,2);














