CREATE TABLE IF NOT EXISTS "customer" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" VARCHAR(50) NOT NULL,
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "phone_number" VARCHAR(50) NOT NULL,
  "role" SMALLINT NOT NULL DEFAULT 2,
  "deleted" BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "zipcode" (
  "zip" VARCHAR(10) PRIMARY KEY,
  "city" VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS "address" (
  "id" SERIAL PRIMARY KEY,
  "zipcode" VARCHAR(10) NOT NULL REFERENCES "zipcode" ("zip"),
  "street" VARCHAR(50) NOT NULL,
  "deleted" BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "customer_address" (
  "customer_id" INTEGER NOT NULL REFERENCES "customer" ("id"),
  "address_id" INTEGER NOT NULL REFERENCES "address" ("id"),
  "primary_address" BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY ("customer_id", "address_id")
);

CREATE TABLE IF NOT EXISTS "category" (
  "name" VARCHAR(50) PRIMARY KEY,
  "description" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "manufacturer" (
  "name" VARCHAR(50) PRIMARY KEY,
  "description" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS "discount" (
  "id" SERIAL PRIMARY KEY,
  "percentage" NUMERIC(4,2) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "start_at" TIMESTAMPTZ NOT NULL,
  "end_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS "product" (
  "id" SERIAL PRIMARY KEY,
  "category_name" VARCHAR(50) NOT NULL REFERENCES "category" ("name"),
  "manufacturer_name" VARCHAR(50) NOT NULL REFERENCES "manufacturer" ("name"),
  "description" VARCHAR(255) NOT NULL,
  "price" NUMERIC(9,2) NOT NULL,
  "stock" INTEGER NOT NULL
);

CREATE INDEX product_description_idx ON product (description);
-- full-text search capabilities. This is much more powerful and flexible than pattern-matching (LIKE)
CREATE INDEX product_description_fts_idx ON product USING gin(to_tsvector('english', description));

CREATE TABLE IF NOT EXISTS "product_discount" (
  "product_id" INTEGER NOT NULL REFERENCES "product" ("id"),
  "discount_id" INTEGER NOT NULL REFERENCES "discount" ("id"),
  PRIMARY KEY ("product_id", "discount_id")
);

CREATE TABLE IF NOT EXISTS "customer_product_review" (
  "customer_id" INTEGER NOT NULL REFERENCES "customer" ("id"),
  "product_id" INTEGER NOT NULL REFERENCES "product" ("id"),
  "stars" NUMERIC(2,1) NOT NULL,
  "comment" VARCHAR(255),
  "deleted" BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY ("customer_id", "product_id")
);

CREATE TABLE IF NOT EXISTS "shopping_order" (
  "id" SERIAL PRIMARY KEY,
  "customer_id" INTEGER NOT NULL REFERENCES "customer" ("id"),
  "placed_at" TIMESTAMPTZ NOT NULL,
  "total_amount" NUMERIC(9,2) NOT NULL,
  "status" VARCHAR(50) NOT NULL DEFAULT 'Pending'
);

CREATE TABLE IF NOT EXISTS "item" (
  "id" SERIAL PRIMARY KEY,
  "shopping_order_id" INTEGER NOT NULL REFERENCES "shopping_order" ("id"),
  "product_id" INTEGER NOT NULL REFERENCES "product" ("id"),
  "quantity" SMALLINT NOT NULL,
  "sub_total" NUMERIC(9,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS "payment_method" (
  "id" SERIAL PRIMARY KEY,
  "method" VARCHAR(50) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "fee" NUMERIC(5,2) NOT NULL,
  "deprecated" BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "payment" (
  "id" SERIAL PRIMARY KEY,
  "shopping_order_id" INTEGER NOT NULL REFERENCES "shopping_order" ("id"),
  "payment_method_id" INTEGER NOT NULL REFERENCES "payment_method" ("id"),
  "status" VARCHAR(50) NOT NULL DEFAULT 'Pending'
);

CREATE TABLE IF NOT EXISTS "shipping_method" (
  "id" SERIAL PRIMARY KEY,
  "method" VARCHAR(50) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "fee" NUMERIC(5,2) NOT NULL,
  "deprecated" BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS "shipping" (
  "id" SERIAL PRIMARY KEY,
  "shopping_order_id" INTEGER NOT NULL REFERENCES "shopping_order" ("id"),
  "shipping_method_id" INTEGER NOT NULL REFERENCES "shipping_method" ("id"),
  "address_id" INTEGER NOT NULL REFERENCES "address" ("id"),
  "shipped_at" TIMESTAMPTZ NOT NULL,
  "delivery_at" TIMESTAMPTZ NOT NULL
);

INSERT INTO "zipcode" ("zip", "city") VALUES ('0010', 'Oslo');
INSERT INTO "zipcode" ("zip", "city") VALUES ('2000', 'Lillestrøm');
INSERT INTO "zipcode" ("zip", "city") VALUES ('2372', 'Brøttum');
INSERT INTO "zipcode" ("zip", "city") VALUES ('3010', 'Drammen');
INSERT INTO "zipcode" ("zip", "city") VALUES ('4010', 'Stavanger');
INSERT INTO "zipcode" ("zip", "city") VALUES ('4609', 'Kardemommeby');
INSERT INTO "zipcode" ("zip", "city") VALUES ('5010', 'Bergen');
INSERT INTO "zipcode" ("zip", "city") VALUES ('6010', 'Ålesund');
INSERT INTO "zipcode" ("zip", "city") VALUES ('7010', 'Trondheim');
INSERT INTO "zipcode" ("zip", "city") VALUES ('8010', 'Bodø');
INSERT INTO "zipcode" ("zip", "city") VALUES ('9010', 'Tromsø');
INSERT INTO "zipcode" ("zip", "city") VALUES ('9170', 'Longyearbyen');

INSERT INTO "category" ("name", "description") VALUES ('Computer equipment', 'Computer equipment is the physical components of a computer system');
INSERT INTO "category" ("name", "description") VALUES ('Gaming', 'Computers, equipment, and accessories designed for gaming');
INSERT INTO "category" ("name", "description") VALUES ('TV, Sound & Image', 'TV, Sound & Image equipment');
INSERT INTO "category" ("name", "description") VALUES ('PC & Tablets', 'Computers, tablets and accessories');
INSERT INTO "category" ("name", "description") VALUES ('Phones & Watches', 'Phones, watches and accessories');
INSERT INTO "category" ("name", "description") VALUES ('Appliances', 'Home appliances');
INSERT INTO "category" ("name", "description") VALUES ('Home & Leisure', 'Home and leisure equipment');

INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('ASUS', 'ASUS is a multinational computer hardware and electronics company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Acer', 'Acer Inc. is a Taiwanese multinational hardware and electronics corporation', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Dell', 'Dell is an American multinational computer technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('HP', 'The Hewlett-Packard Company, commonly shortened to Hewlett-Packard or HP, was an American multinational information technology company', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Lenovo', 'Lenovo Group Limited, often shortened to Lenovo, is a Chinese multinational technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Apple', 'Apple Inc. is an American multinational technology company', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Samsung', 'Samsung is a South Korean multinational conglomerate', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Sony', 'Sony Corporation is a Japanese multinational conglomerate corporation', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Nintendo', 'Nintendo Co., Ltd. is a Japanese multinational consumer electronics and video game company', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Microsoft', 'Microsoft Corporation is an American multinational technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Intel', 'Intel Corporation is an American multinational corporation and technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('AMD', 'Advanced Micro Devices, Inc. is an American multinational semiconductor company', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Nvidia', 'Nvidia Corporation is an American multinational technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('MSI', 'Micro-Star International Co., Ltd is a Taiwanese multinational information technology corporation', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Gigabyte', 'Gigabyte Technology Co., Ltd. is a Taiwanese manufacturer and distributor of computer hardware', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Corsair', 'Corsair Components, Inc. is an American computer peripherals and hardware company', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Kingston', 'Kingston Technology Corporation is an American, privately-held, multinational computer technology corporation', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Crucial', 'Crucial Technology is a brand of Micron Technology, Inc., one of the largest memory manufacturers in the world', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Seagate', 'Seagate Technology PLC is an American data storage company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Sonos', 'Sonos, Inc. is an American developer and manufacturer of audio products', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Bosch', 'Robert Bosch GmbH, or Bosch, is a German multinational engineering and technology company', '1234567890');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('AVA', 'AVA is a Norwegian electronics retailer', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Iiglo', 'Iiglo is a Norwegian electronics retailer', '123456789');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Philips', 'Koninklijke Philips N.V. is a Dutch multinational conglomerate corporation', '0987654321');
INSERT INTO "manufacturer" ("name", "description", "phone_number") VALUES ('Roborock', 'Roborock is a Chinese manufacturer of robotic vacuum cleaners', '1234567890');

INSERT INTO "discount" ("percentage", "description", "start_at", "end_at") VALUES ('0.20', 'Summer sale', '2024-01-06 00:00:00', '2024-07-31 23:59:59');
INSERT INTO "discount" ("percentage", "description", "start_at", "end_at") VALUES ('0.33', 'Autumn sale', '2024-09-01 00:00:00', '2024-11-30 23:59:59');
INSERT INTO "discount" ("percentage", "description", "start_at", "end_at") VALUES ('0.50', 'Black Friday', '2024-11-29 00:00:00', '2024-11-30 23:59:59');
INSERT INTO "discount" ("percentage", "description", "start_at", "end_at") VALUES ('0.10', 'Winter sale', '2024-12-01 00:00:00', '2024-12-31 23:59:59');

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Computer equipment', 'ASUS', 'ASUS 27" ROG Strix XG27AQ', 5995.00, 20);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Computer equipment', 'ASUS', 'ASUS 27" gamingskjerm TUF VG279QM', 3995.00, 10);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Computer equipment', 'MSI', 'MSI 34" Curved MAG342CQR', 5490.00, 10);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Computer equipment', 'Samsung', 'Samsung 57" Odyssey Neo G9 S57CG95', 24990.00, 3);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Gaming', 'Sony', 'Playstation 5 slim', 6790.00, 7);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Gaming', 'Microsoft', 'Xbox Series X', 6349.00, 5);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Gaming', 'Nintendo', 'Nintendo Switch OLED 2021 64GB (white)', 4090.00, 32);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Gaming', 'ASUS', 'ASUS ROG Strix G16CH', 36990.00, 2);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('TV, Sound & Image', 'Samsung', 'Samsung 65" QN85D NEO QLED 4K TV TQ65QN85D', 19990.00, 5);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('TV, Sound & Image', 'Sony', 'Sony 85" LED 4K Google TV XR85X90L', 34990.00, 4);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('TV, Sound & Image', 'Sonos', 'Sonos Arc Soundboard (white)', 9990.00, 17);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('TV, Sound & Image', 'ASUS', 'ASUS LED projektor ZenBeam E2', 3695.00, 1);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('PC & Tablets', 'Apple', 'Macbook Pro 14 M3 Pro (2023) 512GB (black)', 29990.00, 27);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('PC & Tablets', 'Microsoft', 'Surface Pro 8 i5 8GB 256GB (black)', 14990.00, 12);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('PC & Tablets', 'Lenovo', 'Lenovo Tab P11 Pro 128GB (black)', 4990.00, 8);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('PC & Tablets', 'Acer', 'Acer Chromebook Spin 713 256GB (black)', 6990.00, 3);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Phones & Watches', 'Apple', 'iPhone 13 Pro 256GB (black)', 14990.00, 15);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Phones & Watches', 'Samsung', 'Samsung Galaxy S21 Ultra 5G 256GB (black)', 10990.00, 9);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Phones & Watches', 'Apple', 'Apple Watch Ultra 2 49mm LTE Titan (L)', 10990.00, 42);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Phones & Watches', 'Samsung', 'Samsung Galaxy Watch 4 44mm LTE (black)', 4990.00, 18);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Appliances', 'Bosch', 'Bosch Series 6 WAU28PS0SN iDos Washing Machine (white)', 11490.00, 4);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Appliances', 'Bosch', 'Bosch Serie 6 SMU6ZCS00S Dishwasher (steel)', 14990.00, 7);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Appliances', 'Bosch', 'Bosch Series 6 KGN36VLEA Fridge (white)', 9990.00, 3);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Appliances', 'Bosch', 'Bosch Series 6 HBS573BS0S Oven (steel)', 7990.00, 2);

INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Home & Leisure', 'AVA', 'AVA Master P70 Large Bundle', 3990.00, 2);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Home & Leisure', 'Iiglo', 'iiglo IIAC12000W aircondition with WIFI', 4990.00, 6);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Home & Leisure', 'Philips', 'Philips OneBlade Pro 360 Face + Body QP6551/15', 1099.00, 9);
INSERT INTO "product" ("category_name", "manufacturer_name", "description", "price", "stock") VALUES ('Home & Leisure', 'Roborock', 'Roborock S7 Robot Vacuum Cleaner (black)', 8990.00, 3);

INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('Vipps', 'Vipps is a Norwegian mobile payment application', 15.00);
INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('Klarna', 'Klarna is a Swedish bank that provides online financial services', 25.00);
INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('Visa', 'Visa is an American multinational financial services corporation', 10.00);
INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('Mastercard', 'Mastercard is an American multinational financial services corporation', 0.00);
INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('PayPal', 'PayPal Holdings, Inc. is an American company operating an online payments system', 20.00);
INSERT INTO "payment_method" ("method", "description", "fee") VALUES ('Invoice', 'Invoice payment', 40.00);

INSERT INTO "shipping_method" ("method", "description", "fee") VALUES ('Posten', 'Posten Norge AS is the Norwegian postal service', 50.00);
INSERT INTO "shipping_method" ("method", "description", "fee") VALUES ('DHL', 'DHL Express is a division of the German logistics company Deutsche Post DHL', 100.00);
INSERT INTO "shipping_method" ("method", "description", "fee") VALUES ('UPS', 'United Parcel Service, Inc. is an American multinational package delivery and supply chain management company', 75.00);
INSERT INTO "shipping_method" ("method", "description", "fee") VALUES ('FedEx', 'FedEx Corporation is an American multinational delivery services company', 90.00);
INSERT INTO "shipping_method" ("method", "description", "fee") VALUES ('Bring', 'Bring is a Norwegian postal and logistics company', 60.00);
