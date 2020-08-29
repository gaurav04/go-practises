

CREATE DATABASE `inventorydb`;



CREATE TABLE `inventorydb`.`products` (

  `productId` INT NOT NULL AUTO_INCREMENT,

  `manufacturer` VARCHAR(255) NOT NULL,

  `sku` VARCHAR(60) NOT NULL,

  `upc` VARCHAR(60) NOT NULL,

  `pricePerUnit` DECIMAL(13,2) NOT NULL,

  `quantityOnHand` INT NOT NULL,

  `productName` VARCHAR(255) NOT NULL,

  PRIMARY KEY (`productId`));



INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Johns-Jenkins","p5z343vdS","939581000000",497.45,9703,"sticky note");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hessel, Schimmel and Feeney","i7v300kmx","740979000000",282.29,9217,"leg warmers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Swaniawski, Bartoletti and Bruen","q0L657ys7","111730000000",436.26,5905,"lamp shade");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Runolfsdottir, Littel and Dicki","x78426lq1","93986215015",537.90,2642,"flowers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kuhn, Cronin and Spencer","r4X793mdR","260149000000",112.10,6144,"clamp");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Quigley, Casper and Boyer","g2P499xrM","390162000000",593.53,6507,"twister");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Gutmann and Sons","v4j250hbi","465922000000",88.97,4348,"clay pot");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Bins-Hansen","d3E278nk2","20072026056",933.35,4577,"tooth picks");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Jones, Braun and Ratke","o3w911oal","879638000000",426.23,1882,"mirror");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Upton-Mraz","k0R875prt","877387000000",630.61,4036,"rug");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schneider, Douglas and Franecki","h3t822kaB","507019000000",13.67,2285,"headphones");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Weimann, Waelchi and Kassulke","t8A474iuv","669100000000",369.46,5409,"balloon");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Moore-Gibson","c4i321mvh","600535000000",250.98,6219,"lip gloss");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schuppe, Cummerata and Hammes","l9V771xw1","677102000000",577.90,1104,"sidewalk");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ward-Quigley","x7b275hk2","110459000000",642.19,9371,"pen");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Cassin Inc","b27590nys","920302000000",145.19,5382,"outlet");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Shanahan Inc","j4y570kml","850798000000",409.02,4567,"blanket");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Upton-Runolfsdottir","m2K116lkV","10335859487",815.26,7083,"lotion");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Lubowitz Group","w61375szc","286418000000",700.53,2224,"socks");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Padberg, Grady and Mueller","f1D653dwZ","62442197606",176.67,4544,"CD");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Barrows-Gibson","g5F556bj2","667449000000",66.44,138,"model car");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Feeney, Glover and Johns","o72175oxZ","353375000000",454.12,4712,"soda can");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ankunding, Haag and Corwin","n5m968dmo","356116000000",852.15,5325,"fridge");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Armstrong, Mohr and Bradtke","o6O935rpv","451505000000",275.21,7339,"bottle cap");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schmidt and Sons","j32187bai","393839000000",425.45,5434,"photo album");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Fadel, Robel and Turner","g7K279msW","227667000000",774.24,2224,"twezzers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Cassin, Kreiger and Wisoky","o4o723dsy","663345000000",899.61,3985,"glow stick");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Baumbach LLC","m8p061fi3","156939000000",663.10,9644,"truck");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Cummerata, Hills and Anderson","v4d029mtP","130529000000",223.74,4173,"purse");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kautzer, Bradtke and Stracke","i2w704pvs","498143000000",929.37,1465,"bowl");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Bernhard, Fisher and Reinger","p8A041wxj","302419000000",935.96,7069,"sketch pad");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Stamm and Sons","i1q225viC","190041000000",341.29,9364,"sponge");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Senger Group","y6u840ste","981691000000",785.00,9635,"cookie jar");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Conroy-Little","i2b810vfi","151356000000",956.12,8016,"sharpie");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Huels Inc","r2Y361mak","315225000000",78.00,8471,"shawl");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pacocha and Sons","m6f710dvL","54158362672",392.98,5583,"face wash");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Bednar, Bednar and Skiles","d3l280ulV","916805000000",761.49,5408,"toothpaste");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pouros Group","h9z038nu2","213707000000",49.02,8531,"ring");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Feil, Wehner and Russel","j9X185vnN","959689000000",828.92,3896,"door");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Mayert, Barton and Feeney","b9J149udy","377314000000",863.18,1618,"towel");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Lindgren and Sons","d5X474kg4","544254000000",139.86,7135,"vase");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Lueilwitz-Stehr","b81957uq9","766152000000",758.59,352,"sun glasses");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Goldner, Bogisich and Zieme","o1O578rpk","541391000000",146.27,332,"soy sauce packet");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Braun-Auer","o1d743wqj","234837000000",205.50,6497,"thermostat");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Block, Wehner and Mann","a7b451lhB","211640000000",658.75,795,"playing card");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Jenkins, Hickle and Zemlak","a7F511yoM","262647000000",0.95,7017,"camera");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Cassin, Jaskolski and Glover","l2D749trK","507953000000",414.79,2005,"key chain");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hermiston and Sons","u12812ak7","412832000000",848.51,5802,"zipper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ankunding Group","n6K818cdd","726622000000",450.32,4164,"ipod");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Halvorson, Bradtke and O'Connell","q6V805zsg","644787000000",220.37,4611,"plastic fork");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hauck Group","d8N567aas","800212000000",579.81,3962,"shovel");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schamberger, Klein and Kertzmann","q88014dfs","97978011389",922.06,790,"chapter book");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Oberbrunner-Hodkiewicz","z4E342hzS","30785250582",24.97,182,"picture frame");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Tremblay, Kunde and Schoen","s5f791nqQ","200749000000",718.77,3881,"drill press");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Wolff, Lynch and Watsica","q7I787jbj","120736000000",601.53,5210,"toe ring");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Connelly, Parker and Raynor","g1w131bfk","65741106217",208.49,4350,"ice cube tray");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Spencer, Wehner and Schaden","k3l585yjg","17719032131",846.91,1711,"bed");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pfannerstill-Rippin","t4l717zlj","593405000000",764.69,2022,"washing machine");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Crona-Wunsch","f7F573eq3","671359000000",459.69,1506,"couch");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Goyette Inc","t6h083sru","122200000000",400.61,8524,"bananas");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("O'Keefe and Sons","y1O915yp6","564495000000",942.71,4454,"coasters");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Eichmann, Reichert and Von","o8K265mog","437830000000",784.47,1508,"food");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("VonRueden-Nader","j0P337hiR","517934000000",2.79,1139,"puddle");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hills-Price","u46967ytd","184724000000",672.06,9942,"glasses");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Abshire, Kihn and Rodriguez","v9W856avM","647320000000",230.59,2901,"computer");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Miller, Corkery and Eichmann","m7I885qem","33375947636",369.37,3751,"street lights");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Adams, Kuhn and Gerlach","q5y373cuA","129448000000",279.98,3852,"chocolate");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kuphal-Davis","w4u543iwl","366183000000",325.61,5487,"eraser");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kovacek-Kulas","f4i153lo1","397448000000",645.23,9411,"cinder block");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Wyman-Mayer","n0t343fmG","987399000000",343.77,2965,"eye liner");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Oberbrunner-Smitham","m6N173baL","769622000000",505.85,748,"mouse pad");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Frami Group","m3A594rab","572520000000",128.77,8150,"sand paper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ferry and Sons","x5f817sz8","313736000000",559.33,6688,"mop");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Orn-Deckow","s18515vgH","755274000000",203.90,8179,"USB drive");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Donnelly, Cremin and Howell","o19221uqo","272827000000",29.99,3389,"toilet");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Goodwin-Klocko","o0R990dv9","837386000000",839.76,1492,"video games");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hermann, Weissnat and Powlowski","n4k323vb6","675663000000",248.68,7970,"paint brush");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Sanford Group","i2L410spm","9255390626",612.23,4249,"newspaper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Mraz Inc","y8L982qu4","464824000000",427.35,6718,"tree");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("McGlynn, Cronin and Schneider","c8u327yrd","387752000000",931.77,8963,"carrots");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schimmel Group","r01817zs3","241866000000",53.66,6358,"helmet");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Fay-Howell","y7X879jd0","721123000000",413.48,688,"magnet");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Bergnaum Group","z17497rqy","243664000000",913.62,4543,"greeting card");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kilback Group","g4e101fq9","131218000000",267.88,8846,"radio");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hilpert, Lesch and Boyle","i4Q336env","677500000000",91.68,2304,"chalk");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Mohr Group","i5n769shB","183153000000",969.15,9129,"lace");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Simonis LLC","r89646jii","550821000000",830.22,9758,"hair tie");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Braun Inc","w23943mdr","158033000000",613.01,1505,"flag");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pfeffer and Sons","n1M345ck6","812439000000",700.05,7421,"canvas");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Crist-Yost","t5v999mc8","169032000000",410.42,533,"bow");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Powlowski-Olson","s3W460csY","937380000000",848.63,2273,"controller");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Marquardt-Lebsack","t0S420rlp","949602000000",213.18,7188,"bookmark");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pacocha Group","m1o387nbL","352225000000",820.03,8386,"tomato");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Christiansen-Kertzmann","o0S408raF","518584000000",875.72,9501,"chair");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Halvorson-Nicolas","f81054cuk","285137000000",749.54,5822,"plate");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Vandervort, Morar and Funk","q6K118dx7","738902000000",420.06,2892,"window");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Moore Group","r51014gfv","192860000000",955.25,6531,"fork");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schuster, Predovic and Hamill","s3F352ekS","329256000000",849.20,3059,"desk");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Batz, Cormier and Torp","r0g162mf6","688529000000",871.60,7022,"phone");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ruecker, Krajcik and Wilkinson","y98114miZ","573871000000",925.16,3878,"remote");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Stoltenberg Inc","a1Z306tuX","89886202324",220.52,4694,"clock");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Rosenbaum, Muller and Schamberger","d9q172ah3","46321145111",328.86,3128,"pencil");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Funk, Murray and Ledner","w6E748lu8","888163000000",373.36,9986,"apple");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Pacocha Inc","s1V240pzF","285995000000",185.76,7936,"beef");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Miller and Sons","q4Q271kpz","553355000000",120.27,8933,"boom box");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Graham Inc","o4T091vjj","114841000000",834.12,8185,"pants");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schoen Inc","u3d665sxK","130502000000",183.24,561,"hanger");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Koss, Sporer and Cremin","j4n828pvd","611714000000",610.49,6026,"money");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Stoltenberg, Hoppe and Reinger","h51298cmu","988759000000",422.50,8152,"cat");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kuhlman, Kessler and Schamberger","q4t839kpq","958257000000",830.45,1522,"car");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Little, Marquardt and Wilderman","k8C714br7","327578000000",915.42,2350,"white out");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Von, Auer and Hansen","x5J737iaA","813584000000",830.92,930,"candy wrapper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Shields, Krajcik and Gusikowski","l1S070ooK","746699000000",891.08,7693,"button");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kuhlman Inc","a5i730xnw","166336000000",973.88,6074,"scotch tape");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Cruickshank LLC","z8Y652qyj","795739000000",332.22,6131,"speakers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("O'Kon-Doyle","m6H046xqT","276336000000",148.31,8105,"soap");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Rowe, Carroll and Dickinson","s5C945llY","531314000000",111.62,4545,"perfume");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ferry-Leffler","c75832jgw","772628000000",496.34,1388,"monitor");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Gusikowski-Thompson","f9Z903vzn","915383000000",293.01,4830,"credit card");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Johns, Kassulke and Bergstrom","w4B521ot1","22085198947",354.37,9986,"brocolli");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Macejkovic-Nitzsche","f09234xcN","624018000000",911.90,3874,"sofa");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schimmel-Zemlak","t67112jvx","914348000000",456.08,4000,"sailboat");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Runolfsdottir-Willms","b2N103kqm","223169000000",675.66,2808,"floor");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Krajcik Group","c9Y079hme","239920000000",769.42,2057,"cork");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Flatley Inc","q1s356llD","30618129480",78.09,7623,"wagon");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Miller Group","k7M989rkI","973983000000",80.13,1624,"table");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Reichert and Sons","m6U136ueF","226098000000",342.44,6506,"thread");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Ryan-Heaney","k0I251xuI","962369000000",152.72,760,"slipper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Beer Inc","w5g494kmk","684954000000",369.68,1494,"hair brush");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Kuhn, Hagenes and Connelly","f3j155khy","505178000000",273.22,8428,"nail clippers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Collins and Sons","c06175duk","532510000000",942.91,4655,"needle");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Farrell-Becker","s2s744vlb","260918000000",924.29,9277,"cup");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Wolf, Schamberger and Tromp","f0U474udL","412375000000",443.43,1760,"stop sign");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hickle Group","h9B811thn","447973000000",739.63,1789,"candle");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Davis, Stokes and King","j9i386jpv","561484000000",704.29,5333,"piano");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Hauck-Gaylord","c1n385sxC","109058000000",420.65,1502,"buckel");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Davis, Mertz and Walker","d7m478vtd","874484000000",718.56,9262,"keys");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Walker-Hansen","b0h256tb5","98391553823",746.92,8777,"pool stick");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("McKenzie LLC","y16570wxA","230658000000",35.26,7921,"bag");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Fritsch-Rogahn","b1b632mvc","692603000000",267.59,3268,"mp3 player");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Beahan, Kiehn and Hansen","l3y775oqQ","294992000000",690.15,2215,"shampoo");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Prohaska LLC","g69742enX","448980000000",407.03,6258,"rubber band");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("O'Hara-Bartell","i09672gsM","939989000000",702.11,7941,"keyboard");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Wyman-Osinski","v0t980jbV","107171000000",97.10,9843,"tv");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Durgan, Parker and Lowe","j9o570aqw","625801000000",783.26,6651,"rubber duck");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Bruen-Weissnat","z3r510axb","317094000000",51.79,5416,"packing peanuts");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Swaniawski-Carter","x7e598riU","753389000000",392.33,334,"bread");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Daniel-Bahringer","c8g204enF","476957000000",179.62,1791,"checkbook");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Jacobi and Sons","d16352ffQ","451609000000",900.93,8297,"glass");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Goyette-Kulas","k0V573ztB","791237000000",842.54,2655,"bottle");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Graham, Feest and Jerde","j7x764yiQ","93894957591",328.94,2527,"toothbrush");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("McLaughlin, Ankunding and Schowalter","a46529cqG","530809000000",695.46,3139,"shirt");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Schumm, O'Kon and Beer","q4b168njp","51164055605",575.48,9113,"seat belt");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Windler LLC","s9a742biR","101409000000",160.95,8929,"fake flowers");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Reynolds, Zemlak and Marquardt","f7K945cky","884392000000",672.43,9118,"conditioner");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Little, Durgan and Walsh","j2Q575ffi","809558000000",567.63,6227,"paper");  

INSERT INTO `inventorydb`.`products` (`manufacturer`, `sku`, `upc`, `pricePerUnit`, `quantityOnHand`, `productName`) VALUES ("Mohr, Padberg and Dickens","u2O211ora","845643000000",257.46,5025,"knife");  
