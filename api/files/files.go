// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/files.html
package files

import (
  // "strconv"
  "net/http"
  "github.com/activatedgeek/putio/api/commons"
)

const pathPrefix = "/files"

type Files struct {
  Req *commons.Request
}

func BuildNewFile(accessToken string, config *commons.Config) *Files {
  return &Files{
    Req: commons.BuildNewRequest(accessToken, config.Endpoint + pathPrefix),
  }
}

func (f *Files) List(parent_id string) (*http.Response, string, []error) {
  return f.Req.Get("/list").Query("parent_id=" + parent_id).End()
}

func (f *Files) Search() {

}

func (f *Files) SearchPage() {

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
