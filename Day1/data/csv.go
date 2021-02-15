package data

import (
	"errors"
	"io/ioutil"
)

func StringToArray(path string, delim byte) []string {
	var array []string
	var tmp string

	for i := 0; i < len(path); i++ {
		if path[i] == delim {
			if tmp != "" {
				array = append(array, tmp)
				tmp = ""
			}
		} else {
			tmp += string(path[i])
		}
	}
	if tmp != "" {
		array = append(array, tmp)
	}
	return array
}

func ReadFile(path string) ([]string, error) {
	by, err := ioutil.ReadFile(path)

	var array []string
	var tmp string

	for i := 0; i < len(by); i++ {
		if by[i] == '\n' {
			if tmp != "" {
				array = append(array, tmp)
				tmp = ""
			}
		} else {
			tmp += string(by[i])
		}
	}
	if tmp != "" {
		array = append(array, tmp)
	}
	return array, err
}

func LineToCSV(line string) ([]string, error) {
	array := StringToArray(line, ',')
	if len(array) == 1 {
		return array, errors.New("error")
	}
	return array, nil
}
