// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"telegram-bot/internal/ent/migrate"

	"telegram-bot/internal/ent/stardict"
	"telegram-bot/internal/ent/words"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// StarDict is the client for interacting with the StarDict builders.
	StarDict *StarDictClient
	// Words is the client for interacting with the Words builders.
	Words *WordsClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.StarDict = NewStarDictClient(c.config)
	c.Words = NewWordsClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		StarDict: NewStarDictClient(cfg),
		Words:    NewWordsClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		StarDict: NewStarDictClient(cfg),
		Words:    NewWordsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		StarDict.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.StarDict.Use(hooks...)
	c.Words.Use(hooks...)
}

// StarDictClient is a client for the StarDict schema.
type StarDictClient struct {
	config
}

// NewStarDictClient returns a client for the StarDict from the given config.
func NewStarDictClient(c config) *StarDictClient {
	return &StarDictClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stardict.Hooks(f(g(h())))`.
func (c *StarDictClient) Use(hooks ...Hook) {
	c.hooks.StarDict = append(c.hooks.StarDict, hooks...)
}

// Create returns a create builder for StarDict.
func (c *StarDictClient) Create() *StarDictCreate {
	mutation := newStarDictMutation(c.config, OpCreate)
	return &StarDictCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StarDict entities.
func (c *StarDictClient) CreateBulk(builders ...*StarDictCreate) *StarDictCreateBulk {
	return &StarDictCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StarDict.
func (c *StarDictClient) Update() *StarDictUpdate {
	mutation := newStarDictMutation(c.config, OpUpdate)
	return &StarDictUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StarDictClient) UpdateOne(sd *StarDict) *StarDictUpdateOne {
	mutation := newStarDictMutation(c.config, OpUpdateOne, withStarDict(sd))
	return &StarDictUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StarDictClient) UpdateOneID(id int) *StarDictUpdateOne {
	mutation := newStarDictMutation(c.config, OpUpdateOne, withStarDictID(id))
	return &StarDictUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StarDict.
func (c *StarDictClient) Delete() *StarDictDelete {
	mutation := newStarDictMutation(c.config, OpDelete)
	return &StarDictDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StarDictClient) DeleteOne(sd *StarDict) *StarDictDeleteOne {
	return c.DeleteOneID(sd.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StarDictClient) DeleteOneID(id int) *StarDictDeleteOne {
	builder := c.Delete().Where(stardict.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StarDictDeleteOne{builder}
}

// Query returns a query builder for StarDict.
func (c *StarDictClient) Query() *StarDictQuery {
	return &StarDictQuery{
		config: c.config,
	}
}

// Get returns a StarDict entity by its id.
func (c *StarDictClient) Get(ctx context.Context, id int) (*StarDict, error) {
	return c.Query().Where(stardict.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StarDictClient) GetX(ctx context.Context, id int) *StarDict {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StarDictClient) Hooks() []Hook {
	return c.hooks.StarDict
}

// WordsClient is a client for the Words schema.
type WordsClient struct {
	config
}

// NewWordsClient returns a client for the Words from the given config.
func NewWordsClient(c config) *WordsClient {
	return &WordsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `words.Hooks(f(g(h())))`.
func (c *WordsClient) Use(hooks ...Hook) {
	c.hooks.Words = append(c.hooks.Words, hooks...)
}

// Create returns a create builder for Words.
func (c *WordsClient) Create() *WordsCreate {
	mutation := newWordsMutation(c.config, OpCreate)
	return &WordsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Words entities.
func (c *WordsClient) CreateBulk(builders ...*WordsCreate) *WordsCreateBulk {
	return &WordsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Words.
func (c *WordsClient) Update() *WordsUpdate {
	mutation := newWordsMutation(c.config, OpUpdate)
	return &WordsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WordsClient) UpdateOne(w *Words) *WordsUpdateOne {
	mutation := newWordsMutation(c.config, OpUpdateOne, withWords(w))
	return &WordsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WordsClient) UpdateOneID(id int) *WordsUpdateOne {
	mutation := newWordsMutation(c.config, OpUpdateOne, withWordsID(id))
	return &WordsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Words.
func (c *WordsClient) Delete() *WordsDelete {
	mutation := newWordsMutation(c.config, OpDelete)
	return &WordsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WordsClient) DeleteOne(w *Words) *WordsDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WordsClient) DeleteOneID(id int) *WordsDeleteOne {
	builder := c.Delete().Where(words.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WordsDeleteOne{builder}
}

// Query returns a query builder for Words.
func (c *WordsClient) Query() *WordsQuery {
	return &WordsQuery{
		config: c.config,
	}
}

// Get returns a Words entity by its id.
func (c *WordsClient) Get(ctx context.Context, id int) (*Words, error) {
	return c.Query().Where(words.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WordsClient) GetX(ctx context.Context, id int) *Words {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WordsClient) Hooks() []Hook {
	return c.hooks.Words
}
