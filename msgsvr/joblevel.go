// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type JobLevel struct{}

func (m JobLevel) ProtocolId() retroproto.MsgSvrId {
	return retroproto.JobLevel
}

func (m JobLevel) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobLevel) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
