// Code generated by go-enum, DO NOT EDIT.
package day

import (
	"fmt"
	"io"
	"strconv"
)

func (day_enum Day) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(day_enum.String()))
}

func (day_enum *Day) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum value %T must be a string", val)
	}

	enum, err := DayFromString(str)
	if err != nil {
		return err
	}
	 
	*day_enum = *enum
	return nil
}
