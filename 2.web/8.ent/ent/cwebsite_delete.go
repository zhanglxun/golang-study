// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"golangStudy/2.web/8.ent/ent/cwebsite"
	"golangStudy/2.web/8.ent/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CWebsiteDelete is the builder for deleting a CWebsite entity.
type CWebsiteDelete struct {
	config
	hooks    []Hook
	mutation *CWebsiteMutation
}

// Where appends a list predicates to the CWebsiteDelete builder.
func (cd *CWebsiteDelete) Where(ps ...predicate.CWebsite) *CWebsiteDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CWebsiteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CWebsiteMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CWebsiteDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CWebsiteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cwebsite.Table, sqlgraph.NewFieldSpec(cwebsite.FieldID, field.TypeInt64))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CWebsiteDeleteOne is the builder for deleting a single CWebsite entity.
type CWebsiteDeleteOne struct {
	cd *CWebsiteDelete
}

// Where appends a list predicates to the CWebsiteDelete builder.
func (cdo *CWebsiteDeleteOne) Where(ps ...predicate.CWebsite) *CWebsiteDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CWebsiteDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cwebsite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CWebsiteDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
