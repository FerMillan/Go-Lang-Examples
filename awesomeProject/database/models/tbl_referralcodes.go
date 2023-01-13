// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TBLReferralcode is an object representing the database table.
type TBLReferralcode struct {
	UserID       int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Referralcode string      `boil:"referralcode" json:"referralcode" toml:"referralcode" yaml:"referralcode"`
	Timesused    int         `boil:"timesused" json:"timesused" toml:"timesused" yaml:"timesused"`
	Source       null.String `boil:"source" json:"source,omitempty" toml:"source" yaml:"source,omitempty"`

	R *tblReferralcodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tblReferralcodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TBLReferralcodeColumns = struct {
	UserID       string
	Referralcode string
	Timesused    string
	Source       string
}{
	UserID:       "user_id",
	Referralcode: "referralcode",
	Timesused:    "timesused",
	Source:       "source",
}

var TBLReferralcodeTableColumns = struct {
	UserID       string
	Referralcode string
	Timesused    string
	Source       string
}{
	UserID:       "tbl_referralcodes.user_id",
	Referralcode: "tbl_referralcodes.referralcode",
	Timesused:    "tbl_referralcodes.timesused",
	Source:       "tbl_referralcodes.source",
}

// Generated where

var TBLReferralcodeWhere = struct {
	UserID       whereHelperint
	Referralcode whereHelperstring
	Timesused    whereHelperint
	Source       whereHelpernull_String
}{
	UserID:       whereHelperint{field: "`tbl_referralcodes`.`user_id`"},
	Referralcode: whereHelperstring{field: "`tbl_referralcodes`.`referralcode`"},
	Timesused:    whereHelperint{field: "`tbl_referralcodes`.`timesused`"},
	Source:       whereHelpernull_String{field: "`tbl_referralcodes`.`source`"},
}

// TBLReferralcodeRels is where relationship names are stored.
var TBLReferralcodeRels = struct {
}{}

// tblReferralcodeR is where relationships are stored.
type tblReferralcodeR struct {
}

// NewStruct creates a new relationship struct
func (*tblReferralcodeR) NewStruct() *tblReferralcodeR {
	return &tblReferralcodeR{}
}

// tblReferralcodeL is where Load methods for each relationship are stored.
type tblReferralcodeL struct{}

var (
	tblReferralcodeAllColumns            = []string{"user_id", "referralcode", "timesused", "source"}
	tblReferralcodeColumnsWithoutDefault = []string{"user_id", "referralcode", "timesused", "source"}
	tblReferralcodeColumnsWithDefault    = []string{}
	tblReferralcodePrimaryKeyColumns     = []string{"user_id"}
	tblReferralcodeGeneratedColumns      = []string{}
)

type (
	// TBLReferralcodeSlice is an alias for a slice of pointers to TBLReferralcode.
	// This should almost always be used instead of []TBLReferralcode.
	TBLReferralcodeSlice []*TBLReferralcode
	// TBLReferralcodeHook is the signature for custom TBLReferralcode hook methods
	TBLReferralcodeHook func(context.Context, boil.ContextExecutor, *TBLReferralcode) error

	tblReferralcodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tblReferralcodeType                 = reflect.TypeOf(&TBLReferralcode{})
	tblReferralcodeMapping              = queries.MakeStructMapping(tblReferralcodeType)
	tblReferralcodePrimaryKeyMapping, _ = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, tblReferralcodePrimaryKeyColumns)
	tblReferralcodeInsertCacheMut       sync.RWMutex
	tblReferralcodeInsertCache          = make(map[string]insertCache)
	tblReferralcodeUpdateCacheMut       sync.RWMutex
	tblReferralcodeUpdateCache          = make(map[string]updateCache)
	tblReferralcodeUpsertCacheMut       sync.RWMutex
	tblReferralcodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tblReferralcodeAfterSelectHooks []TBLReferralcodeHook

var tblReferralcodeBeforeInsertHooks []TBLReferralcodeHook
var tblReferralcodeAfterInsertHooks []TBLReferralcodeHook

var tblReferralcodeBeforeUpdateHooks []TBLReferralcodeHook
var tblReferralcodeAfterUpdateHooks []TBLReferralcodeHook

var tblReferralcodeBeforeDeleteHooks []TBLReferralcodeHook
var tblReferralcodeAfterDeleteHooks []TBLReferralcodeHook

var tblReferralcodeBeforeUpsertHooks []TBLReferralcodeHook
var tblReferralcodeAfterUpsertHooks []TBLReferralcodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TBLReferralcode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TBLReferralcode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TBLReferralcode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TBLReferralcode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TBLReferralcode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TBLReferralcode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TBLReferralcode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TBLReferralcode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TBLReferralcode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tblReferralcodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTBLReferralcodeHook registers your hook function for all future operations.
func AddTBLReferralcodeHook(hookPoint boil.HookPoint, tblReferralcodeHook TBLReferralcodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tblReferralcodeAfterSelectHooks = append(tblReferralcodeAfterSelectHooks, tblReferralcodeHook)
	case boil.BeforeInsertHook:
		tblReferralcodeBeforeInsertHooks = append(tblReferralcodeBeforeInsertHooks, tblReferralcodeHook)
	case boil.AfterInsertHook:
		tblReferralcodeAfterInsertHooks = append(tblReferralcodeAfterInsertHooks, tblReferralcodeHook)
	case boil.BeforeUpdateHook:
		tblReferralcodeBeforeUpdateHooks = append(tblReferralcodeBeforeUpdateHooks, tblReferralcodeHook)
	case boil.AfterUpdateHook:
		tblReferralcodeAfterUpdateHooks = append(tblReferralcodeAfterUpdateHooks, tblReferralcodeHook)
	case boil.BeforeDeleteHook:
		tblReferralcodeBeforeDeleteHooks = append(tblReferralcodeBeforeDeleteHooks, tblReferralcodeHook)
	case boil.AfterDeleteHook:
		tblReferralcodeAfterDeleteHooks = append(tblReferralcodeAfterDeleteHooks, tblReferralcodeHook)
	case boil.BeforeUpsertHook:
		tblReferralcodeBeforeUpsertHooks = append(tblReferralcodeBeforeUpsertHooks, tblReferralcodeHook)
	case boil.AfterUpsertHook:
		tblReferralcodeAfterUpsertHooks = append(tblReferralcodeAfterUpsertHooks, tblReferralcodeHook)
	}
}

