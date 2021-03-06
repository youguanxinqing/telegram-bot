// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"telegram-bot/internal/ent/predicate"
	"telegram-bot/internal/ent/words"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordsUpdate is the builder for updating Words entities.
type WordsUpdate struct {
	config
	hooks    []Hook
	mutation *WordsMutation
}

// Where appends a list predicates to the WordsUpdate builder.
func (wu *WordsUpdate) Where(ps ...predicate.Words) *WordsUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetWord sets the "word" field.
func (wu *WordsUpdate) SetWord(s string) *WordsUpdate {
	wu.mutation.SetWord(s)
	return wu
}

// SetExplain sets the "explain" field.
func (wu *WordsUpdate) SetExplain(s string) *WordsUpdate {
	wu.mutation.SetExplain(s)
	return wu
}

// SetNillableExplain sets the "explain" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableExplain(s *string) *WordsUpdate {
	if s != nil {
		wu.SetExplain(*s)
	}
	return wu
}

// SetPhonetic sets the "phonetic" field.
func (wu *WordsUpdate) SetPhonetic(s string) *WordsUpdate {
	wu.mutation.SetPhonetic(s)
	return wu
}

// SetNillablePhonetic sets the "phonetic" field if the given value is not nil.
func (wu *WordsUpdate) SetNillablePhonetic(s *string) *WordsUpdate {
	if s != nil {
		wu.SetPhonetic(*s)
	}
	return wu
}

// SetFrequency sets the "frequency" field.
func (wu *WordsUpdate) SetFrequency(i int) *WordsUpdate {
	wu.mutation.ResetFrequency()
	wu.mutation.SetFrequency(i)
	return wu
}

// SetNillableFrequency sets the "frequency" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableFrequency(i *int) *WordsUpdate {
	if i != nil {
		wu.SetFrequency(*i)
	}
	return wu
}

// AddFrequency adds i to the "frequency" field.
func (wu *WordsUpdate) AddFrequency(i int) *WordsUpdate {
	wu.mutation.AddFrequency(i)
	return wu
}

// SetLastShowTime sets the "last_show_time" field.
func (wu *WordsUpdate) SetLastShowTime(t time.Time) *WordsUpdate {
	wu.mutation.SetLastShowTime(t)
	return wu
}

// SetNillableLastShowTime sets the "last_show_time" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableLastShowTime(t *time.Time) *WordsUpdate {
	if t != nil {
		wu.SetLastShowTime(*t)
	}
	return wu
}

// SetIsHide sets the "is_hide" field.
func (wu *WordsUpdate) SetIsHide(b bool) *WordsUpdate {
	wu.mutation.SetIsHide(b)
	return wu
}

// SetNillableIsHide sets the "is_hide" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableIsHide(b *bool) *WordsUpdate {
	if b != nil {
		wu.SetIsHide(*b)
	}
	return wu
}

