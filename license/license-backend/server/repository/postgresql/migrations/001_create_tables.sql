CREATE TABLE customers (
                           id UUID PRIMARY KEY,
                           name VARCHAR(255) DEFAULT 'cName',
                           email VARCHAR(255) DEFAULT 'cEmail@gmail.com',
                           phone VARCHAR(50) DEFAULT '0',
                           created_at BIGINT NOT NULL,
                           updated_at BIGINT DEFAULT -1
);

CREATE TABLE products (
                          id UUID PRIMARY KEY,
                          name VARCHAR(255) DEFAULT 'pName',
                          title VARCHAR(255) DEFAULT 'pTitle',
                          version VARCHAR(50) DEFAULT 'pVersion',
                          created_at BIGINT NOT NULL,
                          updated_at BIGINT DEFAULT -1
);

CREATE TABLE customer_products (
                                   id UUID PRIMARY KEY,
                                   customer_id UUID REFERENCES customers(id),
                                   product_id UUID REFERENCES products(id),
                                   hardware_hash VARCHAR(255) DEFAULT 'hash',
                                   license_type VARCHAR(50) DEFAULT 'unknown',
                                   is_active BOOLEAN DEFAULT FALSE,
                                   expire_at BIGINT DEFAULT -1,
                                   first_confirmed_at BIGINT DEFAULT -1,
                                   last_confirmed_at BIGINT DEFAULT -1,
                                   created_at BIGINT NOT NULL,
                                   updated_at BIGINT DEFAULT -1
);

CREATE TABLE restrictions (
                             id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                             product_id UUID REFERENCES products(id),
                             key VARCHAR(255) DEFAULT 'rKey',
                             value VARCHAR(255) DEFAULT 'rValue',
                             created_at BIGINT NOT NULL,
                             updated_at BIGINT DEFAULT -1
);

CREATE TABLE users (
                       id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       name VARCHAR(255) DEFAULT 'Uname',
                       username VARCHAR(255) DEFAULT 'uUser',
                       password VARCHAR(255) DEFAULT 'uPass',
                       role VARCHAR(50) DEFAULT 'admin' CHECK (role IN ('superAdmin', 'admin')),
                       created_at BIGINT NOT NULL,
                       updated_at BIGINT DEFAULT -1
);

CREATE TABLE logs (
                     id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                     title VARCHAR(255) NOT NULL,
                     message TEXT NOT NULL,
                     created_at BIGINT NOT NULL
);

CREATE TABLE logs_template (
                              id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                              key VARCHAR(255) NOT NULL,
                              value TEXT NOT NULL,
                              language VARCHAR(10) NOT NULL,
                              created_at BIGINT NOT NULL


);
