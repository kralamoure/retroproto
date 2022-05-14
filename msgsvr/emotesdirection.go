package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type EmotesDirection struct {
	Id  int
	Dir int
}

func (m EmotesDirection) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesDirection
}

func (m EmotesDirection) MessageName() string {
	return "EmotesDirection"
}

func (m EmotesDirection) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Dir), nil
}

func (m *EmotesDirection) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	dir, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Dir = int(dir)

	return nil
}
