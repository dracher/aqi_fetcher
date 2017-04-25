package libs

import (
	"time"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

// NewBoltDB is
func NewBoltDB(dbPath string) *bolt.DB {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// SaveAqiData is
func SaveAqiData(db *bolt.DB, aqiData *AqiFlatData) error {
	db.Update(func(tx *bolt.Tx) error {
		buk, err := tx.CreateBucketIfNotExists([]byte(aqiData.CityName))
		if err != nil {
			log.Errorf("Can't create bucket %s", err)
			return err
		}
		val, _ := json.Marshal(aqiData)
		err = buk.Put([]byte(aqiData.Date), val)
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	})
	return nil
}

// DeleteBuckets is
func DeleteBuckets(db *bolt.DB, bucketNames []string) {
	db.Update(func(tx *bolt.Tx) error {
		for _, bn := range bucketNames {
			err := tx.DeleteBucket([]byte(bn))
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	})
}
