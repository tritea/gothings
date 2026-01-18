package thingmodel_test

import (
	"testing"

	"github.com/AtomPod/thingmodel/thingmodel"
	"github.com/stretchr/testify/assert"
)

var (
	dataStr = `
	{
		"name": "sensors",
		"events": [
			{
				"name": "man",
				"description": "have man",
				"type": "alert",
				"data": {
					"type": "string",
					"specs": {
						"length": 30
					}
				}
			}
		],
		"properties": [
			{
				"name": "test_string_5",
				"description": "",
				"data": {
					"type": "string",
					"specs": {
						"length": 5
					}
				}
			},
			{
				"name": "temp",
				"description": "temp",
				"data": {
					"type": "array",
					"specs": {
						"length": 5,
						"data": {
							"type": "number",
							"specs": {
								"min": 50,
								"max": 100,
								"step": 0.01
							}
						}
					}
				}
			},
			{
				"name": "hello",
				"description": "hello world",
				"required": true,
				"data": {
					"type": "struct",
					"specs": {
						"name": {
							"type": "string",
							"specs": {
								"length": 15
							}
						},
						"age": {
							"type": "integer",
							"specs": {
								"min": 0,
								"max": 15,
								"step": 1,
								"unit": "y"
							}
						}
					}
				} 
			}
		]
	}
	`
)

func TestParse(t *testing.T) {
	thm := &thingmodel.ThingModel{}
	err := thm.Parse([]byte(dataStr))
	assert.Nil(t, err)
}

func TestStringValid(t *testing.T) {
	thm := &thingmodel.ThingModel{}
	err := thm.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Str string
		Ok  bool
	}{
		{
			"", true,
		},
		{
			"0", true,
		},
		{
			"123", true,
		},
		{
			"12345", true,
		},
		{
			"123456", false,
		},
	}

	for _, v := range validData {
		ok, err := thm.ValidateProperty("test_string_5", v.Str)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}
