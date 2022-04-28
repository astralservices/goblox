package client

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Exit(m.Run())
}

func Test_NewClient(t *testing.T) {
	c := New()
	if c == nil {
		t.Errorf("New() returned nil")
	}
}

func Test_NewClientWithToken(t *testing.T) {
	c := New(SetToken(os.Getenv("TOKEN")))
	if c == nil {
		t.Errorf("New() returned nil")
	}
}

func Test_NewClientWithTokenLogin(t *testing.T) {
	c := New(SetToken(os.Getenv("TOKEN")))
	authed := c.login()

	if !authed {
		t.Logf("login() returned false")
	} else {
		t.Logf("login() returned true")
	}
}
