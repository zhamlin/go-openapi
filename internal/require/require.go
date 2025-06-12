package require

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func Error(t testing.TB, err error, msgAndArgs ...any) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func NoError(t testing.TB, err error, msgAndArgs ...any) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func ErrorContains(t testing.TB, err error, substring string, msgAndArgs ...any) {
	t.Helper()

	if err == nil || !strings.Contains(err.Error(), substring) {
		t.Fatalf("expected error to contain %q, got %v", substring, err)
	}
}

func JSONEq(t testing.TB, expected, actual string, msgAndArgs ...any) {
	t.Helper()

	var ej, aj any
	if err := json.Unmarshal([]byte(expected), &ej); err != nil {
		t.Fatalf("invalid expected JSON: %v", err)
	}
	if err := json.Unmarshal([]byte(actual), &aj); err != nil {
		t.Fatalf("invalid actual JSON: %v", err)
	}

	if !reflect.DeepEqual(ej, aj) {
		t.Fatalf("JSON not equal\nExpected: %s\nActual: %s", expected, actual)
	}
}

func Nil(t testing.TB, obj any, msgAndArgs ...any) {
	t.Helper()

	if obj != nil && (reflect.ValueOf(obj).Kind() != reflect.Ptr || !reflect.ValueOf(obj).IsNil()) {
		t.Fatalf("expected nil value, got: %#v", obj)
	}
}

func NotNil(t testing.TB, obj any, msgAndArgs ...any) {
	t.Helper()

	if obj == nil || (reflect.ValueOf(obj).Kind() == reflect.Ptr && reflect.ValueOf(obj).IsNil()) {
		t.Fatalf("expected non-nil value, got nil")
	}
}

func Len(t testing.TB, obj any, length int, msgAndArgs ...any) {
	t.Helper()

	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() != length {
			t.Fatalf("expected length %d, got %d", length, v.Len())
		}
	default:
		t.Fatalf("type %T does not have length", obj)
	}
}

func Equal(t testing.TB, expected, actual any, msgAndArgs ...any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("values not equal\nExpected: %#v\nActual: %#v", expected, actual)
	}
}

func Empty(t testing.TB, obj any, msgAndArgs ...any) {
	t.Helper()
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() != 0 {
			t.Fatalf("expected empty, but length is %d", v.Len())
		}
	default:
		if !v.IsZero() {
			t.Fatalf("expected zero value, got: %#v", obj)
		}
	}
}

func NotEmpty(t testing.TB, obj any, msgAndArgs ...any) {
	t.Helper()
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() == 0 {
			t.Fatalf("expected non-empty, but length is 0")
		}
	default:
		if v.IsZero() {
			t.Fatalf("expected non-zero value, got: %#v", obj)
		}
	}
}

func IsType(t testing.TB, expectedType, actual any, msgAndArgs ...any) {
	t.Helper()
	eType := reflect.TypeOf(expectedType)
	aType := reflect.TypeOf(actual)
	if eType != aType {
		t.Fatalf("expected type %v, got type %v", eType, aType)
	}
}

func Truef(t testing.TB, condition bool, format string, args ...any) {
	t.Helper()
	if !condition {
		t.Fatalf(format, args...)
	}
}

