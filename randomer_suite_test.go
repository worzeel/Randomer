package randomer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRandomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Randomer Suite")
}

// this is actually quite cool to use vim to edit code!
