package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EscrowRecord tracks escrowed funds for a specific inquiry
type EscrowRecord struct {
	InquiryId uint64   `json:"inquiry_id"`
	Amount    sdk.Coin `json:"amount"`
	Creator   string   `json:"creator"`
}

var (
	EscrowKeyPrefix = []byte{0x01} // Prefix for escrow records
)

// EscrowKey returns the store key for escrow records
func EscrowKey(inquiryId uint64) []byte {
	key := make([]byte, len(EscrowKeyPrefix)+8)
	copy(key, EscrowKeyPrefix)
	binary.BigEndian.PutUint64(key[len(EscrowKeyPrefix):], inquiryId)
	return key
}

// SetEscrowRecord stores an escrow record for an inquiry
func (k Keeper) SetEscrowRecord(ctx context.Context, inquiryId uint64, amount sdk.Coin, creator string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	record := EscrowRecord{
		InquiryId: inquiryId,
		Amount:    amount,
		Creator:   creator,
	}
	
	bz, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}
	storeAdapter.Set(EscrowKey(inquiryId), bz)
}

// GetEscrowRecord retrieves an escrow record for an inquiry
func (k Keeper) GetEscrowRecord(ctx context.Context, inquiryId uint64) (EscrowRecord, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	bz := storeAdapter.Get(EscrowKey(inquiryId))
	if bz == nil {
		return EscrowRecord{}, false
	}
	
	var record EscrowRecord
	err := json.Unmarshal(bz, &record)
	if err != nil {
		return EscrowRecord{}, false
	}
	return record, true
}

// RemoveEscrowRecord removes an escrow record for an inquiry
func (k Keeper) RemoveEscrowRecord(ctx context.Context, inquiryId uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storeAdapter.Delete(EscrowKey(inquiryId))
}

// GetAllEscrowRecords returns all escrow records
func (k Keeper) GetAllEscrowRecords(ctx context.Context) []EscrowRecord {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	var records []EscrowRecord
	iterator := storetypes.KVStorePrefixIterator(storeAdapter, EscrowKeyPrefix)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		var record EscrowRecord
		err := json.Unmarshal(iterator.Value(), &record)
		if err != nil {
			continue // Skip invalid records
		}
		records = append(records, record)
	}
	
	return records
}

// ProcessExpiredInquiries checks for expired inquiries and refunds escrow automatically
func (k Keeper) ProcessExpiredInquiries(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentBlock := uint64(sdkCtx.BlockHeight())
	
	// Get all hosting inquiries
	allInquiries := k.GetAllHostingInquiry(ctx)
	
	for _, inquiry := range allInquiries {
		// Check if inquiry has expired
		if currentBlock > inquiry.EndTime {
			// Get escrow record for this inquiry
			escrowRecord, found := k.GetEscrowRecord(ctx, inquiry.Id)
			if found {
				// Convert creator address for refund
				creatorAddr, err := sdk.AccAddressFromBech32(escrowRecord.Creator)
				if err != nil {
					k.Logger().Error("invalid creator address in escrow record", 
						"inquiry_id", inquiry.Id, 
						"creator", escrowRecord.Creator, 
						"error", err)
					continue
				}

				// Check if there are still active contracts for this inquiry
				// TODO: Add logic to check active contracts before refunding
				// For now, we'll refund expired inquiries

				// Refund escrowed funds
				err = k.RefundFunds(ctx, creatorAddr, escrowRecord.Amount)
				if err != nil {
					k.Logger().Error("failed to refund expired inquiry escrow", 
						"inquiry_id", inquiry.Id, 
						"amount", escrowRecord.Amount.String(), 
						"error", err)
					continue
				}

				// Remove escrow record
				k.RemoveEscrowRecord(ctx, inquiry.Id)
				
				// Remove expired inquiry
				k.RemoveHostingInquiry(sdkCtx, inquiry.Id)

				k.Logger().Info("processed expired inquiry", 
					"inquiry_id", inquiry.Id, 
					"refunded_amount", escrowRecord.Amount.String())
			}
		}
	}

	return nil
}

// GetEscrowRecordsByCreator returns all escrow records for a specific creator
func (k Keeper) GetEscrowRecordsByCreator(ctx context.Context, creator string) []EscrowRecord {
	allRecords := k.GetAllEscrowRecords(ctx)
	var creatorRecords []EscrowRecord
	
	for _, record := range allRecords {
		if record.Creator == creator {
			creatorRecords = append(creatorRecords, record)
		}
	}
	
	return creatorRecords
}

// GetActiveEscrowRecords returns all escrow records that are still active (not expired)
func (k Keeper) GetActiveEscrowRecords(ctx context.Context) []EscrowRecord {
	inquiries := k.GetAllHostingInquiry(ctx)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	var activeRecords []EscrowRecord
	
	for _, inquiry := range inquiries {
		if currentHeight <= inquiry.EndTime {
			record, found := k.GetEscrowRecord(ctx, inquiry.Id)
			if found {
				activeRecords = append(activeRecords, record)
			}
		}
	}
	
	return activeRecords
}

// UpdateEscrowAmount updates the amount in an existing escrow record
func (k Keeper) UpdateEscrowAmount(ctx context.Context, inquiryId uint64, newAmount sdk.Coin) error {
	record, found := k.GetEscrowRecord(ctx, inquiryId)
	if !found {
		return fmt.Errorf("escrow record not found for inquiry %d", inquiryId)
	}
	
	record.Amount = newAmount
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	
	bz, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal updated escrow record: %w", err)
	}
	storeAdapter.Set(EscrowKey(inquiryId), bz)
	
	return nil
}

// GetTotalEscrowedAmount returns the total amount of tokens currently in escrow
func (k Keeper) GetTotalEscrowedAmount(ctx context.Context, denom string) sdk.Coin {
	allRecords := k.GetAllEscrowRecords(ctx)
	totalAmount := math.ZeroInt()
	
	for _, record := range allRecords {
		if record.Amount.Denom == denom {
			totalAmount = totalAmount.Add(record.Amount.Amount)
		}
	}
	
	return sdk.NewCoin(denom, totalAmount)
}