// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/authorize"
	"chat-backend/internal/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthorizeQuery is the builder for querying Authorize entities.
type AuthorizeQuery struct {
	config
	ctx        *QueryContext
	order      []authorize.OrderOption
	inters     []Interceptor
	predicates []predicate.Authorize
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthorizeQuery builder.
func (aq *AuthorizeQuery) Where(ps ...predicate.Authorize) *AuthorizeQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AuthorizeQuery) Limit(limit int) *AuthorizeQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AuthorizeQuery) Offset(offset int) *AuthorizeQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AuthorizeQuery) Unique(unique bool) *AuthorizeQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AuthorizeQuery) Order(o ...authorize.OrderOption) *AuthorizeQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// First returns the first Authorize entity from the query.
// Returns a *NotFoundError when no Authorize was found.
func (aq *AuthorizeQuery) First(ctx context.Context) (*Authorize, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authorize.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AuthorizeQuery) FirstX(ctx context.Context) *Authorize {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Authorize ID from the query.
// Returns a *NotFoundError when no Authorize ID was found.
func (aq *AuthorizeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authorize.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AuthorizeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Authorize entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Authorize entity is found.
// Returns a *NotFoundError when no Authorize entities are found.
func (aq *AuthorizeQuery) Only(ctx context.Context) (*Authorize, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authorize.Label}
	default:
		return nil, &NotSingularError{authorize.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AuthorizeQuery) OnlyX(ctx context.Context) *Authorize {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Authorize ID in the query.
// Returns a *NotSingularError when more than one Authorize ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AuthorizeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authorize.Label}
	default:
		err = &NotSingularError{authorize.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AuthorizeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Authorizes.
func (aq *AuthorizeQuery) All(ctx context.Context) ([]*Authorize, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Authorize, *AuthorizeQuery]()
	return withInterceptors[[]*Authorize](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AuthorizeQuery) AllX(ctx context.Context) []*Authorize {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Authorize IDs.
func (aq *AuthorizeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(authorize.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AuthorizeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AuthorizeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AuthorizeQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AuthorizeQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AuthorizeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AuthorizeQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthorizeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AuthorizeQuery) Clone() *AuthorizeQuery {
	if aq == nil {
		return nil
	}
	return &AuthorizeQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]authorize.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Authorize{}, aq.predicates...),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JwtToken string `json:"jwt_token,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Authorize.Query().
//		GroupBy(authorize.FieldJwtToken).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AuthorizeQuery) GroupBy(field string, fields ...string) *AuthorizeGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthorizeGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = authorize.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JwtToken string `json:"jwt_token,omitempty"`
//	}
//
//	client.Authorize.Query().
//		Select(authorize.FieldJwtToken).
//		Scan(ctx, &v)
func (aq *AuthorizeQuery) Select(fields ...string) *AuthorizeSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AuthorizeSelect{AuthorizeQuery: aq}
	sbuild.label = authorize.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthorizeSelect configured with the given aggregations.
func (aq *AuthorizeQuery) Aggregate(fns ...AggregateFunc) *AuthorizeSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AuthorizeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !authorize.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AuthorizeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Authorize, error) {
	var (
		nodes = []*Authorize{}
		_spec = aq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Authorize).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Authorize{config: aq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aq *AuthorizeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AuthorizeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authorize.Table, authorize.Columns, sqlgraph.NewFieldSpec(authorize.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authorize.FieldID)
		for i := range fields {
			if fields[i] != authorize.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AuthorizeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(authorize.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = authorize.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthorizeGroupBy is the group-by builder for Authorize entities.
type AuthorizeGroupBy struct {
	selector
	build *AuthorizeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AuthorizeGroupBy) Aggregate(fns ...AggregateFunc) *AuthorizeGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AuthorizeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthorizeQuery, *AuthorizeGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AuthorizeGroupBy) sqlScan(ctx context.Context, root *AuthorizeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthorizeSelect is the builder for selecting fields of Authorize entities.
type AuthorizeSelect struct {
	*AuthorizeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AuthorizeSelect) Aggregate(fns ...AggregateFunc) *AuthorizeSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AuthorizeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthorizeQuery, *AuthorizeSelect](ctx, as.AuthorizeQuery, as, as.inters, v)
}

func (as *AuthorizeSelect) sqlScan(ctx context.Context, root *AuthorizeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}