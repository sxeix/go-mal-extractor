package main

import "testing"

func TestCreateHttpString(t *testing.T) {
	result := createHttpString("user", "status")
	expected := "https://api.jikan.moe/v3/user/user/animelist/status"
	if result != expected {
		t.Errorf("createHttpString('user', 'status')  != %s", expected)
	}
}
