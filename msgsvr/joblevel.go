// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type JobLevel struct{}

func (m JobLevel) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.JobLevel
}

func (m JobLevel) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobLevel) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
