package admin

import (
	"github.com/4kpros/go-api/services/permission"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
	"github.com/4kpros/go-api/services/user/model"
)

type AdminService struct {
	UserRepository       *user.UserRepository
	RoleRepository       *role.RoleRepository
	PermissionRepository *permission.PermissionRepository
}

func NewAdminService(
	userRepository *user.UserRepository,
	roleRepository *role.RoleRepository,
	permissionRepository *permission.PermissionRepository,
) *AdminService {
	return &AdminService{
		UserRepository:       userRepository,
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
	}
}

// Create new admin
func (service *AdminService) Create(token string, user *model.User) (result *model.User, errCode int, err error) {
	// Create admin role (unique)
	// Create admin permission (unique)
	// Create admin user (unique)
	return
}
