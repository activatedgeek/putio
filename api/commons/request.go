// this package is a wrapper for http communication with the api
// it wraps the request into standard operable structure
package commons

import (
  "github.com/parnurzeal/gorequest"
)

func BuildNewRequest(accessToken string) (*gorequest.SuperAgent) {
  return gorequest.New().Set("Accept", "application/json")
}
