package service

import (
	"errors"
	"math/big"
	"net/http"
	"strings"

	"github.com/nlsh710599/and-practice/internal/utils"
)

var ErrInvalidName = errors.New("invalid number name")
var ErrInvalidValue = errors.New("invalid number value")
var ErrCreateDuplicateNumber = errors.New("number already exist")
var ErrKnown = errors.New("unknown error")

func isDuplicateKey(err error) bool {
	return strings.Contains(err.Error(), "23505")
}

func (ctrl *BigNumberComputationService) Create(r *http.Request, args *CreateRequest, reply *string) error {
	name := args.Name
	value := args.Value
	if !utils.ValidName(name) {
		return ErrInvalidName
	}
	if !utils.IsNumeric(value) {
		return ErrInvalidValue
	}
	if err := ctrl.RDS.CreateNumber(name, value); err != nil {
		if isDuplicateKey(err) {
			return ErrCreateDuplicateNumber
		}
		return err
	}
	return nil
}

func (ctrl *BigNumberComputationService) Update(r *http.Request, args *UpdateRequest, reply *string) error {
	name := args.Name
	value := args.Value
	if !utils.ValidName(name) {
		return ErrInvalidName
	}
	if !utils.IsNumeric(value) {
		return ErrInvalidValue
	}
	if err := ctrl.RDS.UpdateNumber(name, value); err != nil {
		return err
	}
	return nil
}

func (ctrl *BigNumberComputationService) Delete(r *http.Request, args *DeleteRequest, reply *string) error {
	name := args.Name
	if !utils.ValidName(name) {
		return ErrInvalidName
	}
	if err := ctrl.RDS.DeleteNumber(name); err != nil {
		return err
	}
	return nil
}

func (ctrl *BigNumberComputationService) Add(r *http.Request, args *AddRequest, reply *string) error {
	addend := new(big.Float)
	summand := new(big.Float)
	var ok bool

	numMap, err := ctrl.RDS.ReadNumber([]string{args.Addend, args.Summand})
	if err != nil {
		return err
	}

	if utils.IsNumeric(args.Addend) {
		addend, ok = addend.SetString(args.Addend)
		if !ok {
			return err
		}
	} else {
		addend, ok = addend.SetString(numMap[args.Addend])
		if !ok {
			return err
		}
	}

	if utils.IsNumeric(args.Summand) {
		summand, ok = summand.SetString(args.Summand)
		if !ok {
			return err
		}
	} else {
		summand, ok = summand.SetString(numMap[args.Summand])
		if !ok {
			return err
		}
	}

	*reply = new(big.Float).Add(addend, summand).String()

	return nil
}

func (ctrl *BigNumberComputationService) Subtract(r *http.Request, args *SubtractRequest, reply *string) error {
	subtrahend := new(big.Float)
	minuend := new(big.Float)
	var ok bool

	numMap, err := ctrl.RDS.ReadNumber([]string{args.Subtrahend, args.Minuend})
	if err != nil {
		return err
	}

	if utils.IsNumeric(args.Subtrahend) {
		subtrahend, ok = subtrahend.SetString(args.Subtrahend)
		if !ok {
			return err
		}
	} else {
		subtrahend, ok = subtrahend.SetString(numMap[args.Subtrahend])
		if !ok {
			return err
		}
	}

	if utils.IsNumeric(args.Minuend) {
		minuend, ok = minuend.SetString(args.Minuend)
		if !ok {
			return err
		}
	} else {
		minuend, ok = minuend.SetString(numMap[args.Minuend])
		if !ok {
			return err
		}
	}

	*reply = new(big.Float).Sub(subtrahend, minuend).String()

	return nil
}

func (ctrl *BigNumberComputationService) Multiply(r *http.Request, args *MultiplyRequest, reply *string) error {
	multiplicand := new(big.Float)
	multiplier := new(big.Float)
	var ok bool

	numMap, err := ctrl.RDS.ReadNumber([]string{args.Multiplicand, args.Multiplier})
	if err != nil {
		return err
	}

	if utils.IsNumeric(args.Multiplicand) {
		multiplicand, ok = multiplicand.SetString(args.Multiplicand)
		if !ok {
			return err
		}
	} else {
		multiplicand, ok = multiplicand.SetString(numMap[args.Multiplicand])
		if !ok {
			return err
		}
	}

	if utils.IsNumeric(args.Multiplier) {
		multiplier, ok = multiplier.SetString(args.Multiplier)
		if !ok {
			return err
		}
	} else {
		multiplier, ok = multiplier.SetString(numMap[args.Multiplier])
		if !ok {
			return err
		}
	}

	*reply = new(big.Float).Mul(multiplicand, multiplier).String()

	return nil
}

func (ctrl *BigNumberComputationService) Divide(r *http.Request, args *DivideRequest, reply *string) error {
	dividend := new(big.Float)
	divisor := new(big.Float)
	var ok bool

	numMap, err := ctrl.RDS.ReadNumber([]string{args.Dividend, args.Divisor})
	if err != nil {
		return err
	}

	if utils.IsNumeric(args.Dividend) {
		dividend, ok = dividend.SetString(args.Dividend)
		if !ok {
			return err
		}
	} else {
		dividend, ok = dividend.SetString(numMap[args.Dividend])
		if !ok {
			return err
		}
	}

	if utils.IsNumeric(args.Divisor) {
		divisor, ok = divisor.SetString(args.Divisor)
		if !ok {
			return err
		}
	} else {
		divisor, ok = divisor.SetString(numMap[args.Divisor])
		if !ok {
			return err
		}
	}

	*reply = new(big.Float).Quo(dividend, divisor).String()

	return nil
}
