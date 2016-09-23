// this module implements the zip related operations
// as documented at https://api.put.io/v2/docs/zip.html
package friends

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/parnurzeal/gorequest"
)

const pathPrefix = "/friends"

type Friends struct {
  ApiEndpoint string
  Request *gorequest.SuperAgent
}

func (f *Friends) List() {

}

func (f *Friends) FriendRequests() {

}

func (f *Friends) FriendRequest() {

}

func (f *Friends) Approve() {

}

func (f *Friends) Deny() {

}

func (f *Friends) Unfriend() {

}

func BuildNewFriend(accessToken string, config *commons.Config) *Friends {
  return &Friends{
    ApiEndpoint: config.Endpoint + pathPrefix,
    Request: commons.BuildNewRequest(accessToken),
  }
}
