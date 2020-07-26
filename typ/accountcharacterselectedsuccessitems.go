package typ

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1"
	"github.com/kralamoure/d1/d1typ"
)

type AccountCharacterSelectedSuccessItem struct {
	Id         int
	TemplateId int
	Qty        int
	Position   d1typ.CharacterItemPosition
	Effects    []d1typ.Effect
}

func (m AccountCharacterSelectedSuccessItem) Serialized() (string, error) {
	return fmt.Sprintf("%x~%x~%x~%x~%s", m.Id, m.TemplateId, m.Qty, m.Position, strings.Join(d1.EncodeItemEffects(m.Effects), ",")), nil
}

// TODO
func (m *AccountCharacterSelectedSuccessItem) Deserialize(extra string) error {
	// sli := strings.Split(extra, "~")
	return nil
}
