package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		authority:     authority,
		logger:        logger,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// CalculateEscrowAmount calculates the total escrow amount needed for a hosting inquiry
func (k Keeper) CalculateEscrowAmount(ctx context.Context, fileSize, duration, replication uint64) (sdk.Coin, error) {
	params := k.GetParams(ctx)
	
	// Convert to math types for calculation
	fileSizeDec := math.LegacyNewDecFromInt(math.NewIntFromUint64(fileSize))
	durationDec := math.LegacyNewDecFromInt(math.NewIntFromUint64(duration))
	replicationDec := math.LegacyNewDecFromInt(math.NewIntFromUint64(replication))
	
	// Calculate: fileSize × duration × replication × basePricePerBytePerBlock
	totalCost := fileSizeDec.Mul(durationDec).Mul(replicationDec).Mul(params.BasePricePerBytePerBlock)
	
	// Convert back to Int (truncating decimals)
	totalCostInt := totalCost.TruncateInt()
	
	// For now, use the first denomination from the bank module (typically the native token)
	// In production, this should be configurable
	return sdk.NewCoin("token", totalCostInt), nil
}

// EscrowFunds locks funds from sender account to the module account for escrow
func (k Keeper) EscrowFunds(ctx context.Context, sender sdk.AccAddress, amount sdk.Coin) error {
	coins := sdk.NewCoins(amount)
	
	// Transfer funds from sender to module account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins)
	if err != nil {
		return fmt.Errorf("failed to escrow funds: %w", err)
	}
	
	k.Logger().Info("funds escrowed", 
		"sender", sender.String(), 
		"amount", amount.String(),
	)
	
	return nil
}

// ReleaseFunds sends escrowed funds from module account to recipient
func (k Keeper) ReleaseFunds(ctx context.Context, recipient sdk.AccAddress, amount sdk.Coin) error {
	coins := sdk.NewCoins(amount)
	
	// Transfer funds from module account to recipient
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)
	if err != nil {
		return fmt.Errorf("failed to release funds: %w", err)
	}
	
	k.Logger().Info("funds released", 
		"recipient", recipient.String(), 
		"amount", amount.String(),
	)
	
	return nil
}

// RefundFunds returns escrowed funds from module account back to original sender
func (k Keeper) RefundFunds(ctx context.Context, recipient sdk.AccAddress, amount sdk.Coin) error {
	coins := sdk.NewCoins(amount)
	
	// Transfer funds from module account back to recipient
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)
	if err != nil {
		return fmt.Errorf("failed to refund funds: %w", err)
	}
	
	k.Logger().Info("funds refunded", 
		"recipient", recipient.String(), 
		"amount", amount.String(),
	)
	
	return nil
}

// StakeForHostingProvider stakes tokens for a hosting provider
func (k Keeper) StakeForHostingProvider(ctx context.Context, provider sdk.AccAddress, amount sdk.Coin) error {
	// Check if provider already has a stake
	existingStake, found := k.GetProviderStake(ctx, provider.String())
	if found {
		// Add to existing stake
		newAmount := existingStake.Amount.Add(amount)
		
		// Transfer additional funds to staking module account
		coins := sdk.NewCoins(amount)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, provider, "hosting_bonded_pool", coins)
		if err != nil {
			return fmt.Errorf("failed to stake additional funds: %w", err)
		}
		
		// Update stake record
		err = k.UpdateProviderStake(ctx, provider.String(), newAmount)
		if err != nil {
			return fmt.Errorf("failed to update provider stake: %w", err)
		}
	} else {
		// Create new stake
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		currentHeight := uint64(sdkCtx.BlockHeight())
		
		// Transfer funds to staking module account
		coins := sdk.NewCoins(amount)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, provider, "hosting_bonded_pool", coins)
		if err != nil {
			return fmt.Errorf("failed to stake funds: %w", err)
		}
		
		// Store stake record
		k.SetProviderStake(ctx, provider.String(), amount, currentHeight)
	}
	
	k.Logger().Info("funds staked for hosting", 
		"provider", provider.String(), 
		"amount", amount.String(),
	)
	
	return nil
}

// GetProviderStakeAmount returns the staked amount for a provider
func (k Keeper) GetProviderStakeAmount(ctx context.Context, provider sdk.AccAddress) (sdk.Coin, error) {
	stake, found := k.GetProviderStake(ctx, provider.String())
	if !found {
		// Return zero coin if no stake found
		return sdk.NewCoin("token", math.ZeroInt()), nil
	}
	return stake.Amount, nil
}

