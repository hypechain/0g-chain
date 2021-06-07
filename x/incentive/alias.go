// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/kava-labs/kava/x/incentive/types
// ALIASGEN: github.com/kava-labs/kava/x/incentive/keeper
package incentive

import (
	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

const (
	USDXMintingClaimType               = types.USDXMintingClaimType
	HardLiquidityProviderClaimType     = types.HardLiquidityProviderClaimType
	BondDenom                          = types.BondDenom
	EventTypeClaim                     = types.EventTypeClaim
	EventTypeRewardPeriod              = types.EventTypeRewardPeriod
	EventTypeClaimPeriod               = types.EventTypeClaimPeriod
	EventTypeClaimPeriodExpiry         = types.EventTypeClaimPeriodExpiry
	AttributeValueCategory             = types.AttributeValueCategory
	AttributeKeyClaimedBy              = types.AttributeKeyClaimedBy
	AttributeKeyClaimAmount            = types.AttributeKeyClaimAmount
	AttributeKeyClaimType              = types.AttributeKeyClaimType
	AttributeKeyRewardPeriod           = types.AttributeKeyRewardPeriod
	AttributeKeyClaimPeriod            = types.AttributeKeyClaimPeriod
	ModuleName                         = types.ModuleName
	StoreKey                           = types.StoreKey
	RouterKey                          = types.RouterKey
	DefaultParamspace                  = types.DefaultParamspace
	QuerierRoute                       = types.QuerierRoute
	Small                              = types.Small
	Medium                             = types.Medium
	Large                              = types.Large
	QueryGetRewards                    = types.QueryGetRewards
	QueryGetHardRewards                = types.QueryGetHardRewards
	QueryGetHardRewardsUnsynced        = types.QueryGetHardRewardsUnsynced
	QueryGetUSDXMintingRewards         = types.QueryGetUSDXMintingRewards
	QueryGetUSDXMintingRewardsUnsynced = types.QueryGetUSDXMintingRewardsUnsynced
	QueryGetRewardFactors              = types.QueryGetRewardFactors
	QueryGetParams                     = types.QueryGetParams
	QueryGetRewardPeriods              = types.QueryGetRewardPeriods
	QueryGetClaimPeriods               = types.QueryGetClaimPeriods
	RestClaimCollateralType            = types.RestClaimCollateralType
	RestClaimOwner                     = types.RestClaimOwner
	RestClaimType                      = types.RestClaimType
	RestUnsynced                       = types.RestUnsynced
	BeginningOfMonth                   = keeper.BeginningOfMonth
	MidMonth                           = keeper.MidMonth
	PaymentHour                        = keeper.PaymentHour
)

var (
	// functions aliases
	GetTotalVestingPeriodLength              = types.GetTotalVestingPeriodLength
	NewUSDXMintingClaim                      = types.NewUSDXMintingClaim
	NewHardLiquidityProviderClaim            = types.NewHardLiquidityProviderClaim
	NewMultiRewardPeriod                     = types.NewMultiRewardPeriod
	NewRewardIndex                           = types.NewRewardIndex
	NewMultiRewardIndex                      = types.NewMultiRewardIndex
	RegisterCodec                            = types.RegisterCodec
	NewGenesisState                          = types.NewGenesisState
	DefaultGenesisState                      = types.DefaultGenesisState
	NewGenesisAccumulationTime               = types.NewGenesisAccumulationTime
	NewMsgClaimUSDXMintingReward             = types.NewMsgClaimUSDXMintingReward
	NewMsgClaimUSDXMintingRewardVVesting     = types.NewMsgClaimUSDXMintingRewardVVesting
	NewMsgClaimHardReward                    = types.NewMsgClaimHardReward
	NewMsgClaimHardRewardVVesting            = types.NewMsgClaimHardRewardVVesting
	NewParams                                = types.NewParams
	DefaultParams                            = types.DefaultParams
	ParamKeyTable                            = types.ParamKeyTable
	NewRewardPeriod                          = types.NewRewardPeriod
	NewMultiplier                            = types.NewMultiplier
	NewPeriod                                = types.NewPeriod
	NewQueryRewardsParams                    = types.NewQueryRewardsParams
	NewQueryUSDXMintingRewardsParams         = types.NewQueryUSDXMintingRewardsParams
	NewQueryUSDXMintingRewardsUnsyncedParams = types.NewQueryUSDXMintingRewardsUnsyncedParams
	NewQueryHardRewardsParams                = types.NewQueryHardRewardsParams
	NewQueryHardRewardsUnsyncedParams        = types.NewQueryHardRewardsUnsyncedParams
	NewQueryRewardFactorsParams              = types.NewQueryRewardFactorsParams
	NewRewardFactor                          = types.NewRewardFactor
	NewKeeper                                = keeper.NewKeeper
	NewQuerier                               = keeper.NewQuerier
	CalculateTimeElapsed                     = keeper.CalculateTimeElapsed

	// variable aliases
	ModuleCdc                                       = types.ModuleCdc
	ErrClaimNotFound                                = types.ErrClaimNotFound
	ErrRewardPeriodNotFound                         = types.ErrRewardPeriodNotFound
	ErrInvalidAccountType                           = types.ErrInvalidAccountType
	ErrNoClaimsFound                                = types.ErrNoClaimsFound
	ErrInsufficientModAccountBalance                = types.ErrInsufficientModAccountBalance
	ErrAccountNotFound                              = types.ErrAccountNotFound
	ErrInvalidMultiplier                            = types.ErrInvalidMultiplier
	ErrZeroClaim                                    = types.ErrZeroClaim
	ErrClaimExpired                                 = types.ErrClaimExpired
	ErrInvalidClaimType                             = types.ErrInvalidClaimType
	ErrInvalidClaimOwner                            = types.ErrInvalidClaimOwner
	USDXMintingClaimKeyPrefix                       = types.USDXMintingClaimKeyPrefix
	USDXMintingRewardFactorKeyPrefix                = types.USDXMintingRewardFactorKeyPrefix
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix   = types.PreviousUSDXMintingRewardAccrualTimeKeyPrefix
	HardLiquidityClaimKeyPrefix                     = types.HardLiquidityClaimKeyPrefix
	HardSupplyRewardIndexesKeyPrefix                = types.HardSupplyRewardIndexesKeyPrefix
	PreviousHardSupplyRewardAccrualTimeKeyPrefix    = types.PreviousHardSupplyRewardAccrualTimeKeyPrefix
	HardBorrowRewardIndexesKeyPrefix                = types.HardBorrowRewardIndexesKeyPrefix
	PreviousHardBorrowRewardAccrualTimeKeyPrefix    = types.PreviousHardBorrowRewardAccrualTimeKeyPrefix
	HardDelegatorRewardFactorKeyPrefix              = types.HardDelegatorRewardFactorKeyPrefix
	PreviousHardDelegatorRewardAccrualTimeKeyPrefix = types.PreviousHardDelegatorRewardAccrualTimeKeyPrefix
	USDXMintingRewardDenom                          = types.USDXMintingRewardDenom
	HardLiquidityRewardDenom                        = types.HardLiquidityRewardDenom
	KeyUSDXMintingRewardPeriods                     = types.KeyUSDXMintingRewardPeriods
	KeyHardSupplyRewardPeriods                      = types.KeyHardSupplyRewardPeriods
	KeyHardBorrowRewardPeriods                      = types.KeyHardBorrowRewardPeriods
	KeyHardDelegatorRewardPeriods                   = types.KeyHardDelegatorRewardPeriods
	KeyClaimEnd                                     = types.KeyClaimEnd
	KeyMultipliers                                  = types.KeyMultipliers
	DefaultActive                                   = types.DefaultActive
	DefaultRewardPeriods                            = types.DefaultRewardPeriods
	DefaultMultiRewardPeriods                       = types.DefaultMultiRewardPeriods
	DefaultMultipliers                              = types.DefaultMultipliers
	DefaultUSDXClaims                               = types.DefaultUSDXClaims
	DefaultHardClaims                               = types.DefaultHardClaims
	DefaultGenesisAccumulationTimes                 = types.DefaultGenesisAccumulationTimes
	DefaultClaimEnd                                 = types.DefaultClaimEnd
	GovDenom                                        = types.GovDenom
	PrincipalDenom                                  = types.PrincipalDenom
	IncentiveMacc                                   = types.IncentiveMacc
)

type (
	Claim                                 = types.Claim
	Claims                                = types.Claims
	BaseClaim                             = types.BaseClaim
	BaseMultiClaim                        = types.BaseMultiClaim
	USDXMintingClaim                      = types.USDXMintingClaim
	USDXMintingClaims                     = types.USDXMintingClaims
	HardLiquidityProviderClaim            = types.HardLiquidityProviderClaim
	HardLiquidityProviderClaims           = types.HardLiquidityProviderClaims
	MultiRewardPeriod                     = types.MultiRewardPeriod
	MultiRewardPeriods                    = types.MultiRewardPeriods
	RewardIndex                           = types.RewardIndex
	RewardIndexes                         = types.RewardIndexes
	MultiRewardIndex                      = types.MultiRewardIndex
	MultiRewardIndexes                    = types.MultiRewardIndexes
	SupplyKeeper                          = types.SupplyKeeper
	StakingKeeper                         = types.StakingKeeper
	CdpKeeper                             = types.CdpKeeper
	HardKeeper                            = types.HardKeeper
	AccountKeeper                         = types.AccountKeeper
	CDPHooks                              = types.CDPHooks
	HARDHooks                             = types.HARDHooks
	GenesisState                          = types.GenesisState
	GenesisAccumulationTime               = types.GenesisAccumulationTime
	GenesisAccumulationTimes              = types.GenesisAccumulationTimes
	MsgClaimUSDXMintingReward             = types.MsgClaimUSDXMintingReward
	MsgClaimUSDXMintingRewardVVesting     = types.MsgClaimUSDXMintingRewardVVesting
	MsgClaimHardReward                    = types.MsgClaimHardReward
	MsgClaimHardRewardVVesting            = types.MsgClaimHardRewardVVesting
	Params                                = types.Params
	RewardPeriod                          = types.RewardPeriod
	RewardPeriods                         = types.RewardPeriods
	Multiplier                            = types.Multiplier
	Multipliers                           = types.Multipliers
	MultiplierName                        = types.MultiplierName
	QueryRewardsParams                    = types.QueryRewardsParams
	QueryUSDXMintingRewardsParams         = types.QueryUSDXMintingRewardsParams
	QueryUSDXMintingRewardsUnsyncedParams = types.QueryUSDXMintingRewardsUnsyncedParams
	QueryHardRewardsParams                = types.QueryHardRewardsParams
	QueryHardRewardsUnsyncedParams        = types.QueryHardRewardsUnsyncedParams
	QueryRewardFactorsParams              = types.QueryRewardFactorsParams
	RewardFactor                          = types.RewardFactor
	RewardFactors                         = types.RewardFactors
	Hooks                                 = keeper.Hooks
	Keeper                                = keeper.Keeper
)
