CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(50),
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    title VARCHAR(255),
    version VARCHAR(50),
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS customer_products (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customers(id),
    product_id UUID REFERENCES products(id),
    hardware_hash VARCHAR(255),
    license_type VARCHAR(50),
    is_active BOOLEAN,
    expire_at BIGINT,
    first_confirmed_at BIGINT,
    last_confirmed_at BIGINT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS restrictions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key VARCHAR(255),
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS customer_products_restrictions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    restriction_id INTEGER REFERENCES restrictions(id),
    customer_product_id UUID REFERENCES customer_products(id),
    value VARCHAR(255),
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255),
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(50) CHECK (role IN ('superAdmin', 'admin')),
    created_at BIGINT NOT NULL,
    updated_at BIGINT
    );

CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at BIGINT NOT NULL
    );

CREATE TABLE IF NOT EXISTS logs_template (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key VARCHAR(255) NOT NULL,
    value TEXT NOT NULL,
    language VARCHAR(10) NOT NULL
    );
