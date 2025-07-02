package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// StakeForHostingProvider handles provider staking operations
func (k msgServer) StakeForHostingProvider(goCtx context.Context, provider string, amount sdk.Coin) error {
	// Convert provider address
	providerAddr, err := sdk.AccAddressFromBech32(provider)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid provider address")
	}

	// Validate minimum stake amount
	params := k.GetParams(goCtx)
	minStakeAmount := sdk.NewCoin(amount.Denom, params.MinProviderStake)
	
	// For new stakes, the amount must meet minimum requirement
	_, found := k.GetProviderStake(goCtx, provider)
	if !found {
		if amount.IsLT(minStakeAmount) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
				fmt.Sprintf("stake amount %s is below minimum required %s", amount.String(), minStakeAmount.String()))
		}
	}

	// Perform staking operation
	err = k.StakeForHosting(goCtx, providerAddr, amount)
	if err != nil {
		return errorsmod.Wrap(err, "failed to stake for hosting")
	}

	return nil
}

// UnstakeFromHostingProvider handles provider unstaking operations
func (k msgServer) UnstakeFromHostingProvider(goCtx context.Context, provider string, amount sdk.Coin) error {
	// Convert provider address
	providerAddr, err := sdk.AccAddressFromBech32(provider)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid provider address")
	}

	// Get current stake
	stake, found := k.GetProviderStake(goCtx, provider)
	if !found {
		return errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "provider has no stake")
	}

	// Validate unstake amount
	if amount.Amount.GT(stake.Amount.Amount) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
			fmt.Sprintf("cannot unstake %s, only %s staked", amount.String(), stake.Amount.String()))
	}

	// Calculate remaining stake after unstaking
	remainingStake := stake.Amount.Sub(amount)
	
	// Check if remaining stake meets minimum requirement (unless unstaking everything)
	if !remainingStake.IsZero() {
		params := k.GetParams(goCtx)
		minStakeAmount := sdk.NewCoin(amount.Denom, params.MinProviderStake)
		
		if remainingStake.IsLT(minStakeAmount) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
				fmt.Sprintf("remaining stake %s would be below minimum required %s", 
					remainingStake.String(), minStakeAmount.String()))
		}
	}

	// Transfer funds from staking module account back to provider
	coins := sdk.NewCoins(amount)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(goCtx, "hosting_bonded_pool", providerAddr, coins)
	if err != nil {
		return errorsmod.Wrap(err, "failed to unstake funds")
	}

	// Update or remove stake record
	if remainingStake.IsZero() {
		k.RemoveProviderStake(goCtx, provider)
	} else {
		err = k.UpdateProviderStake(goCtx, provider, remainingStake)
		if err != nil {
			return errorsmod.Wrap(err, "failed to update stake record")
		}
	}

	k.Logger().Info("funds unstaked from hosting", 
		"provider", provider, 
		"unstaked_amount", amount.String(),
		"remaining_stake", remainingStake.String(),
	)

	return nil
}