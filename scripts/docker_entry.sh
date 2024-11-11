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

# Validate binary exists

if ! command -v filespace-chaind > /dev/null; then
  make install

  if ! command -v filespace-chaind; then
    echo "filespace-chaind binary not found in $PATH"
    exit 1
  fi
fi

# Set parameters
# Directory where `filespace-chaind` expects its configuration
HOME_DIRECTORY="$HOME/.filespace-chain"

if [[ -n "$PERSISTENT_DATA_DIR" ]]; then
  # Persistent directory
  PERSISTENT_HOME="$PERSISTENT_DATA_DIR/.filespace-chain"

    if [[ "$WIPE_DATA" == "true" ]]; then
      rm -r "$PERSISTENT_HOME"
      echo "Wiping old data"
    fi

  # Create the persistent directory if it doesn't exist
  mkdir -p "$PERSISTENT_HOME"

  # Create a symbolic link
  ln -sfn "$PERSISTENT_HOME" "$HOME_DIRECTORY"

  echo "created a symlink from $PERSISTENT_HOME to $HOME_DIRECTORY"
fi



CONFIG_DIRECTORY="$HOME_DIRECTORY/config"
TENDERMINT_CONFIG_FILE="$CONFIG_DIRECTORY/config.toml"
CLIENT_CONFIG_FILE="$CONFIG_DIRECTORY/client.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
CHAIN_ID=${CHAIN_ID:-"filespace-chain"}
MONIKER_NAME=${MONIKER_NAME:-"local"}
WRITE_NEW_GENESIS=${WRITE_NEW_GENESIS:-"true"}
ENABLE_API=${ENABLE_API:-"true"}
OWNER_NECCESSARY=${OWNER_NECCESSARY:-"true"}

# Setting non-default ports to avoid port conflicts when running local rollapp
SETTLEMENT_ADDR=${SETTLEMENT_ADDR:-"0.0.0.0:26657"}
P2P_ADDRESS=${P2P_ADDRESS:-"0.0.0.0:26656"}
GRPC_ADDRESS=${GRPC_ADDRESS:-"0.0.0.0:9090"}
GRPC_WEB_ADDRESS=${GRPC_WEB_ADDRESS:-"0.0.0.0:9091"}
API_ADDRESS=${API_ADDRESS:-"0.0.0.0:1317"}
JSONRPC_ADDRESS=${JSONRPC_ADDRESS:-"0.0.0.0:9545"}
JSONRPC_WS_ADDRESS=${JSONRPC_WS_ADDRESS:-"0.0.0.0:9546"}

TOKEN_AMOUNT=${TOKEN_AMOUNT:-"1000000000000000000000uspace"} #1M SPACE (1e6 SPACE = 1e6 * 1e18 = 1e24 uspace)
STAKING_AMOUNT=${STAKING_AMOUNT:-"670000000000000000000uspace"} #67% is staked (inflation goal)


echo "PERSISTENT_DATA_DIR: $PERSISTENT_DATA_DIR"
echo "TENDERMINT_CONFIG_FILE: $TENDERMINT_CONFIG_FILE"
echo "CLIENT_CONFIG_FILE: $CLIENT_CONFIG_FILE"
echo "APP_CONFIG_FILE: $APP_CONFIG_FILE"

# Check if necessary files exist in DATA_DIRECTORY, init chain if not
if [ "$WIPE_DATA" == "true" ] || [ "$WRITE_NEW_GENESIS" == "true" ] || [ ! -f "$TENDERMINT_CONFIG_FILE" ] || [ ! -f "$CLIENT_CONFIG_FILE" ] || [ ! -f "$APP_CONFIG_FILE" ]; then

    if [[ "$WIPE_DATA" == "true" ]]; then
      echo "Re-Initializing chain"
    elif [[ "$WRITE_NEW_GENESIS" == "true" ]]; then
      echo "Writing new genesis, Re-Initializing chain"
    else
        echo "Necessary files not found in $HOME_DIRECTORY. Initializing chain..."
    fi

    filespace-chaind init "$MONIKER_NAME" --chain-id="$CHAIN_ID" -o

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
    set_bank_denom_metadata
    set_epochs_params
    set_incentives_params
    add_peers

    if [[ "$ENABLE_API" == "true" ]]; then
      enable_api
    fi
else
    echo "Necessary files found. Continuing with existing configuration..."
fi



if [[ "$OWNER_NECCESSARY" == "true" ]]; then

  CONTINUE_LOOP="true"

  check_key() {
      # Run the command and capture the output
      KEY_INFO=$(filespace-chaind keys show owner --keyring-backend test 2>&1)

      # Check if the output contains expected information
      if [[ $KEY_INFO == *"address: space"* ]]; then
          CONTINUE_LOOP="false"  # Exit the loop
      else
          echo "Key 'owner' not found."
          echo "filespace-chaind keys add owner --recover --keyring-backend test"
      fi
  }

  check_key
  # Infinite loop
  while [ "$CONTINUE_LOOP" == "true" ]; do
      echo "Waiting 30s for key 'owner' to be set..."
      sleep 30
      check_key
  done

  echo "Key 'owner' found."

else
  filespace-chaind keys add owner --keyring-backend test
fi


if [[ "$WRITE_NEW_GENESIS" == "true" ]]; then

  echo "Writing new genesis"

  filespace-chaind genesis add-genesis-account "$(filespace-chaind keys show owner -a --keyring-backend test)" "$TOKEN_AMOUNT"
  filespace-chaind genesis add-genesis-account space1uqyssd3xeuaadlyqh6s3z2pszw36e6hxw092d3 "$TOKEN_AMOUNT"
  filespace-chaind genesis add-genesis-account space16ry9r6dqlgv9gfjhzlfflx40345rdzlad2euqu "$TOKEN_AMOUNT"
  filespace-chaind genesis add-genesis-account space1gwqqxk80qdqdejyxzvrjl8yflw22zflrz996nj "$TOKEN_AMOUNT"
  filespace-chaind genesis gentx owner "$TOKEN_AMOUNT" --chain-id "$CHAIN_ID" --keyring-backend test
  filespace-chaind genesis collect-gentxs
else
  echo "Copy existing genesis"
  if [ -f "/tmp/data/genesis/genesis.json" ]; then
       cp /tmp/data/genesis/genesis.json $GENESIS_FILE
   else
       echo "Genesis file not found"
       exit 1
   fi

fi

echo "Starting filespace"

filespace-chaind start