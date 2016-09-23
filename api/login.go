// this package starts a temporary callback HTTP server to handle callbacks
// from the OAuth permissions from Put.io
package api

import (
  "log"
  "os/exec"
  "runtime"
  "github.com/activatedgeek/putio/api/commons"
)

// api authentication path
const authenticationPath = "/oauth2/authenticate"

// construct url as prescribed in https://api.put.io/v2/docs/gettingstarted.html#sign-up
func getTargetAuthUrl(config *commons.Config) string {
  targetAuthURL := config.Endpoint + authenticationPath + "?client_id=" +
    config.ClientId + "&response_type=code&redirect_uri=" +
    config.RedirectUri

  return targetAuthURL
}

func Login(config *commons.Config) {
  // string for BSD open equivalent on different OSes
  var openCommandString string
  switch runtime.GOOS {
    case "darwin":
      openCommandString = "open"
    case "linux":
      openCommandString = "xdg-open"
    default:
      // @TODO error handling, throw error or panic TBD
      panic("Unsupported system " + runtime.GOOS)
  }

  // open Put.io OAuth permissions page
  cmd := exec.Command(openCommandString, getTargetAuthUrl(config))
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}
