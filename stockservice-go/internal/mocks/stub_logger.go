package mocks

type StubLogger struct{}

func (l *StubLogger) Debugf(_ string, _ ...interface{}) {}
func (l *StubLogger) Infof(_ string, _ ...interface{})  {}
func (l *StubLogger) Errorf(_ string, _ ...interface{}) {}
func (l *StubLogger) Info(_ ...interface{})             {}
func (l *StubLogger) Error()                            {}
