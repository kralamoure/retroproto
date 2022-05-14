package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type DialogQuestion struct {
	Question       int
	QuestionParams []string
	Answers        []int
}

func (m DialogQuestion) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogQuestion
}

func (m DialogQuestion) Serialized() (string, error) {
	var questionParams string
	if len(m.QuestionParams) != 0 {
		questionParams = fmt.Sprintf(";%s", strings.Join(m.QuestionParams, ","))
	}

	var answers string
	if len(m.Answers) != 0 {
		answers = ";"
		answersSli := make([]string, len(m.Answers))
		for i, v := range m.Answers {
			answersSli[i] = fmt.Sprintf("%d", v)
		}
		answers = fmt.Sprintf("|%s", strings.Join(answersSli, ";"))
	}

	return fmt.Sprintf("%d%s%s", m.Question, questionParams, answers), nil
}

func (m *DialogQuestion) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
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
		answers := strings.Split(sli[1], ";")
		m.Answers = make([]int, len(answers))
		for i, v := range answers {
			answer, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			m.Answers[i] = int(answer)
		}
	}

	return nil
}
