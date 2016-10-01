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
