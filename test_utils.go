package dmx

type testReadWriteCloser struct {
	written *[][]byte
	read    *[][]byte
}

func createTestReadWriteCloser() *testReadWriteCloser {
	return &testReadWriteCloser{&[][]byte{}, &[][]byte{}}
}

func (t testReadWriteCloser) Write(p []byte) (int, error) {
	*t.written = append(*t.written, p)
	return len(p), nil
}

func (t testReadWriteCloser) Read(b []byte) (int, error) {
	*t.read = append(*t.read, b)
	return len(b), nil
}

func (testReadWriteCloser) Close() error {
	return nil
}
