// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PayPerViewEvent is an object representing the database table.
type PayPerViewEvent struct {
	ID          int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string     `boil:"description" json:"description" toml:"description" yaml:"description"`
	EventType   string     `boil:"event_type" json:"eventType" toml:"eventType" yaml:"eventType"`
	Start       null.Time  `boil:"start" json:"start,omitempty" toml:"start" yaml:"start,omitempty"`
	End         null.Time  `boil:"end" json:"end,omitempty" toml:"end" yaml:"end,omitempty"`
	PriceEth    null.Int64 `boil:"price_eth" json:"priceEth,omitempty" toml:"priceEth" yaml:"priceEth,omitempty"`
	PriceBTC    null.Int64 `boil:"price_btc" json:"priceBTC,omitempty" toml:"priceBTC" yaml:"priceBTC,omitempty"`
	PriceXMR    null.Int64 `boil:"price_xmr" json:"priceXMR,omitempty" toml:"priceXMR" yaml:"priceXMR,omitempty"`
	CreatedAt   null.Time  `boil:"created_at" json:"createdAt,omitempty" toml:"createdAt" yaml:"createdAt,omitempty"`
	UpdatedAt   null.Time  `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`

	R *payPerViewEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L payPerViewEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PayPerViewEventColumns = struct {
	ID          string
	Name        string
	Description string
	EventType   string
	Start       string
	End         string
	PriceEth    string
	PriceBTC    string
	PriceXMR    string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	EventType:   "event_type",
	Start:       "start",
	End:         "end",
	PriceEth:    "price_eth",
	PriceBTC:    "price_btc",
	PriceXMR:    "price_xmr",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PayPerViewEventWhere = struct {
	ID          whereHelperint64
	Name        whereHelperstring
	Description whereHelperstring
	EventType   whereHelperstring
	Start       whereHelpernull_Time
	End         whereHelpernull_Time
	PriceEth    whereHelpernull_Int64
	PriceBTC    whereHelpernull_Int64
	PriceXMR    whereHelpernull_Int64
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	ID:          whereHelperint64{field: "\"pay_per_view_events\".\"id\""},
	Name:        whereHelperstring{field: "\"pay_per_view_events\".\"name\""},
	Description: whereHelperstring{field: "\"pay_per_view_events\".\"description\""},
	EventType:   whereHelperstring{field: "\"pay_per_view_events\".\"event_type\""},
	Start:       whereHelpernull_Time{field: "\"pay_per_view_events\".\"start\""},
	End:         whereHelpernull_Time{field: "\"pay_per_view_events\".\"end\""},
	PriceEth:    whereHelpernull_Int64{field: "\"pay_per_view_events\".\"price_eth\""},
	PriceBTC:    whereHelpernull_Int64{field: "\"pay_per_view_events\".\"price_btc\""},
	PriceXMR:    whereHelpernull_Int64{field: "\"pay_per_view_events\".\"price_xmr\""},
	CreatedAt:   whereHelpernull_Time{field: "\"pay_per_view_events\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"pay_per_view_events\".\"updated_at\""},
}

// PayPerViewEventRels is where relationship names are stored.
var PayPerViewEventRels = struct {
	SmartContracts string
}{
	SmartContracts: "SmartContracts",
}

// payPerViewEventR is where relationships are stored.
type payPerViewEventR struct {
	SmartContracts SmartContractSlice
}

// NewStruct creates a new relationship struct
func (*payPerViewEventR) NewStruct() *payPerViewEventR {
	return &payPerViewEventR{}
}

// payPerViewEventL is where Load methods for each relationship are stored.
type payPerViewEventL struct{}

var (
	payPerViewEventAllColumns            = []string{"id", "name", "description", "event_type", "start", "end", "price_eth", "price_btc", "price_xmr", "created_at", "updated_at"}
	payPerViewEventColumnsWithoutDefault = []string{"name", "description", "event_type", "start", "end", "price_eth", "price_btc", "price_xmr", "created_at", "updated_at"}
	payPerViewEventColumnsWithDefault    = []string{"id"}
	payPerViewEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// PayPerViewEventSlice is an alias for a slice of pointers to PayPerViewEvent.
	// This should generally be used opposed to []PayPerViewEvent.
	PayPerViewEventSlice []*PayPerViewEvent

	payPerViewEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	payPerViewEventType                 = reflect.TypeOf(&PayPerViewEvent{})
	payPerViewEventMapping              = queries.MakeStructMapping(payPerViewEventType)
	payPerViewEventPrimaryKeyMapping, _ = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, payPerViewEventPrimaryKeyColumns)
	payPerViewEventInsertCacheMut       sync.RWMutex
	payPerViewEventInsertCache          = make(map[string]insertCache)
	payPerViewEventUpdateCacheMut       sync.RWMutex
	payPerViewEventUpdateCache          = make(map[string]updateCache)
	payPerViewEventUpsertCacheMut       sync.RWMutex
	payPerViewEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single payPerViewEvent record from the query.
func (q payPerViewEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PayPerViewEvent, error) {
	o := &PayPerViewEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pay_per_view_events")
	}

	return o, nil
}

// All returns all PayPerViewEvent records from the query.
func (q payPerViewEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (PayPerViewEventSlice, error) {
	var o []*PayPerViewEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PayPerViewEvent slice")
	}

	return o, nil
}

// Count returns the count of all PayPerViewEvent records in the query.
func (q payPerViewEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pay_per_view_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q payPerViewEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pay_per_view_events exists")
	}

	return count > 0, nil
}

