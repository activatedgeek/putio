// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/account.html
package account

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/parnurzeal/gorequest"
)

const pathPrefix = "/account"

type Account struct {
  ApiEndpoint string
  Request *gorequest.SuperAgent
}

func (a *Account) Info() {

}

func (a *Account) Settings() {

}

func (a *Account) SetSettings() {

}

func BuildNewAccount(accessToken string, config *commons.Config) *Account {
  return &Account{
    ApiEndpoint: config.Endpoint + pathPrefix,
    Request: commons.BuildNewRequest(accessToken),
  }
}
