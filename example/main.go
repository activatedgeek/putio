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

  _, body, err := session.Files.List("0")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(body)
}
