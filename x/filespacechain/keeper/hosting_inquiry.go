package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetHostingInquiryCount get the total number of hostingInquiry
func (k Keeper) GetHostingInquiryCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingInquiryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHostingInquiryCount set the total number of hostingInquiry
func (k Keeper) SetHostingInquiryCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.HostingInquiryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHostingInquiry appends a hostingInquiry in the store with a new id and update the count
func (k Keeper) AppendHostingInquiry(
	ctx context.Context,
	hostingInquiry types.HostingInquiry,
) uint64 {
	// Create the hostingInquiry
	count := k.GetHostingInquiryCount(ctx)

	// Set the ID of the appended value
	hostingInquiry.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingInquiryKey))
	appendedValue := k.cdc.MustMarshal(&hostingInquiry)
	store.Set(GetHostingInquiryIDBytes(hostingInquiry.Id), appendedValue)

	// Update hostingInquiry count
	k.SetHostingInquiryCount(ctx, count+1)

	return count
}

// SetHostingInquiry set a specific hostingInquiry in the store
func (k Keeper) SetHostingInquiry(ctx context.Context, hostingInquiry types.HostingInquiry) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingInquiryKey))
	b := k.cdc.MustMarshal(&hostingInquiry)
	store.Set(GetHostingInquiryIDBytes(hostingInquiry.Id), b)
}

// GetHostingInquiry returns a hostingInquiry from its id
func (k Keeper) GetHostingInquiry(ctx context.Context, id uint64) (val types.HostingInquiry, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingInquiryKey))
	b := store.Get(GetHostingInquiryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHostingInquiry removes a hostingInquiry from the store
func (k Keeper) RemoveHostingInquiry(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingInquiryKey))
	store.Delete(GetHostingInquiryIDBytes(id))
}

// GetAllHostingInquiry returns all hostingInquiry
func (k Keeper) GetAllHostingInquiry(ctx context.Context) (list []types.HostingInquiry) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingInquiryKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostingInquiry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHostingInquiryIDBytes returns the byte representation of the ID
func GetHostingInquiryIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.HostingInquiryKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
