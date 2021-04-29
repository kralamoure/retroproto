// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismSurvived struct{}

func (m ConquestPrismSurvived) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismSurvived
}

func (m ConquestPrismSurvived) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismSurvived) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
