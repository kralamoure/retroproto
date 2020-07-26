// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismSurvived struct{}

func (m ConquestPrismSurvived) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismSurvived
}

func (m ConquestPrismSurvived) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismSurvived) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
