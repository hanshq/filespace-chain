package types

import (
	"fmt"

	"cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyBasePricePerBytePerBlock = []byte("BasePricePerBytePerBlock")
	KeyMinProviderStake         = []byte("MinProviderStake")
	KeySlashingFraction         = []byte("SlashingFraction")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	basePricePerBytePerBlock math.LegacyDec,
	minProviderStake math.Int,
	slashingFraction math.LegacyDec,
) Params {
	return Params{
		BasePricePerBytePerBlock: basePricePerBytePerBlock,
		MinProviderStake:         minProviderStake,
		SlashingFraction:         slashingFraction,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		math.LegacyNewDecWithPrec(1, 12), // 0.000000000001 (1e-12)
		math.NewInt(1000000),             // 1,000,000 base units
		math.LegacyNewDecWithPrec(5, 2),  // 0.05 (5%)
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBasePricePerBytePerBlock, &p.BasePricePerBytePerBlock, validateBasePricePerBytePerBlock),
		paramtypes.NewParamSetPair(KeyMinProviderStake, &p.MinProviderStake, validateMinProviderStake),
		paramtypes.NewParamSetPair(KeySlashingFraction, &p.SlashingFraction, validateSlashingFraction),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBasePricePerBytePerBlock(p.BasePricePerBytePerBlock); err != nil {
		return err
	}
	if err := validateMinProviderStake(p.MinProviderStake); err != nil {
		return err
	}
	if err := validateSlashingFraction(p.SlashingFraction); err != nil {
		return err
	}
	return nil
}

func validateBasePricePerBytePerBlock(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("base price per byte per block must be positive: %s", v)
	}

	return nil
}

func validateMinProviderStake(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("min provider stake must be positive: %s", v)
	}

	return nil
}

func validateSlashingFraction(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("slashing fraction must be positive: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("slashing fraction must be less than or equal to 1: %s", v)
	}

	return nil
}
