package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("isPrime", func() {
	It("2 is prime", func() {
		Expect(isPrime(2)).To(BeTrue())
	})

	It("4 is not prime", func() {
		Expect(isPrime(4)).To(BeFalse())
	})
})

var _ = Describe("struct", func() {
	It("should be equal", func() {
		type Post struct {
			title string
		}

		type User struct {
			name  string
			email *string
			posts []*Post
		}

		userA := User{name: "userA", email: &[]string{"hoge"}[0], posts: []*Post{{title: "postA"}}}
		userB := User{name: "userA", posts: []*Post{{title: "postB"}}}
		Expect(userA).NotTo(Equal(userB))
	})
})
