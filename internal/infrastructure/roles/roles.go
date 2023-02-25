package roles

import "fmt"

type Role string

const (
	RoleRead  Role = "read"
	RoleWrite Role = "write"
	RoleAll   Role = "all"
)

func (role Role) IsAll() bool {
	return role == RoleAll
}

func ConvertRole(rawRole string) (*Role, error) {
	roles := map[string]Role{
		string(RoleRead):  RoleRead,
		string(RoleWrite): RoleWrite,
		string(RoleAll):   RoleAll,
	}

	env, ok := roles[rawRole]
	if !ok {
		return nil, fmt.Errorf("incorrect role")
	}

	return &env, nil
}

func GetAllRoles() []Role {
	return []Role{
		RoleRead,
		RoleWrite,
	}
}
