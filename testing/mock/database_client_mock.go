package mock

import "github.com/stretchr/testify/mock"

type DatabaseClientMock struct {
	mock.Mock
}

func (c *DatabaseClientMock) Save(statement string, entity Entity) error {
	args := c.Called(statement, entity)
	return args.Error(0)
}
