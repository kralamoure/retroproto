package retroproto

import (
	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountGetServersList struct {
	*retropb.AccountGetServersList
}

func (m AccountGetServersList) MessageId() string {
	return "Ax"
}

func (m AccountGetServersList) String() string {
	return ""
}
