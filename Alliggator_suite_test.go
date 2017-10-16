package alliggator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAlliggator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alliggator Suite")
}
