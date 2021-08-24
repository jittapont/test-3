package main_test

import (
	"sync"
	"testing"

	. "github.com/jittapont/test-3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeEach(func() {
	Wg = sync.WaitGroup{}
})

var _ = AfterEach(func() {
	Wg.Wait()
})

func TestTest3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test3 Suite")
}
