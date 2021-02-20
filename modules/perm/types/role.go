package types

import (
	"fmt"
)

// Auth return the auth of the role
func (r Role) Auth() Auth {
	return 1 << r
}

func GetRolesFromStr(strRoles ...string) (roles []Role, err error) {
	for _, strRole := range strRoles {
		role, err := RoleFromstring(strRole)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// RoleFromstring turns a string into a Auth
func RoleFromstring(str string) (Role, error) {
	option, ok := Role_value[str]
	if !ok {
		return Role(0xff), fmt.Errorf("'%s' is not a valid vote option", str)
	}
	return Role(option), nil
}

// ValidRole returns true if the role is valid and false otherwise.
func ValidRole(role Role) bool {
	if role == RoleRootAdmin ||
		role == RolePermAdmin ||
		role == RoleBlacklistAdmin ||
		role == RoleNodeAdmin ||
		role == RoleParamAdmin ||
		role == RoleIDAdmin ||
		role == RoleBaseM1Admin ||
		role == RolePowerUser ||
		role == RoleRelayerUser {
		return true
	}
	return false
}

// Marshal needed for protobuf compatibility
func (r Role) Marshal() ([]byte, error) {
	return []byte{byte(r)}, nil
}

// Unmarshal needed for protobuf compatibility
func (r *Role) Unmarshal(data []byte) error {
	*r = Role(data[0])
	return nil
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (r Role) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(r.String()))
	default:
		s.Write([]byte(fmt.Sprintf("%v", byte(r))))
	}
}
