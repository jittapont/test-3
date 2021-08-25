package main_test

import (
	"testing"

	. "github.com/jittapont/test-3"
	. "github.com/onsi/ginkgo"
	"go.uber.org/goleak"
)

var _ = Describe("main", func() {
	Main()
})

func TestMain(t *testing.T) {
	defer goleak.VerifyNone(t)
	Test2()
}
