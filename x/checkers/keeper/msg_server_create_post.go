package keeper

import (
	"context"
	"fmt"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("msg::", msg)
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatePostResponse{}, nil
}
