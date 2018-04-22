package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGitclean(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gitclean Suite")
}
