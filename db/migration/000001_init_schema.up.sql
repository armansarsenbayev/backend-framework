CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL
);

CREATE TABLE "restaurants" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL
);

CREATE TABLE "menu_items" (
    "id" BIGSERIAL PRIMARY KEY,
    "restaurant_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "price" NUMERIC NOT NULL
);

CREATE TABLE "orders" (
    "id" BIGSERIAL PRIMARY KEY,
    "restaurant_id" BIGINT NOT NULL,
    "customer_name" VARCHAR(255) NOT NULL,
    "status" VARCHAR(50),
    "total_price" NUMERIC NOT NULL
);
