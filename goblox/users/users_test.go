package users

import (
	"log"
	"testing"
)

// Test_main :
func Test_main(test *testing.T) {
	user := User{}
	user.New()

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