// SmartContracts retrieves all the smart_contract's SmartContracts with an executor.
func (o *PayPerViewEvent) SmartContracts(mods ...qm.QueryMod) smartContractQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"smart_contracts\".\"pay_per_view_event_id\"=?", o.ID),
	)

	query := SmartContracts(queryMods...)
	queries.SetFrom(query.Query, "\"smart_contracts\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"smart_contracts\".*"})
	}

	return query
}

// LoadSmartContracts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (payPerViewEventL) LoadSmartContracts(ctx context.Context, e boil.ContextExecutor, singular bool, maybePayPerViewEvent interface{}, mods queries.Applicator) error {
	var slice []*PayPerViewEvent
	var object *PayPerViewEvent

	if singular {
		object = maybePayPerViewEvent.(*PayPerViewEvent)
	} else {
		slice = *maybePayPerViewEvent.(*[]*PayPerViewEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &payPerViewEventR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &payPerViewEventR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`smart_contracts`), qm.WhereIn(`smart_contracts.pay_per_view_event_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load smart_contracts")
	}

	var resultSlice []*SmartContract
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice smart_contracts")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on smart_contracts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for smart_contracts")
	}

	if singular {
		object.R.SmartContracts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &smartContractR{}
			}
			foreign.R.PayPerViewEvent = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PayPerViewEventID {
				local.R.SmartContracts = append(local.R.SmartContracts, foreign)
				if foreign.R == nil {
					foreign.R = &smartContractR{}
				}
				foreign.R.PayPerViewEvent = local
				break
			}
		}
	}

	return nil
}

// AddSmartContracts adds the given related objects to the existing relationships
// of the pay_per_view_event, optionally inserting them as new records.
// Appends related to o.R.SmartContracts.
// Sets related.R.PayPerViewEvent appropriately.
func (o *PayPerViewEvent) AddSmartContracts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SmartContract) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PayPerViewEventID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"smart_contracts\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pay_per_view_event_id"}),
				strmangle.WhereClause("\"", "\"", 2, smartContractPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PayPerViewEventID = o.ID
		}
	}

	if o.R == nil {
		o.R = &payPerViewEventR{
			SmartContracts: related,
		}
	} else {
		o.R.SmartContracts = append(o.R.SmartContracts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &smartContractR{
				PayPerViewEvent: o,
			}
		} else {
			rel.R.PayPerViewEvent = o
		}
	}
	return nil
}

// PayPerViewEvents retrieves all the records using an executor.
func PayPerViewEvents(mods ...qm.QueryMod) payPerViewEventQuery {
	mods = append(mods, qm.From("\"pay_per_view_events\""))
	return payPerViewEventQuery{NewQuery(mods...)}
}

// FindPayPerViewEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPayPerViewEvent(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PayPerViewEvent, error) {
	payPerViewEventObj := &PayPerViewEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pay_per_view_events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, payPerViewEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pay_per_view_events")
	}

	return payPerViewEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PayPerViewEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pay_per_view_events provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(payPerViewEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	payPerViewEventInsertCacheMut.RLock()
	cache, cached := payPerViewEventInsertCache[key]
	payPerViewEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			payPerViewEventAllColumns,
			payPerViewEventColumnsWithDefault,
			payPerViewEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pay_per_view_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pay_per_view_events\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pay_per_view_events")
	}

	if !cached {
		payPerViewEventInsertCacheMut.Lock()
		payPerViewEventInsertCache[key] = cache
		payPerViewEventInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PayPerViewEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PayPerViewEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	payPerViewEventUpdateCacheMut.RLock()
	cache, cached := payPerViewEventUpdateCache[key]
	payPerViewEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			payPerViewEventAllColumns,
			payPerViewEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pay_per_view_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pay_per_view_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, payPerViewEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, append(wl, payPerViewEventPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update pay_per_view_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pay_per_view_events")
	}

	if !cached {
		payPerViewEventUpdateCacheMut.Lock()
		payPerViewEventUpdateCache[key] = cache
		payPerViewEventUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q payPerViewEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pay_per_view_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pay_per_view_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PayPerViewEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payPerViewEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pay_per_view_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, payPerViewEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in payPerViewEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all payPerViewEvent")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PayPerViewEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pay_per_view_events provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(payPerViewEventColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	payPerViewEventUpsertCacheMut.RLock()
	cache, cached := payPerViewEventUpsertCache[key]
	payPerViewEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			payPerViewEventAllColumns,
			payPerViewEventColumnsWithDefault,
			payPerViewEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			payPerViewEventAllColumns,
			payPerViewEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pay_per_view_events, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(payPerViewEventPrimaryKeyColumns))
			copy(conflict, payPerViewEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pay_per_view_events\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(payPerViewEventType, payPerViewEventMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pay_per_view_events")
	}

	if !cached {
		payPerViewEventUpsertCacheMut.Lock()
		payPerViewEventUpsertCache[key] = cache
		payPerViewEventUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PayPerViewEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PayPerViewEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PayPerViewEvent provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), payPerViewEventPrimaryKeyMapping)
	sql := "DELETE FROM \"pay_per_view_events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pay_per_view_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pay_per_view_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q payPerViewEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no payPerViewEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pay_per_view_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pay_per_view_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PayPerViewEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payPerViewEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pay_per_view_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, payPerViewEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from payPerViewEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pay_per_view_events")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PayPerViewEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPayPerViewEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PayPerViewEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PayPerViewEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), payPerViewEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pay_per_view_events\".* FROM \"pay_per_view_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, payPerViewEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PayPerViewEventSlice")
	}

	*o = slice

	return nil
}

// PayPerViewEventExists checks if the PayPerViewEvent row exists.
func PayPerViewEventExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pay_per_view_events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pay_per_view_events exists")
	}

	return exists, nil
}