// SetCreateTime sets the "create_time" field.
func (wu *WordsUpdate) SetCreateTime(t time.Time) *WordsUpdate {
	wu.mutation.SetCreateTime(t)
	return wu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableCreateTime(t *time.Time) *WordsUpdate {
	if t != nil {
		wu.SetCreateTime(*t)
	}
	return wu
}

// SetDeleteTime sets the "delete_time" field.
func (wu *WordsUpdate) SetDeleteTime(t time.Time) *WordsUpdate {
	wu.mutation.SetDeleteTime(t)
	return wu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableDeleteTime(t *time.Time) *WordsUpdate {
	if t != nil {
		wu.SetDeleteTime(*t)
	}
	return wu
}

// SetIsDelete sets the "is_delete" field.
func (wu *WordsUpdate) SetIsDelete(b bool) *WordsUpdate {
	wu.mutation.SetIsDelete(b)
	return wu
}

// SetNillableIsDelete sets the "is_delete" field if the given value is not nil.
func (wu *WordsUpdate) SetNillableIsDelete(b *bool) *WordsUpdate {
	if b != nil {
		wu.SetIsDelete(*b)
	}
	return wu
}

// Mutation returns the WordsMutation object of the builder.
func (wu *WordsUpdate) Mutation() *WordsMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WordsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WordsUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WordsUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WordsUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WordsUpdate) check() error {
	if v, ok := wu.mutation.Word(); ok {
		if err := words.WordValidator(v); err != nil {
			return &ValidationError{Name: "word", err: fmt.Errorf(`ent: validator failed for field "Words.word": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Frequency(); ok {
		if err := words.FrequencyValidator(v); err != nil {
			return &ValidationError{Name: "frequency", err: fmt.Errorf(`ent: validator failed for field "Words.frequency": %w`, err)}
		}
	}
	return nil
}

func (wu *WordsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   words.Table,
			Columns: words.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: words.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Word(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldWord,
		})
	}
	if value, ok := wu.mutation.Explain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldExplain,
		})
	}
	if value, ok := wu.mutation.Phonetic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldPhonetic,
		})
	}
	if value, ok := wu.mutation.Frequency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: words.FieldFrequency,
		})
	}
	if value, ok := wu.mutation.AddedFrequency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: words.FieldFrequency,
		})
	}
	if value, ok := wu.mutation.LastShowTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldLastShowTime,
		})
	}
	if value, ok := wu.mutation.IsHide(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: words.FieldIsHide,
		})
	}
	if value, ok := wu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldCreateTime,
		})
	}
	if value, ok := wu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldDeleteTime,
		})
	}
	if value, ok := wu.mutation.IsDelete(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: words.FieldIsDelete,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{words.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WordsUpdateOne is the builder for updating a single Words entity.
type WordsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WordsMutation
}

// SetWord sets the "word" field.
func (wuo *WordsUpdateOne) SetWord(s string) *WordsUpdateOne {
	wuo.mutation.SetWord(s)
	return wuo
}

// SetExplain sets the "explain" field.
func (wuo *WordsUpdateOne) SetExplain(s string) *WordsUpdateOne {
	wuo.mutation.SetExplain(s)
	return wuo
}

// SetNillableExplain sets the "explain" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableExplain(s *string) *WordsUpdateOne {
	if s != nil {
		wuo.SetExplain(*s)
	}
	return wuo
}

// SetPhonetic sets the "phonetic" field.
func (wuo *WordsUpdateOne) SetPhonetic(s string) *WordsUpdateOne {
	wuo.mutation.SetPhonetic(s)
	return wuo
}

// SetNillablePhonetic sets the "phonetic" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillablePhonetic(s *string) *WordsUpdateOne {
	if s != nil {
		wuo.SetPhonetic(*s)
	}
	return wuo
}

// SetFrequency sets the "frequency" field.
func (wuo *WordsUpdateOne) SetFrequency(i int) *WordsUpdateOne {
	wuo.mutation.ResetFrequency()
	wuo.mutation.SetFrequency(i)
	return wuo
}

// SetNillableFrequency sets the "frequency" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableFrequency(i *int) *WordsUpdateOne {
	if i != nil {
		wuo.SetFrequency(*i)
	}
	return wuo
}

// AddFrequency adds i to the "frequency" field.
func (wuo *WordsUpdateOne) AddFrequency(i int) *WordsUpdateOne {
	wuo.mutation.AddFrequency(i)
	return wuo
}

// SetLastShowTime sets the "last_show_time" field.
func (wuo *WordsUpdateOne) SetLastShowTime(t time.Time) *WordsUpdateOne {
	wuo.mutation.SetLastShowTime(t)
	return wuo
}

// SetNillableLastShowTime sets the "last_show_time" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableLastShowTime(t *time.Time) *WordsUpdateOne {
	if t != nil {
		wuo.SetLastShowTime(*t)
	}
	return wuo
}

// SetIsHide sets the "is_hide" field.
func (wuo *WordsUpdateOne) SetIsHide(b bool) *WordsUpdateOne {
	wuo.mutation.SetIsHide(b)
	return wuo
}

// SetNillableIsHide sets the "is_hide" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableIsHide(b *bool) *WordsUpdateOne {
	if b != nil {
		wuo.SetIsHide(*b)
	}
	return wuo
}

// SetCreateTime sets the "create_time" field.
func (wuo *WordsUpdateOne) SetCreateTime(t time.Time) *WordsUpdateOne {
	wuo.mutation.SetCreateTime(t)
	return wuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableCreateTime(t *time.Time) *WordsUpdateOne {
	if t != nil {
		wuo.SetCreateTime(*t)
	}
	return wuo
}

// SetDeleteTime sets the "delete_time" field.
func (wuo *WordsUpdateOne) SetDeleteTime(t time.Time) *WordsUpdateOne {
	wuo.mutation.SetDeleteTime(t)
	return wuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableDeleteTime(t *time.Time) *WordsUpdateOne {
	if t != nil {
		wuo.SetDeleteTime(*t)
	}
	return wuo
}

// SetIsDelete sets the "is_delete" field.
func (wuo *WordsUpdateOne) SetIsDelete(b bool) *WordsUpdateOne {
	wuo.mutation.SetIsDelete(b)
	return wuo
}

// SetNillableIsDelete sets the "is_delete" field if the given value is not nil.
func (wuo *WordsUpdateOne) SetNillableIsDelete(b *bool) *WordsUpdateOne {
	if b != nil {
		wuo.SetIsDelete(*b)
	}
	return wuo
}

// Mutation returns the WordsMutation object of the builder.
func (wuo *WordsUpdateOne) Mutation() *WordsMutation {
	return wuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WordsUpdateOne) Select(field string, fields ...string) *WordsUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Words entity.
func (wuo *WordsUpdateOne) Save(ctx context.Context) (*Words, error) {
	var (
		err  error
		node *Words
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WordsUpdateOne) SaveX(ctx context.Context) *Words {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WordsUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WordsUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WordsUpdateOne) check() error {
	if v, ok := wuo.mutation.Word(); ok {
		if err := words.WordValidator(v); err != nil {
			return &ValidationError{Name: "word", err: fmt.Errorf(`ent: validator failed for field "Words.word": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Frequency(); ok {
		if err := words.FrequencyValidator(v); err != nil {
			return &ValidationError{Name: "frequency", err: fmt.Errorf(`ent: validator failed for field "Words.frequency": %w`, err)}
		}
	}
	return nil
}

func (wuo *WordsUpdateOne) sqlSave(ctx context.Context) (_node *Words, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   words.Table,
			Columns: words.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: words.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Words.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, words.FieldID)
		for _, f := range fields {
			if !words.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != words.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Word(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldWord,
		})
	}
	if value, ok := wuo.mutation.Explain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldExplain,
		})
	}
	if value, ok := wuo.mutation.Phonetic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: words.FieldPhonetic,
		})
	}
	if value, ok := wuo.mutation.Frequency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: words.FieldFrequency,
		})
	}
	if value, ok := wuo.mutation.AddedFrequency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: words.FieldFrequency,
		})
	}
	if value, ok := wuo.mutation.LastShowTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldLastShowTime,
		})
	}
	if value, ok := wuo.mutation.IsHide(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: words.FieldIsHide,
		})
	}
	if value, ok := wuo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldCreateTime,
		})
	}
	if value, ok := wuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: words.FieldDeleteTime,
		})
	}
	if value, ok := wuo.mutation.IsDelete(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: words.FieldIsDelete,
		})
	}
	_node = &Words{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{words.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
