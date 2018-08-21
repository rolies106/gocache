package gocache

import (
	"github.com/go-redis/redis"

	"fmt"
)

type StorageConnection struct {
	Connection *redis.Client
}

/**
 * Initiate connection into cache storage
 * @param  	storage_type string	Storage type
 * @param 	address string		Database address include port
 * @param 	password string		Database password
 * @param 	db_no int 			Database No
 * @return {[type]}      [description]
 */
func (this *StorageConnection) Connect(storage_type string, address string, password string, db_no int) {

	switch storage_type {
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db_no,
		})

		pong, err := client.Ping().Result()

		if err != nil {
			panic(err)
		}
		defer redis.Close()

		this.Connection = client
	default:
		fmt.Printf("%s type is not supported.", os)
	}
}
