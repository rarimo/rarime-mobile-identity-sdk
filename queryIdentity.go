package identity

import "encoding/json"

// QueryIdentityInputs represents the inputs for queryIdentity proof.
type QueryIdentityInputs struct {
	Dg1                       []int64  `json:"dg1"`
	EventID                   string   `json:"eventID"`
	EventData                 string   `json:"eventData"`
	IDStateRoot               string   `json:"idStateRoot"`
	IDStateSiblings           []string `json:"idStateSiblings"`
	PkPassportHash            string   `json:"pkPassportHash"`
	Selector                  string   `json:"selector"`
	SkIdentity                string   `json:"skIdentity"`
	Timestamp                 string   `json:"timestamp"`
	IdentityCounter           string   `json:"identityCounter"`
	TimestampLowerbound       string   `json:"timestampLowerbound"`
	TimestampUpperbound       string   `json:"timestampUpperbound"`
	IdentityCounterLowerbound string   `json:"identityCounterLowerbound"`
	IdentityCounterUpperbound string   `json:"identityCounterUpperbound"`
	BirthDateLowerbound       string   `json:"birthDateLowerbound"`
	BirthDateUpperbound       string   `json:"birthDateUpperbound"`
	ExpirationDateLowerbound  string   `json:"expirationDateLowerbound"`
	ExpirationDateUpperbound  string   `json:"expirationDateUpperbound"`
	CitizenshipMask           string   `json:"citizenshipMask"`
}

// Marshal returns the JSON representation of the inputs.
func (r *QueryIdentityInputs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
