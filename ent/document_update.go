// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"docuSync/ent/document"
	"docuSync/ent/predicate"
	"docuSync/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DocumentUpdate is the builder for updating Document entities.
type DocumentUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentUpdate builder.
func (du *DocumentUpdate) Where(ps ...predicate.Document) *DocumentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetTitle sets the "title" field.
func (du *DocumentUpdate) SetTitle(s string) *DocumentUpdate {
	du.mutation.SetTitle(s)
	return du
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableTitle(s *string) *DocumentUpdate {
	if s != nil {
		du.SetTitle(*s)
	}
	return du
}

// ClearTitle clears the value of the "title" field.
func (du *DocumentUpdate) ClearTitle() *DocumentUpdate {
	du.mutation.ClearTitle()
	return du
}

// SetText sets the "text" field.
func (du *DocumentUpdate) SetText(s string) *DocumentUpdate {
	du.mutation.SetText(s)
	return du
}

// SetNillableText sets the "text" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableText(s *string) *DocumentUpdate {
	if s != nil {
		du.SetText(*s)
	}
	return du
}

// ClearText clears the value of the "text" field.
func (du *DocumentUpdate) ClearText() *DocumentUpdate {
	du.mutation.ClearText()
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DocumentUpdate) SetCreatedAt(t time.Time) *DocumentUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableCreatedAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DocumentUpdate) SetUpdatedAt(t time.Time) *DocumentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// AddEditorIDs adds the "editors" edge to the User entity by IDs.
func (du *DocumentUpdate) AddEditorIDs(ids ...int) *DocumentUpdate {
	du.mutation.AddEditorIDs(ids...)
	return du
}

