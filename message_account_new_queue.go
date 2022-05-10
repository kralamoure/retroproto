package retroproto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountNewQueue struct {
	*retropb.AccountNewQueue
}

func NewAccountNewQueue(extra string) (AccountNewQueue, error) {
	var m retropb.AccountNewQueue

	sli := strings.Split(extra, "|")
	if len(sli) < 5 {
		return AccountNewQueue{}, errInvalidMessage
	}

	if sli[0] != "" {
		position, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return AccountNewQueue{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Position = int32(position)
	}

	if sli[1] != "" {
		totalAbo, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return AccountNewQueue{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.Subscribers = int32(totalAbo)
	}

	if sli[2] != "" {
		totalNonAbo, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return AccountNewQueue{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.NonSubscribers = int32(totalNonAbo)
	}

	if sli[3] != "" {
		subscriber, err := strconv.ParseBool(sli[3])
		if err != nil {
			return AccountNewQueue{}, fmt.Errorf("could not parse bool: %w", err)
		}
		m.Subscriber = subscriber
	}

	if sli[4] != "" {
		queueId, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return AccountNewQueue{}, fmt.Errorf("could not parse int: %w", err)
		}
		m.QueueId = int32(queueId)
	}

	return AccountNewQueue{AccountNewQueue: &m}, nil
}

func (m AccountNewQueue) MessageId() string {
	return "AX"
}

func (m AccountNewQueue) String() string {
	var subscriber string
	if m.GetSubscriber() {
		subscriber = "1"
	}

	return fmt.Sprintf("%d|%d|%d|%s|%d", m.GetPosition(), m.GetSubscribers(), m.GetNonSubscribers(), subscriber, m.GetQueueId())
}
