#!/bin/bash

# export env
export DATA_DIR="/tmp/da"
# Set the necessary environment variables
CLI_NAME="store"
INSTALL_PATH="/usr/local/bin"
CLI_URL="https://github.com/Mangaal/store-cli/raw/main/store"

# Download the latest release of your CLI
curl -Lo "$CLI_NAME" "$CLI_URL"

# Make the CLI executable
chmod +x "$CLI_NAME"

# Move the CLI to the installation path
sudo mv "$CLI_NAME" "$INSTALL_PATH/"

echo "Installation complete. You can now use '$CLI_NAME' to run the CLI."