// AddEditors adds the "editors" edges to the User entity.
func (du *DocumentUpdate) AddEditors(u ...*User) *DocumentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddEditorIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (du *DocumentUpdate) SetOwnerID(id int) *DocumentUpdate {
	du.mutation.SetOwnerID(id)
	return du
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (du *DocumentUpdate) SetNillableOwnerID(id *int) *DocumentUpdate {
	if id != nil {
		du = du.SetOwnerID(*id)
	}
	return du
}

// SetOwner sets the "owner" edge to the User entity.
func (du *DocumentUpdate) SetOwner(u *User) *DocumentUpdate {
	return du.SetOwnerID(u.ID)
}

// AddAllowedUserIDs adds the "allowed_users" edge to the User entity by IDs.
func (du *DocumentUpdate) AddAllowedUserIDs(ids ...int) *DocumentUpdate {
	du.mutation.AddAllowedUserIDs(ids...)
	return du
}

// AddAllowedUsers adds the "allowed_users" edges to the User entity.
func (du *DocumentUpdate) AddAllowedUsers(u ...*User) *DocumentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddAllowedUserIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (du *DocumentUpdate) Mutation() *DocumentMutation {
	return du.mutation
}

// ClearEditors clears all "editors" edges to the User entity.
func (du *DocumentUpdate) ClearEditors() *DocumentUpdate {
	du.mutation.ClearEditors()
	return du
}

// RemoveEditorIDs removes the "editors" edge to User entities by IDs.
func (du *DocumentUpdate) RemoveEditorIDs(ids ...int) *DocumentUpdate {
	du.mutation.RemoveEditorIDs(ids...)
	return du
}

// RemoveEditors removes "editors" edges to User entities.
func (du *DocumentUpdate) RemoveEditors(u ...*User) *DocumentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveEditorIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (du *DocumentUpdate) ClearOwner() *DocumentUpdate {
	du.mutation.ClearOwner()
	return du
}

// ClearAllowedUsers clears all "allowed_users" edges to the User entity.
func (du *DocumentUpdate) ClearAllowedUsers() *DocumentUpdate {
	du.mutation.ClearAllowedUsers()
	return du
}

// RemoveAllowedUserIDs removes the "allowed_users" edge to User entities by IDs.
func (du *DocumentUpdate) RemoveAllowedUserIDs(ids ...int) *DocumentUpdate {
	du.mutation.RemoveAllowedUserIDs(ids...)
	return du
}

// RemoveAllowedUsers removes "allowed_users" edges to User entities.
func (du *DocumentUpdate) RemoveAllowedUsers(u ...*User) *DocumentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveAllowedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DocumentUpdate) Save(ctx context.Context) (int, error) {
	if err := du.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DocumentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DocumentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DocumentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DocumentUpdate) defaults() error {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		if document.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized document.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := document.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (du *DocumentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(document.Table, document.Columns, sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Title(); ok {
		_spec.SetField(document.FieldTitle, field.TypeString, value)
	}
	if du.mutation.TitleCleared() {
		_spec.ClearField(document.FieldTitle, field.TypeString)
	}
	if value, ok := du.mutation.Text(); ok {
		_spec.SetField(document.FieldText, field.TypeString, value)
	}
	if du.mutation.TextCleared() {
		_spec.ClearField(document.FieldText, field.TypeString)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(document.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(document.FieldUpdatedAt, field.TypeTime, value)
	}
	if du.mutation.EditorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedEditorsIDs(); len(nodes) > 0 && !du.mutation.EditorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EditorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   document.OwnerTable,
			Columns: []string{document.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   document.OwnerTable,
			Columns: []string{document.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AllowedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAllowedUsersIDs(); len(nodes) > 0 && !du.mutation.AllowedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AllowedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DocumentUpdateOne is the builder for updating a single Document entity.
type DocumentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentMutation
}

// SetTitle sets the "title" field.
func (duo *DocumentUpdateOne) SetTitle(s string) *DocumentUpdateOne {
	duo.mutation.SetTitle(s)
	return duo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableTitle(s *string) *DocumentUpdateOne {
	if s != nil {
		duo.SetTitle(*s)
	}
	return duo
}

// ClearTitle clears the value of the "title" field.
func (duo *DocumentUpdateOne) ClearTitle() *DocumentUpdateOne {
	duo.mutation.ClearTitle()
	return duo
}

// SetText sets the "text" field.
func (duo *DocumentUpdateOne) SetText(s string) *DocumentUpdateOne {
	duo.mutation.SetText(s)
	return duo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableText(s *string) *DocumentUpdateOne {
	if s != nil {
		duo.SetText(*s)
	}
	return duo
}

// ClearText clears the value of the "text" field.
func (duo *DocumentUpdateOne) ClearText() *DocumentUpdateOne {
	duo.mutation.ClearText()
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DocumentUpdateOne) SetCreatedAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableCreatedAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DocumentUpdateOne) SetUpdatedAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// AddEditorIDs adds the "editors" edge to the User entity by IDs.
func (duo *DocumentUpdateOne) AddEditorIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.AddEditorIDs(ids...)
	return duo
}

// AddEditors adds the "editors" edges to the User entity.
func (duo *DocumentUpdateOne) AddEditors(u ...*User) *DocumentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddEditorIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (duo *DocumentUpdateOne) SetOwnerID(id int) *DocumentUpdateOne {
	duo.mutation.SetOwnerID(id)
	return duo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableOwnerID(id *int) *DocumentUpdateOne {
	if id != nil {
		duo = duo.SetOwnerID(*id)
	}
	return duo
}

// SetOwner sets the "owner" edge to the User entity.
func (duo *DocumentUpdateOne) SetOwner(u *User) *DocumentUpdateOne {
	return duo.SetOwnerID(u.ID)
}

// AddAllowedUserIDs adds the "allowed_users" edge to the User entity by IDs.
func (duo *DocumentUpdateOne) AddAllowedUserIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.AddAllowedUserIDs(ids...)
	return duo
}

// AddAllowedUsers adds the "allowed_users" edges to the User entity.
func (duo *DocumentUpdateOne) AddAllowedUsers(u ...*User) *DocumentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddAllowedUserIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (duo *DocumentUpdateOne) Mutation() *DocumentMutation {
	return duo.mutation
}

// ClearEditors clears all "editors" edges to the User entity.
func (duo *DocumentUpdateOne) ClearEditors() *DocumentUpdateOne {
	duo.mutation.ClearEditors()
	return duo
}

// RemoveEditorIDs removes the "editors" edge to User entities by IDs.
func (duo *DocumentUpdateOne) RemoveEditorIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.RemoveEditorIDs(ids...)
	return duo
}

// RemoveEditors removes "editors" edges to User entities.
func (duo *DocumentUpdateOne) RemoveEditors(u ...*User) *DocumentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveEditorIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (duo *DocumentUpdateOne) ClearOwner() *DocumentUpdateOne {
	duo.mutation.ClearOwner()
	return duo
}

// ClearAllowedUsers clears all "allowed_users" edges to the User entity.
func (duo *DocumentUpdateOne) ClearAllowedUsers() *DocumentUpdateOne {
	duo.mutation.ClearAllowedUsers()
	return duo
}

// RemoveAllowedUserIDs removes the "allowed_users" edge to User entities by IDs.
func (duo *DocumentUpdateOne) RemoveAllowedUserIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.RemoveAllowedUserIDs(ids...)
	return duo
}

// RemoveAllowedUsers removes "allowed_users" edges to User entities.
func (duo *DocumentUpdateOne) RemoveAllowedUsers(u ...*User) *DocumentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveAllowedUserIDs(ids...)
}

// Where appends a list predicates to the DocumentUpdate builder.
func (duo *DocumentUpdateOne) Where(ps ...predicate.Document) *DocumentUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DocumentUpdateOne) Select(field string, fields ...string) *DocumentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Document entity.
func (duo *DocumentUpdateOne) Save(ctx context.Context) (*Document, error) {
	if err := duo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DocumentUpdateOne) SaveX(ctx context.Context) *Document {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DocumentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DocumentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DocumentUpdateOne) defaults() error {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		if document.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized document.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := document.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (duo *DocumentUpdateOne) sqlSave(ctx context.Context) (_node *Document, err error) {
	_spec := sqlgraph.NewUpdateSpec(document.Table, document.Columns, sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Document.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, document.FieldID)
		for _, f := range fields {
			if !document.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != document.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Title(); ok {
		_spec.SetField(document.FieldTitle, field.TypeString, value)
	}
	if duo.mutation.TitleCleared() {
		_spec.ClearField(document.FieldTitle, field.TypeString)
	}
	if value, ok := duo.mutation.Text(); ok {
		_spec.SetField(document.FieldText, field.TypeString, value)
	}
	if duo.mutation.TextCleared() {
		_spec.ClearField(document.FieldText, field.TypeString)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(document.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(document.FieldUpdatedAt, field.TypeTime, value)
	}
	if duo.mutation.EditorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedEditorsIDs(); len(nodes) > 0 && !duo.mutation.EditorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EditorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.EditorsTable,
			Columns: document.EditorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   document.OwnerTable,
			Columns: []string{document.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   document.OwnerTable,
			Columns: []string{document.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AllowedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAllowedUsersIDs(); len(nodes) > 0 && !duo.mutation.AllowedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AllowedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   document.AllowedUsersTable,
			Columns: document.AllowedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Document{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
