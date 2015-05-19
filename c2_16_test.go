package matasano

import "testing"

func TestCookData(t *testing.T) {

	c := CookData(";admin=true;")
	m := UncookData(c)

	if _, ok := m["admin"]; ok == true {
		t.Error("Allowed admin to be created")
	}

	if m["userdata"] != "%3Badmin%3Dtrue%3B" {
		t.Log(m["userdata"])
		t.Error("Did not escape values correctly")
	}

	if m["comment1"] != "cooking%20MCs" {
		t.Log(m["comment1"])
		t.Error("Did not set comment1 correctly")
	}

	if m["comment2"] != "%20like%20a%20pound%20of%20bacon" {
		t.Log(m["comment2"])
		t.Error("Did not set comment2 correctly")
	}

}

func TestFlipTheBits(t *testing.T) {
	c := FlipTheBits()
	m := UncookData(c)

	if m["admin"] != "true" {
		t.Error("Did not set admin correctly")
	}

	if m["comment1"] != "cooking%20MCs" {
		t.Log(m["comment1"])
		t.Error("Did not set comment1 correctly")
	}

	if m["comment2"] != "%20like%20a%20pound%20of%20bacon" {
		t.Log(m["comment2"])
		t.Error("Did not set comment2 correctly")
	}
}
