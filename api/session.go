// this packa
package api

import (
  "github.com/activatedgeek/putio/api/commons"
  "github.com/activatedgeek/putio/api/account"
  "github.com/activatedgeek/putio/api/files"
  "github.com/activatedgeek/putio/api/friends"
  "github.com/activatedgeek/putio/api/transfers"
  "github.com/activatedgeek/putio/api/zips"
)

type Session struct {
  Account *account.Account
  Files *files.Files
  Friends *friends.Friends
  Transfers *transfers.Transfers
  Zip *zips.Zip
}

// @NOTE persistence of access token is out of scope
func NewSession(accessToken string, config *commons.Config) *Session {
  return &Session{
    Account: account.BuildNewAccount(accessToken, config),
    Files: files.BuildNewFile(accessToken, config),
    Friends: friends.BuildNewFriend(accessToken, config),
    Transfers: transfers.BuildNewTransfer(accessToken, config),
    Zip: zips.BuildNewZip(accessToken, config),
  }
}
