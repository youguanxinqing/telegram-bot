// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"telegram-bot/internal/ent/predicate"
	"telegram-bot/internal/ent/words"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeWords = "Words"
)

// WordsMutation represents an operation that mutates the Words nodes in the graph.
type WordsMutation struct {
	config
	op             Op
	typ            string
	id             *int
	word           *string
	explain        *string
	phonetic       *string
	frequency      *int
	addfrequency   *int
	last_show_time *time.Time
	is_hide        *bool
	create_time    *time.Time
	delete_time    *time.Time
	is_delete      *bool
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Words, error)
	predicates     []predicate.Words
}

var _ ent.Mutation = (*WordsMutation)(nil)

// wordsOption allows management of the mutation configuration using functional options.
type wordsOption func(*WordsMutation)

// newWordsMutation creates new mutation for the Words entity.
func newWordsMutation(c config, op Op, opts ...wordsOption) *WordsMutation {
	m := &WordsMutation{
		config:        c,
		op:            op,
		typ:           TypeWords,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWordsID sets the ID field of the mutation.
func withWordsID(id int) wordsOption {
	return func(m *WordsMutation) {
		var (
			err   error
			once  sync.Once
			value *Words
		)
		m.oldValue = func(ctx context.Context) (*Words, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Words.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWords sets the old Words of the mutation.
func withWords(node *Words) wordsOption {
	return func(m *WordsMutation) {
		m.oldValue = func(context.Context) (*Words, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WordsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WordsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WordsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WordsMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Words.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetWord sets the "word" field.
func (m *WordsMutation) SetWord(s string) {
	m.word = &s
}

// Word returns the value of the "word" field in the mutation.
func (m *WordsMutation) Word() (r string, exists bool) {
	v := m.word
	if v == nil {
		return
	}
	return *v, true
}

// OldWord returns the old "word" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldWord(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWord is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWord requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWord: %w", err)
	}
	return oldValue.Word, nil
}

// ResetWord resets all changes to the "word" field.
func (m *WordsMutation) ResetWord() {
	m.word = nil
}

// SetExplain sets the "explain" field.
func (m *WordsMutation) SetExplain(s string) {
	m.explain = &s
}

// Explain returns the value of the "explain" field in the mutation.
func (m *WordsMutation) Explain() (r string, exists bool) {
	v := m.explain
	if v == nil {
		return
	}
	return *v, true
}

// OldExplain returns the old "explain" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldExplain(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExplain is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExplain requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExplain: %w", err)
	}
	return oldValue.Explain, nil
}

// ResetExplain resets all changes to the "explain" field.
func (m *WordsMutation) ResetExplain() {
	m.explain = nil
}

// SetPhonetic sets the "phonetic" field.
func (m *WordsMutation) SetPhonetic(s string) {
	m.phonetic = &s
}

// Phonetic returns the value of the "phonetic" field in the mutation.
func (m *WordsMutation) Phonetic() (r string, exists bool) {
	v := m.phonetic
	if v == nil {
		return
	}
	return *v, true
}

// OldPhonetic returns the old "phonetic" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldPhonetic(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhonetic is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhonetic requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhonetic: %w", err)
	}
	return oldValue.Phonetic, nil
}

// ResetPhonetic resets all changes to the "phonetic" field.
func (m *WordsMutation) ResetPhonetic() {
	m.phonetic = nil
}

// SetFrequency sets the "frequency" field.
func (m *WordsMutation) SetFrequency(i int) {
	m.frequency = &i
	m.addfrequency = nil
}

// Frequency returns the value of the "frequency" field in the mutation.
func (m *WordsMutation) Frequency() (r int, exists bool) {
	v := m.frequency
	if v == nil {
		return
	}
	return *v, true
}

// OldFrequency returns the old "frequency" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldFrequency(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFrequency is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFrequency requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrequency: %w", err)
	}
	return oldValue.Frequency, nil
}

// AddFrequency adds i to the "frequency" field.
func (m *WordsMutation) AddFrequency(i int) {
	if m.addfrequency != nil {
		*m.addfrequency += i
	} else {
		m.addfrequency = &i
	}
}

// AddedFrequency returns the value that was added to the "frequency" field in this mutation.
func (m *WordsMutation) AddedFrequency() (r int, exists bool) {
	v := m.addfrequency
	if v == nil {
		return
	}
	return *v, true
}

// ResetFrequency resets all changes to the "frequency" field.
func (m *WordsMutation) ResetFrequency() {
	m.frequency = nil
	m.addfrequency = nil
}

// SetLastShowTime sets the "last_show_time" field.
func (m *WordsMutation) SetLastShowTime(t time.Time) {
	m.last_show_time = &t
}

// LastShowTime returns the value of the "last_show_time" field in the mutation.
func (m *WordsMutation) LastShowTime() (r time.Time, exists bool) {
	v := m.last_show_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastShowTime returns the old "last_show_time" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldLastShowTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastShowTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastShowTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastShowTime: %w", err)
	}
	return oldValue.LastShowTime, nil
}

// ResetLastShowTime resets all changes to the "last_show_time" field.
func (m *WordsMutation) ResetLastShowTime() {
	m.last_show_time = nil
}

// SetIsHide sets the "is_hide" field.
func (m *WordsMutation) SetIsHide(b bool) {
	m.is_hide = &b
}

