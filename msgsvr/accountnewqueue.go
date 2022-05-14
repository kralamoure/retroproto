package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountNewQueue struct {
	Position    int
	TotalAbo    int
	TotalNonAbo int
	Subscriber  bool
	QueueId     int
}

func NewAccountNewQueue(extra string) (AccountNewQueue, error) {
	var m AccountNewQueue

	if err := m.Deserialize(extra); err != nil {
		return AccountNewQueue{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountNewQueue) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountNewQueue
}

func (m AccountNewQueue) MessageName() string {
	return "AccountNewQueue"
}

func (m AccountNewQueue) Serialized() (string, error) {
	var subscriber string
	if m.Subscriber {
		subscriber = "1"
	}

	return fmt.Sprintf("%d|%d|%d|%s|%d", m.Position, m.TotalAbo, m.TotalNonAbo, subscriber, m.QueueId), nil
}

func (m *AccountNewQueue) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 5 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		position, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Position = int(position)
	}

	if sli[1] != "" {
		totalAbo, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.TotalAbo = int(totalAbo)
	}

	if sli[2] != "" {
		totalNonAbo, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.TotalNonAbo = int(totalNonAbo)
	}

	if sli[3] != "" {
		subscriber, err := strconv.ParseBool(sli[3])
		if err != nil {
			return err
		}
		m.Subscriber = subscriber
	}

	if sli[4] != "" {
		queueId, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return err
		}
		m.QueueId = int(queueId)
	}

	return nil
}
