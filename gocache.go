package gocache

import (
	"fmt"
)

type GoCache struct{}

func (this *GoCache) SetMultiKeysCache(cache_type string) {

	fmt.Println("%v", StorageConnection.Connection)
}
