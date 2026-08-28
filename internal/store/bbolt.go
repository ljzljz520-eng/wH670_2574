package store

import (
	"examarchive/internal/codec"
	"examarchive/internal/model"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
)

var bucket = []byte("exam_records")

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bbolt.Tx) error { _, e := tx.CreateBucketIfNotExists(bucket); return e })
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Store{db: db, path: path}, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) Path() string { return filepath.Clean(s.path) }
func (s *Store) Put(r model.ExamRecord) error {
	data, err := codec.Encode(r)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(r.Candidate.ID), data) })
}
func (s *Store) Get(id string) (model.ExamRecord, error) {
	var out model.ExamRecord
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucket).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("record %s not found", id)
		}
		var e error
		out, e = codec.Decode(v)
		return e
	})
	return out, err
}
func (s *Store) All() ([]model.ExamRecord, error) {
	out := []model.ExamRecord{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(_, v []byte) error {
			r, e := codec.Decode(v)
			if e == nil {
				out = append(out, r)
			}
			return e
		})
	})
	return out, err
}
func (s *Store) Delete(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(id)) })
}
