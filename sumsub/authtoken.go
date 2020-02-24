package sumsub

import (
	"time"
)

type AuthToken struct {
	token *string
}

func (self *AuthToken) IsValid() bool {
	return self.token != nil
}

func (self *AuthToken) setToken(token *string) {
	self.token = token

	go func() {
		select {
		case <-time.After(time.Second * 595):
			self.token = nil
		}
	}()
}
