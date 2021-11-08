package wcog

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type WCOGRow struct {
	Block string
	PosNo string
	Count int
}

type WCOG struct {
	v []*WCOGRow
	sync.RWMutex
}

func (w *WCOG) addOrInc(block, posno string) {
	w.Lock()
	defer w.Unlock()

	found := false

	for _, v := range w.v {
		if v.Block == block && v.PosNo == posno {
			found = true
			v.Count++
			break
		}
	}

	if !found {
		cur := &WCOGRow{
			Block: block,
			PosNo: posno,
			Count: 1,
		}

		w.v = append(w.v, cur)
	}
}

func ReadWCOGs(paths []string) (*WCOG, error) {
	m := &WCOG{
		v: make([]*WCOGRow, 0),
	}

	g := new(errgroup.Group)

	for _, path := range paths {
		path := path
		g.Go(func() error {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			r := csv.NewReader(bufio.NewReader(f))
			r.FieldsPerRecord = -1

			// header
			_, err = r.Read()
			if err != nil {
				return err
			}
			_, err = r.Read()
			if err != nil {
				return err
			}

			for {
				rec, err := r.Read()

				if errors.Is(err, io.EOF) {
					break
				} else if err != nil {
					return err
				}

				block := rec[6]
				posno := filterPos(rec[0])

				m.addOrInc(block, posno)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return m, nil
}

func (w *WCOG) GetQuantity(block, posno string) int {
	for _, v := range w.v {
		if v.Block == block && v.PosNo == posno {
			w.RLock()
			defer w.RUnlock()

			return v.Count
		}
	}

	return 1
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
