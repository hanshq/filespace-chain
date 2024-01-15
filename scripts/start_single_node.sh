#!/bin/sh

set -o errexit -o nounset

CHAINID=$1
GENACCT=$2


if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account address..."
  exit 1
fi

FILE=~/.filespace-chain/config/genesis.json
if test -f "$FILE"; then
    echo "$FILE exists. Starting node"
else

  echo "Init chain"

  # Build genesis file incl account for passed address
  coins="1000space"
  coinsfaucet="1000space"
  coinsuser="100space"
  filespace-chaind init --chain-id $CHAINID $CHAINID
  filespace-chaind keys add validator --keyring-backend="test"
  filespace-chaind keys add faucet --keyring-backend="test"
  filespace-chaind genesis add-genesis-account $(filespace-chaind keys show validator -a --keyring-backend="test") $coins
  filespace-chaind genesis add-genesis-account $(filespace-chaind keys show faucet -a --keyring-backend="test") $coinsfaucet
  filespace-chaind genesis add-genesis-account $GENACCT $coinsuser
  filespace-chaind genesis gentx validator 1000space --keyring-backend="test" --chain-id $CHAINID
  filespace-chaind genesis collect-gentxs

  # Set proper defaults and change ports
  sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.filespace-chain/config/config.toml
  sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.filespace-chain/config/config.toml
  sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.filespace-chain/config/config.toml
  sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.filespace-chain/config/config.toml
  sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' ~/.filespace-chain/config/config.toml

  cat ~/.filespace-chain/config/config.toml

  echo "App.toml"

  sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0space"/g' ~/.filespace-chain/config/app.toml

  cat ~/.filespace-chain/config/app.toml
fi

# Start the filespace
filespace-chaind start
# --pruning=nothing

#faucet executable needs to be copied to /build
#start faucet in detached mode on same container
#faucet --cli-name filespace-chaind --port 4500 --max-credit 30 --denoms space --keyring-backend test