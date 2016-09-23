// this package is a wrapper for http communication with the api
// it wraps the response into standard operable structure
package commons

type Response struct {
  // status string, "OK" or "ERROR"
  Status string

  // only in case of "ERROR"
  ErrorType string
  ErrorMessage string

  //
  Data string

  // raw response from http request
  Raw string
}

func BuildNewResponse() (*Response) {
  return &Response{};
}
