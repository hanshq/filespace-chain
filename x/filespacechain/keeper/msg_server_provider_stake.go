package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
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
	err = k.Keeper.StakeForHostingProvider(goCtx, providerAddr, amount)
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

// StakeForHosting handles MsgStakeForHosting transaction
func (k msgServer) StakeForHosting(goCtx context.Context, msg *types.MsgStakeForHosting) (*types.MsgStakeForHostingResponse, error) {
	// Convert creator address
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	// Validate minimum stake amount for new stakes
	params := k.GetParams(goCtx)
	minStakeAmount := sdk.NewCoin(msg.Amount.Denom, params.MinProviderStake)
	
	_, found := k.GetProviderStake(goCtx, msg.Creator)
	if !found && msg.Amount.IsLT(minStakeAmount) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
			fmt.Sprintf("stake amount %s is below minimum required %s", msg.Amount.String(), minStakeAmount.String()))
	}

	// Perform staking operation using the keeper method
	err = k.Keeper.StakeForHostingProvider(goCtx, creatorAddr, msg.Amount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to stake for hosting")
	}

	return &types.MsgStakeForHostingResponse{}, nil
}

// UnstakeFromHosting handles MsgUnstakeFromHosting transaction
func (k msgServer) UnstakeFromHosting(goCtx context.Context, msg *types.MsgUnstakeFromHosting) (*types.MsgUnstakeFromHostingResponse, error) {
	// Convert creator address
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	// Get current stake
	stake, found := k.GetProviderStake(goCtx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "provider has no stake")
	}

	// Validate unstake amount
	if msg.Amount.Amount.GT(stake.Amount.Amount) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
			fmt.Sprintf("cannot unstake %s, only %s staked", msg.Amount.String(), stake.Amount.String()))
	}

	// Calculate remaining stake after unstaking
	remainingStake := stake.Amount.Sub(msg.Amount)
	
	// Check if remaining stake meets minimum requirement (unless unstaking everything)
	if !remainingStake.IsZero() {
		params := k.GetParams(goCtx)
		minStakeAmount := sdk.NewCoin(msg.Amount.Denom, params.MinProviderStake)
		
		if remainingStake.IsLT(minStakeAmount) {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, 
				fmt.Sprintf("remaining stake %s would be below minimum required %s", 
					remainingStake.String(), minStakeAmount.String()))
		}
	}

	// Transfer funds from staking module account back to provider
	coins := sdk.NewCoins(msg.Amount)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(goCtx, "hosting_bonded_pool", creatorAddr, coins)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to unstake funds")
	}

	// Update or remove stake record
	if remainingStake.IsZero() {
		k.RemoveProviderStake(goCtx, msg.Creator)
	} else {
		err = k.UpdateProviderStake(goCtx, msg.Creator, remainingStake)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to update stake record")
		}
	}

	k.Logger().Info("funds unstaked from hosting", 
		"provider", msg.Creator, 
		"unstaked_amount", msg.Amount.String(),
		"remaining_stake", remainingStake.String(),
	)

	return &types.MsgUnstakeFromHostingResponse{}, nil
}