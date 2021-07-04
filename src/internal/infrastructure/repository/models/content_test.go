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

func testContents(t *testing.T) {
	t.Parallel()

	query := Contents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testContentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
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

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Contents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ContentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Content exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContentExists to return true, but got false.")
	}
}

func testContentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	contentFound, err := FindContent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if contentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testContentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Contents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testContentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Contents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentOne := &Content{}
	contentTwo := &Content{}
	if err = randomize.Struct(seed, contentOne, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}
	if err = randomize.Struct(seed, contentTwo, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Contents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contentOne := &Content{}
	contentTwo := &Content{}
	if err = randomize.Struct(seed, contentOne, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}
	if err = randomize.Struct(seed, contentTwo, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func contentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func contentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Content) error {
	*o = Content{}
	return nil
}

func testContentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Content{}
	o := &Content{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Content object: %s", err)
	}

	AddContentHook(boil.BeforeInsertHook, contentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contentBeforeInsertHooks = []ContentHook{}

	AddContentHook(boil.AfterInsertHook, contentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contentAfterInsertHooks = []ContentHook{}

	AddContentHook(boil.AfterSelectHook, contentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contentAfterSelectHooks = []ContentHook{}

	AddContentHook(boil.BeforeUpdateHook, contentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contentBeforeUpdateHooks = []ContentHook{}

	AddContentHook(boil.AfterUpdateHook, contentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contentAfterUpdateHooks = []ContentHook{}

	AddContentHook(boil.BeforeDeleteHook, contentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contentBeforeDeleteHooks = []ContentHook{}

	AddContentHook(boil.AfterDeleteHook, contentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contentAfterDeleteHooks = []ContentHook{}

	AddContentHook(boil.BeforeUpsertHook, contentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contentBeforeUpsertHooks = []ContentHook{}

	AddContentHook(boil.AfterUpsertHook, contentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contentAfterUpsertHooks = []ContentHook{}
}

func testContentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(contentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentToManyBrowses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c Browse

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, browseDBTypes, false, browseColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, browseDBTypes, false, browseColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ContentID = a.ID
	c.ContentID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Browses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ContentID == b.ContentID {
			bFound = true
		}
		if v.ContentID == c.ContentID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ContentSlice{&a}
	if err = a.L.LoadBrowses(ctx, tx, false, (*[]*Content)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Browses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Browses = nil
	if err = a.L.LoadBrowses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Browses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testContentToManyFavorites(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
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

	b.ContentID = a.ID
	c.ContentID = a.ID

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
		if v.ContentID == b.ContentID {
			bFound = true
		}
		if v.ContentID == c.ContentID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ContentSlice{&a}
	if err = a.L.LoadFavorites(ctx, tx, false, (*[]*Content)(&slice), nil); err != nil {
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

func testContentToManyAddOpBrowses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c, d, e Browse

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Browse{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, browseDBTypes, false, strmangle.SetComplement(browsePrimaryKeyColumns, browseColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Browse{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBrowses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ContentID {
			t.Error("foreign key was wrong value", a.ID, first.ContentID)
		}
		if a.ID != second.ContentID {
			t.Error("foreign key was wrong value", a.ID, second.ContentID)
		}

		if first.R.Content != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Content != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Browses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Browses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Browses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testContentToManyAddOpFavorites(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c, d, e Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
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

		if a.ID != first.ContentID {
			t.Error("foreign key was wrong value", a.ID, first.ContentID)
		}
		if a.ID != second.ContentID {
			t.Error("foreign key was wrong value", a.ID, second.ContentID)
		}

		if first.R.Content != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Content != &a {
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
func testContentToOneCategoryUsingCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Content
	var foreign Category

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CategoryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Category().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentSlice{&local}
	if err = local.L.LoadCategory(ctx, tx, false, (*[]*Content)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Category = nil
	if err = local.L.LoadCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Content
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Content)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentToOneSetOpCategoryUsingCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Category{&b, &c} {
		err = a.SetCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Category != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Contents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CategoryID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CategoryID))
		reflect.Indirect(reflect.ValueOf(&a.CategoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CategoryID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryID, x.ID)
		}
	}
}
func testContentToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Contents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testContentToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Content
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.UserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Contents) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testContentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
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

func testContentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Contents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contentDBTypes = map[string]string{`ID`: `int`, `UserID`: `int`, `CategoryID`: `int`, `Title`: `varchar`, `Description`: `text`, `ThumbnailURL`: `text`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_              = bytes.MinRead
)

func testContentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(contentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(contentAllColumns) == len(contentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentDBTypes, true, contentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testContentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contentAllColumns) == len(contentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Content{}
	if err = randomize.Struct(seed, o, contentDBTypes, true, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentDBTypes, true, contentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contentAllColumns, contentPrimaryKeyColumns) {
		fields = contentAllColumns
	} else {
		fields = strmangle.SetComplement(
			contentAllColumns,
			contentPrimaryKeyColumns,
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

	slice := ContentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testContentsUpsert(t *testing.T) {
	t.Parallel()

	if len(contentAllColumns) == len(contentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLContentUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Content{}
	if err = randomize.Struct(seed, &o, contentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Content: %s", err)
	}

	count, err := Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, contentDBTypes, false, contentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Content: %s", err)
	}

	count, err = Contents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
