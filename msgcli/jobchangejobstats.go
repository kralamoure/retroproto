// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type JobChangeJobStats struct{}

func (m JobChangeJobStats) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.JobChangeJobStats
}

func (m JobChangeJobStats) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobChangeJobStats) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
