package users

import (
	"github.com/astralservices/goblox/goblox/groups"
	"github.com/astralservices/goblox/goblox/network"
)

type User struct {
	network.NetworkRequest
}

type UserGroups func() []groups.IGroup

type IUser struct {
	Description            string `json:"description"`
	Created                string `json:"created"`
	IsBanned               bool   `json:"isBanned"`
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
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
