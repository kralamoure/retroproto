package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type DialogResponse struct {
	Question int
	Answer   int
}

func NewDialogResponse(extra string) (DialogResponse, error) {
	var m DialogResponse

	if err := m.Deserialize(extra); err != nil {
		return DialogResponse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogResponse) MessageId() retroproto.MsgCliId {
	return retroproto.DialogResponse
}

func (m DialogResponse) MessageName() string {
	return "DialogResponse"
}

func (m DialogResponse) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Question, m.Answer), nil
}

func (m *DialogResponse) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
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
