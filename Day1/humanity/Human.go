package humanity

import (
	"SofwareGoDay1/data"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

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
	_data, err := data.ReadFile(path)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(_data); i++ {
		tmp, errtmp := data.LineToCSV(_data[i])
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

func NewHumansFromJsonFile(path string) ([]*Human, error) {
	var array []*Human
	_data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(_data, &array)
	return array, nil
}
