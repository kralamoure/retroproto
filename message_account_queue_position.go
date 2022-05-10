package retroproto

import (
	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountQueuePosition struct {
	*retropb.AccountQueuePosition
}

func (m AccountQueuePosition) MessageId() string {
	return "Af"
}

func (m AccountQueuePosition) String() string {
	return ""
}
