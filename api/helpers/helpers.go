package helpers

import (
	"os"
	"strings"
)

/*
The Go program can enter into an infinite loop if a user submits localhost:3000 as the URL to be shortened (3:59 - 4:11).

Since the service itself runs on localhost:3000, the application would attempt to resolve its own domain,
resulting in a recursive cycle. To prevent this, the creator implements a helper function called removeDomainError
(4:17 - 4:24) within the helpers package, which checks the input against the environment's DOMAIN variable (8:16 - 8:39)
and strips prefixes like http, https, and www to ensure the user isn't trying to abuse the system (9:08 - 10:39).
*/

// both created in shorten.go file
func EnforceHTTP(url string) string {
    url = strings.TrimSpace(url)
    if url == "" {
        return ""
    }

    // Already has http or https protocol
    if strings.HasPrefix(strings.ToLower(url), "http://") ||
       strings.HasPrefix(strings.ToLower(url), "https://") {
        return url
    }

    // Add http:// by default
    return "http://" + url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN"){
		return false
	}

	// if the user enter like https://localhost:3000 or www.localhost:3000 to prevent this we are implementing this
	newURL := strings.Replace(url,"http://","",1)
	newURL = strings.Replace(newURL,"https://","",1)
	newURL = strings.Replace(newURL,"www.","",1)
	newURL = strings.Replace(newURL,"http://","",1)
	newURL= strings.Split(newURL,"/")[0]

	if newURL == os.Getenv("DOMAIN"){
		return false
	}
	return true
}
