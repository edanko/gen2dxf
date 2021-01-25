package gen

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type ProfileGen struct {
	CommonData *CommonData
	Profiles   []*ProfileData
	RestData   *RestData
}

func ParseProfileFile(fname string) *ProfileGen {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	return readListedProfile(f)
}

func readListedProfile(r io.Reader) *ProfileGen {
	s := bufio.NewScanner(r)

	// type
	s.Scan()
	// version
	s.Scan()
	// usage
	s.Scan()

	prof := &ProfileGen{}

	for s.Scan() {
		t := s.Text()
		switch t {
		case "COMMON_DATA":
			cd := readCommonData(s)
			prof.CommonData = cd

		case "PROFILE_DATA":
			pd := readProfileData(s)
			prof.Profiles = append(prof.Profiles, pd)

		case "LEFT_END":
			prof.Profiles[len(prof.Profiles)-1].LeftEnd = readEnd(s)

		case "RIGHT_END":
			prof.Profiles[len(prof.Profiles)-1].RightEnd = readEnd(s)

		case "HOLES_NOTCHES_CUTOUTS":
			prof.Profiles[len(prof.Profiles)-1].HolesNotchesCutouts = append(prof.Profiles[len(prof.Profiles)-1].HolesNotchesCutouts, readHolesNotchesCutouts(s))

		case "CONNECTION_TRACE":
			prof.Profiles[len(prof.Profiles)-1].ConnectionTrace = readConnectionTrace(s)

		case "GEOMETRY_DATA":
			prof.Profiles[len(prof.Profiles)-1].GeometryData = append(prof.Profiles[len(prof.Profiles)-1].GeometryData, readGeometryData(s))

		case "STRING_DATA":
			prof.Profiles[len(prof.Profiles)-1].StringData = append(prof.Profiles[len(prof.Profiles)-1].StringData, readStringData(s))

		case "REST_DATA":
			prof.RestData = readRestData(s)

		default:
			fmt.Println("unknown section", t)
		}
	}

	return prof
}
