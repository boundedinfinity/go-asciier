//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package utf8

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/commons/slices"
)

var (
	All = []Utf8Char{
		NULL,
		START_OF_HEADER,
		START_OF_TEXT,
		END_OF_TEXT,
		END_OF_TRANSMISSION,
		ENQUIRY,
		ACKNOWLEDGE,
		BELL,
		BACKSPACE,
		HORIZONTAL_TAB,
		LINE_FEED,
		VERTICAL_TAB,
		FORM_FEED,
		CARRIAGE_RETURN,
		SHIFT_OUT,
		SHIFT_IN,
		DATA_LINK_ESCAPE,
		DEVICE_CONTROL_1,
		DEVICE_CONTROL_2,
		DEVICE_CONTROL_3,
		DEVICE_CONTROL_4,
		NEGATIVE_ACKNOLEDGE,
		SYNCHRONIZE,
		END_OF_TRANSMISSION_BLOCK,
		CANCEL,
		END_OF_MEDIUM,
		SUBSTITUE,
		ESCAPE,
		FILE_SEPARATOR,
		GROUP_SEPARATOR,
		RECORD_SEPARATOR,
		UNIT_SEPARATOR,
		SPACE,
		EXCLAMATION_MARK,
		DOUBLE_QUOTE,
		NUMBER,
		DOLLAR,
		PERCENT,
		AMPERSAND,
		SINGLE_QUOTE,
		LEFT_PARENTHESIS,
		RIGHT_PARENTHESIS,
		ASTERISK,
		PLUS,
		COMMA,
		MINUS,
		PERIOD,
		SLASH,
		NUMBER_ZERO,
		NUMBER_ONE,
		NUMBER_TWO,
		NUMBER_THREE,
		NUMBER_FOUR,
		NUMBER_FIVE,
		NUMBER_SIX,
		NUMBER_SEVEN,
		NUMBER_EIGHT,
		NUMBER_NINE,
		COLON,
		SEMICOLON,
		LESS_THAN,
		EQUAL_SIGN,
		GREATER_THAN,
		QUESTION_MARK,
		AT_SIGN,
		UPPERCASE_A,
		UPPERCASE_B,
		UPPERCASE_C,
		UPPERCASE_D,
		UPPERCASE_E,
		UPPERCASE_F,
		UPPERCASE_G,
		UPPERCASE_H,
		UPPERCASE_I,
		UPPERCASE_J,
		UPPERCASE_K,
		UPPERCASE_L,
		UPPERCASE_M,
		UPPERCASE_N,
		UPPERCASE_O,
		UPPERCASE_P,
		UPPERCASE_Q,
		UPPERCASE_R,
		UPPERCASE_S,
		UPPERCASE_T,
		UPPERCASE_U,
		UPPERCASE_V,
		UPPERCASE_W,
		UPPERCASE_X,
		UPPERCASE_Y,
		UPPERCASE_Z,
		LEFT_SQUARE_BRACKET,
		BACKSLASH,
		RIGHT_SQUARE_BRACKET,
		CARET,
		UNDERSCORE,
		GRAVE,
		LOWERCASE_A,
		LOWERCASE_B,
		LOWERCASE_C,
		LOWERCASE_D,
		LOWERCASE_E,
		LOWERCASE_F,
		LOWERCASE_G,
		LOWERCASE_H,
		LOWERCASE_I,
		LOWERCASE_J,
		LOWERCASE_K,
		LOWERCASE_L,
		LOWERCASE_M,
		LOWERCASE_N,
		LOWERCASE_O,
		LOWERCASE_P,
		LOWERCASE_Q,
		LOWERCASE_R,
		LOWERCASE_S,
		LOWERCASE_T,
		LOWERCASE_U,
		LOWERCASE_V,
		LOWERCASE_W,
		LOWERCASE_X,
		LOWERCASE_Y,
		LOWERCASE_Z,
		LEFT_CURLY_BRACKET,
		VERTICAL_BAR,
		RIGHT_CURLY_BRACKET,
		TILDE,
		DELETE,
	}
)

func (t Utf8Char) String() string {
	return string(t)
}

func Parse(v byte) (Utf8Char, error) {
	f, ok := slices.FindFn(All, func(x Utf8Char) bool {
		return Utf8Char(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s byte) bool {
	return slices.ContainsFn(All, func(v Utf8Char) bool {
		return byte(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v byte) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slices.Join(All, ","),
	)
}

func (t Utf8Char) MarshalJSON() ([]byte, error) {
	return json.Marshal(byte(t))
}

func (t *Utf8Char) UnmarshalJSON(data []byte) error {
	var v byte

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t Utf8Char) MarshalYAML() (interface{}, error) {
	return byte(t), nil
}

func (t *Utf8Char) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v byte

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
