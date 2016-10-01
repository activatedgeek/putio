// this module implements the zip related operations
// as documented at https://api.put.io/v2/docs/zip.html
package friends

import (
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/friends"

type Friends struct {
  Req *commons.Request
}

func BuildNewFriend(accessToken string, config *commons.Config) *Friends {
  return &Friends{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (f *Friends) List() (*http.Response, string, []error) {
  return f.Req.Get("/list").End()
}

func (f *Friends) ListRequests() (*http.Response, string, []error) {
  return f.Req.Get("/waiting-requests").End()
}

func (f *Friends) Request(username string) (*http.Response, string, []error) {
  return f.Req.Post(username + "/request").End()
}

func (f *Friends) Approve(username string) (*http.Response, string, []error) {
  return f.Req.Post(username + "/approve").End()
}

func (f *Friends) Deny(username string) (*http.Response, string, []error) {
  return f.Req.Post(username + "/deny").End()
}

func (f *Friends) Unfriend(username string) (*http.Response, string, []error) {
  return f.Req.Post(username + "/unfriend").End()
}
