package goblox

type IAuthenticatedUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type IPagedResponse[T any] struct {
	PreviousPageCursor string `json:"previousPageCursor"`
	NextPageCursor     string `json:"nextPageCursor"`
	Data               []T    `json:"data"`
}

type DatumedResponse[T any] struct {
	Data []T `json:"data"`
}

type UsersHandlerStruct struct{}
