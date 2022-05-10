package retroproto

import (
	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AksPing struct {
	*retropb.AksPing
}

func (m AksPing) MessageId() string {
	return "ping"
}

func (m AksPing) String() string {
	return ""
}
