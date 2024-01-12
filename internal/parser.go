package internal

import (
	"log"
	"net/url"
)

/*
ParseURL - Parse a valid URL else return an error
A valid URL should follow the below
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
*/
func ParseURL(URL string) error {
	_, parseErr := url.Parse(URL)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	return nil
}
