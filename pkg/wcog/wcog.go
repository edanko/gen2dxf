package wcog

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
	"sync"
)

func ReadWCOG(r io.Reader, m *sync.Map) error {
	c := csv.NewReader(r)
	c.FieldsPerRecord = -1

	// header
	_, err := c.Read()
	if err != nil {
		return err
	}
	_, err = c.Read()
	if err != nil {
		return err
	}

	for {
		rec, err := c.Read()

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		block := rec[6]
		posno := filterPos(rec[0])

		val, exists := m.Load(block + posno)
		if exists {
			if n, ok := val.(uint); ok {
				n++
				m.Store(block+posno, n)
			}
		} else {
			m.Store(block+posno, uint(1))
		}
	}
	return nil
}

func filterPos(s string) string {
	posSplit := strings.Split(s, "-")
	posno := posSplit[len(posSplit)-1]
	posno = strings.ReplaceAll(posno, "P", "")
	posno = strings.ReplaceAll(posno, "S", "")
	posno = strings.ReplaceAll(posno, "B", "")
	posno = strings.ReplaceAll(posno, "C", "")

	return posno
}
