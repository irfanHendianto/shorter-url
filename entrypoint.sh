#!/bin/sh

# Define the path to the config file
CONFIG_FILE="/app/config/config.json"

if [ -f "$CONFIG_FILE" ]; then
    echo "Loading configuration from $CONFIG_FILE"

    # Read and export each key-value pair in the JSON file as an environment variable
    export $(cat $CONFIG_FILE | jq -r 'to_entries | .[] | "\(.key)=\(.value)"')

    echo "Environment variables set:"
    env | grep -E 'PORT|SHORT_URL_OPTIONS|SHORT_URL_LENGTH|MONGO_URL|DATABASE_URL'
else
    echo "Config file not found at $CONFIG_FILE"
    exit 1
fi

# Set default MongoDB credentials (hardcoded admin username)
MONGO_INITDB_ROOT_USERNAME="admin"

# Dynamically construct DATABASE_URL
export DATABASE_URL="mongodb://${MONGO_INITDB_ROOT_USERNAME}:${MONGO_INITDB_ROOT_PASSWORD}@${MONGO_URL}"

echo "DATABASE_URL set to: $DATABASE_URL"

# Execute the main application
exec "$@"
