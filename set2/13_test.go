package matasano

import "testing"

func TestProfileFor(t *testing.T) {

	profile := ProfileFor("john@gmail.com")
	expected := "email=john@gmail.com&id=10&role=user"

	if profile != expected {
		t.Log(profile)
		t.Log(expected)
		t.Fatal("Did not encode correctly")
	}

}

func TestEncryptProfile(t *testing.T) {
	c := EncryptProfile("john@gmail.com")
	p := DecryptProfile(c)

	expected := "email=john@gmail.com&id=10&role=user"

	if p != expected {
		t.Log(p)
		t.Log(expected)
		t.Fatal("Did not encode correctly")
	}
}

func TestCreateAdminProfile(t *testing.T) {
	expected := "email=jimi@gmail.com&id=10&role=admin"
	admin := CreateAdminProfile()
	if admin != expected {
		t.Fatal("Could not create admin profile")
	}
}
