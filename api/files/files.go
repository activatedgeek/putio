// this module implements the file related operations
// as documented at https://api.put.io/v2/docs/files.html
package files

import (
  "strconv"
  "strings"
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

func (f *Files) Search(query string, page int, params interface{}) (*http.Response, string, []error) {
  return f.Req.Get("/search/" + query + "/page/" + strconv.Itoa(page)).Query(params).End()
}

func (f *Files) Upload() {

}

func (f *Files) CreateFolder(name string, parentId string) (*http.Response, string, []error) {
  return f.Req.Post("/create-folder").Send("name=" + name).Send("parent_id=" + parentId).End()
}

func (f *Files) Get(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id).End()
}

func (f *Files) Delete(fileIds []string) (*http.Response, string, []error) {
  return f.Req.Post("/delete").Send("file_ids=" + strings.Join(fileIds, ",")).End()
}

func (f *Files) Rename(fileId string, name string) (*http.Response, string, []error) {
  return f.Req.Post("/rename").Send("file_id=" + fileId).Send("name=" + name).End()
}

func (f *Files) Move(fileIds []string, parentId string) (*http.Response, string, []error) {
  return f.Req.Post("/move").Send("file_ids=" + strings.Join(fileIds, ",")).Send("parent_id=" + parentId).End()
}

func (f *Files) ConvertToMP4(id string) (*http.Response, string, []error) {
  return f.Req.Post("/" + id + "/mp4").End()
}

func (f *Files) GetMP4Status(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/mp4").End()
}

func (f *Files) Download() {

}

func (f *Files) Share(fileIds []string, friends []string) (*http.Response, string, []error) {
  return f.Req.Post("/share").Send("file_ids=" + strings.Join(fileIds, ",")).Send("friends=" + strings.Join(friends, ",")).End()
}

func (f *Files) Shared() (*http.Response, string, []error) {
  return f.Req.Get("/shared").End()
}

func (f *Files) SharedWith(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/shared-with").End()
}

func (f *Files) Unshare(id string, friends []string) (*http.Response, string, []error) {
  return f.Req.Post("/" + id + "/unshare").Send("friends=" + strings.Join(friends, ",")).End()
}

func (f *Files) Subtitles(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/subtitles").End()
}

func (f *Files) DefaultSubtitle(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/subtitles/default").End()
}

func (f *Files) Subtitle(id string, key string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/subtitles/" + key).End()
}

func (f *Files) DefaultHLSPlaylist(id string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/hls/media.m3u8").End()
}

func (f *Files) HLSPlaylist(id string, subtitleKey string) (*http.Response, string, []error) {
  return f.Req.Get("/" + id + "/hls/media.m3u8").Query("subtitle_key=" + subtitleKey).End()
}

func (f *Files) SetVideoPosition(id string, time int) (*http.Response, string, []error) {
  return f.Req.Post("/" + id + "/start-from").Send("time=" + strconv.Itoa(time)).End()
}

func (f *Files) DeleteVideoPosition(id string) (*http.Response, string, []error) {
  return f.Req.Post("/" + id + "/start-from/delete").End()
}
