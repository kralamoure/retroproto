// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobChangeJobStats struct{}

func NewJobChangeJobStats(extra string) (JobChangeJobStats, error) {
	var m JobChangeJobStats

	if err := m.Deserialize(extra); err != nil {
		return JobChangeJobStats{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobChangeJobStats) MessageId() retroproto.MsgCliId {
	return retroproto.JobChangeJobStats
}

func (m JobChangeJobStats) MessageName() string {
	return "JobChangeJobStats"
}

func (m JobChangeJobStats) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobChangeJobStats) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
