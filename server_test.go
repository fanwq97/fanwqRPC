package fanwqRPC

import (
	"testing"

	"github.com/fanwq97/fanwqRPC/testdata"

	"github.com/stretchr/testify/assert"
)

func TestRegisterService(t *testing.T) {
	s := NewServer()
	err := s.RegisterService("helloworld", new(testdata.Service))
	assert.Nil(t, err)
}
