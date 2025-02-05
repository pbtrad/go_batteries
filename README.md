# Go Batteries

**Go Batteries** is a Go(Golang) application designed to communicate with household battery storage systems, enabling efficient energy monitoring, data collection, and optimization.

## Overview

Modern homes equipped with battery storage systems (such as Tesla Powerwall, Sonnen, or Enphase batteries) store energy from renewable sources like solar panels and discharge it when needed. This project:
- **Collects real-time power data** from household batteries.
- **Analyzes energy flow** between the grid, home, and battery.
- **Stores and processes energy data** to optimize energy usage.

## Data Management

### **Fast Power Data**
The system captures real-time power metrics, logging the following:
- **Battery Charge State (SOC%)** – Monitors battery capacity.
- **Grid Active Power** – Measures power drawn from the grid.
- **Battery Active & Apparent Power** – Tracks energy movement in/out of the battery.
- **Solar Generation Power** – Analyzes self-consumption vs. grid export.

### **Half-Hourly Energy Data**
This captures aggregated energy data in 30-minute intervals, including:
- **PV (Solar) to Home/Battery/Grid** – Where solar energy is being used.
- **Grid to Home/Battery** – How much grid power is consumed.
- **Battery to Home/Grid** – Battery discharge behavior.

## Energy Optimization

The collected data is useful for:
- **Grid Flexibility** – Understanding when to charge/discharge batteries based on grid demand.
- **Solar Self-Consumption** – Maximizing solar energy usage at home instead of exporting to the grid.
- **Cost Savings** – Reducing reliance on expensive peak-hour electricity.
- **Demand Response Integration** – Enabling automated battery adjustments based on energy tariffs.
