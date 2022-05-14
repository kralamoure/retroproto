// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EmotesUseSuccess struct{}

func (m EmotesUseSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesUseSuccess
}

func (m EmotesUseSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
