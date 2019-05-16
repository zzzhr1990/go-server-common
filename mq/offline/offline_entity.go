package offline

//  github.com/zzzhr1990/go-server-common/mq/offline

import (
	"encoding/json"
)

const (
	// MetaDataEvent when meta change
	MetaDataEvent int32 = 2005
	// ProgressChangeEvent change
	ProgressChangeEvent int32 = 2010
	// FileCompleteEvent change
	FileCompleteEvent int32 = 2020
	// StatusChangeEvent new status
	StatusChangeEvent int32 = 2050
	// UserFileCopyEvent when task start
	UserFileCopyEvent int32 = 2060
	// TaskCompleteEvent when task complete
	TaskCompleteEvent int32 = 2080
)

// MqEntity for MQ
type MqEntity struct {
	Type     int32
	Identity string
	Data     json.RawMessage
}

// ToJSONByte info
func (mqEntity *MqEntity) ToJSONByte() []byte {
	i, _ := json.Marshal(mqEntity)
	return i
}

/*
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
*/
