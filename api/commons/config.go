// this module contains the configuration needed to
// operate with the Put.io API
package commons

import (
  "strconv"
)

// local callback server related constants
const defaultLocalHost = "http://localhost"
const defaultLocalPort = 12725
const defaultLocalPath = "/auth/callback"

// Put.io related constants
const apiVersion = "v2"
const apiEndpoint = "https://api.put.io"

type Config struct {
  // version of the API to be used (cannot be changed)
  ApiVersion string

  // base url endpoint for Put.io API
  Endpoint string

  // client id for the application
  ClientId string

  // client secret for the application
  ClientSecret string

  // local server base uri (used for auth callbacks)
  RedirectUri string

  // local server port
  Port int
}

// get empty config, use when access token already available
func NewEmptyConfig() *Config {
  return &Config{
    ApiVersion: apiVersion,
    Endpoint: apiEndpoint + "/" + apiVersion,
  }
}

// initialize a new Config with default values
func NewDefaultConfig(clientId string, clientSecret string) *Config {
  return &Config{
    ApiVersion: apiVersion,
    Endpoint: apiEndpoint + "/" + apiVersion,
    ClientId: clientId,
    ClientSecret: clientSecret,
    RedirectUri: defaultLocalHost + ":" + strconv.Itoa(defaultLocalPort) + defaultLocalPath,
    Port: defaultLocalPort,
  }
}

// supply new Config with values
func NewConfig(clientId string, clientSecret string, hostUri string, port int) *Config {
  return &Config{
    ApiVersion: apiVersion,
    Endpoint: apiEndpoint + "/" + apiVersion,
    ClientId: clientId,
    ClientSecret: clientSecret,
    RedirectUri: hostUri + ":" + strconv.Itoa(port) + defaultLocalPath,
    Port: port,
  }
}
