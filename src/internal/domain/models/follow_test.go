// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFollows(t *testing.T) {
	t.Parallel()

	query := Follows()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFollowsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
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

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Follows().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FollowExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Follow exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FollowExists to return true, but got false.")
	}
}

func testFollowsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	followFound, err := FindFollow(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if followFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFollowsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Follows().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFollowsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Follows().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFollowsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	followOne := &Follow{}
	followTwo := &Follow{}
	if err = randomize.Struct(seed, followOne, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}
	if err = randomize.Struct(seed, followTwo, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Follows().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFollowsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	followOne := &Follow{}
	followTwo := &Follow{}
	if err = randomize.Struct(seed, followOne, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}
	if err = randomize.Struct(seed, followTwo, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func followBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func followAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Follow) error {
	*o = Follow{}
	return nil
}

func testFollowsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Follow{}
	o := &Follow{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, followDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Follow object: %s", err)
	}

	AddFollowHook(boil.BeforeInsertHook, followBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	followBeforeInsertHooks = []FollowHook{}

	AddFollowHook(boil.AfterInsertHook, followAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	followAfterInsertHooks = []FollowHook{}

	AddFollowHook(boil.AfterSelectHook, followAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	followAfterSelectHooks = []FollowHook{}

	AddFollowHook(boil.BeforeUpdateHook, followBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	followBeforeUpdateHooks = []FollowHook{}

	AddFollowHook(boil.AfterUpdateHook, followAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	followAfterUpdateHooks = []FollowHook{}

	AddFollowHook(boil.BeforeDeleteHook, followBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	followBeforeDeleteHooks = []FollowHook{}

	AddFollowHook(boil.AfterDeleteHook, followAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	followAfterDeleteHooks = []FollowHook{}

	AddFollowHook(boil.BeforeUpsertHook, followBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	followBeforeUpsertHooks = []FollowHook{}

	AddFollowHook(boil.AfterUpsertHook, followAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	followAfterUpsertHooks = []FollowHook{}
}

func testFollowsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(followColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowToOneUserUsingFollower(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Follow
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FollowerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Follower().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FollowSlice{&local}
	if err = local.L.LoadFollower(ctx, tx, false, (*[]*Follow)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Follower == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Follower = nil
	if err = local.L.LoadFollower(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Follower == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFollowToOneUserUsingFollowing(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Follow
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, followDBTypes, false, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FollowingID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Following().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FollowSlice{&local}
	if err = local.L.LoadFollowing(ctx, tx, false, (*[]*Follow)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Following == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Following = nil
	if err = local.L.LoadFollowing(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Following == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFollowToOneSetOpUserUsingFollower(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Follow
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, followDBTypes, false, strmangle.SetComplement(followPrimaryKeyColumns, followColumnsWithoutDefault)...); err != nil {
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
		err = a.SetFollower(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Follower != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FollowerFollows[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FollowerID != x.ID {
			t.Error("foreign key was wrong value", a.FollowerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FollowerID))
		reflect.Indirect(reflect.ValueOf(&a.FollowerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FollowerID != x.ID {
			t.Error("foreign key was wrong value", a.FollowerID, x.ID)
		}
	}
}
func testFollowToOneSetOpUserUsingFollowing(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Follow
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, followDBTypes, false, strmangle.SetComplement(followPrimaryKeyColumns, followColumnsWithoutDefault)...); err != nil {
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
		err = a.SetFollowing(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Following != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FollowingFollows[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FollowingID != x.ID {
			t.Error("foreign key was wrong value", a.FollowingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FollowingID))
		reflect.Indirect(reflect.ValueOf(&a.FollowingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FollowingID != x.ID {
			t.Error("foreign key was wrong value", a.FollowingID, x.ID)
		}
	}
}

func testFollowsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
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

func testFollowsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFollowsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Follows().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	followDBTypes = map[string]string{`ID`: `int`, `FollowingID`: `int`, `FollowerID`: `int`}
	_             = bytes.MinRead
)

func testFollowsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(followPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(followAllColumns) == len(followPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followDBTypes, true, followPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFollowsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(followAllColumns) == len(followPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Follow{}
	if err = randomize.Struct(seed, o, followDBTypes, true, followColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followDBTypes, true, followPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(followAllColumns, followPrimaryKeyColumns) {
		fields = followAllColumns
	} else {
		fields = strmangle.SetComplement(
			followAllColumns,
			followPrimaryKeyColumns,
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

	slice := FollowSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFollowsUpsert(t *testing.T) {
	t.Parallel()

	if len(followAllColumns) == len(followPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFollowUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Follow{}
	if err = randomize.Struct(seed, &o, followDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Follow: %s", err)
	}

	count, err := Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, followDBTypes, false, followPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follow struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Follow: %s", err)
	}

	count, err = Follows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
