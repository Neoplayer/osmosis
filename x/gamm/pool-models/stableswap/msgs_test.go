package stableswap_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	appParams "github.com/osmosis-labs/osmosis/v12/app/params"
	stableswap "github.com/osmosis-labs/osmosis/v12/x/gamm/pool-models/stableswap"
	"github.com/osmosis-labs/osmosis/v12/x/gamm/types"
)

func TestMsgCreateStableswapPool(t *testing.T) {
	appParams.SetAddressPrefixes()
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address()).String()
	invalidAddr := sdk.AccAddress("invalid")
	_, invalidAccErr := sdk.AccAddressFromBech32(invalidAddr.String())

	createMsg := func(after func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
		testPoolAsset := sdk.Coins{
			sdk.NewCoin("osmo", sdk.NewInt(100)),
			sdk.NewCoin("atom", sdk.NewInt(100)),
		}

		poolParams := &stableswap.PoolParams{
			SwapFee: sdk.NewDecWithPrec(1, 2),
			ExitFee: sdk.NewDecWithPrec(1, 2),
		}

		msg := &stableswap.MsgCreateStableswapPool{
			Sender:               addr1,
			PoolParams:           poolParams,
			InitialPoolLiquidity: testPoolAsset,
			ScalingFactors:       []uint64{1, 1},
			FuturePoolGovernor:   "",
		}

		return after(*msg)
	}

	default_msg := createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
		// Do nothing
		return msg
	})

	require.Equal(t, default_msg.Route(), types.RouterKey)
	require.Equal(t, default_msg.Type(), "create_stableswap_pool")
	signers := default_msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1)

	tests := []struct {
		name      string
		msg       stableswap.MsgCreateStableswapPool
		expectErr error
	}{
		{
			name: "proper msg",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				// Do nothing
				return msg
			}),
		},
		{
			name: "invalid sender",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.Sender = invalidAddr.String()
				return msg
			}),
			expectErr: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", invalidAccErr),
		},
		{
			name: "has nil InitialPoolLiquidity ",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.InitialPoolLiquidity = nil
				return msg
			}),
			expectErr: types.ErrTooFewPoolAssets,
		},
		{
			name: "has one coin in InitialPoolLiquidity",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.InitialPoolLiquidity = sdk.Coins{
					sdk.NewCoin("osmo", sdk.NewInt(100)),
				}
				return msg
			}),
			expectErr: types.ErrTooFewPoolAssets,
		},
		{
			name: "have assets in excess of cap",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.InitialPoolLiquidity = sdk.Coins{
					sdk.NewCoin("osmo", sdk.NewInt(100)),
					sdk.NewCoin("atom", sdk.NewInt(100)),
					sdk.NewCoin("usdt", sdk.NewInt(100)),
					sdk.NewCoin("usdc", sdk.NewInt(100)),
					sdk.NewCoin("juno", sdk.NewInt(100)),
					sdk.NewCoin("akt", sdk.NewInt(100)),
					sdk.NewCoin("regen", sdk.NewInt(100)),
					sdk.NewCoin("band", sdk.NewInt(100)),
					sdk.NewCoin("evmos", sdk.NewInt(100)),
				}
				return msg
			}),
			expectErr: types.ErrTooManyPoolAssets,
		},
		{
			name: "negative swap fee with zero exit fee",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.PoolParams = &stableswap.PoolParams{
					SwapFee: sdk.NewDecWithPrec(-1, 2),
					ExitFee: sdk.NewDecWithPrec(0, 0),
				}
				return msg
			}),
			expectErr: types.ErrNegativeSwapFee,
		},
		{
			name: "scaling factors with invalid lenght",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.ScalingFactors = []uint64{1, 2, 3}
				return msg
			}),
			expectErr: types.ErrInvalidStableswapScalingFactors,
		},
		{
			name: "invalid governor",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "invalid_cosmos_address"
				return msg
			}),
			expectErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid future governor: %s", "invalid_cosmos_address")),
		},
		{
			name: "invalid governor : len governor > 2",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "lptoken,1000h,invalid_cosmos_address"
				return msg
			}),
			expectErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid future governor: %s", "lptoken,1000h,invalid_cosmos_address")),
		},
		{
			name: "invalid governor : len governor > 2",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "lptoken,1000h,invalid_cosmos_address"
				return msg
			}),
			expectErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid future governor: %s", "lptoken,1000h,invalid_cosmos_address")),
		},
		{
			name: "valid governor: err when parse duration ",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "lptoken, invalid_duration"
				return msg
			}),
			expectErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid future governor: %s", "lptoken, invalid_duration")),
		},
		{
			name: "valid governor: just lock duration for pool token",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "1000h"
				return msg
			}),
		},
		{
			name: "valid governor: address",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = "osmo1fqlr98d45v5ysqgp6h56kpujcj4cvsjnjq9nck"
				return msg
			}),
		},
		{
			name: "valid governor: address",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.FuturePoolGovernor = ""
				return msg
			}),
		},
		{
			name: "zero swap fee, zero exit fee",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.PoolParams = &stableswap.PoolParams{
					ExitFee: sdk.NewDecWithPrec(0, 0),
					SwapFee: sdk.NewDecWithPrec(0, 0),
				}
				return msg
			}),
		},
		{
			name: "multi assets pool",
			msg: createMsg(func(msg stableswap.MsgCreateStableswapPool) stableswap.MsgCreateStableswapPool {
				msg.InitialPoolLiquidity = sdk.Coins{
					sdk.NewCoin("osmo", sdk.NewInt(100)),
					sdk.NewCoin("atom", sdk.NewInt(100)),
					sdk.NewCoin("usdt", sdk.NewInt(100)),
					sdk.NewCoin("usdc", sdk.NewInt(100)),
				}
				msg.ScalingFactors = []uint64{1, 1, 1, 1}
				return msg
			}),
		},
	}

	for _, test := range tests {
		err := test.msg.ValidateBasic()
		if test.expectErr == nil {
			require.NoError(t, err, "test: %v", test.name)
		} else {
			require.Error(t, err, "test: %v", test.name)
			require.ErrorAs(t, test.expectErr, &err)
		}
	}
}
