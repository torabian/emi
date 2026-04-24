package queries

import "database/sql"

type Manifest struct {
	DB            *sql.DB
	FilterResolver func(string) (string, error)
}


func (m *Manifest) BasicSelect(ctx BasicSelectContext) ([]BasicSelectRow, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return []BasicSelectRow{}, err
		}
		ctx.Filter = filter
	}
	return BasicSelect(m.DB, ctx)
}

func (m *Manifest) CountUsers(ctx CountUsersContext) ([]CountUsersRow, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return []CountUsersRow{}, err
		}
		ctx.Filter = filter
	}
	return CountUsers(m.DB, ctx)
}

func (m *Manifest) CreateUserTable(ctx CreateUserTableContext) (sql.Result, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return nil, err
		}
		ctx.Filter = filter
	}
	return CreateUserTable(m.DB, ctx)
}

func (m *Manifest) SelectAllUsers(ctx SelectAllUsersContext) ([]SelectAllUsersRow, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return []SelectAllUsersRow{}, err
		}
		ctx.Filter = filter
	}
	return SelectAllUsers(m.DB, ctx)
}

func (m *Manifest) SimpleInsert(ctx SimpleInsertContext) (sql.Result, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return nil, err
		}
		ctx.Filter = filter
	}
	return SimpleInsert(m.DB, ctx)
}

func (m *Manifest) Transaction(ctx TransactionContext) (sql.Result, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return nil, err
		}
		ctx.Filter = filter
	}
	return Transaction(m.DB, ctx)
}

func (m *Manifest) UpdateStatement(ctx UpdateStatementContext) (sql.Result, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return nil, err
		}
		ctx.Filter = filter
	}
	return UpdateStatement(m.DB, ctx)
}

func (m *Manifest) VirtualUserOrder(ctx VirtualUserOrderContext) ([]VirtualUserOrderRow, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return []VirtualUserOrderRow{}, err
		}
		ctx.Filter = filter
	}
	return VirtualUserOrder(m.DB, ctx)
}

