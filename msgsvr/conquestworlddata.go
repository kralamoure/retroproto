// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestWorldData struct{}

func (m ConquestWorldData) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestWorldData
}

func (m ConquestWorldData) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestWorldData) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
