// this module implements the zip related operations
// as documented at https://api.put.io/v2/docs/zip.html
package zips

import (
  "strings"
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/zips"

type Zips struct {
  Req *commons.Request
}

func BuildNewZip(accessToken string, config *commons.Config) *Zips {
  return &Zips{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (z *Zips) Create(fileIds []string) (*http.Response, string, []error) {
  return z.Req.Post("/create").Send("file_ids=" + strings.Join(fileIds, ",")).End()
}

func (z *Zips) List() (*http.Response, string, []error) {
  return z.Req.Get("/list").End()
}

func (z *Zips) Get(zipId string) (*http.Response, string, []error) {
  return z.Req.Get("/" + zipId).End()
}
