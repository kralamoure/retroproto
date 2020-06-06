// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsFileCheckAnswer struct{}

func (m BasicsFileCheckAnswer) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsFileCheckAnswer
}

func (m BasicsFileCheckAnswer) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsFileCheckAnswer) Deserialize(extra string) error {
	return nil
}
