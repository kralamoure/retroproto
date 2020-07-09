package typ

import (
	"fmt"
)

type AccountCharacterSelectedSuccessItem struct {
	Id         int
	TemplateId int
	Qty        int
	Position   int // -1 for no position
	Effects    string
}

func (m AccountCharacterSelectedSuccessItem) Serialized() (string, error) {
	return fmt.Sprintf("%x~%x~%x~%x~%s", m.Id, m.TemplateId, m.Qty, m.Position, m.Effects), nil
}

// TODO
func (m *AccountCharacterSelectedSuccessItem) Deserialize(extra string) error {
	// sli := strings.Split(extra, "~")
	return nil
}
