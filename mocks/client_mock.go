package mocks

type ClientMock struct {
	CallsCounter
	result []byte
	error  error
}

func NewClientMock(expectedCallsCount int, result []byte, error error) *ClientMock {
	return &ClientMock{
		CallsCounter: CallsCounter{
			expectedCallsCount: expectedCallsCount,
		},
		result: result,
		error:  error,
	}
}

func (c *ClientMock) DoRequest(url string) ([]byte, error) {
	defer c.incCallsCount()
	return c.result, c.error
}
