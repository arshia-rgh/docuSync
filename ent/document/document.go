// Code generated by ent, DO NOT EDIT.

package document

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the document type in the database.
	Label = "document"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeEditors holds the string denoting the editors edge name in mutations.
	EdgeEditors = "editors"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeAllowedUsers holds the string denoting the allowed_users edge name in mutations.
	EdgeAllowedUsers = "allowed_users"
	// Table holds the table name of the document in the database.
	Table = "documents"
	// EditorsTable is the table that holds the editors relation/edge. The primary key declared below.
	EditorsTable = "document_editors"
	// EditorsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	EditorsInverseTable = "users"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "documents"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_owned_documents"
	// AllowedUsersTable is the table that holds the allowed_users relation/edge. The primary key declared below.
	AllowedUsersTable = "user_allowed_documents"
	// AllowedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AllowedUsersInverseTable = "users"
)

// Columns holds all SQL columns for document fields.
var Columns = []string{
	FieldID,
	FieldTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "documents"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_owned_documents",
}

var (
	// EditorsPrimaryKey and EditorsColumn2 are the table columns denoting the
	// primary key for the editors relation (M2M).
	EditorsPrimaryKey = []string{"document_id", "user_id"}
	// AllowedUsersPrimaryKey and AllowedUsersColumn2 are the table columns denoting the
	// primary key for the allowed_users relation (M2M).
	AllowedUsersPrimaryKey = []string{"user_id", "document_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "docuSync/ent/runtime"
var (
	Hooks [1]ent.Hook
)

// OrderOption defines the ordering options for the Document queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByEditorsCount orders the results by editors count.
func ByEditorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEditorsStep(), opts...)
	}
}

// ByEditors orders the results by editors terms.
func ByEditors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEditorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByAllowedUsersCount orders the results by allowed_users count.
func ByAllowedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAllowedUsersStep(), opts...)
	}
}

// ByAllowedUsers orders the results by allowed_users terms.
func ByAllowedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAllowedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEditorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EditorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EditorsTable, EditorsPrimaryKey...),
	)
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newAllowedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AllowedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AllowedUsersTable, AllowedUsersPrimaryKey...),
	)
}
