package merkle

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTreeHash_CalculateRoot(t *testing.T) {
	t.Parallel()

	testTable := map[string]struct {
		in      [][]byte
		want    string
		wantErr bool
	}{
		"even input": {
			in:   [][]byte{[]byte("1"), []byte("2")},
			want: "4295f72eeb1e3507b8461e240e3b8d18c1e7bd2f1122b11fc9ec40a65894031a",
		},
		"odd input": {
			in:   [][]byte{[]byte("1"), []byte("2"), []byte("3")},
			want: "0932f1d2e98219f7d7452801e2b64ebd9e5c005539db12d9b1ddabe7834d9044",
		},
		"file input": {
			in: func() [][]byte {
				b := testdata(t, "input.txt")

				return bytes.Split(b, []byte("\n"))
			}(),
			want: "136953c1b6c435fbcf07e979553bd07209d89078dad1aa9ae8d66d81032a21fe",
		},
		"error - empty data": {
			in:      [][]byte{},
			wantErr: true,
		},
	}

	for name, tt := range testTable {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			h := NewTreeHash(&SHA256Hasher{})

			got, err := h.CalculateRoot(tt.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err, "unexpected error")
			assert.Equal(t, tt.want, got)
		})
	}
}

func testdata(t *testing.T, name string) []byte {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(wd, "testdata", name)

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return b
}

func BenchmarkTreeHash_CalculateRoot(b *testing.B) {
	b.ReportAllocs()

	f, err := os.ReadFile("./testdata/input.txt")
	if err != nil {
		b.Fatal(err)
	}

	fb := bytes.Split(f, []byte("\n"))
	t := NewTreeHash(&SHA256Hasher{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = t.CalculateRoot(fb)
	}
}
