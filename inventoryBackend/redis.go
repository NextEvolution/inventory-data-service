package inventoryBackend

import (
	. "nextevolution/inventory-data-service/types"
)

type RedisInventoryBackend struct {

}

func (ris *RedisInventoryBackend) CreateItem(items []Item) (string, error) {
	return "", nil
}

func NewRedisInventoryBackend (protocol string, address string, poolSize int) (RedisInventoryBackend){

	return nil
}