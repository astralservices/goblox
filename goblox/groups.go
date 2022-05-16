package goblox

import (
	"encoding/json"
	"log"
	"strconv"
)

type GroupsHandler struct {
	client *Client
}

// Creates a new group handler with the given client.
//
// A group handler is used to fetch groups by ID.
func NewGroupsHandler(client Client) *GroupsHandler {
	return &GroupsHandler{
		client: &client,
	}
}

// Gets a group by ID.
func (ref *GroupsHandler) GetGroupById(groupId int64) (*Group, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	log.Println("sending request")
	read, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/groups/"+strconv.Itoa(int(groupId)), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IGroup
	err = json.Unmarshal([]byte(read), &r)

	group := Group{
		IGroup: r,
	}

	g := group.New(&r, ref.client)

	return g, err
}
