// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type JobSkills struct{}

func (m JobSkills) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.JobSkills
}

func (m JobSkills) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobSkills) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
