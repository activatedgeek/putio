// this module implements the zip related operations
// as documented at https://api.put.io/v2/docs/zip.html
package zip

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/parnurzeal/gorequest"
)

const pathPrefix = "/zip"

type Zip struct {
  ApiEndpoint string
  Request *gorequest.SuperAgent
}

func (z *Zip) Create() {

}

func (z *Zip) List() {

}

func (z *Zip) Get() {

}

func BuildNewZip(accessToken string, config *commons.Config) *Zip {
  return &Zip{
    ApiEndpoint: config.Endpoint + pathPrefix,
    Request: commons.BuildNewRequest(accessToken),
  }
}
