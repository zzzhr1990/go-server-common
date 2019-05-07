package httprpc

import (
	"encoding/json"
)

//StandardResponseEntity .
type StandardResponseEntity struct {
	Code    string           `json:"code,omitempty"`
	Status  int              `json:"status,omitempty"`
	Success bool             `json:"success,omitempty"`
	Message string           `json:"message,omitempty"`
	Result  *json.RawMessage `json:"result,omitempty"`
}

// DecodeRaw decode
func (entity *StandardResponseEntity) DecodeRaw(result interface{}) error {
	if entity.Result == nil {
		return nil
	}
	return json.Unmarshal(*entity.Result, result)
}
