# Put.io API v2

This is the implementation of Put.io API in Golang.

The documentation for the OAuth2 HTTP API can be found [here](https://api.put.io/v2/docs/gettingstarted.html).

# Docs

## Usage

A basic usage of the SDK is provided below. All other calls follow similar pattern.

```
package main

import (
  "os"
  "fmt"
  "log"
  "github.com/activatedgeek/putio/api"
  "github.com/activatedgeek/putio/api/commons"
)

func main() {
  config := commons.NewDefaultConfig(os.Getenv("PUTIO_CLIENT_ID"), os.Getenv("PUTIO_CLIENT_SECRET"))
  session := api.NewSession(os.Getenv("PUTIO_ACCESS_TOKEN"), config)

  // list all the files in root folder
  _, body, err := session.Files.List("0")
  if err != nil {
    log.Fatal(err)
  }

  // print the response JSON
  fmt.Println(body)
}
```

# Coverage

* **Files**
  - [x] List
  - [ ] Search
  - [ ] Upload
  - [ ] Create Folder
  - [ ] Get
  - [ ] Delete
  - [ ] Rename
  - [ ] Move
  - [ ] Convert to MP4
  - [ ] Get MP4 Status
  - [ ] Download
  - [ ] Sharing
  - [ ] Subtitles
  - [ ] Download Subtitle
  - [ ] HLS Playlist
  - [ ] Events
  - [ ] Set Video Position
  - [ ] Get Video Position

* **Transfers**
  - [x] List
  - [x] Add
  - [x] Get
  - [x] Retry
  - [x] Cancel
  - [x] Clean

* **Zip**
  - [x] Create
  - [x] List
  - [x] Get

* **Friends**
  - [x] List
  - [x] List Requests
  - [x] Request
  - [x] Approve
  - [x] Deny
  - [x] Unfriend

* **Account**
  - [x] Info
  - [x] Settings
  - [x] Set Settings

# Development

**NOTE**: `glide` is used for dependency management

Keep this package in your `GOPATH` and run `glide install`.
