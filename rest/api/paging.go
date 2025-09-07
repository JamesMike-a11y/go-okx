package api

type PagingParam struct {
	After  int64 `url:"after,omitempty"`
	Before int64 `url:"before,omitempty"`
	Begin  int64 `url:"begin,omitempty"`
	End    int64 `url:"end,omitempty"`
	Limit  int   `url:"limit,omitempty"`
}
