package matasano

import "strings"

type Values struct {
	params []string
}

func ParseQueryValues(input string) Values {
	v := Values{}

	for _, param := range strings.Split(input, "&") {
		data := strings.Split(param, "=")

		key := data[0]
		value := data[1]
		v.Add(key, value)
	}

	return v
}

func (v *Values) Add(key, value string) {
	if v.params == nil {
		v.params = make([]string, 0)
	}
	v.params = append(v.params, key, value)
}

func (v Values) Encode() string {
	out := ""
	size := len(v.params)
	for i := 0; i < size; i += 2 {
		key := v.params[i]
		value := v.params[i+1]
		key = strings.Replace(key, "&", "", -1)
		key = strings.Replace(key, "=", "", -1)
		value = strings.Replace(value, "&", "", -1)
		value = strings.Replace(value, "=", "", -1)

		out += key + "=" + value + "&"
	}
	return out[:len(out)-1]
}

func ProfileFor(email string) string {
	v := Values{}
	v.Add("email", email)
	v.Add("id", "10")
	v.Add("role", "user")
	return v.Encode()
}

func EncryptProfile(email string) []byte {
	key := []byte("criminologically")
	profile := ProfileFor(email)
	p := PadBytes([]byte(profile), 16)
	c, _ := EncryptAESInECBMode(key, p)
	return c
}

func DecryptProfile(profile []byte) string {
	key := []byte("criminologically")
	p, _ := DecryptAESInECBMode(key, profile)
	p = UnpadBytes(p, 16)
	values := ParseQueryValues(string(p))
	return values.Encode()
}

func CreateAdminProfile() string {
	p1 := EncryptProfile("jimi@gmail.com")
	p2 := EncryptProfile(string([]byte{
		'_', '_', '_', '_', '_', '_', '_', '_', '_', '_',
		'a', 'd', 'm', 'i', 'n', 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	}))

	copy(p1[32:], p2[16:32])

	return DecryptProfile(p1)
}
