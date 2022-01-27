package main

import (
	UnitTests "mLibAPI/tests"
	testing "testing"
)

func Test(t *testing.T) {
	UnitTests.TestBooksRoute(t)
}
