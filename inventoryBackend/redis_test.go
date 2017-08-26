package inventoryBackend_test

import (
	. "nextevolution/inventory-data-service/inventoryBackend"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redis", func() {
	redisAddress := "localhost:6379"
	redisProtocol := "tcp"
	redisPoolSize := 1
	rib := NewRedisInventoryBackend(redisProtocol, redisAddress, redisPoolSize)

	Context("CRUD", func(){
		It("Creates an inventory item", func(){

		})
	})

})
