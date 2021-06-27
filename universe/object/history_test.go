package object

import (
	"testing"
	"time"

	"github.com/mkuchenbecker/storage/universe/vector"
	"github.com/stretchr/testify/assert"
)

func TestInsertGet(t *testing.T) {
	t.Parallel()

	id := ObjectIdentifier("foo")
	hist := NewHistory(id)

	objTime1 := New(id, vector.New(1, 0, 0))
	objTime2 := New(id, vector.New(2, 0, 0))
	objTime3 := New(id, vector.New(3, 0, 0))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	time3 := time2.Add(time.Second)

	assert.NoError(t, hist.Insert(time1, objTime1))
	assert.NoError(t, hist.Insert(time2, objTime2))
	assert.NoError(t, hist.Insert(time3, objTime3))

	actual, _ := hist.Get(time1)
	assert.Equal(t, objTime1, actual)

	actual, _ = hist.Get(time2)
	assert.Equal(t, objTime2, actual)

	actual, _ = hist.Get(time3)
	assert.Equal(t, objTime3, actual)
}

func TestInsertIdMismatch(t *testing.T) {
	t.Parallel()

	id := ObjectIdentifier("foo")
	hist := NewHistory(id)

	obj := New(ObjectIdentifier("bar"), vector.New(0, 0, 0))
	assert.Error(t, hist.Insert(time.Now(), obj))
}

func TestIdentifier(t *testing.T) {
	t.Parallel()

	id := ObjectIdentifier("foo")
	hist := NewHistory(id)
	assert.Equal(t, id, hist.Identifier())
}

func TestPrune(t *testing.T) {
	t.Parallel()

	id := ObjectIdentifier("foo")
	hist := NewHistory(id)

	objTime1 := New(id, vector.New(1, 0, 0))
	objTime2 := New(id, vector.New(2, 0, 0))
	objTime3 := New(id, vector.New(3, 0, 0))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	time3 := time2.Add(time.Second)

	assert.NoError(t, hist.Insert(time1, objTime1))
	assert.NoError(t, hist.Insert(time2, objTime2))
	assert.NoError(t, hist.Insert(time3, objTime3))

	_, ok := hist.Get(time1)
	assert.True(t, ok)

	hist.Prune(time2)

	_, ok = hist.Get(time1)
	assert.False(t, ok)
	_, ok = hist.Get(time2)
	assert.True(t, ok)
	_, ok = hist.Get(time3)
	assert.True(t, ok)

}
