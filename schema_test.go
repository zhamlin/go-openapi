package openapi_test

import (
	"encoding/json"
	"testing"

	"github.com/sv-tools/openapi"
	"github.com/sv-tools/openapi/internal/require"
)

func TestSchema_Marshal_Unmarshal(t *testing.T) {
	for _, tt := range []struct {
		name            string
		data            string
		expected        string
		emptyExtensions bool
	}{
		{
			name:            "spec only",
			data:            `{"title": "foo"}`,
			emptyExtensions: true,
		},
		{
			name:            "spec with extension field",
			data:            `{"title": "foo", "b": "bar"}`,
			emptyExtensions: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("json", func(t *testing.T) {
				var v *openapi.Schema
				require.NoError(t, json.Unmarshal([]byte(tt.data), &v))
				if tt.emptyExtensions {
					require.Empty(t, v.Extensions)
				} else {
					require.NotEmpty(t, v.Extensions)
				}
				data, err := json.Marshal(&v)
				require.NoError(t, err)
				if tt.expected == "" {
					tt.expected = tt.data
				}
				require.JSONEq(t, tt.expected, string(data))
			})
		})
	}
}

func TestSchema_AddExt(t *testing.T) {
	for _, tt := range []struct {
		name     string
		key      string
		value    any
		expected map[string]any
	}{
		{
			name:  "without prefix",
			key:   "foo",
			value: 42,
			expected: map[string]any{
				"foo": 42,
			},
		},
		{
			name:  "with prefix",
			key:   "x-foo",
			value: 43,
			expected: map[string]any{
				"x-foo": 43,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ext := openapi.Schema{}
			ext.AddExt(tt.key, tt.value)
			require.Equal(t, tt.expected, ext.Extensions)
		})
	}
}
