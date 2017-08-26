package inventoryBackend_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInventoryBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InventoryBackend Suite")
}
