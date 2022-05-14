// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EmotesAdd struct{}

func (m EmotesAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesAdd
}

func (m EmotesAdd) MessageName() string {
	return "EmotesAdd"
}

func (m EmotesAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
