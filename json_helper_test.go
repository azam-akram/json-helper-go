package jsonhelper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testEmpStr = `{
    "id": "The ID",
    "name": "The User",
    "designation": "CEO",
    "address":
    [
        {
            "doorNumber": 1,
            "street": "The office street 1",
            "city": "The office city 1",
            "country": "The office country 1"
        },
        {
            "doorNumber": 2,
            "street": "The home street 2",
            "city": "The home city 2",
            "country": "The home country 2"
        }
    ]
}`

func Test_StringToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := NewJsonHelper()

	var emp Employee
	err := jh.StringToStruct(testEmpStr, &emp)
	assertThat.Nil(err)

	assertThat.Equal(emp.ID, "The ID")
	assertThat.Equal(emp.Name, "The User")
}

func Test_StringToStruct_EmployeeShort_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := NewJsonHelper()

	var emp EmployeeShort
	err := jh.StringToStruct(testEmpStr, &emp)
	assertThat.Nil(err)

	assertThat.Equal(emp.ID, "The ID")
	assertThat.Equal(emp.Name, "The User")
}

func Test_StructToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}

	jh := NewJsonHelper()
	str, err := jh.StructToString(employee)
	assertThat.Nil(err)

	expectedRes := `{"id":"The ID","name":"The User"}`

	assertThat.Equal(expectedRes, str)
}

func Test_StringToMap_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := NewJsonHelper()
	jMap, _ := jh.StringToMap(testEmpStr)

	id := jMap["id"].(string)
	user := jMap["name"].(string)

	assertThat.Equal(id, "The ID")
	assertThat.Equal(user, "The User")
}

func Test_MapToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	expectedRes := "{\"id\":\"The ID\",\"name\":\"The User\"}"

	mapData := map[string]interface{}{
		"id":   "The ID",
		"name": "The User",
	}

	jh := NewJsonHelper()
	jsonStr, err := jh.MapToString(mapData)

	assertThat.Nil(err)
	assertThat.Equal(jsonStr, expectedRes)
}

func Test_StringToBytes_Success(t *testing.T) {
	jh := NewJsonHelper()

	jh.StringToBytes(testEmpStr)

	assert.NotNil(t, testEmpStr)
}

func Test_BytesToString_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHelper()

	inputBytes := []byte(`{"id": "The ID", "name": "The User"}`)
	outputString := jh.BytesToString(inputBytes)

	actualBytes := []byte(outputString)

	assertThat.Equal(inputBytes, actualBytes)
}

func Test_StructToBytes_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHelper()

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}
	actualBytes, err := jh.StructToBytes(employee)

	assertThat.Nil(err)
	assertThat.NotNil(actualBytes)
}

func Test_BytesToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHelper()

	byteValue, err := os.ReadFile("testdata/employee.json")
	assertThat.Nil(err)

	var emp *Employee
	err = jh.BytesToStruct(byteValue, &emp)

	assertThat.Nil(err)
	assertThat.Equal(emp.ID, "The ID")
}

func Test_ModifyInputJson_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHelper()

	modifiedEmpMap, err := jh.ModifyInputJson(testEmpStr)

	assertThat.Nil(err)
	assert.NotNil(t, modifiedEmpMap)
	assertThat.Equal(modifiedEmpMap["degree"], "phd")
	assertThat.Equal(modifiedEmpMap["name"], "The User")
}
