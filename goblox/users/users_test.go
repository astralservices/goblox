package users

import (
	"log"
	"testing"
)

// Test_main :
func Test_main(test *testing.T) {
	u, err := user.GetUserById(59692622)

	if err != nil {
		test.Errorf(err.Error())
	}

	log.Println(u)
}

func Test_ByUsername(t *testing.T) {
	user := User{}
	user.New()

	u, err := user.GetUserByUsername("AmusedCrepe")

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Println(u)
}

func Test_GetGroupsForUser(t *testing.T) {
	user := User{}
	user.New()

	groups, err := user.GetGroupsForUser(59692622)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Println(groups)
}

func Test_GetGroupsForUserFromUserInterface(t *testing.T) {
	user := User{}
	user.New()

	u, err := user.GetUserById(59692622)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Println(u.groups)
}

func Test_PopulateUser(t *testing.T) {
	user := User{}
	user.New()

	u, err := user.GetUserAndPopulate(UserParams{
		id: 59692622,
		populate: &UserPopulate{
			groups: true,
		},
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Println(u)
}
