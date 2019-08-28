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

	db := GenericOpenDB("./storage.db")

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucket))
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(keyName), []byte(keyValue))
		return err
	})
}

func (c *GeneralDB) ReadKey(bucket string, keyName string) []byte {

	db := GenericOpenDB("./storage.db")

	v := []byte{0}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v = b.Get([]byte(keyName))
		return nil
	})
	return v
}

//func (c *GeneralDB) IterateOverKeysInBucketReturnBoth() [][]string {

//}

func GenericOpenDB(dbName string) *bolt.DB {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return db
}
