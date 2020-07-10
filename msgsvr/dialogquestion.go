package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type DialogQuestion struct {
	Question       int
	QuestionParams []string
	Responses      []int
}

func (m DialogQuestion) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DialogQuestion
}

func (m DialogQuestion) Serialized() (string, error) {
	var questionParams string
	if len(m.QuestionParams) != 0 {
		questionParams = ";"
		questionParams += strings.Join(m.QuestionParams, ",")
	}

	var responses string
	if len(m.Responses) != 0 {
		responsesSli := make([]string, len(m.Responses))
		for i, v := range m.Responses {
			responsesSli[i] = fmt.Sprintf("%d", v)
		}
	}

	return fmt.Sprintf("%d%s%s", m.Question, questionParams, responses), nil
}

func (m *DialogQuestion) Deserialize(extra string) error {
	if extra == "" {
		return d1proto.ErrInvalidMsg
	}

	sli := strings.Split(extra, "|")

	sli2 := strings.Split(sli[0], ";")

	question, err := strconv.ParseInt(sli2[0], 10, 32)
	if err != nil {
		return err
	}
	m.Question = int(question)

	if len(sli2) > 1 && sli2[1] != "" {
		m.QuestionParams = strings.Split(sli2[1], ",")
	}

	if len(sli) > 1 && sli[1] != "" {
		responses := strings.Split(sli[1], ";")
		m.Responses = make([]int, len(responses))
		for i, v := range responses {
			response, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			m.Responses[i] = int(response)
		}
	}

	return nil
}
