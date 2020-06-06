// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EmotesUseSuccess struct{}

func (m EmotesUseSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EmotesUseSuccess
}

func (m EmotesUseSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *EmotesUseSuccess) Deserialize(extra string) error {
	return nil
}
