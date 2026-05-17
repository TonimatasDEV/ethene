package auth

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://sessionserver.mojang.com/session/minecraft/hasJoined"

type Auth struct {
	UUID string `json:"id"`
	Name string `json:"name"`
	Prop []Prop `json:"properties"`
}

type Prop struct {
	Name string  `json:"name"`
	Data string  `json:"value"`
	Sign *string `json:"signature"`
}

func Authenticate(secret []byte, name string) *Auth {
	return execute(generateAuthURL(name, generateAuthSHA(secret)))
}

func execute(url string) *Auth {
	out, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer out.Body.Close()

	bdy, err := io.ReadAll(out.Body)
	if err != nil {
		return nil
	}

	var auth Auth

	err = json.Unmarshal(bdy, &auth)
	return &auth
}

func generateAuthURL(name, hash string) string {
	return fmt.Sprintf("%s?username=%s&serverId=%s", url, name, hash)
}

func generateAuthSHA(secret []byte) string {
	sha := sha1.New()

	_, public := NewCrypt()

	sha.Write(secret)
	sha.Write(public)

	hash := sha.Sum(nil)

	negative := (hash[0] & 0x80) == 0x80

	if negative {
		carry := true

		for i := len(hash) - 1; i >= 0; i-- {
			hash[i] = ^hash[i]
			if carry {
				carry = hash[i] == 0xff
				hash[i]++
			}
		}
	}

	res := strings.TrimLeft(fmt.Sprintf("%x", hash), "0")
	if negative {
		res = "-" + res
	}

	return res
}
