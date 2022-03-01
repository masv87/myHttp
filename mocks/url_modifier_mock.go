package mocks

type UrlModifierMock struct {
	CallsCounter
	result string
	error  error
}

func NewUrlModifierMock(expectedCallsCount int, result string, error error) *UrlModifierMock {
	return &UrlModifierMock{
		CallsCounter: CallsCounter{
			expectedCallsCount: expectedCallsCount,
		},
		result: result,
		error:  error,
	}
}

func (u *UrlModifierMock) Modify(url string) (string, error) {
	defer u.incCallsCount()
	return u.result, u.error
}
