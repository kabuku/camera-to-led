package client

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
)

func CreateDigestAuthHeader(res runtime.ClientResponse, method, path, username, password string) string {

	authorization := createWwwAuthenticateMap(res)

	realmHeader := authorization["realm"]
	qopHeader := authorization["qop"]
	nonceHeader := authorization["nonce"]
	opaqueHeader := authorization["opaque"]
	realm := realmHeader
	// A1
	h := md5.New()
	A1 := fmt.Sprintf("%s:%s:%s", username, realm, password)
	io.WriteString(h, A1)
	HA1 := fmt.Sprintf("%x", h.Sum(nil))

	// A2
	h = md5.New()
	A2 := fmt.Sprintf("%s:%s", method, path)
	io.WriteString(h, A2)
	HA2 := fmt.Sprintf("%x", h.Sum(nil))

	// response
	cnonce := RandomKey()
	response := H(strings.Join([]string{HA1, nonceHeader, "00000001", cnonce, qopHeader, HA2}, ":"))

	// now make header
	AuthHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=00000001, qop=%s, response="%s", opaque="%s", algorithm=MD5`,
		username, realmHeader, nonceHeader, path, cnonce, qopHeader, response, opaqueHeader)
	return AuthHeader
}

func RandomKey() string {
	k := make([]byte, 12)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}
		bytes += n
	}
	return base64.StdEncoding.EncodeToString(k)
}

/*
 H function for MD5 algorithm (returns a lower-case hex MD5 digest)
*/
func H(data string) string {
	digest := md5.New()
	digest.Write([]byte(data))
	return fmt.Sprintf("%x", digest.Sum(nil))
}

func createWwwAuthenticateMap(res runtime.ClientResponse) map[string]string {
	s := strings.SplitN(res.GetHeader("Www-Authenticate"), " ", 2)
	if len(s) != 2 || s[0] != "Digest" {
		return nil
	}

	result := map[string]string{}
	for _, kv := range strings.Split(s[1], ",") {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			continue
		}
		result[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}
	return result
}
