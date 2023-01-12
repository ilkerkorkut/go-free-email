package gofreeemail_test

import (
	"testing"

	gofreeemail "github.com/ilkerkorkut/go-free-email"
	"github.com/stretchr/testify/assert"
)

func TestIsFreeEmail(t *testing.T) {
	assert.True(t, gofreeemail.NewGoFreeEmails().IsFreeEmail("gmail.com"))
}
