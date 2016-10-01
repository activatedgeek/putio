// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/account.html
package account

import (
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/account"

type Account struct {
  Req *commons.Request
}

func BuildNewAccount(accessToken string, config *commons.Config) *Account {
  return &Account{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (a *Account) Info() (*http.Response, string, []error) {
  return a.Req.Get("/info").End()
}

func (a *Account) Settings() (*http.Response, string, []error) {
  return a.Req.Get("/settings").End()
}

func (a *Account) SetSettings(newSettings interface{}) (*http.Response, string, []error) {
  return a.Req.Post("/settings").Send(newSettings).End()
}
