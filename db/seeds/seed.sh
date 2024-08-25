#!/bin/bash

HOST="pg-db"
USER="postgres"
PASSWORD="postgres"
DB="seryu_cargo_db"

run_copy_command() {
    local table=$1
    local file=$2

    psql postgresql://$USER:$PASSWORD@localhost:5432/$DB "\copy $table from '/docker-entrypoint-initdb.d/$file' delimiter ',' CSV HEADER;"
}

run_copy_command "drivers" "drivers.csv"
run_copy_command "driver_attendances" "driver_attendances.csv"
run_copy_command "shipments" "shipments.csv"
run_copy_command "shipment_costs" "shipment_costs.csv"
run_copy_command "variable_configs" "variable_configs.csv"
