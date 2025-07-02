package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

// GetPaymentHistory returns the payment history for a specific contract
func (k Keeper) GetPaymentHistory(ctx context.Context, contractId uint64) (types.PaymentHistory, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PaymentHistoryKey))
	
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, contractId)
	
	val := store.Get(bz)
	if val == nil {
		return types.PaymentHistory{}, false
	}
	
	var paymentHistory types.PaymentHistory
	k.cdc.MustUnmarshal(val, &paymentHistory)
	return paymentHistory, true
}

// SetPaymentHistory sets the payment history for a specific contract
func (k Keeper) SetPaymentHistory(ctx context.Context, paymentHistory types.PaymentHistory) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PaymentHistoryKey))
	
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, paymentHistory.ContractId)
	
	val := k.cdc.MustMarshal(&paymentHistory)
	store.Set(bz, val)
}

// RemovePaymentHistory removes the payment history for a specific contract
func (k Keeper) RemovePaymentHistory(ctx context.Context, contractId uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PaymentHistoryKey))
	
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, contractId)
	
	store.Delete(bz)
}

// GetAllPaymentHistory returns all payment history records
func (k Keeper) GetAllPaymentHistory(ctx context.Context) (list []types.PaymentHistory) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PaymentHistoryKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHostingOffersByInquiry returns all hosting offers for a specific inquiry
func (k Keeper) GetHostingOffersByInquiry(ctx context.Context, inquiryId uint64) (list []types.HostingOffer) {
	allOffers := k.GetAllHostingOffer(ctx)
	
	for _, offer := range allOffers {
		if offer.InquiryId == inquiryId {
			list = append(list, offer)
		}
	}
	
	return
}