package conversion

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Conversion houses various built in conversion transforms (XML, JSON, CSV)
type Conversion struct {
}

// NewConversion creates, initializes and returns a new instance of Conversion
func NewConversion() Conversion {
	return Conversion{}
}

// TransformToXML transforms an EdgeX event to XML.
// It will return an error and stop the pipeline if a non-edgex event is received or if no data is received.
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

// TransformToJSON transforms an EdgeX event to JSON.
// It will return an error and stop the pipeline if a non-edgex event is received or if no data is received.
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
