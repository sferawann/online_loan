CREATE TABLE "user" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "nik" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(20) NOT NULL,
  "limit" decimal(10,2) NOT NULL,
  "id_role" INT,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
  FOREIGN KEY (id_role) REFERENCES "role"(id)
);

select * from products;

CREATE TABLE "product" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "installment" INT NOT NULL,
  "bunga" decimal(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "role" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "transaction" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_user" INT,
  "id_product" INT,
  "status" bool,
  "amount" decimal(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
  "due_date" timestamptz NOT NULL DEFAULT 'NOW()',
  FOREIGN KEY (id_user) REFERENCES user(id),
  FOREIGN KEY (id_product) REFERENCES product(id)
);

CREATE TABLE "accept_status" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_transaction" INT,
  "status" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
  FOREIGN KEY (id_transaction) REFERENCES transaction(id)
);

CREATE TABLE "payment" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_transaction" INT,
  "payment_amount" decimal(10,2) NOT NULL,
  "payment_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "id_payment_method" INT,
  FOREIGN KEY (id_transaction) REFERENCES transaction(id),
  FOREIGN KEY (id_payment_method) REFERENCES payment_method(id)
);

CREATE TABLE "payment_method" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);