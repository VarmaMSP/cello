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

func TestInsertQuery(t *testing.T) {
	assert := assert.New(t)

	x := &TestModel{}
	actual := InsertQuery("test", x, 2)
	expected := "INSERT INTO test(x,y) VALUES (?,?),(?,?)"

	assert.Equal(expected, actual)
}

func TestUpdateQuery(t *testing.T) {
	assert := assert.New(t)

	x := &TestModel{"pavan", "varma"}
	y := *x
	y.x = "varma"
	y.y = "pavan"
	actual, values, _ := UpdateQuery("test", x, &y, "")
	expected := "UPDATE test SET x = ?,y = ?"

	p, _ := values[0].(string)
	q, _ := values[1].(string)

	assert.Equal("varma", p)
	assert.Equal("pavan", q)
	assert.Equal(expected, actual)
}

func TestValuesFromAddrs(t *testing.T) {
	assert := assert.New(t)

	x := &TestModel{"field1", "field2"}
	actual := ValuesFromAddrs(x.FieldAddrs())
	expected := []interface{}{"field1", "field2"}

	assert.ElementsMatch(expected, actual)
}
