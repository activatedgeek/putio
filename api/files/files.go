// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/files.html
package files

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/parnurzeal/gorequest"
)

const pathPrefix = "/files"

type Files struct {
  ApiEndpoint string
  Request *gorequest.SuperAgent
}

func (f *Files) List() {

}

func (f *Files) Search() {

}

func (f *Files) Upload() {

}

func (f *Files) CreateFolder() {

}

func (f *Files) Get() {

}

func (f *Files) Delete() {

}

func (f *Files) Rename() {

}

func (f *Files) Move() {

}

func (f *Files) ConvertToMP4() {

}

func (f *Files) GetMP4Status() {

}

func (f *Files) Download() {

}

func (f *Files) Share() {

}

func (f *Files) ListShared() {

}

func (f *Files) ListSharedWith() {

}

func (f *Files) Unshare() {

}

func (f *Files) ListSubtitles() {

}

func (f *Files) GetSubtitles() {

}

func (f *Files) DownloadSubtitles() {

}

func (f *Files) HLSPlaylist() {

}

func (f *Files) Events() {

}

func (f *Files) DeleteEvent() {

}

func (f *Files) SetVideoPosition() {

}

func (f *Files) DeleteVideoPosition() {

}

func BuildNewFile(accessToken string, config *commons.Config) *Files {
  return &Files{
    ApiEndpoint: config.Endpoint + pathPrefix,
    Request: commons.BuildNewRequest(accessToken),
  }
}
