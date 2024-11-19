package mocks

type StubLogger struct{}

func (l *StubLogger) Debugf(format string, args ...interface{}) {}
func (l *StubLogger) Infof(format string, args ...interface{})  {}
func (l *StubLogger) Errorf(format string, args ...interface{}) {}
func (l *StubLogger) Info(args ...interface{})                  {}
func (l *StubLogger) Error(args ...interface{})                 {}
