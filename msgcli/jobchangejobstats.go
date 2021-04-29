// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type JobChangeJobStats struct{}

func (m JobChangeJobStats) ProtocolId() d1proto.MsgCliId {
	return d1proto.JobChangeJobStats
}

func (m JobChangeJobStats) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *JobChangeJobStats) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
