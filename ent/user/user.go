// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeAllowedDocuments holds the string denoting the allowed_documents edge name in mutations.
	EdgeAllowedDocuments = "allowed_documents"
	// EdgeOwnedDocuments holds the string denoting the owned_documents edge name in mutations.
	EdgeOwnedDocuments = "owned_documents"
	// EdgeEditedDocuments holds the string denoting the edited_documents edge name in mutations.
	EdgeEditedDocuments = "edited_documents"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AllowedDocumentsTable is the table that holds the allowed_documents relation/edge. The primary key declared below.
	AllowedDocumentsTable = "user_allowed_documents"
	// AllowedDocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	AllowedDocumentsInverseTable = "documents"
	// OwnedDocumentsTable is the table that holds the owned_documents relation/edge.
	OwnedDocumentsTable = "documents"
	// OwnedDocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	OwnedDocumentsInverseTable = "documents"
	// OwnedDocumentsColumn is the table column denoting the owned_documents relation/edge.
	OwnedDocumentsColumn = "user_owned_documents"
	// EditedDocumentsTable is the table that holds the edited_documents relation/edge. The primary key declared below.
	EditedDocumentsTable = "document_editors"
	// EditedDocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	EditedDocumentsInverseTable = "documents"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
}

var (
	// AllowedDocumentsPrimaryKey and AllowedDocumentsColumn2 are the table columns denoting the
	// primary key for the allowed_documents relation (M2M).
	AllowedDocumentsPrimaryKey = []string{"user_id", "document_id"}
	// EditedDocumentsPrimaryKey and EditedDocumentsColumn2 are the table columns denoting the
	// primary key for the edited_documents relation (M2M).
	EditedDocumentsPrimaryKey = []string{"document_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAllowedDocumentsCount orders the results by allowed_documents count.
func ByAllowedDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAllowedDocumentsStep(), opts...)
	}
}

// ByAllowedDocuments orders the results by allowed_documents terms.
func ByAllowedDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAllowedDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnedDocumentsCount orders the results by owned_documents count.
func ByOwnedDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnedDocumentsStep(), opts...)
	}
}

// ByOwnedDocuments orders the results by owned_documents terms.
func ByOwnedDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnedDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEditedDocumentsCount orders the results by edited_documents count.
func ByEditedDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEditedDocumentsStep(), opts...)
	}
}

// ByEditedDocuments orders the results by edited_documents terms.
func ByEditedDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEditedDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAllowedDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AllowedDocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AllowedDocumentsTable, AllowedDocumentsPrimaryKey...),
	)
}
func newOwnedDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnedDocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OwnedDocumentsTable, OwnedDocumentsColumn),
	)
}
func newEditedDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EditedDocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, EditedDocumentsTable, EditedDocumentsPrimaryKey...),
	)
}
