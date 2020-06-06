// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type JobXP struct{}

func (m JobXP) ProtocolId() d1proto.MsgSvrId {
	return d1proto.JobXP
}

func (m JobXP) Serialized() (string, error) {
	return "", nil
}

func (m *JobXP) Deserialize(extra string) error {
	return nil
}
