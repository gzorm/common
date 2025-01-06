// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gzorm/common/ent/predicate"
	"github.com/gzorm/common/ent/winuser"
)

// WinUserQuery is the builder for querying WinUser entities.
type WinUserQuery struct {
	config
	ctx        *QueryContext
	order      []winuser.OrderOption
	inters     []Interceptor
	predicates []predicate.WinUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WinUserQuery builder.
func (wuq *WinUserQuery) Where(ps ...predicate.WinUser) *WinUserQuery {
	wuq.predicates = append(wuq.predicates, ps...)
	return wuq
}

// Limit the number of records to be returned by this query.
func (wuq *WinUserQuery) Limit(limit int) *WinUserQuery {
	wuq.ctx.Limit = &limit
	return wuq
}

// Offset to start from.
func (wuq *WinUserQuery) Offset(offset int) *WinUserQuery {
	wuq.ctx.Offset = &offset
	return wuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wuq *WinUserQuery) Unique(unique bool) *WinUserQuery {
	wuq.ctx.Unique = &unique
	return wuq
}

// Order specifies how the records should be ordered.
func (wuq *WinUserQuery) Order(o ...winuser.OrderOption) *WinUserQuery {
	wuq.order = append(wuq.order, o...)
	return wuq
}

// First returns the first WinUser entity from the query.
// Returns a *NotFoundError when no WinUser was found.
func (wuq *WinUserQuery) First(ctx context.Context) (*WinUser, error) {
	nodes, err := wuq.Limit(1).All(setContextOp(ctx, wuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{winuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wuq *WinUserQuery) FirstX(ctx context.Context) *WinUser {
	node, err := wuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WinUser ID from the query.
// Returns a *NotFoundError when no WinUser ID was found.
func (wuq *WinUserQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wuq.Limit(1).IDs(setContextOp(ctx, wuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{winuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wuq *WinUserQuery) FirstIDX(ctx context.Context) int32 {
	id, err := wuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WinUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WinUser entity is found.
// Returns a *NotFoundError when no WinUser entities are found.
func (wuq *WinUserQuery) Only(ctx context.Context) (*WinUser, error) {
	nodes, err := wuq.Limit(2).All(setContextOp(ctx, wuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{winuser.Label}
	default:
		return nil, &NotSingularError{winuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wuq *WinUserQuery) OnlyX(ctx context.Context) *WinUser {
	node, err := wuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WinUser ID in the query.
// Returns a *NotSingularError when more than one WinUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (wuq *WinUserQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wuq.Limit(2).IDs(setContextOp(ctx, wuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{winuser.Label}
	default:
		err = &NotSingularError{winuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wuq *WinUserQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := wuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WinUsers.
func (wuq *WinUserQuery) All(ctx context.Context) ([]*WinUser, error) {
	ctx = setContextOp(ctx, wuq.ctx, "All")
	if err := wuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WinUser, *WinUserQuery]()
	return withInterceptors[[]*WinUser](ctx, wuq, qr, wuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wuq *WinUserQuery) AllX(ctx context.Context) []*WinUser {
	nodes, err := wuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WinUser IDs.
func (wuq *WinUserQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if wuq.ctx.Unique == nil && wuq.path != nil {
		wuq.Unique(true)
	}
	ctx = setContextOp(ctx, wuq.ctx, "IDs")
	if err = wuq.Select(winuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wuq *WinUserQuery) IDsX(ctx context.Context) []int32 {
	ids, err := wuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wuq *WinUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wuq.ctx, "Count")
	if err := wuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wuq, querierCount[*WinUserQuery](), wuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wuq *WinUserQuery) CountX(ctx context.Context) int {
	count, err := wuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wuq *WinUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wuq.ctx, "Exist")
	switch _, err := wuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wuq *WinUserQuery) ExistX(ctx context.Context) bool {
	exist, err := wuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WinUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wuq *WinUserQuery) Clone() *WinUserQuery {
	if wuq == nil {
		return nil
	}
	return &WinUserQuery{
		config:     wuq.config,
		ctx:        wuq.ctx.Clone(),
		order:      append([]winuser.OrderOption{}, wuq.order...),
		inters:     append([]Interceptor{}, wuq.inters...),
		predicates: append([]predicate.WinUser{}, wuq.predicates...),
		// clone intermediate query.
		sql:  wuq.sql.Clone(),
		path: wuq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WinUser.Query().
//		GroupBy(winuser.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wuq *WinUserQuery) GroupBy(field string, fields ...string) *WinUserGroupBy {
	wuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WinUserGroupBy{build: wuq}
	grbuild.flds = &wuq.ctx.Fields
	grbuild.label = winuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.WinUser.Query().
//		Select(winuser.FieldUsername).
//		Scan(ctx, &v)
func (wuq *WinUserQuery) Select(fields ...string) *WinUserSelect {
	wuq.ctx.Fields = append(wuq.ctx.Fields, fields...)
	sbuild := &WinUserSelect{WinUserQuery: wuq}
	sbuild.label = winuser.Label
	sbuild.flds, sbuild.scan = &wuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WinUserSelect configured with the given aggregations.
func (wuq *WinUserQuery) Aggregate(fns ...AggregateFunc) *WinUserSelect {
	return wuq.Select().Aggregate(fns...)
}

func (wuq *WinUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wuq); err != nil {
				return err
			}
		}
	}
	for _, f := range wuq.ctx.Fields {
		if !winuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wuq.path != nil {
		prev, err := wuq.path(ctx)
		if err != nil {
			return err
		}
		wuq.sql = prev
	}
	return nil
}

func (wuq *WinUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WinUser, error) {
	var (
		nodes = []*WinUser{}
		_spec = wuq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WinUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WinUser{config: wuq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wuq *WinUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wuq.querySpec()
	_spec.Node.Columns = wuq.ctx.Fields
	if len(wuq.ctx.Fields) > 0 {
		_spec.Unique = wuq.ctx.Unique != nil && *wuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wuq.driver, _spec)
}

func (wuq *WinUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(winuser.Table, winuser.Columns, sqlgraph.NewFieldSpec(winuser.FieldID, field.TypeInt32))
	_spec.From = wuq.sql
	if unique := wuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wuq.path != nil {
		_spec.Unique = true
	}
	if fields := wuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, winuser.FieldID)
		for i := range fields {
			if fields[i] != winuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wuq *WinUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wuq.driver.Dialect())
	t1 := builder.Table(winuser.Table)
	columns := wuq.ctx.Fields
	if len(columns) == 0 {
		columns = winuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wuq.sql != nil {
		selector = wuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wuq.ctx.Unique != nil && *wuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wuq.predicates {
		p(selector)
	}
	for _, p := range wuq.order {
		p(selector)
	}
	if offset := wuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (wuq *WinUserQuery) SqlQuery(ctx context.Context) (r1 string, r2 []any) {
	builder := sql.Dialect(wuq.driver.Dialect())
	t1 := builder.Table(winuser.Table)
	columns := wuq.ctx.Fields
	if len(columns) == 0 {
		columns = winuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wuq.sql != nil {
		selector = wuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wuq.ctx.Unique != nil && *wuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wuq.predicates {
		p(selector)
	}
	for _, p := range wuq.order {
		p(selector)
	}
	if offset := wuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	r1, r2 = selector.Query()
	//r3 = make([]*WinUser,0)

	return
}

// WinUserGroupBy is the group-by builder for WinUser entities.
type WinUserGroupBy struct {
	selector
	build *WinUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wugb *WinUserGroupBy) Aggregate(fns ...AggregateFunc) *WinUserGroupBy {
	wugb.fns = append(wugb.fns, fns...)
	return wugb
}

// Scan applies the selector query and scans the result into the given value.
func (wugb *WinUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wugb.build.ctx, "GroupBy")
	if err := wugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinUserQuery, *WinUserGroupBy](ctx, wugb.build, wugb, wugb.build.inters, v)
}

func (wugb *WinUserGroupBy) sqlScan(ctx context.Context, root *WinUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wugb.fns))
	for _, fn := range wugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wugb.flds)+len(wugb.fns))
		for _, f := range *wugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WinUserSelect is the builder for selecting fields of WinUser entities.
type WinUserSelect struct {
	*WinUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wus *WinUserSelect) Aggregate(fns ...AggregateFunc) *WinUserSelect {
	wus.fns = append(wus.fns, fns...)
	return wus
}

// Scan applies the selector query and scans the result into the given value.
func (wus *WinUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wus.ctx, "Select")
	if err := wus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WinUserQuery, *WinUserSelect](ctx, wus.WinUserQuery, wus, wus.inters, v)
}

func (wus *WinUserSelect) sqlScan(ctx context.Context, root *WinUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wus.fns))
	for _, fn := range wus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
