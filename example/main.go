package main

import (
  "os"
  "github.com/activatedgeek/putio/api"
)

// @TODO
func main() {
  config := api.NewDefaultConfig(os.Getenv("PUTIO_CLIENT_ID"), os.Getenv("PUTIO_CLIENT_SECRET"))
  api.Login(config)
}
