package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/runtime"
	"sort"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hanshq/filespace-chain/x/filespacechain/types"
)

func (k msgServer) CreateHostingInquiry(goCtx context.Context, msg *types.MsgCreateHostingInquiry) (*types.MsgCreateHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingInquiry = types.HostingInquiry{
		Creator:         msg.Creator,
		FileEntryCid:    msg.FileEntryCid,
		ReplicationRate: msg.ReplicationRate,
		EscrowAmount:    msg.EscrowAmount,
		EndTime:         msg.EndTime,
	}

	id := k.AppendHostingInquiry(
		ctx,
		hostingInquiry,
	)

	// Find the lowest hosting offers
	lowestOffers, err := k.GetLowestHostingOffers(ctx, hostingInquiry)

	if err != nil {
		fmt.Printf("Could not find lowest offers %s\n")
	}

	for i, offer := range lowestOffers {
		if uint64(i) >= hostingInquiry.ReplicationRate {
			break // Do not create more contracts than the replication rate
		}

		// Create the hosting contract
		contract := types.HostingContract{
			InquiryId: id,
			OfferId:   offer.Id,
			Creator:   offer.Creator,
			// ... other contract fields
		}
		k.AppendHostingContract(ctx, contract)

		// Emit event for each contract creation
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"createHostingContract", // This is the event type
				sdk.NewAttribute("InquiryId", strconv.FormatUint(id, 10)),
				sdk.NewAttribute("OfferId", strconv.FormatUint(offer.Id, 10)),
				sdk.NewAttribute("FileEntryCID", msg.FileEntryCid),
				sdk.NewAttribute("HosterId", offer.Creator),
				sdk.NewAttribute("Test", strconv.Itoa(1)),
				// ... other attributes
			),
		)
	}

	return &types.MsgCreateHostingInquiryResponse{
		Id: id,
	}, nil
}

// GetLowestHostingOffers fetches and returns the lowest hosting offers based on the inquiry criteria.
func (k Keeper) GetLowestHostingOffers(ctx sdk.Context, inquiry types.HostingInquiry) ([]types.HostingOffer, error) {
	var offers []types.HostingOffer
	var lowestOffers []types.HostingOffer

	// Example: Fetch all offers from the store
	// This is a simplified example. You should replace it with the actual method to query offers from your store.
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.HostingOfferKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostingOffer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		offers = append(offers, val)
	}

	// Sort offers based on your criteria, e.g., price
	sort.Slice(offers, func(i, j int) bool {
		return offers[i].PricePerBlock.IsLT(offers[j].PricePerBlock) // Assuming 'Price' is a field in your HostingOffer type
	})

	// Filter and select top 'n' offers based on inquiry's replication rate
	for _, offer := range offers {
		if uint64(len(lowestOffers)) >= inquiry.ReplicationRate {
			break
		}
		// Additional filtering criteria can be applied here
		lowestOffers = append(lowestOffers, offer)
	}

	return lowestOffers, nil
}

func (k msgServer) UpdateHostingInquiry(goCtx context.Context, msg *types.MsgUpdateHostingInquiry) (*types.MsgUpdateHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hostingInquiry = types.HostingInquiry{
		Creator:         msg.Creator,
		Id:              msg.Id,
		FileEntryCid:    msg.FileEntryCid,
		ReplicationRate: msg.ReplicationRate,
		EscrowAmount:    msg.EscrowAmount,
		EndTime:         msg.EndTime,
	}

	// Checks that the element exists
	val, found := k.GetHostingInquiry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHostingInquiry(ctx, hostingInquiry)

	return &types.MsgUpdateHostingInquiryResponse{}, nil
}

func (k msgServer) DeleteHostingInquiry(goCtx context.Context, msg *types.MsgDeleteHostingInquiry) (*types.MsgDeleteHostingInquiryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHostingInquiry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHostingInquiry(ctx, msg.Id)

	return &types.MsgDeleteHostingInquiryResponse{}, nil
}
