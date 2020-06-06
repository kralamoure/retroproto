package typ

import (
	"fmt"
	"strconv"
	"strings"
)

type CommonAccessories struct {
	Weapon CommonAccessoriesAccessory
	Hat    CommonAccessoriesAccessory
	Cape   CommonAccessoriesAccessory
	Pet    CommonAccessoriesAccessory
	Shield CommonAccessoriesAccessory
}

type CommonAccessoriesAccessory struct {
	TemplateId int
	Type       int
	Frame      int
}

func (m CommonAccessories) Serialized() (string, error) {
	var weapon string
	if m.Weapon.TemplateId != 0 {
		weapon = fmt.Sprintf("%x", m.Weapon.TemplateId)
		if m.Weapon.Type != 0 {
			weapon += fmt.Sprintf("~%d", m.Weapon.Type)
		}
		if m.Weapon.Frame != 0 {
			weapon += fmt.Sprintf("~%d", m.Weapon.Frame)
		}
	}

	var hat string
	if m.Hat.TemplateId != 0 {
		hat = fmt.Sprintf("%x", m.Hat.TemplateId)
		if m.Hat.Type != 0 {
			hat += fmt.Sprintf("~%d", m.Hat.Type)
		}
		if m.Hat.Frame != 0 {
			hat += fmt.Sprintf("~%d", m.Hat.Frame)
		}
	}

	var cape string
	if m.Cape.TemplateId != 0 {
		cape = fmt.Sprintf("%x", m.Cape.TemplateId)
		if m.Cape.Type != 0 {
			cape += fmt.Sprintf("~%d", m.Cape.Type)
		}
		if m.Cape.Frame != 0 {
			cape += fmt.Sprintf("~%d", m.Cape.Frame)
		}
	}

	var pet string
	if m.Pet.TemplateId != 0 {
		pet = fmt.Sprintf("%x", m.Pet.TemplateId)
		if m.Pet.Type != 0 {
			pet += fmt.Sprintf("~%d", m.Pet.Type)
		}
		if m.Pet.Frame != 0 {
			pet += fmt.Sprintf("~%d", m.Pet.Frame)
		}
	}

	var shield string
	if m.Shield.TemplateId != 0 {
		shield = fmt.Sprintf("%x", m.Shield.TemplateId)
		if m.Shield.Type != 0 {
			shield += fmt.Sprintf("~%d", m.Shield.Type)
		}
		if m.Shield.Frame != 0 {
			shield += fmt.Sprintf("~%d", m.Shield.Frame)
		}
	}

	accessories := fmt.Sprintf("%s,%s,%s,%s,%s",
		weapon, hat, cape, pet, shield,
	)

	return accessories, nil
}

func (m *CommonAccessories) Deserialize(extra string) error {
	sli := strings.Split(extra, ",")

	if len(sli) >= 1 {
		if sli[0] != "" {
			accessorySli := strings.Split(sli[0], "~")

			id, err := strconv.ParseInt(accessorySli[0], 16, 32)
			if err != nil {
				return err
			}
			m.Weapon.TemplateId = int(id)

			if len(accessorySli) >= 3 {
				typeId, err := strconv.ParseInt(accessorySli[1], 10, 32)
				if err != nil {
					return err
				}
				m.Weapon.Type = int(typeId)

				frame, err := strconv.ParseInt(accessorySli[2], 10, 32)
				if err != nil {
					return err
				}
				m.Weapon.Frame = int(frame)
			}
		}
	}

	if len(sli) >= 2 {
		if sli[1] != "" {
			accessorySli := strings.Split(sli[1], "~")

			id, err := strconv.ParseInt(accessorySli[0], 16, 32)
			if err != nil {
				return err
			}
			m.Hat.TemplateId = int(id)

			if len(accessorySli) >= 3 {
				typeId, err := strconv.ParseInt(accessorySli[1], 10, 32)
				if err != nil {
					return err
				}
				m.Hat.Type = int(typeId)

				frame, err := strconv.ParseInt(accessorySli[2], 10, 32)
				if err != nil {
					return err
				}
				m.Hat.Frame = int(frame)
			}
		}
	}

	if len(sli) >= 3 {
		if sli[2] != "" {
			accessorySli := strings.Split(sli[2], "~")

			id, err := strconv.ParseInt(accessorySli[0], 16, 32)
			if err != nil {
				return err
			}
			m.Cape.TemplateId = int(id)

			if len(accessorySli) >= 3 {
				typeId, err := strconv.ParseInt(accessorySli[1], 10, 32)
				if err != nil {
					return err
				}
				m.Cape.Type = int(typeId)

				frame, err := strconv.ParseInt(accessorySli[2], 10, 32)
				if err != nil {
					return err
				}
				m.Cape.Frame = int(frame)
			}
		}
	}

	if len(sli) >= 4 {
		if sli[3] != "" {
			accessorySli := strings.Split(sli[3], "~")

			id, err := strconv.ParseInt(accessorySli[0], 16, 32)
			if err != nil {
				return err
			}
			m.Pet.TemplateId = int(id)

			if len(accessorySli) >= 3 {
				typeId, err := strconv.ParseInt(accessorySli[1], 10, 32)
				if err != nil {
					return err
				}
				m.Pet.Type = int(typeId)

				frame, err := strconv.ParseInt(accessorySli[2], 10, 32)
				if err != nil {
					return err
				}
				m.Pet.Frame = int(frame)
			}
		}
	}

	if len(sli) >= 5 {
		if sli[4] != "" {
			accessorySli := strings.Split(sli[4], "~")

			id, err := strconv.ParseInt(accessorySli[0], 16, 32)
			if err != nil {
				return err
			}
			m.Shield.TemplateId = int(id)

			if len(accessorySli) >= 3 {
				typeId, err := strconv.ParseInt(accessorySli[1], 10, 32)
				if err != nil {
					return err
				}
				m.Shield.Type = int(typeId)

				frame, err := strconv.ParseInt(accessorySli[2], 10, 32)
				if err != nil {
					return err
				}
				m.Shield.Frame = int(frame)
			}
		}
	}

	return nil
}
