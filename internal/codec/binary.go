package codec

import (
	"bytes"
	"encoding/gob"
	"examarchive/internal/model"
)

func Encode(r model.ExamRecord) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(r)
	return b.Bytes(), err
}
func Decode(data []byte) (model.ExamRecord, error) {
	var r model.ExamRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&r)
	return r, err
}
func EncodeBatch(rs []model.ExamRecord) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(rs)
	return b.Bytes(), err
}
func DecodeBatch(data []byte) ([]model.ExamRecord, error) {
	var rs []model.ExamRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&rs)
	return rs, err
}
