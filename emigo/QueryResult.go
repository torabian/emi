package emigo

// QueryResultMeta accompanies a filtered/paged list result: how many rows matched the
// filter (ignoring paging), and an opaque cursor for fetching the next page - nil once
// there's nothing further to fetch (e.g. the page came back empty). Loosely mirrors
// fireback's own QueryResultMeta (see fireback's definitions.go/CrudCoreActions.go),
// minus the parts (TotalAvailableItems, permission/workspace scoping) that don't have
// an emi equivalent yet.
//
// A cursor's actual encoding is an implementation detail of whatever produced it (see
// emigorm.BuildQueryCursor/ApplyQueryCursor for the gorm-backed one) - callers should
// only ever treat it as an opaque token to send back on the next request, never parse
// or construct one themselves.
type QueryResultMeta struct {
	TotalItems int64   `json:"totalItems" yaml:"totalItems"`
	Cursor     *string `json:"cursor" yaml:"cursor"`
}
