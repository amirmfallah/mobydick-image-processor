package lib

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/jwt"
)

func GetHeaderValue(headers http.Header, headerKey string) (string, error) {
	headerValue, exists := headers[headerKey]
	if !exists {
		err := errors.New("Couldn't find " + headerKey + " in request headers.")
		fmt.Print("GetHeaderValue-")
		return "", err
	}
	if len(headerValue) == 0 {
		err := errors.New("Empty value " + headerKey + " in request headers.")
		fmt.Print("GetHeaderValue-")
		return "", err
	}
	return headerValue[0], nil
}

func GetBearerUser(headers http.Header) (*BearerUser, error) {
	bearerToken, err := GetHeaderValue(headers, "Authorization")
	if err != nil {
		fmt.Print("GetBearerUser-")
		return nil, err
	}

	token := strings.Split(bearerToken, " ")
	if !(len(token) > 0 && len(token[1]) > 0) {
		fmt.Print("Split-")
		return nil, err
	}

	var sharedKey = []byte(JWT_SECRET)
	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token[1]))
	if err != nil {
		fmt.Print("jwtVerify-")
		return nil, err
	}

	var claims BearerUser
	err = verifiedToken.Claims(&claims)
	if err != nil {
		fmt.Print("verifiedToken.Claims-")
		return nil, err
	}

	return &claims, nil
}