// SlashProvider slashes a portion of provider's stake for service failures
func (k Keeper) SlashProvider(ctx context.Context, provider sdk.AccAddress, fraction math.LegacyDec) error {
	stake, found := k.GetProviderStake(ctx, provider.String())
	if !found {
		return fmt.Errorf("provider %s has no stake to slash", provider.String())
	}
	
	// Calculate slash amount
	slashAmountDec := math.LegacyNewDecFromInt(stake.Amount.Amount).Mul(fraction)
	slashAmount := sdk.NewCoin(stake.Amount.Denom, slashAmountDec.TruncateInt())
	
	// Calculate remaining stake
	remainingAmount := stake.Amount.Sub(slashAmount)
	
	// Note: In a real implementation, you might want to burn the slashed tokens
	// For now, we'll leave the tokens in the module account but reduce the tracked stake
	
	if remainingAmount.IsZero() {
		// Remove stake entirely if slashed to zero
		k.RemoveProviderStake(ctx, provider.String())
	} else {
		// Update stake with remaining amount
		err := k.UpdateProviderStake(ctx, provider.String(), remainingAmount)
		if err != nil {
			return fmt.Errorf("failed to update slashed stake: %w", err)
		}
	}
	
	k.Logger().Info("provider stake slashed", 
		"provider", provider.String(), 
		"slashed_amount", slashAmount.String(),
		"remaining_stake", remainingAmount.String(),
		"slash_fraction", fraction.String(),
	)
	
	return nil
}

// ValidateProviderStake checks if a provider has sufficient stake
func (k Keeper) ValidateProviderStake(ctx context.Context, provider sdk.AccAddress) error {
	params := k.GetParams(ctx)
	
	stake, found := k.GetProviderStake(ctx, provider.String())
	if !found {
		return fmt.Errorf("provider %s has no stake", provider.String())
	}
	
	// Convert minimum stake to same denomination as provider's stake
	minStakeAmount := sdk.NewCoin(stake.Amount.Denom, params.MinProviderStake)
	
	if stake.Amount.IsLT(minStakeAmount) {
		return fmt.Errorf("provider %s stake %s is below minimum required %s", 
			provider.String(), stake.Amount.String(), minStakeAmount.String())
	}
	
	return nil
}

// ProcessPeriodicPayments processes periodic payments for all active contracts
// This should be called in BeginBlock to distribute payments every block
func (k Keeper) ProcessPeriodicPayments(ctx context.Context) error {
	// Get all hosting contracts
	contracts := k.GetAllHostingContract(ctx)
	
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentHeight := uint64(sdkCtx.BlockHeight())
	
	for _, contract := range contracts {
		// Skip if contract hasn't started yet
		if currentHeight < contract.StartBlock {
			continue
		}
		
		// Skip if contract has ended
		if currentHeight > contract.EndBlock {
			continue
		}
		
		// Get the associated inquiry to find escrow information
		inquiry, found := k.GetHostingInquiry(ctx, contract.InquiryId)
		if !found {
			k.Logger().Error("inquiry not found for active contract",
				"contract_id", contract.Id,
				"inquiry_id", contract.InquiryId,
			)
			continue
		}
		
		// Get escrow record for this inquiry
		escrowRecord, found := k.GetEscrowRecord(ctx, inquiry.Id)
		if !found {
			k.Logger().Error("escrow record not found for inquiry",
				"inquiry_id", inquiry.Id,
			)
			continue
		}
		
		// Calculate payment per block per provider
		// 50% of total escrow is distributed as periodic payments
		periodicTotal := escrowRecord.Amount.Amount.Quo(math.NewInt(2))
		
		// Calculate blocks elapsed and total blocks
		blocksElapsed := currentHeight - contract.StartBlock + 1
		totalBlocks := contract.EndBlock - contract.StartBlock + 1
		
		// Calculate cumulative payment that should have been made by now
		cumulativePaymentDec := math.LegacyNewDecFromInt(periodicTotal).
			Mul(math.LegacyNewDec(int64(blocksElapsed))).
			Quo(math.LegacyNewDec(int64(totalBlocks)))
		cumulativePayment := cumulativePaymentDec.TruncateInt()
		
		// Get payment history to determine what's already been paid
		paymentHistory, found := k.GetPaymentHistory(ctx, contract.Id)
		if !found {
			// Initialize payment history if not found
			paymentHistory = types.PaymentHistory{
				ContractId:    contract.Id,
				TotalPaid:     sdk.NewCoin(escrowRecord.Amount.Denom, math.ZeroInt()),
				LastPaymentBlock: contract.StartBlock - 1,
			}
		}
		
		// Calculate payment due (cumulative - already paid)
		paymentDue := cumulativePayment.Sub(paymentHistory.TotalPaid.Amount)
		
		// Only process if payment is due
		if paymentDue.IsPositive() {
			// Get hosting offers associated with this contract
			offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
			if len(offers) == 0 {
				k.Logger().Error("no offers found for contract",
					"contract_id", contract.Id,
					"inquiry_id", contract.InquiryId,
				)
				continue
			}
			
			// Distribute payment equally among providers
			numProviders := int64(len(offers))
			paymentPerProviderDec := math.LegacyNewDecFromInt(paymentDue).Quo(math.LegacyNewDec(numProviders))
			paymentPerProvider := sdk.NewCoin(escrowRecord.Amount.Denom, paymentPerProviderDec.TruncateInt())
			
			// Process payment to each provider
			for _, offer := range offers {
				providerAddr, err := sdk.AccAddressFromBech32(offer.Creator)
				if err != nil {
					k.Logger().Error("invalid provider address",
						"address", offer.Creator,
						"error", err,
					)
					continue
				}
				
				// Release payment to provider
				err = k.ReleaseFunds(ctx, providerAddr, paymentPerProvider)
				if err != nil {
					k.Logger().Error("failed to release periodic payment",
						"provider", offer.Creator,
						"amount", paymentPerProvider.String(),
						"error", err,
					)
					continue
				}
			}
			
			// Update payment history
			totalPaidAmount := paymentHistory.TotalPaid.Amount.Add(paymentDue)
			paymentHistory.TotalPaid = sdk.NewCoin(escrowRecord.Amount.Denom, totalPaidAmount)
			paymentHistory.LastPaymentBlock = currentHeight
			k.SetPaymentHistory(ctx, paymentHistory)
			
			k.Logger().Info("processed periodic payments",
				"contract_id", contract.Id,
				"payment_per_provider", paymentPerProvider.String(),
				"num_providers", numProviders,
				"total_paid", paymentHistory.TotalPaid.String(),
			)
		}
	}
	
	return nil
}

