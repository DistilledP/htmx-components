package test

import "strings"

type TestWriter struct {
	rendered []byte
}

func NewTestWriter() *TestWriter {
	return &TestWriter{}
}

func (tw *TestWriter) Write(str []byte) (int, error) {
	tw.rendered = str

	return len(str), nil
}

func (p *TestWriter) ToString() string {
	return string(p.rendered)
}

func (p *TestWriter) ContainsString(testStr string) bool {
	return strings.Contains(p.ToString(), testStr)
}
