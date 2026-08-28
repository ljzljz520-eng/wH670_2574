package codec

import (
	"encoding/json"
	"examarchive/internal/model"
)

func MarshalJSON(r model.ExamRecord) ([]byte, error) { return json.Marshal(r) }
func UnmarshalJSON(data []byte) (model.ExamRecord, error) {
	var r model.ExamRecord
	e := json.Unmarshal(data, &r)
	return r, e
}
func Clone(r model.ExamRecord) (model.ExamRecord, error) {
	b, e := Encode(r)
	if e != nil {
		return model.ExamRecord{}, e
	}
	return Decode(b)
}
