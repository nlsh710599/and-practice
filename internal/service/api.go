package service

import (
	"errors"
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
	return nil
}

func (ctrl *BigNumberComputationService) Subtract(r *http.Request, args *SubtractRequest, reply *string) error {
	return nil
}

func (ctrl *BigNumberComputationService) Multiply(r *http.Request, args *MultiplyRequest, reply *string) error {
	return nil
}

func (ctrl *BigNumberComputationService) Divide(r *http.Request, args *DivideRequest, reply *string) error {
	return nil
}
