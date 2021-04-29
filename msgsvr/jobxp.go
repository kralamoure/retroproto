// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type JobXP struct{}

func (m JobXP) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.JobXP
}

func (m JobXP) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobXP) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
