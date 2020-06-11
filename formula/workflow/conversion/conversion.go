package conversion

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

type Conversion struct {
}

func NewConversion() Conversion {
	return Conversion{}
}

func (f Conversion) TransformToXML(params ...interface{}) (continuePipeline bool, stringType interface{}) {
	if len(params) < 1 {
		return false, errors.New("No msg Received")
	}
	if result, ok := params[0].(string); ok {
		b, err := xml.Marshal(result)
		if err != nil {
			return false, errors.New("Incorrect type received, expecting string")
		}

		return true, string(b)
	}
	return false, errors.New("Unexpected type received")
}

func (f Conversion) TransformToJSON(params ...interface{}) (continuePipeline bool, stringType interface{}) {
	if len(params) < 1 {
		return false, errors.New("No msg Received")
	}
	if result, ok := params[0].(string); ok {
		b, err := json.Marshal(result)
		if err != nil {
			return false, errors.New("Error marshalling JSON")
		}
		return true, string(b)
	}
	return false, errors.New("Unexpected type received")
}
