package faker

import (
	"github.com/bxcodec/faker/v3"
)

type Generator interface {
	GenerateName() string
}

type FakerGenerator struct{}

func (f *FakerGenerator) GenerateName() string {
	return faker.Name()
}
