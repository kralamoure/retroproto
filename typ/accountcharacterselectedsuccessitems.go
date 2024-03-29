package typ

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterSelectedSuccessItem struct {
	Id         int
	TemplateId int
	Qty        int
	Position   retrotyp.CharacterItemPosition
	Effects    []retrotyp.Effect
}

func NewAccountCharacterSelectedSuccessItem(extra string) (AccountCharacterSelectedSuccessItem, error) {
	var m AccountCharacterSelectedSuccessItem

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterSelectedSuccessItem{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterSelectedSuccessItem) Serialized() (string, error) {
	return fmt.Sprintf("%x~%x~%x~%x~%s", m.Id, m.TemplateId, m.Qty, int(m.Position), strings.Join(retro.EncodeItemEffects(m.Effects), ",")), nil
}

// TODO
func (m *AccountCharacterSelectedSuccessItem) Deserialize(extra string) error {
	// sli := strings.Split(extra, "~")
	return retroproto.ErrNotImplemented
}
