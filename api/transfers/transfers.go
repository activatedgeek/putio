// this module implements the transfers related operations
// as documented at https://api.put.io/v2/docs/transfers.html
package transfers

import (
  "strings"
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/transfers"

type Transfers struct {
  Req *commons.Request
}

func BuildNewTransfer(accessToken string, config *commons.Config) *Transfers {
  return &Transfers{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (t *Transfers) List() (*http.Response, string, []error) {
  return t.Req.Get("/list").End()
}

func (t *Transfers) Add(url string, saveParentId string) (*http.Response, string, []error) {
  return t.Req.Post("/add").Send("url=" + url).Send("save_parent_id=" + saveParentId).End()
}

func (t *Transfers) Get(id string) (*http.Response, string, []error) {
  return t.Req.Get("/" + id).End()
}

func (t *Transfers) Retry(id string) (*http.Response, string, []error) {
  return t.Req.Post("/retry").Send("id=" + id).End()
}

func (t *Transfers) Cancel(transferIds []string) (*http.Response, string, []error) {
  return t.Req.Post("/cancel").Send("transfer_ids=" + strings.Join(transferIds, ",")).End()
}

func (t *Transfers) Clean() (*http.Response, string, []error) {
  return t.Req.Post("/clean").End()
}
