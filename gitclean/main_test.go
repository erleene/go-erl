package main_test

import (

	//. "github.com/go-erl/gitclean"
	repo "github.com/go-erl/gitclean/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Running gitclean in a directory", func() {

	dir, err := repo.CheckRepository()

	Context("when it detects a .git folder", func() {
		It("It should return the path to the repository", func() {
			Expect(dir)
			Expect(err).To(BeNil())
		})
	})

	// Context("When it doesn't detect a .git folder", func() {
	// 	It("It should thrown an error", func() {
	// 		Expect(dir)
	// 	})
	// })
})
