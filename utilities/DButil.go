package utilities

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	bolt "go.etcd.io/bbolt"
)

type GeneralDB struct {
	session discordgo.Session
	message discordgo.MessageCreate
}

func (c *GeneralDB) WriteToDB(bucket string, keyName string, keyValue string) {

	db, err := bolt.Open("./storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(keyName), []byte(keyValue))
		return err
	})
	db.Close()
}

func (c *GeneralDB) DeleteKey(bucket string, keyName string) {
	db, err := bolt.Open("./storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Delete([]byte(keyName))
		return nil
	})
	db.Close()
}

func (c *GeneralDB) ReadKey(bucket string, keyName string) []byte {

	db, err := bolt.Open("./storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	v := []byte{0}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v = b.Get([]byte(keyName))
		return nil
	})
	db.Close()
	return v
}

func (c *GeneralDB) IterateOverKeysInBucketReturnBoth(bucket string) ([]string, []string) {
	db, err := bolt.Open("./storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	nameValuesArray := []string{}
	valueValuesArray := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			nameValuesArray = append(nameValuesArray, string(k))
			valueValuesArray = append(valueValuesArray, string(v))
		}
		return nil
	})
	db.Close()
	return nameValuesArray, valueValuesArray
}

func (c *GeneralDB) EnsureBucketsExist(buckets string) {
	db, err := bolt.Open("storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(([]byte(buckets)))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Close()
}
