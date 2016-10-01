// this package is a wrapper for http communication with the api
// it wraps the request into standard operable structure
package commons

import (
  "github.com/parnurzeal/gorequest"
)

type Request struct {
  AccessToken string
  BaseUrl string
}

func BuildNewRequest(accessToken string, baseUrl string) (*Request) {
  return &Request{
    AccessToken: accessToken,
    BaseUrl: baseUrl,
  }
}

func (r *Request) Get(path string) (*gorequest.SuperAgent) {
  return gorequest.New().Set("Accept", "application/json").Get(r.BaseUrl + path).Query("oauth_token=" + r.AccessToken)
}

func (r *Request) Post(path string) (*gorequest.SuperAgent) {
  return gorequest.New().Set("Accept", "application/json").Post(r.BaseUrl + path).Query("oauth_token=" + r.AccessToken)
}
