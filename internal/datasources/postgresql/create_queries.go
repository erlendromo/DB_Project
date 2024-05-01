package postgresql

var CREATE_DB string = `CREATE DATABASE electromart_db;`

var CreateTablesList []string = []string{
	createUsersTable,
	createAddressesTable,
	createUserAddressesTable,
	createCategoriesTable,
	createBrandsTable,
	createProductsTable,
	createDiscountsTable,
	createProductsDiscountsTable,
	createOrdersTable,
	createOrderItemsTable,
	createPaymentsTable}

var createUsersTable string = `
  CREATE TABLE IF NOT EXISTS "users" (
	"id" SERIAL PRIMARY KEY,
	"username" VARCHAR(50) NOT NULL,
	"password" VARCHAR(50) NOT NULL,
	"first_name" VARCHAR(50) NOT NULL,
	"last_name" VARCHAR(50) NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"phone_number" VARCHAR(20) NOT NULL,
	"role" SMALLINT NOT NULL DEFAULT 1
  );
`

var createAddressesTable string = `
  CREATE TABLE IF NOT EXISTS "addresses" (
	"id" SERIAL PRIMARY KEY,
	"postal_code" SMALLINT NOT NULL,
	"city" VARCHAR(50) NOT NULL,
	"street_name" VARCHAR(50) NOT NULL
  );
`

var createUserAddressesTable string = `
  CREATE TABLE IF NOT EXISTS "user_addresses" (
	"user_id" INTEGER NOT NULL REFERENCES "users" ("id"),
	"address_id" INTEGER NOT NULL REFERENCES "addresses" ("id"),
	PRIMARY KEY ("user_id", "address_id")
  );
`

var createCategoriesTable string = `
  CREATE TABLE IF NOT EXISTS "categories" (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(50) NOT NULL,
	"description" VARCHAR(255)
  );
`

var createBrandsTable string = `
  CREATE TABLE IF NOT EXISTS "brands" (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(50) NOT NULL,
	"description" VARCHAR(255)
  );
`

var createProductsTable string = `
  CREATE TABLE IF NOT EXISTS "products" (
	"id" SERIAL PRIMARY KEY,
	"category_id" INTEGER NOT NULL REFERENCES "categories" ("id"),
	"brand_id" INTEGER NOT NULL REFERENCES "brands" ("id"),
	"name" VARCHAR(50) NOT NULL,
	"price" NUMERIC(9,2) NOT NULL,
	"description" VARCHAR(255),
	"stock" INTEGER
  );
`

var createDiscountsTable string = `
  CREATE TABLE IF NOT EXISTS "discounts" (
	"id" SERIAL PRIMARY KEY,
	"percentage" NUMERIC(4,2) NOT NULL,
	"start_at" TIMESTAMPTZ NOT NULL,
	"end_at" TIMESTAMPTZ NOT NULL
  );
`

var createProductsDiscountsTable string = `
  CREATE TABLE IF NOT EXISTS "products_discounts" (
	"product_id" INTEGER NOT NULL REFERENCES "products" ("id"),
	"discounts_id" INTEGER NOT NULL REFERENCES "discounts" ("id"),
	PRIMARY KEY ("product_id", "discounts_id")
  );
`

var createOrdersTable string = `
  CREATE TABLE IF NOT EXISTS "orders" (
	"id" SERIAL PRIMARY KEY,
	"user_id" INTEGER NOT NULL REFERENCES "users" ("id"),
	"created_at" TIMESTAMPTZ NOT NULL,
	"total_amount" NUMERIC(9,2) NOT NULL,
	"status" VARCHAR(50) NOT NULL
  );
`

var createOrderItemsTable string = `
  CREATE TABLE IF NOT EXISTS "order_items" (
	"id" SERIAL PRIMARY KEY,
	"product_id" INTEGER NOT NULL REFERENCES "products" ("id"),
	"order_id" INTEGER NOT NULL REFERENCES "orders" ("id"),
	"quantity" SMALLINT NOT NULL,
	"subtotal" NUMERIC(9,2) NOT NULL
  );
`

var createPaymentsTable string = `
  CREATE TABLE IF NOT EXISTS "payments" (
	"id" SERIAL PRIMARY KEY,
	"order_id" INTEGER NOT NULL UNIQUE REFERENCES "orders" ("id"),
	"method" VARCHAR(50) NOT NULL,
	"amount" NUMERIC(9,2) NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL,
	"status" VARCHAR(50) NOT NULL
  );
`
