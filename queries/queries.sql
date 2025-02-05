-- name: InsertFastPowerData :exec
INSERT INTO fast_power_data (
    requested_time, received_time, write_time, timestamp, manufacturer, manufacturer_serial,
    generation_active_power, grid_active_power, grid_reactive_power, battery_active_power,
    battery_apparent_power, soc
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetFastPowerData :one
SELECT * FROM fast_power_data WHERE manufacturer = ? AND manufacturer_serial = ? AND timestamp = ?;
