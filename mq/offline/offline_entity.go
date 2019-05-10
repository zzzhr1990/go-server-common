package offline

import (
	"encoding/json"
)

const (
	// ProgressChangeEvent change
	ProgressChangeEvent int32 = 2010
	// FileCompleteEvent change
	FileCompleteEvent int32 = 2020
)

// MqEntity for MQ
type MqEntity struct {
	Type int32
	Data json.RawMessage
}

// ProgressChangeInfo infp
type ProgressChangeInfo struct {
	CurrentSize int64 `json:"currentSize"`
	Size        int64
}

// ToJSON get JSON obj
func (info *ProgressChangeInfo) ToJSON() []byte {
	i, _ := json.Marshal(info)
	return i
}
