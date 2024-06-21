package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockLoggerInterface is a mock type for the LoggerInterface
type MockLoggerInterface struct {
	mock.Mock
}

// Info mocks the Info method of LoggerInterface
func (m *MockLoggerInterface) Info(msg string, args ...interface{}) {
	m.Called(args...)
}

// Error mocks the Error method of LoggerInterface
func (m *MockLoggerInterface) Error(msg string, args ...interface{}) {
	m.Called(args...)
}

// Fatal mocks the Fatal method of LoggerInterface
func (m *MockLoggerInterface) Fatal(msg string, args ...interface{}) {
	m.Called(args...)
}

// Close mocks the Close method of LoggerInterface
func (m *MockLoggerInterface) Close() {
	m.Called()
}
