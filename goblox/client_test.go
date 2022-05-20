package goblox

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Exit(m.Run())
}

func Test_NewClient(t *testing.T) {
	c := New()

	if c.Token == "" {
		t.Logf("token is empty")
	} else {
		t.Logf("token is not empty")
	}
}

func Test_NewClientWithToken(t *testing.T) {
	c := New(SetToken(os.Getenv("TOKEN")))

	if c.Token == "" {
		t.Logf("Token is empty")
	}
}

func Test_NewClientWithTokenLogin(t *testing.T) {
	c := New()
	
	authed := c.Login(os.Getenv("TOKEN"))

	if !authed {
		t.Logf("login() returned false")
	} else {
		t.Logf("login() returned true")
	}
}

func Test_FetchUser(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	user, err := c.Users.GetUserByUsername("AmusedCrepe")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}

func Test_GetUsernameHistory(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	user, err := c.Users.GetUserByUsername("AmusedCrepe")

	if err != nil {
		log.Fatal(err)
	}

	history, err := user.GetUsernameHistory()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(history)
}

func Test_GetUserGroups(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	user, err := c.Users.GetUserByUsername("AmusedCrepe")

	if err != nil {
		log.Fatal(err)
	}

	groups, err := user.GetGroups()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(groups)
}

func Test_GetGroupRoles(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	group, err := c.Groups.GetGroupById(13619939)

	if err != nil {
		log.Println(group)
		log.Fatal(err)
	}

	roles, err := group.GetRoles()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(roles)
}

func Test_GetGroupMembers(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	group, err := c.Groups.GetGroupById(13619939)

	if err != nil {
		log.Fatal(err)
	}

	members, err := group.GetMembers()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(members)
}

func Test_GetGroupIconURL(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	group, err := c.Groups.GetGroupById(13619939)

	if err != nil {
		log.Fatal(err)
	}

	iconURL, err := group.GetIconURL()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(iconURL)
}

func Test_GetGroupAuditLog(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(SetToken(token))

	group, err := c.Groups.GetGroupById(13619939)

	if err != nil {
		log.Fatal(err)
	}

	data, err := group.GetAuditLog()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(data)
}
