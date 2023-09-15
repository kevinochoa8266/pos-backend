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
		unitPrice INTEGER NOT NULL,
		inventory INTEGER,
		storeId INTEGER,
		FOREIGN KEY(storeId) REFERENCES store(id)
);

CREATE TABLE IF NOT EXISTS bulk(
		productId TEXT,
		bulkPrice INTEGER NOT NULL,
		itemsInPacket INTEGER NOT NULL,
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
		unitPrice,
		inventory,
		storeId
		)
		VALUES("1","chocolate", 5, 10, 1);
		
INSERT INTO product (
		id,
		name,
		unitPrice,
		inventory,
		storeId
		)
		VALUES("2","lolipop", 5, 1, 1);


INSERT INTO bulk(productId, bulkPrice, itemsInPacket)
		VALUES("1", 10, 5);




SELECT * FROM product;

SELECT * FROM product p JOIN bulk b ON p.id=b.productId;

SELECT p.id, p.name, p.unitPrice, b.bulkPrice, p.inventory, b.itemsInPacket, p.storeId FROM product p LEFT JOIN bulk b ON p.id = b.productId;


