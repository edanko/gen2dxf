package wcog

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type WCOGRow struct {
	Block string
	PosNo string
	Count int
}

type WCOG struct {
	v []*WCOGRow
	sync.Mutex
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

func ReadWCOGs(paths []string) *WCOG {
	m := &WCOG{
		v: make([]*WCOGRow, 0),
	}

	for _, path := range paths {

		f, err := os.Open(path)
		if err != nil {
			fmt.Println("[!] failed to open", path)
			continue
		}
		defer f.Close()

		r := csv.NewReader(bufio.NewReader(f))
		r.FieldsPerRecord = -1

		// header
		_, err = r.Read()
		if err != nil {
			log.Fatal(err)
		}
		_, err = r.Read()
		if err != nil {
			log.Fatal(err)
		}

		for {
			rec, err := r.Read()

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			block := rec[6]
			posno := filterPos(rec[0])

			m.addOrInc(block, posno)

		}

	}

	return m
}

func (w *WCOG) GetQuantity(block, posno string) int {
	for _, v := range w.v {
		if v.Block == block && v.PosNo == posno {
			w.Lock()
			defer w.Unlock()

			return v.Count
		}
	}

	fmt.Println("[x] block", block, "and pos", posno, "not found")
	return -1
}

func filterPos(s string) string {
	posSplit := strings.Split(s, "-")
	posno := posSplit[len(posSplit)-1]
	posno = strings.Replace(posno, "P", "", -1)
	posno = strings.Replace(posno, "S", "", -1)
	posno = strings.Replace(posno, "B", "", -1)
	posno = strings.Replace(posno, "C", "", -1)

	return posno
}
