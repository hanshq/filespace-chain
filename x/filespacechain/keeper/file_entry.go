package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetFileEntryCount get the total number of fileEntry
func (k Keeper) GetFileEntryCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FileEntryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFileEntryCount set the total number of fileEntry
func (k Keeper) SetFileEntryCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FileEntryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFileEntry appends a fileEntry in the store with a new id and update the count
func (k Keeper) AppendFileEntry(
	ctx context.Context,
	fileEntry types.FileEntry,
) uint64 {
	// Create the fileEntry
	count := k.GetFileEntryCount(ctx)

	// Set the ID of the appended value
	fileEntry.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	appendedValue := k.cdc.MustMarshal(&fileEntry)
	store.Set(GetFileEntryIDBytes(fileEntry.Id), appendedValue)

	// Update fileEntry count
	k.SetFileEntryCount(ctx, count+1)

	return count
}

// SetFileEntry set a specific fileEntry in the store
func (k Keeper) SetFileEntry(ctx context.Context, fileEntry types.FileEntry) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	b := k.cdc.MustMarshal(&fileEntry)
	store.Set(GetFileEntryIDBytes(fileEntry.Id), b)
}

// GetFileEntry returns a fileEntry from its id
func (k Keeper) GetFileEntry(ctx context.Context, id uint64) (val types.FileEntry, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	b := store.Get(GetFileEntryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFileEntry removes a fileEntry from the store
func (k Keeper) RemoveFileEntry(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	store.Delete(GetFileEntryIDBytes(id))
}

// GetAllFileEntry returns all fileEntry
func (k Keeper) GetAllFileEntry(ctx context.Context) (list []types.FileEntry) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FileEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFileEntryIDBytes returns the byte representation of the ID
func GetFileEntryIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.FileEntryKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

// GetFileEntryByCid returns a fileEntry from its CID
func (k Keeper) GetFileEntryByCid(ctx context.Context, cid string) (val types.FileEntry, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileEntryKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var fileEntry types.FileEntry
		k.cdc.MustUnmarshal(iterator.Value(), &fileEntry)
		if fileEntry.Cid == cid {
			return fileEntry, true
		}
	}

	return val, false
}
