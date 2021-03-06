// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
)

// Street is the model entity for the Street schema.
type Street struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StreetQuery when eager-loading is set.
	Edges        StreetEdges `json:"edges"`
	city_streets *int
}

// StreetEdges holds the relations/edges for other nodes in the graph.
type StreetEdges struct {
	// City holds the value of the city edge.
	City *City
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CityOrErr returns the City value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StreetEdges) CityOrErr() (*City, error) {
	if e.loadedTypes[0] {
		if e.City == nil {
			// The edge city was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: city.Label}
		}
		return e.City, nil
	}
	return nil, &NotLoadedError{edge: "city"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Street) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Street) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // city_streets
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Street fields.
func (s *Street) assignValues(values ...interface{}) error {
	if m, n := len(values), len(street.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		s.Name = value.String
	}
	values = values[1:]
	if len(values) == len(street.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field city_streets", value)
		} else if value.Valid {
			s.city_streets = new(int)
			*s.city_streets = int(value.Int64)
		}
	}
	return nil
}

// QueryCity queries the city edge of the Street.
func (s *Street) QueryCity() *CityQuery {
	return (&StreetClient{config: s.config}).QueryCity(s)
}

// Update returns a builder for updating this Street.
// Note that, you need to call Street.Unwrap() before calling this method, if this Street
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Street) Update() *StreetUpdateOne {
	return (&StreetClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Street) Unwrap() *Street {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Street is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Street) String() string {
	var builder strings.Builder
	builder.WriteString("Street(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Streets is a parsable slice of Street.
type Streets []*Street

func (s Streets) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
