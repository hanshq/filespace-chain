package keeper

import (
	"context"
	"encoding/binary"
	"fmt"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// CompleteHostingContract marks a contract as completed and processes completion bonus
func (k Keeper) CompleteHostingContract(ctx context.Context, contractId uint64) error {
	contract, found := k.GetHostingContract(ctx, contractId)
	if !found {
		return fmt.Errorf("contract %d not found", contractId)
	}
	
	// Check if contract has already reached its end block
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	if currentHeight < contract.EndBlock {
		return fmt.Errorf("contract %d has not reached end block yet (current: %d, end: %d)", 
			contractId, currentHeight, contract.EndBlock)
	}
	
	// Check if completion bonus has already been paid
	paymentHistory, found := k.GetPaymentHistory(ctx, contractId)
	if found && paymentHistory.CompletionBonusPaid {
		k.Logger().Info("completion bonus already paid for contract", "contract_id", contractId)
		return nil
	}
	
	// Process completion bonus
	err := k.ProcessCompletionBonus(ctx, contractId)
	if err != nil {
		return fmt.Errorf("failed to process completion bonus for contract %d: %w", contractId, err)
	}
	
	k.Logger().Info("hosting contract completed successfully", 
		"contract_id", contractId,
		"end_block", contract.EndBlock,
		"current_block", currentHeight,
	)
	
	return nil
}

// ProcessExpiredContracts processes all contracts that have passed their end block
// This should be called periodically to clean up expired contracts and pay completion bonuses
func (k Keeper) ProcessExpiredContracts(ctx context.Context) error {
	contracts := k.GetAllHostingContract(ctx)
	
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	for _, contract := range contracts {
		// Skip contracts that haven't ended yet
		if currentHeight < contract.EndBlock {
			continue
		}
		
		// Check if completion bonus has already been paid
		paymentHistory, found := k.GetPaymentHistory(ctx, contract.Id)
		if found && paymentHistory.CompletionBonusPaid {
			continue
		}
		
		// Process completion bonus for expired contracts
		err := k.ProcessCompletionBonus(ctx, contract.Id)
		if err != nil {
			k.Logger().Error("failed to process completion bonus for expired contract",
				"contract_id", contract.Id,
				"error", err,
			)
			continue
		}
		
		k.Logger().Info("processed completion bonus for expired contract",
			"contract_id", contract.Id,
			"end_block", contract.EndBlock,
			"current_block", currentHeight,
		)
	}
	
	return nil
}

// GetActiveContracts returns all contracts that are currently active (within their duration)
func (k Keeper) GetActiveContracts(ctx context.Context) (list []types.HostingContract) {
	contracts := k.GetAllHostingContract(ctx)
	
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	for _, contract := range contracts {
		if currentHeight >= contract.StartBlock && currentHeight <= contract.EndBlock {
			list = append(list, contract)
		}
	}
	
	return
}

// GetExpiredContracts returns all contracts that have passed their end block
func (k Keeper) GetExpiredContracts(ctx context.Context) (list []types.HostingContract) {
	contracts := k.GetAllHostingContract(ctx)
	
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	for _, contract := range contracts {
		if currentHeight > contract.EndBlock {
			list = append(list, contract)
		}
	}
	
	return
}
