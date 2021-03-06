package goblox

import (
	"encoding/json"
	"strconv"
)

// Initialize a new group with prefetched data.
func (ref *Group) New(data *IGroup, client *Client) *Group {
	group := &Group{
		IGroup: *data,
		client: client,
	}

	return group
}

// Gets all the roles in the group.
func (ref *Group) GetRoles() ([]IGroupRole, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/groups/"+strconv.Itoa(int(ref.ID))+"/roles", map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	var r IGroupRoles
	err = json.Unmarshal([]byte(data), &r)

	ref.roles = r.Roles

	return ref.roles, err
}

// Gets all the members in the group.
func (ref *Group) GetMembers() ([]IGroupMember, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/groups/"+strconv.Itoa(int(ref.ID))+"/users?limit=100", map[string]interface{}{})

	if err != nil {
		return []IGroupMember{}, err
	}

	var r IPagedResponse[IGroupMemberResponse]
	err = json.Unmarshal([]byte(data), &r)

	for _, v := range r.Data {
		ref.members = append(ref.members, v.User)
	}

	return ref.members, err
}

// Gets the Icon URL for the group.
func (ref *Group) GetIconURL() (string, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://thumbnails.roblox.com/v1/groups/icons?groupIds="+strconv.Itoa(int(ref.ID))+"&size=420x420&format=png&isCircular=false", map[string]interface{}{})

	if err != nil {
		return "", err
	}

	var r DatumedResponse[IGroupIconURL]
	err = json.Unmarshal([]byte(data), &r)

	ref.iconUrl = r.Data[0].ImageURL

	return ref.iconUrl, err
}

// Gets the Audit Log for the group.
//
// Pagination is coming soon.
func (ref *Group) GetAuditLog() (IPagedResponse[IAuditLog], error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/groups/"+strconv.Itoa(int(ref.ID))+"/audit-log?limit=100", map[string]interface{}{})

	if err != nil {
		return IPagedResponse[IAuditLog]{}, err
	}

	var r IPagedResponse[IAuditLog]
	err = json.Unmarshal([]byte(data), &r)

	return r, err
}

// Gets the group wall for the group.
//
// Pagination is coming soon.
func (ref *Group) GetGroupWall() (IPagedResponse[IWallPost], error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/groups/"+strconv.Itoa(int(ref.ID))+"/wall?limit=100", map[string]interface{}{})

	if err != nil {
		return IPagedResponse[IWallPost]{}, err
	}

	var r IPagedResponse[IWallPost]
	err = json.Unmarshal([]byte(data), &r)

	return r, err
}
