#!/bin/bash

tmp=$(mktemp)
# Common commands
genesis_config_cmds="$(dirname "$0")/src/genesis_config_commands.sh"

if [ -f "$genesis_config_cmds" ]; then
  . "$genesis_config_cmds"
else
  echo "Error: header file not found" >&2
  exit 1
fi

# Set parameters
DATA_DIRECTORY="$HOME/.filespace-chain"
CONFIG_DIRECTORY="$DATA_DIRECTORY/config"
TENDERMINT_CONFIG_FILE="$CONFIG_DIRECTORY/config.toml"
CLIENT_CONFIG_FILE="$CONFIG_DIRECTORY/client.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
CHAIN_ID=${CHAIN_ID:-"filespace-chain"}
MONIKER_NAME=${MONIKER_NAME:-"local"}
KEY_NAME=${KEY_NAME:-"local-user"}

# Setting non-default ports to avoid port conflicts when running local rollapp
SETTLEMENT_ADDR=${SETTLEMENT_ADDR:-"0.0.0.0:36657"}
P2P_ADDRESS=${P2P_ADDRESS:-"0.0.0.0:36656"}
GRPC_ADDRESS=${GRPC_ADDRESS:-"0.0.0.0:8090"}
GRPC_WEB_ADDRESS=${GRPC_WEB_ADDRESS:-"0.0.0.0:8091"}
API_ADDRESS=${API_ADDRESS:-"0.0.0.0:1318"}
JSONRPC_ADDRESS=${JSONRPC_ADDRESS:-"0.0.0.0:9545"}
JSONRPC_WS_ADDRESS=${JSONRPC_WS_ADDRESS:-"0.0.0.0:9546"}

TOKEN_AMOUNT=${TOKEN_AMOUNT:-"1000000000000000000000000uspace"} #1M SPACE (1e6 SPACE = 1e6 * 1e18 = 1e24 uspace)
STAKING_AMOUNT=${STAKING_AMOUNT:-"670000000000000000000000uspace"} #67% is staked (inflation goal)

# Validate binary exists

if ! command -v filespace-chaind > /dev/null; then
  make install

  if ! command -v filespace-chaind; then
    echo "filespace-chaind binary not found in $PATH"
    exit 1
  fi
fi



# Create and init dymension chain
filespace-chaind init "$MONIKER_NAME" --chain-id="$CHAIN_ID"

# ---------------------------------------------------------------------------- #
#                              Set configurations                              #
# ---------------------------------------------------------------------------- #
# [Remaining configuration commands remain unchanged]

# Execute configuration scripts
set_consenus_params
set_gov_params
set_hub_params
set_misc_params
set_EVM_params
set_bank_denom_metadata
set_epochs_params
set_incentives_params


filespace-chaind keys add "$KEY_NAME" --keyring-backend test
filespace-chaind genesis add-genesis-account "$(filespace-chaind keys show "$KEY_NAME" -a --keyring-backend test)" "$TOKEN_AMOUNT"

filespace-chaind genesis gentx "$KEY_NAME" "$STAKING_AMOUNT" --chain-id "$CHAIN_ID" --keyring-backend test
filespace-chaind genesis collect-gentxs

filespace-chaind start --minimum-gas-prices 0space