CREATE DATABASE `inventorydb`;

CREATE TABLE `inventorydb`.`products` (

  `productId` INT NOT NULL AUTO_INCREMENT,

  `manufacturer` VARCHAR(255) NOT NULL,

  `pricePerUnit` DECIMAL(13,2) NOT NULL,

  `quantityOnHand` INT NOT NULL,

  `productName` VARCHAR(255) NOT NULL,

  PRIMARY KEY (`productId`));


INSERT INTO `inventorydb`.`products` (`manufacturer`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Amul",100,20,"Chocalte");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Patanjali",220,10,"Soap");  
