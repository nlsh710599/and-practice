package method

import (
	"errors"
	"log"
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

func (ctrl *Controller) Create(args *[]string, reply *int) error {
	name := (*args)[0]
	value := (*args)[1]
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

func (ctrl *Controller) Update(args *[]string, reply *int) error {
	name := (*args)[0]
	value := (*args)[1]
	if !utils.ValidName(name) {
		return ErrInvalidName
	}
	if !utils.IsNumeric(value) {
		return ErrInvalidValue
	}

	if err := ctrl.RDS.UpdateNumber(name, value); err != nil {
		if isDuplicateKey(err) {
			return ErrCreateDuplicateNumber
		}
		return err
	}
	return nil
}

func (ctrl *Controller) Delete(args *[]string, reply *int) error {
	log.Println(args)
	return nil
}
