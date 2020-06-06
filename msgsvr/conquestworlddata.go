// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestWorldData struct{}

func (m ConquestWorldData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestWorldData
}

func (m ConquestWorldData) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestWorldData) Deserialize(extra string) error {
	return nil
}
