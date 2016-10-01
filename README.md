# Put.io API v2

This is the implementation of Put.io API in Golang.

The documentation for the OAuth2 HTTP API can be found [here](https://api.put.io/v2/docs/gettingstarted.html).

# Docs

## Usage

A basic usage of the SDK is provided below. All other calls follow similar pattern.

```go
package main

import (
  "os"
  "fmt"
  "log"
  "github.com/activatedgeek/putio/api"
  "github.com/activatedgeek/putio/api/commons"
)

func main() {
  // create a new API session
  session := api.NewSession(os.Getenv("PUTIO_ACCESS_TOKEN"), commons.NewEmptyConfig())

  // list all the files in root folder
  _, body, err := session.Files.List("0")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(body)
}
```

# Coverage

* **Files**
  - [x] List
  - [x] Search
  - [ ] Upload
  - [x] Create Folder
  - [x] Get
  - [x] Delete
  - [x] Rename
  - [x] Move
  - [x] Convert to MP4
  - [x] Get MP4 Status
  - [ ] Download
  - [x] Share
  - [x] Shared With
  - [x] Unshare
  - [x] Subtitles
  - [x] Default Subtitle
  - [x] Subtitle
  - [x] Default HLS Playlist
  - [x] HLS Playlist
  - [x] Set Video Position
  - [x] Delete Video Position

* **Events**
  - [x] List
  - [x] Delete

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
