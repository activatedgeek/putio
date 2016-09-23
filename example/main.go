package main

import (
  "os"
  "fmt"
  "github.com/activatedgeek/putio/api"
  "github.com/activatedgeek/putio/api/commons"
)

// @TODO
func main() {
  config := commons.NewDefaultConfig(os.Getenv("PUTIO_CLIENT_ID"), os.Getenv("PUTIO_CLIENT_SECRET"))
  session := api.NewSession(os.Getenv("PUTIO_ACCESS_TOKEN"), config)
  fmt.Println(session)
}
