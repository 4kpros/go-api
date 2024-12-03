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

	// Create roles
	roleRepo.Create(&roleModel.Role{
		Name:        config.Env.RoleAdmin,
		Description: "Administrator role",
	})
	roleRepo.Create(&roleModel.Role{
		Name:        config.Env.RoleDefault,
		Description: "Default role",
	})
	roleAdmin, _ := roleRepo.GetByName(config.Env.RoleAdmin)

	// Find admin user
	foundAdmin, _ := userRepo.GetByEmail(config.Env.UserAdminEmail)
	if foundAdmin != nil && foundAdmin.Email == config.Env.UserAdminEmail {
		return
	}

	// Create new admin
	userInfo, _ := userRepo.CreateUserInfo(&userModel.UserInfo{
		Username: "Admin",
		Language: "en",
	})
	userMfa, _ := userRepo.CreateUserMfa(&userModel.UserMfa{})
	tmpActivatedAt := time.Now()
	_, err = userRepo.Create(&userModel.User{
		Email:    config.Env.UserAdminEmail,
		Password: config.Env.UserAdminPassword,
		RoleID:   roleAdmin.ID,

		LoginMethod: constants.AuthLoginMethodDefault,
		IsActivated: true,
		ActivatedAt: &tmpActivatedAt,

		UserInfoID: userInfo.ID,
		UserMfaID:  userMfa.ID,
	})

	// Add admin permissions
	permissionRepo.UpdatePermissionFeature(
		roleAdmin.ID,
		constants.FeatureAdmin,
	)
	permissionRepo.UpdatePermissionTable(
		roleAdmin.ID,
		"*",
		&permissionModel.PermissionTable{
			Create: true,
			Read:   true,
			Update: true,
			Delete: true,
		},
	)

	return
}
