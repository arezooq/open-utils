package security

import (
	"github.com/arezooq/open-utils/errors"
)

type RoleClientInterface interface {
	GetRolesByIDs([]string) ([]Role, error)
}

type Role struct {
	ID   string
	Name string
}

func ValidateRoleIDs(roleClient RoleClientInterface, roleIds []string) error {
	if len(roleIds) == 0 {
		return nil
	}
	roles, err := roleClient.GetRolesByIDs(roleIds)
	if err != nil {
		return errors.ErrInternal
	}

	if len(roles) != len(roleIds) {
		return errors.ErrValidation
	}
	return nil
}
