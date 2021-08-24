package main_test

import (
	. "github.com/onsi/ginkgo"
	"sync"
)

var _ = Describe("main", func() {
	BeforeEach(func() {
		Wg = sync.WaitGroup{}
	})

})
