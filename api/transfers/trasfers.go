// this module implements the transfers related operations
// as documented at https://api.put.io/v2/docs/transfers.html
package transfers

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/parnurzeal/gorequest"
)

const pathPrefix = "/transfers"

type Transfers struct {
  ApiEndpoint string
  Request *gorequest.SuperAgent
}

func (t *Transfers) List() {

}

func (t *Transfers) Add() {

}

func (t *Transfers) Get() {

}

func (t *Transfers) Retry() {

}

func (t *Transfers) Cancel() {

}

func (t *Transfers) Clean() {

}

func BuildNewTransfer(accessToken string, config *commons.Config) *Transfers {
  return &Transfers{
    ApiEndpoint: config.Endpoint + pathPrefix,
    Request: commons.BuildNewRequest(accessToken),
  }
}
