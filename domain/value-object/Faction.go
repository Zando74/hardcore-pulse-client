package value_object

type Faction string

type InvalidFactionError struct {
	faction string
}

func (e *InvalidFactionError) Error() string {
	return "Invalid faction : " + e.faction
}

func isValidFaction(value string) bool {
	return value == "Horde" || value == "Alliance"
}

func NewFaction(value string) (*Faction, error) {
	if isValidFaction(value) {
		faction := Faction(value)
		return &faction, nil
	}
	return nil, &InvalidFactionError{
		faction: value,
	}
}