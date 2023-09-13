-- SQLite

DROP TABLE IF EXISTS store;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS bulk;

CREATE TABLE IF NOT EXISTS store(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		address TEXT
);

CREATE TABLE IF NOT EXISTS product(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		inventory INTEGER,
		storeId INTEGER,
		FOREIGN KEY(storeId) REFERENCES store(id)
);

CREATE TABLE IF NOT EXISTS bulk(
		productId TEXT,
		individualPrice INTEGER NOT NULL,
		bulkQuantity INTEGER,
		FOREIGN KEY (productId) REFERENCES product(id)
		);

INSERT INTO store (
		id,
		name,
		address
		)
		VALUES(1, "CASA DULCE", "575 lionstone");


INSERT INTO product (
		id,
		name,
		price,
		inventory,
		storeId
		)
		VALUES("1","chocolate", 5, 10, 1);
INSERT INTO product (
		id,
		name,
		price,
		inventory,
		storeId
		)
		VALUES("2","lolipop", 5, 10, 1);


INSERT INTO bulk(productId, individualPrice, bulkQuantity)
		VALUES("1", 1, 5);




SELECT * FROM product;

SELECT * FROM product p JOIN bulk b ON p.id=b.productId;

SELECT * FROM product p LEFT JOIN bulk b ON p.id = b.productId;


