package app

type IRequestHeader interface {
	SetValues(accountID string, requestUser string,
		requestApp string, requestSystem string, requestService string,
		requestSession string)

	ValidateHeader() ErrorCollection

	ToHeader() RequestHeader
}
