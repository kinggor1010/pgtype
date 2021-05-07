package pgtype_test

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgtype/testutil"
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
	}
	for i, tt := range successfulTests {
		var r pgtype.Int4range
		err := r.UnmarshalJSON([]byte(tt.source))
		if err != nil {
			t.Errorf("%d: %v", i, err)
		}

		if  r.Status != tt.result.Status {
			t.Errorf("%d: expected %v to convert to %v, but it was %v", i, tt.source, tt.result, r)
		}
	}
}
