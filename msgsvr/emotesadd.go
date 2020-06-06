// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EmotesAdd struct{}

func (m EmotesAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EmotesAdd
}

func (m EmotesAdd) Serialized() (string, error) {
	return "", nil
}

func (m *EmotesAdd) Deserialize(extra string) error {
	return nil
}
