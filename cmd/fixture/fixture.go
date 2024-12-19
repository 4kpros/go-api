package fixture

import (
	"api/common/constants"
	"api/config"
	"api/services/user/permission"
	permissionModel "api/services/user/permission/model"
	"api/services/user/role"
	roleModel "api/services/user/role/model"
	"api/services/user/user"
	userModel "api/services/user/user/model"
	"time"
)

// Load loads initial database values
func Load() (err error) {
	var roleRepo = role.NewRepository(config.DB)
	var userRepo = user.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)

	// Add role admin
	roleAdmin, _ := roleRepo.GetByName(config.Env.RoleAdmin)
	if !(roleAdmin != nil && roleAdmin.Name == config.Env.RoleAdmin) {
		roleAdmin, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.RoleAdmin,
			Feature:     constants.FeatureAdmin,
			Description: "Administrator role",
		})
	}
	// Add role default
	roleDefault, _ := roleRepo.GetByName(config.Env.RoleDefault)
	if !(roleDefault != nil && roleDefault.Name == config.Env.RoleDefault) {
		_, _ = roleRepo.Create(&roleModel.Role{
			Name:        config.Env.RoleDefault,
			Feature:     constants.FeatureDefault,
			Description: "Default role",
		})
	}

	// Add user admin
	userAdmin, _ := userRepo.GetByEmail(config.Env.UserAdminEmail)
	if !(userAdmin != nil && userAdmin.Email == config.Env.UserAdminEmail) {
		userInfoAdmin, _ := userRepo.CreateUserInfo(&userModel.UserInfo{
			Username: "Admin",
			Language: "en",
		})
		userMfaAdmin, _ := userRepo.CreateUserMfa(&userModel.UserMfa{})

		tmpActivatedAt := time.Now()
		userAdmin, _ = userRepo.Create(&userModel.User{
			Email:    config.Env.UserAdminEmail,
			Password: config.Env.UserAdminPassword,

			LoginMethod: constants.AuthLoginMethodDefault,
			IsActivated: true,
			ActivatedAt: &tmpActivatedAt,

			RoleID:     roleAdmin.ID,
			UserInfoID: userInfoAdmin.ID,
			UserMfaID:  userMfaAdmin.ID,
		})
	}

	if userAdmin == nil || roleAdmin == nil {
		return
	}

	// Add permissions for admin
	foundPermission, err := permissionRepo.GetByRoleIDTableName(roleAdmin.ID, "*")
	if err != nil && (foundPermission == nil || foundPermission.RoleID != roleAdmin.ID) {
		_, _ = permissionRepo.Create(&permissionModel.Permission{
			RoleID:    roleAdmin.ID,
			TableName: "*",
			Create:    true,
			Read:      true,
			Update:    true,
			Delete:    true,
		})
		return
	}

	return
}
