package seeder

import (
	model "rest/model"
)

var roles = []model.Role{
	{
		Name:        "superadmin",
		DisplayName: "Super Admin",
	},
	{
		Name:        "user",
		DisplayName: "Normal User",
	},
}
