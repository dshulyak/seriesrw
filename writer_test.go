package seriesrw

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	tmp, err := ioutil.TempFile("", "testing-seriesrw-")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	w, err := NewFileWriter(tmp.Name(), 1024)
	require.NoError(t, err)
	for i := uint64(0); i < 10; i++ {
		require.NoError(t, w.Write(i))
	}
	require.NoError(t, w.Close())

	buf, err := ioutil.ReadAll(tmp)
	require.NoError(t, err)
	for i := 0; i < 10; i++ {
		r := binary.BigEndian.Uint64(buf[i*8:])
		require.Equal(t, uint64(i), r)
	}

}