// ProcessCompletionBonus processes the completion bonus for a contract
// This distributes the remaining 50% of escrow when contract successfully completes
func (k Keeper) ProcessCompletionBonus(ctx context.Context, contractId uint64) error {
	contract, found := k.GetHostingContract(ctx, contractId)
	if !found {
		return fmt.Errorf("contract %d not found", contractId)
	}
	
	// Get the associated inquiry
	inquiry, found := k.GetHostingInquiry(ctx, contract.InquiryId)
	if !found {
		return fmt.Errorf("inquiry %d not found for contract %d", contract.InquiryId, contractId)
	}
	
	// Get escrow record
	escrowRecord, found := k.GetEscrowRecord(ctx, inquiry.Id)
	if !found {
		return fmt.Errorf("escrow record not found for inquiry %d", inquiry.Id)
	}
	
	// Get payment history to see what's already been paid
	paymentHistory, found := k.GetPaymentHistory(ctx, contractId)
	if !found {
		return fmt.Errorf("payment history not found for contract %d", contractId)
	}
	
	// Calculate completion bonus (total escrow - already paid)
	completionBonus := escrowRecord.Amount.Amount.Sub(paymentHistory.TotalPaid.Amount)
	
	if !completionBonus.IsPositive() {
		k.Logger().Info("no completion bonus due, all funds already distributed",
			"contract_id", contractId,
		)
		return nil
	}
	
	// Get hosting offers to distribute bonus
	offers := k.GetHostingOffersByInquiry(ctx, contract.InquiryId)
	if len(offers) == 0 {
		return fmt.Errorf("no offers found for contract %d", contractId)
	}
	
	// Distribute bonus equally among providers
	numProviders := int64(len(offers))
	bonusPerProviderDec := math.LegacyNewDecFromInt(completionBonus).Quo(math.LegacyNewDec(numProviders))
	bonusPerProvider := sdk.NewCoin(escrowRecord.Amount.Denom, bonusPerProviderDec.TruncateInt())
	
	// Process bonus to each provider
	for _, offer := range offers {
		providerAddr, err := sdk.AccAddressFromBech32(offer.Creator)
		if err != nil {
			k.Logger().Error("invalid provider address",
				"address", offer.Creator,
				"error", err,
			)
			continue
		}
		
		// Release bonus to provider
		err = k.ReleaseFunds(ctx, providerAddr, bonusPerProvider)
		if err != nil {
			k.Logger().Error("failed to release completion bonus",
				"provider", offer.Creator,
				"amount", bonusPerProvider.String(),
				"error", err,
			)
			continue
		}
	}
	
	// Update payment history to reflect full payment
	paymentHistory.TotalPaid = escrowRecord.Amount
	paymentHistory.CompletionBonusPaid = true
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	paymentHistory.LastPaymentBlock = uint64(sdkCtx.BlockHeight())
	k.SetPaymentHistory(ctx, paymentHistory)
	
	// Remove escrow record since contract is complete
	k.RemoveEscrowRecord(ctx, inquiry.Id)
	
	k.Logger().Info("processed completion bonus",
		"contract_id", contractId,
		"bonus_per_provider", bonusPerProvider.String(),
		"num_providers", numProviders,
		"total_bonus", sdk.NewCoin(escrowRecord.Amount.Denom, completionBonus).String(),
	)
	
	return nil
}
