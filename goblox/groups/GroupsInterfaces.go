package groups

import "github.com/astralservices/goblox/goblox/network"

type IUserGroups struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	Group          IGroup     `json:"group"`
	Role           IGroupRole `json:"role"`
	IsPrimaryGroup bool       `json:"isPrimaryGroup"`
}

type Group struct {
	network.NetworkRequest
}

type IGroup struct {
	ID                 int64        `json:"id"`
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	Owner              IGroupMember `json:"owner"`
	Shout              IGroupShout  `json:"shout"`
	MemberCount        int64        `json:"memberCount"`
	IsBuildersClubOnly bool         `json:"isBuildersClubOnly"`
	PublicEntryAllowed bool         `json:"publicEntryAllowed"`
	IsLocked           bool         `json:"isLocked"`
}

type IGroupMember struct {
	BuildersClubMembershipType string `json:"buildersClubMembershipType"`
	UserID                     int64  `json:"userId"`
	Username                   string `json:"username"`
	DisplayName                string `json:"displayName"`
}

type IGroupShout struct {
	Body    string       `json:"body"`
	Poster  IGroupMember `json:"poster"`
	Created string       `json:"created"`
	Updated string       `json:"updated"`
}

type IGroupRole struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rank        int64  `json:"rank"`
	MemberCount int64  `json:"memberCount"`
}
