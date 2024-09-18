-- init-db.sql

CREATE DATABASE users_db;
CREATE DATABASE products_db;
CREATE DATABASE orders_db;

-- Optionally, you can also grant privileges to the user:
GRANT ALL PRIVILEGES ON DATABASE users_db TO ecommerce;
GRANT ALL PRIVILEGES ON DATABASE products_db TO ecommerce;
GRANT ALL PRIVILEGES ON DATABASE orders_db TO ecommerce;
