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
SETTLEMENT_ADDR=${SETTLEMENT_ADDR:-"0.0.0.0:26657"}
P2P_ADDRESS=${P2P_ADDRESS:-"0.0.0.0:26656"}
GRPC_ADDRESS=${GRPC_ADDRESS:-"0.0.0.0:9090"}
GRPC_WEB_ADDRESS=${GRPC_WEB_ADDRESS:-"0.0.0.0:9091"}
API_ADDRESS=${API_ADDRESS:-"0.0.0.0:1317"}
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

rm -rf "$DATA_DIRECTORY"

# Create and init dymension chain
filespace-chaind init "$MONIKER_NAME" --chain-id="$CHAIN_ID"

# ---------------------------------------------------------------------------- #
#                              Set configurations                              #
# ---------------------------------------------------------------------------- #
sed -i'' -e "/\[rpc\]/,+3 s/laddr *= .*/laddr = \"tcp:\/\/$SETTLEMENT_ADDR\"/" "$TENDERMINT_CONFIG_FILE"
sed -i'' -e "/\[p2p\]/,+3 s/laddr *= .*/laddr = \"tcp:\/\/$P2P_ADDRESS\"/" "$TENDERMINT_CONFIG_FILE"

sed -i'' -e "/\[grpc\]/,+6 s/address *= .*/address = \"$GRPC_ADDRESS\"/" "$APP_CONFIG_FILE"
sed -i'' -e "/\[grpc-web\]/,+7 s/address *= .*/address = \"$GRPC_WEB_ADDRESS\"/" "$APP_CONFIG_FILE"
sed -i'' -e "/\[json-rpc\]/,+6 s/address *= .*/address = \"$JSONRPC_ADDRESS\"/" "$APP_CONFIG_FILE"
sed -i'' -e "/\[json-rpc\]/,+9 s/address *= .*/address = \"$JSONRPC_WS_ADDRESS\"/" "$APP_CONFIG_FILE"
sed -i'' -e '/\[api\]/,+3 s/enable *= .*/enable = true/' "$APP_CONFIG_FILE"
sed -i'' -e "/\[api\]/,+9 s/address *= .*/address = \"tcp:\/\/$API_ADDRESS\"/" "$APP_CONFIG_FILE"

sed -i'' -e 's/^minimum-gas-prices *= .*/minimum-gas-prices = "0uspace"/' "$APP_CONFIG_FILE"

sed -i'' -e "s/^chain-id *= .*/chain-id = \"$CHAIN_ID\"/" "$CLIENT_CONFIG_FILE"
sed -i'' -e "s/^keyring-backend *= .*/keyring-backend = \"test\"/" "$CLIENT_CONFIG_FILE"
sed -i'' -e "s/^node *= .*/node = \"tcp:\/\/$SETTLEMENT_ADDR\"/" "$CLIENT_CONFIG_FILE"

# Execute configuration scripts
set_consenus_params
set_gov_params
set_hub_params
set_misc_params
set_EVM_params
set_bank_denom_metadata
set_epochs_params
set_incentives_params
enable_api

cat $APP_CONFIG_FILE

filespace-chaind keys add "$KEY_NAME" --keyring-backend test
filespace-chaind genesis add-genesis-account "$(filespace-chaind keys show "$KEY_NAME" -a --keyring-backend test)" "$TOKEN_AMOUNT"

filespace-chaind genesis gentx "$KEY_NAME" "$STAKING_AMOUNT" --chain-id "$CHAIN_ID" --keyring-backend test
filespace-chaind genesis collect-gentxs

filespace-chaind start --minimum-gas-prices 0space