// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/files.html#events
package events

import (
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/events"

type Events struct {
  Req *commons.Request
}

func BuildNewEvent(accessToken string, config *commons.Config) *Events {
  return &Events{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (e *Events) List() (*http.Response, string, []error) {
  return e.Req.Get("/list").End()
}

func (e *Events) Delete() (*http.Response, string, []error) {
  return e.Req.Post("/delete").End()
}
