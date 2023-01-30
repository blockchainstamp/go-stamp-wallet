package comm

import (
	"encoding/binary"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type SDataBase struct {
	*leveldb.DB
}

func (db *SDataBase) SaveJsonObj(key []byte, v interface{}) error {

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	wo := &opt.WriteOptions{
		Sync: true,
	}

	return db.Put(key, data, wo)
}

func (db *SDataBase) GetJsonObj(key []byte, v interface{}) error {

	data, err := db.Get(key, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

func (db *SDataBase) SaveInt(key string, val uint64) error {
	wo := &opt.WriteOptions{
		Sync: true,
	}

	data := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(data, val)
	return db.Put([]byte(key), data, wo)
}

func (db *SDataBase) GetInt(key string) uint64 {
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(data)
}

func (db *SDataBase) AppendString(key string, sep byte, newV string) error {
	var data []byte
	data, _ = db.Get([]byte(key), nil)
	data = append(data, sep)
	data = append(data, []byte(newV)...)

	return db.Put([]byte(key), data, nil)
}
func (db *SDataBase) SaveString(key string, val string) error {
	wo := &opt.WriteOptions{
		Sync: true,
	}
	return db.Put([]byte(key), []byte(val), wo)
}

func (db *SDataBase) GetString(key string) string {
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return ""
	}
	return string(data)
}

func (db *SDataBase) Del(key []byte) error {
	wo := &opt.WriteOptions{
		Sync: true,
	}
	return db.Delete(key, wo)
}
