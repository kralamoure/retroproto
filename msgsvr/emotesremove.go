// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EmotesRemove struct{}

func (m EmotesRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EmotesRemove
}

func (m EmotesRemove) Serialized() (string, error) {
	return "", nil
}

func (m *EmotesRemove) Deserialize(extra string) error {
	return nil
}
