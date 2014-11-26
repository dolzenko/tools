package iterator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dolzenko/tools/ldbcat/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunDefer()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Iterator Suite")
}
