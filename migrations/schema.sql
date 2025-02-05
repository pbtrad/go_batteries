CREATE TABLE fast_power_data (
    requested_time TIMESTAMP NOT NULL,
    received_time TIMESTAMP NOT NULL,
    write_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timestamp TIMESTAMP NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    manufacturer_serial VARCHAR(255) NOT NULL,
    generation_active_power DECIMAL(10,2),
    grid_active_power DECIMAL(10,2),
    grid_reactive_power DECIMAL(10,2),
    battery_active_power DECIMAL(10,2),
    battery_apparent_power DECIMAL(10,2),
    soc DECIMAL(5,2),
    PRIMARY KEY (manufacturer, manufacturer_serial, timestamp)
);

CREATE TABLE half_hourly_energy_data (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    requested_time TIMESTAMP NOT NULL,
    received_time TIMESTAMP NOT NULL,
    write_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timestamp TIMESTAMP NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    manufacturer_serial VARCHAR(255) NOT NULL,
    pv_to_home DECIMAL(10,2),
    pv_to_battery DECIMAL(10,2),
    pv_to_grid DECIMAL(10,2),
    grid_to_home DECIMAL(10,2),
    grid_to_battery DECIMAL(10,2),
    battery_to_home DECIMAL(10,2),
    battery_to_grid DECIMAL(10,2),
    UNIQUE INDEX (manufacturer, manufacturer_serial, timestamp)
);
