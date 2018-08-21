package gocache

import (
	"github.com/go-redis/redis"

	"fmt"
)

var (
	Connection *redis.Client
)

/**
 * Initiate connection into cache storage
 * @param  	storage_type string	Storage type
 * @param 	address string		Database address include port
 * @param 	password string		Database password
 * @param 	db_no int 			Database No
 * @return {[type]}      [description]
 */
func Connect(storage_type string, address string, password string, db_no int) {

	switch storage_type {
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db_no,
		})

		_, err := client.Ping().Result()

		if err != nil {
			panic(err)
		}

		Connection = client
	default:
		fmt.Printf("%s type is not supported.", storage_type)
	}
}
