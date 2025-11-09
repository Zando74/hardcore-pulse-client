package value_object

type Realm string

type InvalidRealmError struct{}

func (e *InvalidRealmError) Error() string {
	return "Invalid Realm"
}

func NewRealm(value string) (*Realm, error) {
	if len(value) <= 255 {
		Realm := Realm(value)
		return &Realm, nil
	}
	return nil, &InvalidRealmError{}
}