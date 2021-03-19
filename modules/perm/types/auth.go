package types

const (
	AuthDefault = Auth(0)
)

type Auth int32

func (a Auth) Roles() (rs *RoleSet) {
	rs = &RoleSet{
		[]Role{},
	}
	if a.Access(RoleRootAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RoleRootAdmin)
	}
	if a.Access(RolePermAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RolePermAdmin)
	}
	if a.Access(RoleBlacklistAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RoleBlacklistAdmin)
	}
	if a.Access(RoleNodeAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RoleNodeAdmin)
	}
	if a.Access(RoleParamAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RoleParamAdmin)
	}
	if a.Access(RoleIDAdmin.Auth()) {
		rs.Roles = append(rs.Roles, RoleIDAdmin)
	}
	if a.Access(RoleBaseM1Admin.Auth()) {
		rs.Roles = append(rs.Roles, RoleBaseM1Admin)
	}
	if a.Access(RolePowerUser.Auth()) {
		rs.Roles = append(rs.Roles, RolePowerUser)
	}
	if a.Access(RoleRelayerUser.Auth()) {
		rs.Roles = append(rs.Roles, RoleRelayerUser)
	}

	return rs
}

func (a Auth) Access(auth Auth) bool {
	return (a & auth) > 0
}
