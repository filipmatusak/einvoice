#!/bin/bash

set -ex

# Default values of arguments
DROP_DB=false
CREATE_DB=false
SETUP_DB=false
CONNECT_DB=false
OTHER_ARGUMENTS=()

# Loop through arguments and process them
for arg in "$@"
do
    case $arg in
        --drop)
        DROP_DB=true
        shift # Remove --drop from processing
        ;;
        --create)
        CREATE_DB=true
        shift # Remove --create from processing
        ;;
        --setup)
        SETUP_DB=true
        shift # Remove --setup from processing
        ;;
        --connect)
        CONNECT_DB=true
        shift # Remove --connect from processing
        ;;
        *)
        OTHER_ARGUMENTS+=("$1")
        shift # Remove generic argument from processing
        ;;
    esac
done

echo "# Drop db: $DROP_DB"
echo "# Create db: $CREATE_DB"
echo "# Setup db: $SETUP_DB"
echo "# Connect db: $CONNECT_DB"
echo "# Other arguments: ${OTHER_ARGUMENTS[*]}"

if [ $DROP_DB == "true" ]; then
  psql postgres -h 127.0.0.1 -d postgres -f ./sql/drop.sql
  echo "db dropped"
fi

if [ $CREATE_DB == "true" ]; then
  psql postgres -h 127.0.0.1 -d postgres -f ./sql/create.sql
  echo "db created"
fi

if [ $SETUP_DB == "true" ]; then
  psql postgres -h 127.0.0.1 -d einvoice -f ./sql/setup.sql
  echo "db setup"
fi

if [ $CONNECT_DB == "true" ]; then
  echo "connecting to db ..."
  psql postgres -h 127.0.0.1 -d einvoice
fi
