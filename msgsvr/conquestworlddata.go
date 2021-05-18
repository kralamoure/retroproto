// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestWorldData struct{}

func (m ConquestWorldData) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestWorldData
}

func (m ConquestWorldData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
