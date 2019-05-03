package common

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	db := InitDB()
	assert.NoError(t, db.DB().Ping())
	assert.Equal(t, os.Getenv("MYSQL_DATABASE"), db.Dialect().CurrentDatabase())
}

func TestInitTestDB(t *testing.T) {
	db := InitTestDB()
	assert.NoError(t, db.DB().Ping())
	assert.Equal(t, os.Getenv("MYSQL_TEST_DATABASE"), db.Dialect().CurrentDatabase())
}

func TestGetDB(t *testing.T) {
	db1 := InitDB()
	db2 := GetDB()
	assert.Equal(t, fmt.Sprintf("%p", db1), fmt.Sprintf("%p", db2))
	assert.NoError(t, db2.DB().Ping())
}
