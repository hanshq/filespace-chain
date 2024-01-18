
#--------------------validator---- --------------------------#
# Get the validator key
validator_key=$(filespace-chaind tendermint show-validator)

# Check if validator_key is obtained successfully
if [ -z "$validator_key" ]; then
  echo "Failed to obtain validator key"
  exit 1
fi

sed -i "s|\"pubkey\": {.*}|\"pubkey\": ${validator_key}|" /app/data/genesis/validator.json

#-------------------------------------------------------------#

filespace-chaind tx staking create-validator /app/data/genesis/validator.json --from=owner