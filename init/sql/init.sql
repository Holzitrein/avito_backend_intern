SELECT 'CREATE DATABASE avito_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'avito_db')\gexec
set names 'utf8';


CREATE TABLE Users
(
 IdUser  int NOT NULL,
 Balance money NOT NULL,
 Reserve money NOT NULL,
 CONSTRAINT PK_1 PRIMARY KEY ( IdUser )
);

CREATE TABLE History_add
(
 IdAdd  int8range NOT NULL,
 Amount money NOT NULL,
 IdUser int NOT NULL,
 "Date"   date NOT NULL,
 Note   varchar(50) NOT NULL,
 CONSTRAINT PK_1 PRIMARY KEY ( IdAdd ),
 CONSTRAINT FK_1 FOREIGN KEY ( IdUser ) REFERENCES Users ( IdUser )
);


CREATE TABLE Services
(
 IdService int NOT NULL,
 Name      varchar(50) NOT NULL,
 CONSTRAINT PK_1 PRIMARY KEY ( IdService )
);

CREATE TABLE Orders
(
 IdOrder   int8range NOT NULL,
 IdUser    int NOT NULL,
 IdService int NOT NULL,
 Price     money NOT NULL,
 "Date"      date NOT NULL,
 CONSTRAINT PK_1 PRIMARY KEY ( IdOrder ),
 CONSTRAINT FK_2 FOREIGN KEY ( IdUser ) REFERENCES Users ( IdUser ),
 CONSTRAINT FK_3 FOREIGN KEY ( IdService ) REFERENCES Services ( IdService )
);


