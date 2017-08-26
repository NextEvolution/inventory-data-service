package inventoryBackend

import (
  . "nextevolution/inventory-data-service/types"
)

type InventoryBackend interface {
	// Create an item in the inventory. Return it's ID
	CreateItem ([]Item) (string, error)
}