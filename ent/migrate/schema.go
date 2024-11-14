// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DocumentsColumns holds the columns for the "documents" table.
	DocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "user_owned_documents", Type: field.TypeInt, Nullable: true},
	}
	// DocumentsTable holds the schema information for the "documents" table.
	DocumentsTable = &schema.Table{
		Name:       "documents",
		Columns:    DocumentsColumns,
		PrimaryKey: []*schema.Column{DocumentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "documents_users_owned_documents",
				Columns:    []*schema.Column{DocumentsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// DocumentEditorsColumns holds the columns for the "document_editors" table.
	DocumentEditorsColumns = []*schema.Column{
		{Name: "document_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// DocumentEditorsTable holds the schema information for the "document_editors" table.
	DocumentEditorsTable = &schema.Table{
		Name:       "document_editors",
		Columns:    DocumentEditorsColumns,
		PrimaryKey: []*schema.Column{DocumentEditorsColumns[0], DocumentEditorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "document_editors_document_id",
				Columns:    []*schema.Column{DocumentEditorsColumns[0]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "document_editors_user_id",
				Columns:    []*schema.Column{DocumentEditorsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserAllowedDocumentsColumns holds the columns for the "user_allowed_documents" table.
	UserAllowedDocumentsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "document_id", Type: field.TypeInt},
	}
	// UserAllowedDocumentsTable holds the schema information for the "user_allowed_documents" table.
	UserAllowedDocumentsTable = &schema.Table{
		Name:       "user_allowed_documents",
		Columns:    UserAllowedDocumentsColumns,
		PrimaryKey: []*schema.Column{UserAllowedDocumentsColumns[0], UserAllowedDocumentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_allowed_documents_user_id",
				Columns:    []*schema.Column{UserAllowedDocumentsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_allowed_documents_document_id",
				Columns:    []*schema.Column{UserAllowedDocumentsColumns[1]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DocumentsTable,
		UsersTable,
		DocumentEditorsTable,
		UserAllowedDocumentsTable,
	}
)

func init() {
	DocumentsTable.ForeignKeys[0].RefTable = UsersTable
	DocumentEditorsTable.ForeignKeys[0].RefTable = DocumentsTable
	DocumentEditorsTable.ForeignKeys[1].RefTable = UsersTable
	UserAllowedDocumentsTable.ForeignKeys[0].RefTable = UsersTable
	UserAllowedDocumentsTable.ForeignKeys[1].RefTable = DocumentsTable
}
