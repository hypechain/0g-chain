package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgClaimUSDXMintingReward{}
var _ sdk.Msg = &MsgClaimHardReward{}
var _ sdk.Msg = &MsgClaimHardRewardVVesting{}
var _ sdk.Msg = &MsgClaimUSDXMintingRewardVVesting{}

// MsgClaimUSDXMintingReward message type used to claim USDX minting rewards
type MsgClaimUSDXMintingReward struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimUSDXMintingReward returns a new MsgClaimUSDXMintingReward.
func NewMsgClaimUSDXMintingReward(sender sdk.AccAddress, multiplierName string) MsgClaimUSDXMintingReward {
	return MsgClaimUSDXMintingReward{
		Sender:         sender,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimUSDXMintingReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimUSDXMintingReward) Type() string { return "claim_usdx_minting_reward" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimUSDXMintingReward) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimUSDXMintingReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimUSDXMintingReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgClaimUSDXMintingRewardVVesting message type used to claim USDX minting rewards for validator vesting accounts
type MsgClaimUSDXMintingRewardVVesting struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver       sdk.AccAddress `json:"receiver" yaml:"receiver"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimUSDXMintingRewardVVesting returns a new MsgClaimUSDXMintingReward.
func NewMsgClaimUSDXMintingRewardVVesting(sender, receiver sdk.AccAddress, multiplierName string) MsgClaimUSDXMintingRewardVVesting {
	return MsgClaimUSDXMintingRewardVVesting{
		Sender:         sender,
		Receiver:       receiver,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimUSDXMintingRewardVVesting) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimUSDXMintingRewardVVesting) Type() string {
	return "claim_usdx_minting_reward_vvesting"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimUSDXMintingRewardVVesting) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimUSDXMintingRewardVVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimUSDXMintingRewardVVesting) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgClaimHardReward message type used to claim Hard liquidity provider rewards
type MsgClaimHardReward struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimHardReward returns a new MsgClaimHardReward.
func NewMsgClaimHardReward(sender sdk.AccAddress, multiplierName string) MsgClaimHardReward {
	return MsgClaimHardReward{
		Sender:         sender,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimHardReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimHardReward) Type() string {
	return "claim_hard_reward"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimHardReward) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimHardReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimHardReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgClaimHardRewardVVesting message type used to claim Hard liquidity provider rewards for validator vesting accounts
type MsgClaimHardRewardVVesting struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver       sdk.AccAddress `json:"receiver" yaml:"receiver"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimHardRewardVVesting returns a new MsgClaimHardReward.
func NewMsgClaimHardRewardVVesting(sender, receiver sdk.AccAddress, multiplierName string) MsgClaimHardRewardVVesting {
	return MsgClaimHardRewardVVesting{
		Sender:         sender,
		Receiver:       receiver,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimHardRewardVVesting) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimHardRewardVVesting) Type() string {
	return "claim_hard_reward_vvesting"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimHardRewardVVesting) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimHardRewardVVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimHardRewardVVesting) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
