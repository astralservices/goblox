package goblox

type IUserGroups struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	Group          IGroup     `json:"group"`
	Role           IGroupRole `json:"role"`
	IsPrimaryGroup bool       `json:"isPrimaryGroup"`
}

type Group struct {
	IGroup

	iconUrl string

	roles   []IGroupRole
	members []IGroupMember

	client Client
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

type IGroupRoles struct {
	GroupID int64        `json:"groupId"`
	Roles   []IGroupRole `json:"roles"`
}

type IGroupMemberResponse struct {
	User IGroupMember `json:"user"`
	Role IGroupRole   `json:"role"`
}

type IGroupIconURL struct {
	TargetID int64  `json:"targetId"`
	State    string `json:"state"`
	ImageURL string `json:"imageUrl"`
}

type IAuditLog struct {
	Actor       IAuditLogActor `json:"actor"`
	ActionType  string         `json:"actionType"`
	Description any            `json:"description"`
	Created     string         `json:"created"`
}

type IAuditLogActor struct {
	User IGroupMember `json:"user"`
	Role IGroupRole   `json:"role"`
}

type IWallPost struct {
	ID      int64        `json:"id"`
	Body    string       `json:"body"`
	Poster  IGroupMember `json:"poster"`
	Created string       `json:"created"`
	Updated string       `json:"updated"`
}
