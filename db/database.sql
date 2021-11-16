DROP DATABASE IF EXISTS gopher_rentals;
CREATE DATABASE IF NOT EXISTS gopher_rentals;
USE gopher_rentals;

CREATE TABLE IF NOT EXISTS users
(
    id       VARCHAR(128)   NOT NULL,
    email    VARCHAR(128) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS customers
(
    id                    VARCHAR(128) NOT NULL,
    first_name            VARCHAR(128),
    last_name             VARCHAR(128),
    nationality           VARCHAR(128),
    identification_number VARCHAR(128),
    identification_type   VARCHAR(128),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cars
(
    id                VARCHAR(128) NOT NULL,
    model             VARCHAR(45),
    year              INTEGER,
    license_plate     VARCHAR(45),
    current_km        DECIMAL(10, 2),
    max_kg            DECIMAL(10, 2),
    fuel_type         VARCHAR(45),
    hire_price        DECIMAL(10, 2),
    hire_availability TINYINT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS customer_hire_car
(
    id          VARCHAR(128) NOT NULL,
    customer_id VARCHAR(128),
    car_id      VARCHAR(128),
    hire_date   DATETIME,
    return_date DATETIME,
    PRIMARY KEY (id),
    CONSTRAINT fk_customer_id_customer_hire_car FOREIGN KEY (customer_id) REFERENCES customers (id)
        ON DELETE CASCADE ,
    CONSTRAINT fk_car_id_customer_hire_car FOREIGN KEY (car_id) REFERENCES cars (id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS locations
(
    id                        VARCHAR(128) NOT NULL,
    car_id                    VARCHAR(128),
    latitude                  DECIMAL(10, 4),
    longitude                 DECIMAL(10, 4),
    current_location_datetime DATETIME,
    PRIMARY KEY (id),
    CONSTRAINT fk_car_id_locations FOREIGN KEY (car_id) REFERENCES cars (id) ON DELETE CASCADE 
);
