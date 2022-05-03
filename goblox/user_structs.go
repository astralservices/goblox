package goblox

type UserPopulate struct {
	groups bool
}

type UserParams struct {
	id       int
	populate *UserPopulate
}

type IUser struct {
	Description            string `json:"description"`
	Created                string `json:"created"`
	IsBanned               bool   `json:"isBanned"`
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`

	groups *[]IGroup
}

type IUserByUsername struct {
	Data []IUserByUsernameDatum `json:"data"`
}

type IUserByUsernameDatum struct {
	RequestedUsername string `json:"requestedUsername"`
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
}

type IUsernameHistory struct {
	Name string `json:"name"`
}
