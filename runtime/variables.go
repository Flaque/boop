package runtime

import (
	"strconv"
)

type Type int

const (
	STRING Type = iota
	BOOL
	INT
	FLOAT
	UNKNOWN
)

type Variable struct {
	Type  Type
	Value interface{}
}

func NewVariable(t Type, val interface{}) *Variable {
	return &Variable{
		t,
		val,
	}
}

func GetVariable(val interface{}) *Variable {

	if val == nil {
		return NewVariable(UNKNOWN, nil)
	}

	if v, ok := val.(*Variable); ok {
		return v
	}

	if v, ok := val.(string); ok {
		return NewVariable(STRING, v)
	}

	if v, ok := val.(bool); ok {
		return NewVariable(BOOL, v)
	}

	if v, ok := val.(int64); ok {
		return NewVariable(INT, v)
	}

	if v, ok := val.(float64); ok {
		return NewVariable(FLOAT, v)
	}

	return NewVariable(UNKNOWN, nil)
}

func GetVariableFromString(str string) *Variable {

	if str == "true" {
		return NewVariable(BOOL, true)
	}

	if str == "false" {
		return NewVariable(BOOL, false)
	}

	if val, err := strconv.Atoi(str); err == nil {
		return NewVariable(INT, val)
	}

	// TODO parse for float

	return NewVariable(STRING, str)
}

// Is returns true if t is the same type as we've stored
func (v *Variable) Is(t Type) bool {
	return v.Type == t
}

// Equals returns true if the interface matches the stored val
func (v *Variable) Equals(val interface{}) bool {

	// Unknowns never equal each other
	if v.Type == UNKNOWN {
		return false
	}

	return v.Value == val
}

// AsString forces the variable to be a string.
// Not a good idea to call unless you're sure.
func (v *Variable) AsString() string {
	str, _ := v.Value.(string)
	return str
}

func (v *Variable) Plus(other *Variable) *Variable {

	// TODO: Allow FLOAT to add too
	if v.Type != INT || other.Type != INT {
		return NewVariable(UNKNOWN, nil)
	}

	// Add ints
	aval, _ := v.Value.(int)
	bval, _ := v.Value.(int)
	return NewVariable(INT, aval+bval)
}
