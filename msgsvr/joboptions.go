// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type JobOptions struct{}

func (m JobOptions) ProtocolId() retroproto.MsgSvrId {
	return retroproto.JobOptions
}

func (m JobOptions) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobOptions) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
