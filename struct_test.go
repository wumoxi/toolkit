package toolkit

import (
	"encoding/json"
	"sort"
	"testing"
)

// student struct type
type student struct {
	Id      int    `json:"id" canChangeMethod:"-"`
	Name    string `json:"name" canChangeMethod:"change"`
	Email   string `json:"email" canChangeMethod:"change"`
	Phone   string `json:"phone" canChangeMethod:"change"`
	Sex     string `json:"sex" canChangeMethod:"change"`
	Age     int    `json:"age" canChangeMethod:"change,modify"`
	Address string `json:"address" canChangeMethod:"change,modify"`
}

func TestGetAllStructTag(t *testing.T) {
	expected := map[string][]map[string]string{
		"Id":      {{"json": "id"}},
		"Name":    {{"json": "name"}, {"canChangeMethod": "change"}},
		"Email":   {{"json": "email"}, {"canChangeMethod": "change"}},
		"Phone":   {{"json": "phone"}, {"canChangeMethod": "change"}},
		"Sex":     {{"json": "sex"}, {"canChangeMethod": "change"}},
		"Age":     {{"json": "age"}, {"canChangeMethod": "change,modify"}},
		"Address": {{"json": "address"}, {"canChangeMethod": "change,modify"}},
	}
	expectedStr, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("json marshal error: %s\n", err)
	}

	actual, err := GetAllStructTags(student{})
	if err != nil {
		t.Errorf("get all struct tags error: %s\n", err)
	}
	actualStr, err := json.Marshal(actual)
	if err != nil {
		t.Errorf("json marshal error: %s\n", err)
	}

	if string(expectedStr) != string(actualStr) {
		t.Errorf("result should: %s, got: %s\n", string(expectedStr), string(actualStr))
	}
}

func TestGetFieldsNameByTag(t *testing.T) {
	expectedC := []string{"Name", "Email", "Phone", "Sex", "Age", "Address"}
	sort.Slice(expectedC, func(i, j int) bool {
		return expectedC[i] < expectedC[j]
	})
	actualC, err := GetFieldsNameByTag(student{}, "change")
	if err != nil {
		t.Errorf("get field name slice by tag value error: %s\n", err)
	}
	for index, expectedV := range expectedC {
		if expectedV != actualC[index] {
			t.Errorf("expected: %s, got: %s\n", expectedV, actualC[index])
		}
	}

	expectedM := []string{"Age", "Address"}
	sort.Slice(expectedM, func(i, j int) bool {
		return expectedM[i] < expectedM[j]
	})
	actualM, err := GetFieldsNameByTag(student{}, "modify")
	if err != nil {
		t.Errorf("get field name slice by tag value error: %s\n", err)
	}
	for index, expectedV := range expectedM {
		if expectedV != actualM[index] {
			t.Errorf("expected: %s, got: %s\n", expectedV, actualM[index])
		}
	}
}
