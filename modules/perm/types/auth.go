package types

const (
	AuthDefault = Auth(0)
)

type Auth int32

func (a Auth) Roles() (rs []Role) {
	if a.Access(RoleRootAdmin.Auth()) {
		rs = append(rs, RoleRootAdmin)
	}
	if a.Access(RolePermAdmin.Auth()) {
		rs = append(rs, RolePermAdmin)
	}
	if a.Access(RoleBlacklistAdmin.Auth()) {
		rs = append(rs, RoleBlacklistAdmin)
	}
	if a.Access(RoleNodeAdmin.Auth()) {
		rs = append(rs, RoleNodeAdmin)
	}
	if a.Access(RoleParamAdmin.Auth()) {
		rs = append(rs, RoleParamAdmin)
	}
	if a.Access(RoleIDAdmin.Auth()) {
		rs = append(rs, RoleIDAdmin)
	}
	if a.Access(RoleBaseM1Admin.Auth()) {
		rs = append(rs, RoleBaseM1Admin)
	}
	if a.Access(RolePowerUser.Auth()) {
		rs = append(rs, RolePowerUser)
	}
	if a.Access(RoleRelayerUser.Auth()) {
		rs = append(rs, RoleRelayerUser)
	}

	return rs
}

func (a Auth) Access(auth Auth) bool {
	return (a & auth) > 0
}
