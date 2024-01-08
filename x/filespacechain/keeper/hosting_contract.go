package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetHostingContractCount get the total number of hostingContract
func (k Keeper) GetHostingContractCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingContractCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHostingContractCount set the total number of hostingContract
func (k Keeper) SetHostingContractCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingContractCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHostingContract appends a hostingContract in the store with a new id and update the count
func (k Keeper) AppendHostingContract(
	ctx context.Context,
	hostingContract types.HostingContract,
) uint64 {
	// Create the hostingContract
	count := k.GetHostingContractCount(ctx)

	// Set the ID of the appended value
	hostingContract.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingContractKey))
	appendedValue := k.cdc.MustMarshal(&hostingContract)
	store.Set(GetHostingContractIDBytes(hostingContract.Id), appendedValue)

	// Update hostingContract count
	k.SetHostingContractCount(ctx, count+1)

	return count
}

// SetHostingContract set a specific hostingContract in the store
func (k Keeper) SetHostingContract(ctx context.Context, hostingContract types.HostingContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingContractKey))
	b := k.cdc.MustMarshal(&hostingContract)
	store.Set(GetHostingContractIDBytes(hostingContract.Id), b)
}

// GetHostingContract returns a hostingContract from its id
func (k Keeper) GetHostingContract(ctx context.Context, id uint64) (val types.HostingContract, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingContractKey))
	b := store.Get(GetHostingContractIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHostingContract removes a hostingContract from the store
func (k Keeper) RemoveHostingContract(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingContractKey))
	store.Delete(GetHostingContractIDBytes(id))
}

// GetAllHostingContract returns all hostingContract
func (k Keeper) GetAllHostingContract(ctx context.Context) (list []types.HostingContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingContractKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostingContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHostingContractIDBytes returns the byte representation of the ID
func GetHostingContractIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.HostingContractKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