// IsHide returns the value of the "is_hide" field in the mutation.
func (m *WordsMutation) IsHide() (r bool, exists bool) {
	v := m.is_hide
	if v == nil {
		return
	}
	return *v, true
}

// OldIsHide returns the old "is_hide" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldIsHide(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsHide is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsHide requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsHide: %w", err)
	}
	return oldValue.IsHide, nil
}

// ResetIsHide resets all changes to the "is_hide" field.
func (m *WordsMutation) ResetIsHide() {
	m.is_hide = nil
}

// SetCreateTime sets the "create_time" field.
func (m *WordsMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *WordsMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *WordsMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetDeleteTime sets the "delete_time" field.
func (m *WordsMutation) SetDeleteTime(t time.Time) {
	m.delete_time = &t
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *WordsMutation) DeleteTime() (r time.Time, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldDeleteTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *WordsMutation) ResetDeleteTime() {
	m.delete_time = nil
}

// SetIsDelete sets the "is_delete" field.
func (m *WordsMutation) SetIsDelete(b bool) {
	m.is_delete = &b
}

// IsDelete returns the value of the "is_delete" field in the mutation.
func (m *WordsMutation) IsDelete() (r bool, exists bool) {
	v := m.is_delete
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDelete returns the old "is_delete" field's value of the Words entity.
// If the Words object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WordsMutation) OldIsDelete(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDelete is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDelete requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDelete: %w", err)
	}
	return oldValue.IsDelete, nil
}

// ResetIsDelete resets all changes to the "is_delete" field.
func (m *WordsMutation) ResetIsDelete() {
	m.is_delete = nil
}

// Where appends a list predicates to the WordsMutation builder.
func (m *WordsMutation) Where(ps ...predicate.Words) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *WordsMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Words).
func (m *WordsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WordsMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.word != nil {
		fields = append(fields, words.FieldWord)
	}
	if m.explain != nil {
		fields = append(fields, words.FieldExplain)
	}
	if m.phonetic != nil {
		fields = append(fields, words.FieldPhonetic)
	}
	if m.frequency != nil {
		fields = append(fields, words.FieldFrequency)
	}
	if m.last_show_time != nil {
		fields = append(fields, words.FieldLastShowTime)
	}
	if m.is_hide != nil {
		fields = append(fields, words.FieldIsHide)
	}
	if m.create_time != nil {
		fields = append(fields, words.FieldCreateTime)
	}
	if m.delete_time != nil {
		fields = append(fields, words.FieldDeleteTime)
	}
	if m.is_delete != nil {
		fields = append(fields, words.FieldIsDelete)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WordsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case words.FieldWord:
		return m.Word()
	case words.FieldExplain:
		return m.Explain()
	case words.FieldPhonetic:
		return m.Phonetic()
	case words.FieldFrequency:
		return m.Frequency()
	case words.FieldLastShowTime:
		return m.LastShowTime()
	case words.FieldIsHide:
		return m.IsHide()
	case words.FieldCreateTime:
		return m.CreateTime()
	case words.FieldDeleteTime:
		return m.DeleteTime()
	case words.FieldIsDelete:
		return m.IsDelete()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WordsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case words.FieldWord:
		return m.OldWord(ctx)
	case words.FieldExplain:
		return m.OldExplain(ctx)
	case words.FieldPhonetic:
		return m.OldPhonetic(ctx)
	case words.FieldFrequency:
		return m.OldFrequency(ctx)
	case words.FieldLastShowTime:
		return m.OldLastShowTime(ctx)
	case words.FieldIsHide:
		return m.OldIsHide(ctx)
	case words.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case words.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	case words.FieldIsDelete:
		return m.OldIsDelete(ctx)
	}
	return nil, fmt.Errorf("unknown Words field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WordsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case words.FieldWord:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWord(v)
		return nil
	case words.FieldExplain:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExplain(v)
		return nil
	case words.FieldPhonetic:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhonetic(v)
		return nil
	case words.FieldFrequency:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrequency(v)
		return nil
	case words.FieldLastShowTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastShowTime(v)
		return nil
	case words.FieldIsHide:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsHide(v)
		return nil
	case words.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case words.FieldDeleteTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	case words.FieldIsDelete:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDelete(v)
		return nil
	}
	return fmt.Errorf("unknown Words field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WordsMutation) AddedFields() []string {
	var fields []string
	if m.addfrequency != nil {
		fields = append(fields, words.FieldFrequency)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WordsMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case words.FieldFrequency:
		return m.AddedFrequency()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WordsMutation) AddField(name string, value ent.Value) error {
	switch name {
	case words.FieldFrequency:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFrequency(v)
		return nil
	}
	return fmt.Errorf("unknown Words numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WordsMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WordsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WordsMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Words nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WordsMutation) ResetField(name string) error {
	switch name {
	case words.FieldWord:
		m.ResetWord()
		return nil
	case words.FieldExplain:
		m.ResetExplain()
		return nil
	case words.FieldPhonetic:
		m.ResetPhonetic()
		return nil
	case words.FieldFrequency:
		m.ResetFrequency()
		return nil
	case words.FieldLastShowTime:
		m.ResetLastShowTime()
		return nil
	case words.FieldIsHide:
		m.ResetIsHide()
		return nil
	case words.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case words.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	case words.FieldIsDelete:
		m.ResetIsDelete()
		return nil
	}
	return fmt.Errorf("unknown Words field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WordsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WordsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WordsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WordsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WordsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WordsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WordsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Words unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WordsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Words edge %s", name)
}