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
  // client id and secret only needed for login (@TODO)
  // alternatively, commons.NewEmptyConfig() which requires no arguments
  config := commons.NewDefaultConfig(os.Getenv("PUTIO_CLIENT_ID"), os.Getenv("PUTIO_CLIENT_SECRET"))

  // session is required
  session := api.NewSession(os.Getenv("PUTIO_ACCESS_TOKEN"), config)

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
