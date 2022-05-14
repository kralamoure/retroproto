// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type JobChangeJobStats struct{}

func (m JobChangeJobStats) MessageId() retroproto.MsgCliId {
	return retroproto.JobChangeJobStats
}

func (m JobChangeJobStats) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobChangeJobStats) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
