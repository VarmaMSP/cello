package sqlstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestModel struct {
	x string
	y string
}

func (t *TestModel) DbColumns() []string {
	return []string{"x", "y"}
}

func (t *TestModel) FieldAddrs() []interface{} {
	return []interface{}{&t.x, &t.y}
}

func TestValuesFromAddrs(t *testing.T) {
	assert := assert.New(t)

	x := &TestModel{"field1", "field2"}
	actual := ValuesFromAddrs(x.FieldAddrs())
	expected := []interface{}{"field1", "field2"}

	assert.ElementsMatch(expected, actual)
}
