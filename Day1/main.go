package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
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

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	val, atoierr := strconv.Atoi(csv[1])
	if atoierr != nil {
		return nil, atoierr
	}
	hum := &Human{Name: csv[0], Age: val, Country: csv[2]}

	return hum, nil
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	var array []*Human
	data, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		tmp, errtmp := LineToCSV(data[i])
		if errtmp != nil || len(tmp) != 3 {
			return nil, errtmp
		}
		val, atoierr := strconv.Atoi(tmp[1])
		if atoierr != nil {
			return nil, atoierr
		}
		p := &Human{Name: tmp[0], Age: val, Country: tmp[2]}
		array = append(array, p)
	}
	return array, nil
}

func main() {

	file, err := NewHumansFromCsvFile("file.csv")

	if err != nil {
		fmt.Println("marche pas")
	}
	for i := 0; i < len(file); i++ {
		fmt.Println(file[i].Name)
		fmt.Println(file[i].Age)
		fmt.Println(file[i].Country)
	}
}