// One returns a single tblReferralcode record from the query.
func (q tblReferralcodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TBLReferralcode, error) {
	o := &TBLReferralcode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tbl_referralcodes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TBLReferralcode records from the query.
func (q tblReferralcodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (TBLReferralcodeSlice, error) {
	var o []*TBLReferralcode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TBLReferralcode slice")
	}

	if len(tblReferralcodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TBLReferralcode records in the query.
func (q tblReferralcodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tbl_referralcodes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tblReferralcodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tbl_referralcodes exists")
	}

	return count > 0, nil
}

// TBLReferralcodes retrieves all the records using an executor.
func TBLReferralcodes(mods ...qm.QueryMod) tblReferralcodeQuery {
	mods = append(mods, qm.From("`tbl_referralcodes`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tbl_referralcodes`.*"})
	}

	return tblReferralcodeQuery{q}
}

// FindTBLReferralcode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTBLReferralcode(ctx context.Context, exec boil.ContextExecutor, userID int, selectCols ...string) (*TBLReferralcode, error) {
	tblReferralcodeObj := &TBLReferralcode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tbl_referralcodes` where `user_id`=?", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, tblReferralcodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tbl_referralcodes")
	}

	if err = tblReferralcodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tblReferralcodeObj, err
	}

	return tblReferralcodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TBLReferralcode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tbl_referralcodes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tblReferralcodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tblReferralcodeInsertCacheMut.RLock()
	cache, cached := tblReferralcodeInsertCache[key]
	tblReferralcodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tblReferralcodeAllColumns,
			tblReferralcodeColumnsWithDefault,
			tblReferralcodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tbl_referralcodes` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tbl_referralcodes` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tbl_referralcodes` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tblReferralcodePrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tbl_referralcodes")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UserID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for tbl_referralcodes")
	}

CacheNoHooks:
	if !cached {
		tblReferralcodeInsertCacheMut.Lock()
		tblReferralcodeInsertCache[key] = cache
		tblReferralcodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TBLReferralcode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TBLReferralcode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tblReferralcodeUpdateCacheMut.RLock()
	cache, cached := tblReferralcodeUpdateCache[key]
	tblReferralcodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tblReferralcodeAllColumns,
			tblReferralcodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tbl_referralcodes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tbl_referralcodes` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tblReferralcodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, append(wl, tblReferralcodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tbl_referralcodes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tbl_referralcodes")
	}

	if !cached {
		tblReferralcodeUpdateCacheMut.Lock()
		tblReferralcodeUpdateCache[key] = cache
		tblReferralcodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tblReferralcodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tbl_referralcodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tbl_referralcodes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TBLReferralcodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tblReferralcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tbl_referralcodes` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tblReferralcodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tblReferralcode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tblReferralcode")
	}
	return rowsAff, nil
}

var mySQLTBLReferralcodeUniqueColumns = []string{
	"user_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TBLReferralcode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tbl_referralcodes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tblReferralcodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTBLReferralcodeUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tblReferralcodeUpsertCacheMut.RLock()
	cache, cached := tblReferralcodeUpsertCache[key]
	tblReferralcodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tblReferralcodeAllColumns,
			tblReferralcodeColumnsWithDefault,
			tblReferralcodeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tblReferralcodeAllColumns,
			tblReferralcodePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert tbl_referralcodes, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`tbl_referralcodes`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tbl_referralcodes` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for tbl_referralcodes")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tblReferralcodeType, tblReferralcodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for tbl_referralcodes")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for tbl_referralcodes")
	}

CacheNoHooks:
	if !cached {
		tblReferralcodeUpsertCacheMut.Lock()
		tblReferralcodeUpsertCache[key] = cache
		tblReferralcodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TBLReferralcode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TBLReferralcode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TBLReferralcode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tblReferralcodePrimaryKeyMapping)
	sql := "DELETE FROM `tbl_referralcodes` WHERE `user_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tbl_referralcodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tbl_referralcodes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tblReferralcodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tblReferralcodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tbl_referralcodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tbl_referralcodes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TBLReferralcodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tblReferralcodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tblReferralcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tbl_referralcodes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tblReferralcodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tblReferralcode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tbl_referralcodes")
	}

	if len(tblReferralcodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TBLReferralcode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTBLReferralcode(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TBLReferralcodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TBLReferralcodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tblReferralcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tbl_referralcodes`.* FROM `tbl_referralcodes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tblReferralcodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TBLReferralcodeSlice")
	}

	*o = slice

	return nil
}

// TBLReferralcodeExists checks if the TBLReferralcode row exists.
func TBLReferralcodeExists(ctx context.Context, exec boil.ContextExecutor, userID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tbl_referralcodes` where `user_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tbl_referralcodes exists")
	}

	return exists, nil
}

// Exists checks if the TBLReferralcode row exists.
func (o *TBLReferralcode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TBLReferralcodeExists(ctx, exec, o.UserID)
}
