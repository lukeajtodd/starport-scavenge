package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lukeajtodd/scavenge/x/scavenge/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) RevealSolution(goCtx context.Context, msg *types.MsgRevealSolution) (*types.MsgRevealSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	solutionScavengerBytes := []byte(msg.Solution + msg.Creator)

	solutionScavengerHash := sha256.Sum256(solutionScavengerBytes)

	solutionScavengerHashString := hex.EncodeToString(solutionScavengerHash[:])

	_, isFound := k.GetCommit(ctx, solutionScavengerHashString)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Commit with that hash doesn't exist")
	}

	solutionHash := sha256.Sum256([]byte(msg.Solution))

	solutionHashString := hex.EncodeToString(solutionHash[:])

	scavenge, isFound := k.GetScavenge(ctx, solutionHashString)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Solution with that hash doesn't exist")
	}

	_, err := sdk.AccAddressFromBech32(scavenge.Scavenger)

	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge has already been completed")
	}

	scavenge.Scavenger = msg.Creator
	scavenge.Solution = msg.Solution

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	scavenger, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err != nil {
		panic(err)
	}

	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, scavenger, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetScavenge(ctx, scavenge)

	return &types.MsgRevealSolutionResponse{}, nil
}
