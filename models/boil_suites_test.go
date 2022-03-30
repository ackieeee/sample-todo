// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Tasks", testTasks)
}

func TestDelete(t *testing.T) {
	t.Run("Tasks", testTasksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Tasks", testTasksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Tasks", testTasksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Tasks", testTasksExists)
}

func TestFind(t *testing.T) {
	t.Run("Tasks", testTasksFind)
}

func TestBind(t *testing.T) {
	t.Run("Tasks", testTasksBind)
}

func TestOne(t *testing.T) {
	t.Run("Tasks", testTasksOne)
}

func TestAll(t *testing.T) {
	t.Run("Tasks", testTasksAll)
}

func TestCount(t *testing.T) {
	t.Run("Tasks", testTasksCount)
}

func TestHooks(t *testing.T) {
	t.Run("Tasks", testTasksHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Tasks", testTasksInsert)
	t.Run("Tasks", testTasksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Tasks", testTasksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Tasks", testTasksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Tasks", testTasksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Tasks", testTasksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Tasks", testTasksSliceUpdateAll)
}
