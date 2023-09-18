package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Hash a method to create hash of response
func Hash(arr interface{}) [16]byte {
	jsonBytes, _ := json.Marshal(arr)
	return md5.Sum(jsonBytes)
}

// SetEtag this method is used to add Etag header if the response has been updated else return 304 to reduce data out
func SetEtag(context echo.Context, response interface{}) error {

	hash := fmt.Sprintf("%x", Hash(response))
	clientETag := context.Request().Header.Get("If-None-Match")
	if len(clientETag) >= 4 {
		clientETag = strings.Replace(clientETag, "W/", "", 1)
		clientETag = clientETag[1 : len(clientETag)-1]
	}
	if hash == clientETag {
		return context.NoContent(http.StatusNotModified)
	}
	context.Response().Header().Set("Etag", strconv.Quote(hash))
	return context.JSON(http.StatusOK, response)
}
