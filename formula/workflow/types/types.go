package types

type Message struct {
	ContentType string
	Payload     []byte
}
type MessageError struct {
	Err       error
	ErrorCode int
}

type Event struct {
	ID          string    `json:"id" codec:"id,omitempty"`             // ID uniquely identifies an event, for example a UUID
	Pushed      int64     `json:"pushed" codec:"pushed,omitempty"`     // Pushed is a timestamp indicating when the event was exported. If unexported, the value is zero.
	Device      string    `json:"device" codec:"device,omitempty"`     // Device identifies the source of the event, can be a device name or id. Usually the device name.
	Created     int64     `json:"created" codec:"created,omitempty"`   // Created is a timestamp indicating when the event was created.
	Modified    int64     `json:"modified" codec:"modified,omitempty"` // Modified is a timestamp indicating when the event was last modified.
	Origin      int64     `json:"origin" codec:"origin,omitempty"`     // Origin is a timestamp that can communicate the time of the original reading, prior to event creation
	Readings    []Reading `json:"readings" codec:"readings,omitempty"` // Readings will contain zero to many entries for the associated readings of a given event.
	isValidated bool      // internal member used for validation check
}

type Reading struct {
	Id          string `json:"id" codec:"id,omitempty"`
	Pushed      int64  `json:"pushed" codec:"pushed,omitempty"`   // When the data was pushed out of EdgeX (0 - not pushed yet)
	Created     int64  `json:"created" codec:"created,omitempty"` // When the reading was created
	Origin      int64  `json:"origin" codec:"origin,omitempty"`
	Modified    int64  `json:"modified" codec:"modified,omitempty"`
	Device      string `json:"device" codec:"device,omitempty"`
	Name        string `json:"name" codec:"name,omitempty"`
	Value       string `json:"value"  codec:"value,omitempty"`            // Device sensor data value
	BinaryValue []byte `json:"binaryValue" codec:"binaryValue,omitempty"` // Binary data payload
	isValidated bool   // internal member used for validation check
}