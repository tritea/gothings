package property_test

import (
	"testing"

	"github.com/AtomPod/thingmodel/thingmodel/property"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	validDatas := []struct {
		Data string
		Ok   bool
	}{
		{
			`{
				"name": "test_string_5",
				"description": "",
				"data": {
					"type": "string",
					"specs": {
						"length": 5
					}
				}
			}`,
			true,
		},
		{
			`{
				"name": "test_int_0_5",
				"description": "",
				"data": {
					"type": "integer",
					"specs": {
						"min": 0,
						"max": 5
					}
				}
			}`,
			true,
		},
		{
			`{
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
			}`,
			true,
		},
		{
			`{
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
			}`,
			true,
		},
		{
			`{
				"name": "test_string_5",
				"description": "",
				"data": {}
			}`,
			false,
		},
		{
			"", false,
		},
	}

	for _, v := range validDatas {
		d := property.PropertyDescription{}
		err := d.Parse([]byte(v.Data))
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestStringValidate(t *testing.T) {
	dataStr := `{
				"name": "test_string_5",
				"description": "",
				"data": {
					"type": "string",
					"specs": {
						"length": 5
					}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
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
		ok, err := d.Validate(v.Str)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestIntValidate(t *testing.T) {
	dataStr := `{
				"name": "test_int",
				"description": "",
				"data": {
					"type": "integer",
					"specs": {
						"min": 0,
						"max": 5
					}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			0, true,
		},
		{
			1, true,
		},
		{
			2, true,
		},
		{
			3, true,
		},
		{
			4, true,
		},
		{
			5, true,
		},
		{
			-1, false,
		},
		{
			6, false,
		},
		{
			"1", false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestIntStepValidate(t *testing.T) {
	dataStr := `{
				"name": "test_int",
				"description": "",
				"data": {
					"type": "integer",
					"specs": {
						"min": 0,
						"max": 5,
						"step": 2
					}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			0, true,
		},
		{
			1, false,
		},
		{
			2, true,
		},
		{
			3, false,
		},
		{
			4, true,
		},
		{
			5, false,
		},
		{
			-1, false,
		},
		{
			6, false,
		},
		{
			"1", false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestFloatValidate(t *testing.T) {
	dataStr := `{
				"name": "test_float",
				"description": "",
				"data": {
					"type": "number",
					"specs": {
						"min": 0,
						"max": 5
					}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			0.0, true,
		},
		{
			1.0, true,
		},
		{
			2.0, true,
		},
		{
			3.0, true,
		},
		{
			4.0, true,
		},
		{
			5.0, true,
		},
		{
			-1.0, false,
		},
		{
			6.0, false,
		},
		{
			"1", false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestFloatStepValidate(t *testing.T) {
	dataStr := `{
				"name": "test_float",
				"description": "",
				"data": {
					"type": "number",
					"specs": {
						"min": 0,
						"max": 5,
						"step": 0.5
					}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			0, true,
		},
		{
			1.2, false,
		},
		{
			2, true,
		},
		{
			3.2, false,
		},
		{
			4, true,
		},
		{
			4.8, false,
		},
		{
			-1, false,
		},
		{
			6, false,
		},
		{
			"1", false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestBoolValidate(t *testing.T) {
	dataStr := `{
				"name": "test_bool",
				"description": "",
				"data": {
					"type": "boolean",
					"specs": {}
				}
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			true, true,
		},
		{
			false, true,
		},
		{
			2, false,
		},
		{
			3.2, false,
		},
		{
			"1", false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestArrayValidate(t *testing.T) {
	dataStr := `{
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
						"max": 100
					}
				}
			}
		}
	}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value interface{}
		Ok    bool
	}{
		{
			[]float32{50, 60, 70}, true,
		},
		{
			[]int32{30, 50, 70}, false,
		},
		{
			[]string{"a", "b", "c"}, false,
		},
		{
			[]interface{}{50, "a", "c"}, false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestStructValidate(t *testing.T) {
	dataStr := `
				{
				"name": "hello",
				"description": "",
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
			}`

	d := property.PropertyDescription{}
	err := d.Parse([]byte(dataStr))
	assert.Nil(t, err)

	validData := []struct {
		Value map[string]interface{}
		Ok    bool
	}{
		{
			Value: map[string]interface{}{
				"name": "123456",
				"age":  0,
			},
			Ok: true,
		},
		{
			Value: map[string]interface{}{
				"name": "12",
				"be":   0,
			},
			Ok: false,
		},
		{
			Value: map[string]interface{}{
				"test": "12",
				"age":  0,
			},
			Ok: false,
		},
		{
			Value: map[string]interface{}{
				"name": 12,
				"age":  "name",
			},
			Ok: false,
		},
	}

	for _, v := range validData {
		ok, err := d.Validate(v.Value)
		assert.Equal(t, ok, v.Ok)
		if v.Ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}
