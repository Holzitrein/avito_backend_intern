CREATE TABLE Users
(
 idUser  int NOT NULL,
 balance DECIMAL(18,2) NOT NULL,
 reserve DECIMAL(18,2) NOT NULL,
 CONSTRAINT PK_Users PRIMARY KEY ( idUser )
);

CREATE TABLE History_add
(
 idAdd  SERIAL,
 amount DECIMAL(18,2) NOT NULL,
 idUser int NOT NULL,
 timeAdd   timestamp NOT NULL,
 note   varchar(50) NOT NULL,
 CONSTRAINT PK_IdAdd PRIMARY KEY ( idAdd ),
 FOREIGN KEY ( idUser ) REFERENCES Users ( idUser )
);


CREATE TABLE Services
(
 idService int NOT NULL,
 nameSer      varchar(50) NOT NULL,
 CONSTRAINT PK_IdService PRIMARY KEY ( idService )
);

CREATE TABLE Orders
(
 idOrder   int8range NOT NULL,
 idUser    int NOT NULL,
 idService int NOT NULL,
 price     DECIMAL(18,2) NOT NULL,
 created   timestamp NOT NULL,
 statusOrder int NOT NULL,
 CONSTRAINT PK_IdOrder PRIMARY KEY ( idOrder ),
 FOREIGN KEY ( idUser ) REFERENCES Users ( idUser ),
 FOREIGN KEY ( idService ) REFERENCES Services ( idService )
);

INSERT INTO users (iduser,balance,reserve) VALUES (
  '2',
  '250',
  '0'
);