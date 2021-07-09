package main

import "testing"

func TestCreateHttpString(t *testing.T) {
	result := createHttpString("user", "status")
	expected := "https://api.jikan.moe/v3/user/user/animelist/status"
	if result != expected {
		t.Errorf("createHttpString('user', 'status')  != %s", expected)
	}
}

func TestStartMessage(t *testing.T) {
	user := "user"
	result := startMessage(user)
	expected := "Searching for user user"
	if result != expected {
		t.Errorf("startMessage('user')  != %s", expected)
	}
}
