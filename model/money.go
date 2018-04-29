package model

import(
	"encoding/json"
)

type Money int64

func (m *Money) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*m = Money(v * 100)
	return nil
}

// MarshalJSON marshals money
func (m Money) MarshalJSON() ([]byte, error) {
	v := float64(m) / float64(100)

	return json.Marshal(v)
}
