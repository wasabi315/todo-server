package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	common = "common"
	ex     = "ex"
)

var (
	repositories = map[string]*MapRepository{}
)

func TestMain(m *testing.M) {
	repos := []string{
		common,
		ex,
	}

	for _, key := range repos {
		repo, err := NewMapRepository()
		if err != nil {
			panic(err)
		}
		repositories[key] = repo
	}

	code := m.Run()

	os.Exit(code)
}

func assertAndRequire(t *testing.T) (*assert.Assertions, *require.Assertions) {
	return assert.New(t), require.New(t)
}

func setup(t *testing.T, repo string) (Repository, *assert.Assertions, *require.Assertions) {
	t.Helper()
	r, ok := repositories[repo]
	if !ok {
		t.FailNow()
	}
	assert, require := assertAndRequire(t)
	return r, assert, require
}

func mustMakeTodo(t *testing.T, repo Repository) *Todo {
	t.Helper()
	todoArg := TodoArg{}
	todo, err := repo.CreateTodo(todoArg)
	require.NoError(t, err)
	return todo
}
