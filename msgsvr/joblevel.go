// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type JobLevel struct{}

func (m JobLevel) ProtocolId() d1proto.MsgSvrId {
	return d1proto.JobLevel
}

func (m JobLevel) Serialized() (string, error) {
	return "", nil
}

func (m *JobLevel) Deserialize(extra string) error {
	return nil
}
