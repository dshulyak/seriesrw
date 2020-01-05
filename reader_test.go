package seriesrw

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	tmp, err := ioutil.TempFile("", "testing-seriesrw-")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	for i := uint64(0); i < 10; i++ {
		require.NoError(t, binary.Write(tmp, binary.BigEndian, i))
	}

	reader, err := NewFileReader(tmp.Name(), 80)
	require.NoError(t, err)

	var (
		i, j uint64
	)

	for err = reader.Read(&j); !errors.Is(err, io.EOF); err = reader.Read(&j) {
		require.Equal(t, i, j)
		i++
	}
}
