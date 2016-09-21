// this packa
package api

import (
  "github.com/activatedgeek/putio/api/account"
  "github.com/activatedgeek/putio/api/files"
  "github.com/activatedgeek/putio/api/friends"
  "github.com/activatedgeek/putio/api/transfers"
  "github.com/activatedgeek/putio/api/zip"
)

type Session struct {
  Account *account.Account
  Files *files.Files
  Friends *friends.Friends
  Transfers *transfers.Transfers
  Zip *zip.Zip
}

// Create a new Put.io session using the access token
// @NOTE persistence of access token is out of scope
// and should be handled separately
func NewSession(accessToken string) *Session {
  return &Session{
    Account: &account.Account{
      AccessToken: accessToken,
    },
    Files: &files.Files{
      AccessToken: accessToken,
    },
    Friends: &friends.Friends{
      AccessToken: accessToken,
    },
    Transfers: &transfers.Transfers{
      AccessToken: accessToken,
    },
    Zip: &zip.Zip{
      AccessToken: accessToken,
    },
  }
}
