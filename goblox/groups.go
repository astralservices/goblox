package goblox

import (
	"encoding/json"
	"log"
	"strconv"
)

type GroupsHandler struct {
	fetchById func(id int64) (group *Group, err error)
}

type Groups struct {
	client *Client
}

func (ref *GroupsHandler) New(client *Client) *GroupsHandler {
	g := &GroupsHandler{}

	g.fetchById = func(id int64) (user *Group, err error) {
		ref := &Groups{
			client: client,
		}
		return ref.GetGroupById(int64(id))
	}

	return g
}

func (ref *Groups) GetGroupById(groupId int64) (*Group, error) {
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

	g := group.New(&r, &ref.client.http, ref.client)

	return g, err
}
