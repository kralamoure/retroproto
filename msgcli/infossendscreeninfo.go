// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type InfosSendScreenInfo struct{}

func (m InfosSendScreenInfo) ProtocolId() d1proto.MsgCliId {
	return d1proto.InfosSendScreenInfo
}

func (m InfosSendScreenInfo) Serialized() (string, error) {
	return "", nil
}

func (m *InfosSendScreenInfo) Deserialize(extra string) error {
	return nil
}
