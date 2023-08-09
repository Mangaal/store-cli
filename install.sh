#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <platform>"
    exit 1
fi


# Set the necessary environment variables
CLI_NAME="store"
INSTALL_PATH="/usr/local/bin"

case "$1" in
    windows)
        CLI_URL="https://github.com/Mangaal/store-cli/raw/main/downloads/store-windows-amd64.exe"
        ;;
    linux)
        CLI_URL="https://github.com/Mangaal/store-cli/raw/main/downloads/store-linux-amd64"
        ;;
    darwin)
        CLI_URL="https://github.com/Mangaal/store-cli/raw/main/downloads/store-darwin-amd64"
        ;;
    *)
        echo "Invalid argument. Supported options: latest, windows, linux, darwin."
        exit 1
        ;;
esac


# Download the latest release of your CLI
curl -Lo "$CLI_NAME" "$CLI_URL"

# Make the CLI executable
chmod +x "$CLI_NAME"

# Move the CLI to the installation path
sudo mv "$CLI_NAME" "$INSTALL_PATH/"


echo "Installation complete. You can now use '$CLI_NAME' to run the CLI."
