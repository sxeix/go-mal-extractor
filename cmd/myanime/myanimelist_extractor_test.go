package main

import "testing"

func TestCreateHttpString(t *testing.T) {
	result := createHttpString("user", "status")
	expected := "https://api.jikan.moe/v3/user/user/animelist/status"
	if result != expected {
		t.Errorf("createHttpString('user', 'status')  != %s", expected)
	}
}

func TestGetFlags(t *testing.T) {
	user, status, score, sort := getFlags()
	expectUser := "DEFAULT"
	expectStatus := "all"
	expectScore := false
	expectSort := false
	if expectUser != user {
		t.Errorf("getFlags() user != %s", expectUser)
	}
	if expectStatus != status {
		t.Errorf("getFlags() status != %s", expectStatus)
	}
	if expectScore != score {
		t.Errorf("getFlags() score != %t", expectScore)
	}
	if expectSort != sort {
		t.Errorf("getFlags() sort != %t", expectSort)
	}

}
