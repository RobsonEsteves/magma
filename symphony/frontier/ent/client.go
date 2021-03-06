// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/symphony/frontier/ent/migrate"

	"github.com/facebookincubator/symphony/frontier/ent/auditlog"
	"github.com/facebookincubator/symphony/frontier/ent/tenant"
	"github.com/facebookincubator/symphony/frontier/ent/token"
	"github.com/facebookincubator/symphony/frontier/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuditLog is the client for interacting with the AuditLog builders.
	AuditLog *AuditLogClient
	// Tenant is the client for interacting with the Tenant builders.
	Tenant *TenantClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:   c,
		Schema:   migrate.NewSchema(c.driver),
		AuditLog: NewAuditLogClient(c),
		Tenant:   NewTenantClient(c),
		Token:    NewTokenClient(c),
		User:     NewUserClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:   cfg,
		AuditLog: NewAuditLogClient(cfg),
		Tenant:   NewTenantClient(cfg),
		Token:    NewTokenClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuditLog.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:   cfg,
		Schema:   migrate.NewSchema(cfg.driver),
		AuditLog: NewAuditLogClient(cfg),
		Tenant:   NewTenantClient(cfg),
		Token:    NewTokenClient(cfg),
		User:     NewUserClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// AuditLogClient is a client for the AuditLog schema.
type AuditLogClient struct {
	config
}

// NewAuditLogClient returns a client for the AuditLog from the given config.
func NewAuditLogClient(c config) *AuditLogClient {
	return &AuditLogClient{config: c}
}

// Create returns a create builder for AuditLog.
func (c *AuditLogClient) Create() *AuditLogCreate {
	return &AuditLogCreate{config: c.config}
}

// Update returns an update builder for AuditLog.
func (c *AuditLogClient) Update() *AuditLogUpdate {
	return &AuditLogUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuditLogClient) UpdateOne(al *AuditLog) *AuditLogUpdateOne {
	return c.UpdateOneID(al.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *AuditLogClient) UpdateOneID(id int) *AuditLogUpdateOne {
	return &AuditLogUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for AuditLog.
func (c *AuditLogClient) Delete() *AuditLogDelete {
	return &AuditLogDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuditLogClient) DeleteOne(al *AuditLog) *AuditLogDeleteOne {
	return c.DeleteOneID(al.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuditLogClient) DeleteOneID(id int) *AuditLogDeleteOne {
	return &AuditLogDeleteOne{c.Delete().Where(auditlog.ID(id))}
}

// Create returns a query builder for AuditLog.
func (c *AuditLogClient) Query() *AuditLogQuery {
	return &AuditLogQuery{config: c.config}
}

// Get returns a AuditLog entity by its id.
func (c *AuditLogClient) Get(ctx context.Context, id int) (*AuditLog, error) {
	return c.Query().Where(auditlog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuditLogClient) GetX(ctx context.Context, id int) *AuditLog {
	al, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return al
}

// TenantClient is a client for the Tenant schema.
type TenantClient struct {
	config
}

// NewTenantClient returns a client for the Tenant from the given config.
func NewTenantClient(c config) *TenantClient {
	return &TenantClient{config: c}
}

// Create returns a create builder for Tenant.
func (c *TenantClient) Create() *TenantCreate {
	return &TenantCreate{config: c.config}
}

// Update returns an update builder for Tenant.
func (c *TenantClient) Update() *TenantUpdate {
	return &TenantUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TenantClient) UpdateOne(t *Tenant) *TenantUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TenantClient) UpdateOneID(id int) *TenantUpdateOne {
	return &TenantUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Tenant.
func (c *TenantClient) Delete() *TenantDelete {
	return &TenantDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TenantClient) DeleteOne(t *Tenant) *TenantDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TenantClient) DeleteOneID(id int) *TenantDeleteOne {
	return &TenantDeleteOne{c.Delete().Where(tenant.ID(id))}
}

// Create returns a query builder for Tenant.
func (c *TenantClient) Query() *TenantQuery {
	return &TenantQuery{config: c.config}
}

// Get returns a Tenant entity by its id.
func (c *TenantClient) Get(ctx context.Context, id int) (*Tenant, error) {
	return c.Query().Where(tenant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TenantClient) GetX(ctx context.Context, id int) *Tenant {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Create returns a create builder for Token.
func (c *TokenClient) Create() *TokenCreate {
	return &TokenCreate{config: c.config}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	return &TokenUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id int) *TokenUpdateOne {
	return &TokenUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	return &TokenDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TokenClient) DeleteOneID(id int) *TokenDeleteOne {
	return &TokenDeleteOne{c.Delete().Where(token.ID(id))}
}

// Create returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{config: c.config}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id int) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id int) *Token {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryUser queries the user edge of a Token.
func (c *TokenClient) QueryUser(t *Token) *UserQuery {
	query := &UserQuery{config: c.config}
	id := t.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(token.Table, token.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, token.UserTable, token.UserColumn),
	)
	query.sql = sqlgraph.Neighbors(t.driver.Dialect(), step)

	return query
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	return &UserCreate{config: c.config}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	return &UserUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	return &UserDeleteOne{c.Delete().Where(user.ID(id))}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryTokens queries the tokens edge of a User.
func (c *UserClient) QueryTokens(u *User) *TokenQuery {
	query := &TokenQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(token.Table, token.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, user.TokensTable, user.TokensColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}
