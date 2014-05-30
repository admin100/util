package properties

// +----------------------------------------------------------------------
// | properties [ Golang read the properties file ]
// +----------------------------------------------------------------------
// | Author: lemontea <36634584@qq.com> <https://github.com/admin100>
// +----------------------------------------------------------------------
import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type properties struct {
	maps map[string]string
}

//Loading the properties file into memory
func Load(file string) (*properties, error) {
	f, err := os.Open(file)
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

//According to the key to return to the corresponding value
func (p *properties) Get(key string) string {
	return p.maps[key]
}
