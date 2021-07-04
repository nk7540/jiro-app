// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testComments(t *testing.T) {
	t.Parallel()

	query := Comments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCommentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Comments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CommentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Comment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CommentExists to return true, but got false.")
	}
}

func testCommentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	commentFound, err := FindComment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if commentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCommentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Comments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCommentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Comments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCommentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Comments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCommentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	commentOne := &Comment{}
	commentTwo := &Comment{}
	if err = randomize.Struct(seed, commentOne, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}
	if err = randomize.Struct(seed, commentTwo, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func commentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func commentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Comment) error {
	*o = Comment{}
	return nil
}

func testCommentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Comment{}
	o := &Comment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, commentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Comment object: %s", err)
	}

	AddCommentHook(boil.BeforeInsertHook, commentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterInsertHook, commentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	commentAfterInsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterSelectHook, commentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	commentAfterSelectHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpdateHook, commentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpdateHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpdateHook, commentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	commentAfterUpdateHooks = []CommentHook{}

	AddCommentHook(boil.BeforeDeleteHook, commentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	commentBeforeDeleteHooks = []CommentHook{}

	AddCommentHook(boil.AfterDeleteHook, commentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	commentAfterDeleteHooks = []CommentHook{}

	AddCommentHook(boil.BeforeUpsertHook, commentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	commentBeforeUpsertHooks = []CommentHook{}

	AddCommentHook(boil.AfterUpsertHook, commentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	commentAfterUpsertHooks = []CommentHook{}
}

func testCommentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(commentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommentToManyFavorites(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Comment
	var b, c Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CommentID, a.ID)
	queries.Assign(&c.CommentID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Favorites().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.CommentID, b.CommentID) {
			bFound = true
		}
		if queries.Equal(v.CommentID, c.CommentID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CommentSlice{&a}
	if err = a.L.LoadFavorites(ctx, tx, false, (*[]*Comment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Favorites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Favorites = nil
	if err = a.L.LoadFavorites(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Favorites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCommentToManyAddOpFavorites(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Comment
	var b, c, d, e Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Favorite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Favorite{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFavorites(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CommentID) {
			t.Error("foreign key was wrong value", a.ID, first.CommentID)
		}
		if !queries.Equal(a.ID, second.CommentID) {
			t.Error("foreign key was wrong value", a.ID, second.CommentID)
		}

		if first.R.Comment != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Comment != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Favorites[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Favorites[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Favorites().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCommentToManySetOpFavorites(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Comment
	var b, c, d, e Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Favorite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetFavorites(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Favorites().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetFavorites(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Favorites().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CommentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CommentID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CommentID) {
		t.Error("foreign key was wrong value", a.ID, d.CommentID)
	}
	if !queries.Equal(a.ID, e.CommentID) {
		t.Error("foreign key was wrong value", a.ID, e.CommentID)
	}

	if b.R.Comment != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Comment != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Comment != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Comment != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Favorites[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Favorites[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCommentToManyRemoveOpFavorites(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Comment
	var b, c, d, e Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Favorite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddFavorites(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Favorites().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveFavorites(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Favorites().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CommentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CommentID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Comment != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Comment != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Comment != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Comment != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Favorites) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Favorites[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Favorites[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCommentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCommentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCommentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Comments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	commentDBTypes = map[string]string{`ID`: `int`, `UserID`: `int`, `ContentID`: `int`, `Body`: `text`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_              = bytes.MinRead
)

func testCommentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commentDBTypes, true, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCommentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Comment{}
	if err = randomize.Struct(seed, o, commentDBTypes, true, commentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commentDBTypes, true, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(commentAllColumns, commentPrimaryKeyColumns) {
		fields = commentAllColumns
	} else {
		fields = strmangle.SetComplement(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CommentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCommentsUpsert(t *testing.T) {
	t.Parallel()

	if len(commentAllColumns) == len(commentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCommentUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Comment{}
	if err = randomize.Struct(seed, &o, commentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err := Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, commentDBTypes, false, commentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Comment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Comment: %s", err)
	}

	count, err = Comments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
