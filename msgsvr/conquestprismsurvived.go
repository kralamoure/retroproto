// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismSurvived struct{}

func (m ConquestPrismSurvived) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismSurvived
}

func (m ConquestPrismSurvived) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismSurvived) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
