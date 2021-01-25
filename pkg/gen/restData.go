package gen

import (
	"bufio"
	"fmt"
	"strings"
)

type RestData struct {
	RestName string
}

func readRestData(s *bufio.Scanner) *RestData {
	r := &RestData{}
	for s.Scan() {
		c := s.Text()

		switch c {
		case "END_OF_REST_DATA":
			return r

		default:
			l := strings.Split(c, "=")
			switch l[0] {

			case "REST_NAME":
				r.RestName = l[1]

			default:
				fmt.Println("unknown field in rest data:", l[0])
			}
		}
	}
	return r
}
