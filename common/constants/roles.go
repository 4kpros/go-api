package constants

import "slices"

const ROLE_SUPER_ADMIN = "superadmin"
const ROLE_ADMIN = "admin"
const ROLE_MANAGER = "manager"
const ROLE_DEFAULT = "default"

var AllRoles = []string{
	ROLE_SUPER_ADMIN,
	ROLE_ADMIN,
	ROLE_MANAGER,
	ROLE_DEFAULT,
}

func IsRoleValid(role string) bool {
	return slices.Contains(AllRoles, role)
}
