package main

import (
	b64 "encoding/base64"
	"fmt"
	"net/url"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println(sDec)
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	data = uEnc

	s := fmt.Sprintf(`postgres://user:pass@host.com:5432/path?token=%s`, data)
	fmt.Println(s)

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// fmt.Println(u.RawQuery)

	url := "postgres://user:pass@host.com:5432/path"
	query := u.Query()
	fmt.Println(query)
	query.Set("token", data)
	u.RawQuery = query.Encode()
	fmt.Println(u.String())
	fmt.Println(url)

	// m, _ := url.ParseQuery(u.RawQuery)
	// fmt.Println(m["token"])
	//
	// x := m["token"]
	// tokenValue := x[0]
	// fmt.Println(x[0])
	//
	// // enc := b64.StdEncoding.EncodeToString([]byte(u.RawQuery))
	// // fmt.Println(enc)
	//
	// dnc, _ := b64.URLEncoding.DecodeString(tokenValue)
	// fmt.Println(string(dnc))

}
