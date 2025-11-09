package value_object

type LastWord string

type InvalidLastWordError struct {}

func (e *InvalidLastWordError) Error() string {
	return "Invalid last word"
}

func NewLastWord(value string) (*LastWord, error) {
	if len(value) <= 255 {
		lastWord := LastWord(value)
		return &lastWord, nil
	}
	return nil, &InvalidLastWordError{}
}