package randomer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/worzeel/randomer"
)

var _ = Describe("Randomer", func() {
	Describe("Using default settings", func() {
		Context("Initial method calls", func() {
			It("Should be able to be created", func() {
				r := new(Randomer)

				Expect(r).NotTo(BeNil())
			})

			It("Should be able to return a string when calling Random method", func() {
				r := new(Randomer)

				str := r.Random()

				Expect(str).NotTo(BeZero())
			})

			It("Should be able to return a unique string when calling Random method twice", func() {
				r := new(Randomer)

				str1 := r.Random()
				str2 := r.Random()

				Expect(str1).NotTo(Equal(str2))
			})
		})
	})
})
