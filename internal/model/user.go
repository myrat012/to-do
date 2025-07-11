package model

import "fmt"

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleOperator UserRole = "OPERATOR"
)

func GetUserRoleByString(role string) (UserRole, error) {
	switch role {
	case "ADMIN":
		return UserRoleAdmin, nil
	case "OPERATOR":
		return UserRoleOperator, nil
	}
	return "", fmt.Errorf(" method : \"GetUserRoleByString\" cannot convert value = %v to UserRole", role)
}
