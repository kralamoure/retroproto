package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type DialogResponse struct {
	Question int
	Answer   int
}

func (m DialogResponse) ProtocolId() d1proto.MsgCliId {
	return d1proto.DialogResponse
}

func (m DialogResponse) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Question, m.Answer), nil
}

func (m *DialogResponse) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
	}

	question, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Question = int(question)

	answer, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Answer = int(answer)

	return nil
}
