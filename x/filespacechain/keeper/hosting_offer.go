package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetHostingOfferCount get the total number of hostingOffer
func (k Keeper) GetHostingOfferCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingOfferCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHostingOfferCount set the total number of hostingOffer
func (k Keeper) SetHostingOfferCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingOfferCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHostingOffer appends a hostingOffer in the store with a new id and update the count
func (k Keeper) AppendHostingOffer(
	ctx context.Context,
	hostingOffer types.HostingOffer,
) uint64 {
	// Create the hostingOffer
	count := k.GetHostingOfferCount(ctx)

	// Set the ID of the appended value
	hostingOffer.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	appendedValue := k.cdc.MustMarshal(&hostingOffer)
	store.Set(GetHostingOfferIDBytes(hostingOffer.Id), appendedValue)

	// Update hostingOffer count
	k.SetHostingOfferCount(ctx, count+1)

	return count
}

// SetHostingOffer set a specific hostingOffer in the store
func (k Keeper) SetHostingOffer(ctx context.Context, hostingOffer types.HostingOffer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	b := k.cdc.MustMarshal(&hostingOffer)
	store.Set(GetHostingOfferIDBytes(hostingOffer.Id), b)
}

// GetHostingOffer returns a hostingOffer from its id
func (k Keeper) GetHostingOffer(ctx context.Context, id uint64) (val types.HostingOffer, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	b := store.Get(GetHostingOfferIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHostingOffer removes a hostingOffer from the store
func (k Keeper) RemoveHostingOffer(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	store.Delete(GetHostingOfferIDBytes(id))
}

// GetAllHostingOffer returns all hostingOffer
func (k Keeper) GetAllHostingOffer(ctx context.Context) (list []types.HostingOffer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostingOffer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHostingOfferIDBytes returns the byte representation of the ID
func GetHostingOfferIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.HostingOfferKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
