package caster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Author struct {
	FirstName     string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	PublishedBook int32  `protobuf:"varint,3,opt,name=PublishedBook,proto3" json:"PublishedBook,omitempty"`
	Followers     int32  `protobuf:"varint,4,opt,name=Followers,proto3" json:"Followers,omitempty"`
}

type Person struct {
	FirstName string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func TestCast(t *testing.T) {
	a := &Author{
		FirstName:     "J.K.",
		LastName:      "Rowling",
		PublishedBook: 10,
		Followers:     1000,
	}

	p := &Person{}
	err := Cast(a, p)

	assert.NoError(t, err)
	assert.Equal(t, a.FirstName, p.FirstName)
	assert.Equal(t, a.LastName, p.LastName)
}

func BenchmarkTestCastSimpleStruct(b *testing.B) {

	a := &Author{
		FirstName:     "J.K.",
		LastName:      "Rowling",
		PublishedBook: 10,
		Followers:     1000,
	}

	for n := 0; n < b.N; n++ {
		p := &Person{}
		Cast(a, p)
	}
}
