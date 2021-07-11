package pgtype_test

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgtype/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt4rangeTranscode(t *testing.T) {
	testutil.TestSuccessfulTranscode(t, "int4range", []interface{}{
		&pgtype.Int4range{LowerType: pgtype.Empty, UpperType: pgtype.Empty, Status: pgtype.Present},
		&pgtype.Int4range{Lower: pgtype.Int4{Int: 1, Status: pgtype.Present}, Upper: pgtype.Int4{Int: 10, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Exclusive, Status: pgtype.Present},
		&pgtype.Int4range{Lower: pgtype.Int4{Int: -42, Status: pgtype.Present}, Upper: pgtype.Int4{Int: -5, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Exclusive, Status: pgtype.Present},
		&pgtype.Int4range{Lower: pgtype.Int4{Int: 1, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Unbounded, Status: pgtype.Present},
		&pgtype.Int4range{Upper: pgtype.Int4{Int: 1, Status: pgtype.Present}, LowerType: pgtype.Unbounded, UpperType: pgtype.Exclusive, Status: pgtype.Present},
		&pgtype.Int4range{Status: pgtype.Null},
	})
}

func TestInt4rangeNormalize(t *testing.T) {
	testutil.TestSuccessfulNormalize(t, []testutil.NormalizeTest{
		{
			SQL:   "select int4range(1, 10, '(]')",
			Value: pgtype.Int4range{Lower: pgtype.Int4{Int: 2, Status: pgtype.Present}, Upper: pgtype.Int4{Int: 11, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Exclusive, Status: pgtype.Present},
		},
	})
}

func TestInt4rangeUnmarshalJSON(t *testing.T) {
	successfulTests := []struct {
		source string
		result pgtype.Int4range
	}{
		{source: `"empty"`, result: pgtype.Int4range{LowerType: pgtype.Empty, UpperType: pgtype.Empty, Status: pgtype.Present}},
		{source: `"[10,20]"`, result: pgtype.Int4range{Lower: pgtype.Int4{Int: 10, Status: pgtype.Present}, LowerType: pgtype.Inclusive, Upper: pgtype.Int4{Int: 20, Status: pgtype.Present}, UpperType: pgtype.Inclusive, Status: pgtype.Present}},
		{source: `"[-10,)"`, result: pgtype.Int4range{Lower: pgtype.Int4{Int: -10, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Unbounded, Status: pgtype.Present}},
		{source: `null`, result: pgtype.Int4range{Status: pgtype.Null}},
	}
	for _, tt := range successfulTests {
		var r pgtype.Int4range
		err := r.UnmarshalJSON([]byte(tt.source))
		assert.NoError(t, err)
		assert.ObjectsAreEqualValues(tt.result, r)
	}
}

func TestInt4rangeMarshalJSON(t *testing.T) {
	successfulTests := []struct {
		source pgtype.Int4range
		result string
	}{
		{source: pgtype.Int4range{LowerType: pgtype.Empty, UpperType: pgtype.Empty, Status: pgtype.Present}, result: `"empty"`},
		{source: pgtype.Int4range{Lower: pgtype.Int4{Int: 10, Status: pgtype.Present}, LowerType: pgtype.Inclusive, Upper: pgtype.Int4{Int: 20, Status: pgtype.Present}, UpperType: pgtype.Inclusive, Status: pgtype.Present}, result: `"[10,20]"`},
		{source: pgtype.Int4range{Lower: pgtype.Int4{Int: -10, Status: pgtype.Present}, LowerType: pgtype.Inclusive, UpperType: pgtype.Unbounded, Status: pgtype.Present}, result: `"[-10,)"`},
		{source: pgtype.Int4range{Status: pgtype.Null}, result: `null`},
		{source: pgtype.Int4range{}, result: `null`},
	}
	for _, tt := range successfulTests {
		b, err := tt.source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, tt.result, string(b))
	}
}
