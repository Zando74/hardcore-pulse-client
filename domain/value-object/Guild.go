package value_object

type Guild string

type InvalidGuildError struct {}

func (e *InvalidGuildError) Error() string {
	return "Invalid guild"
}

func NewGuild(value string) (*Guild, error) {
	if len(value) <= 255 {
		guild := Guild(value)
		return &guild, nil
	}
	return nil, &InvalidGuildError{}
}