#!/bin/bash

# Exit on error
set -e

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <node_number>"
    echo "Example: $0 0"
    exit 1
fi

NODE_NUMBER=$1
HOME_PREFIX="/data/lokad"
NODE_HOME="${HOME_PREFIX}${NODE_NUMBER}"

# Start the node
lokad start \
    --home "${NODE_HOME}" \
    --chain-id loka_567000-1 \
    --log_level info
