package properties

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type properties struct {
	maps map[string]string
}

func Load(file string) (*properties, error) {
	f, err := os.Open("db.properties")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	maps := make(map[string]string)
	for scanner.Scan() {
		reg := regexp.MustCompile("\\s*#.*") //comment
		line := scanner.Text()
		if !reg.MatchString(line) {
			reg = regexp.MustCompile("\\s*\\w+\\s*=.*")
			if reg.MatchString(line) {
				arr := strings.SplitN(line, "=", 2)
				key := strings.Trim(arr[0], " ")
				value := strings.Trim(arr[1], " ")
				maps[key] = value
			}
		}
	}
	prop := properties{maps: maps}
	return &prop, scanner.Err()
}

func (p *properties) Get(key string) string {
	return p.maps[key]
}
