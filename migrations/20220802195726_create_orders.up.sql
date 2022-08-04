CREATE TABLE orders(
order_uid varchar(40) NOT NULL UNIQUE,
	json_data JSONB
);
