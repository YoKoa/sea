package rt

import (
	"encoding/json"
	"fmt"
	"github.com/YoKoa/sea/formula/workflow/constant"
	"github.com/YoKoa/sea/formula/workflow/types"
	"github.com/ugorji/go/codec"
	"net/http"
	"reflect"
	"sync"
)

const unmarshalErrorMessage = "Unable to unmarshal message payload as %s"

type FlowRuntime struct {
	TargetType    interface{}
	transforms    []AppFunction
	isBusyCopying sync.Mutex
}

func (rt *FlowRuntime) SetTransforms(transforms []AppFunction) {
	rt.transforms = transforms
}

func (rt *FlowRuntime) ProcessMessage(msg types.Message) *types.MessageError {

	if rt.TargetType == nil {
		rt.TargetType = []byte{}
	}

	if reflect.TypeOf(rt.TargetType).Kind() != reflect.Ptr {
		err := fmt.Errorf("TargetType must be a pointer, not a value of the target type.")
		return &types.MessageError{Err: err, ErrorCode: http.StatusInternalServerError}
	}

	// Must make a copy of the type so that data isn't retained between calls.
	target := reflect.New(reflect.ValueOf(rt.TargetType).Elem().Type()).Interface()
	var contentType string

	switch target.(type) {
	case *[]byte:
		target = &msg.Payload
		contentType = msg.ContentType

	default:
		switch msg.ContentType {
		case constant.ContentTypeJSON:

			if err := json.Unmarshal([]byte(msg.Payload), target); err != nil {
				message := fmt.Sprintf(unmarshalErrorMessage, "JSON")
				err = fmt.Errorf("%s : %s", message, err.Error())
				return &types.MessageError{Err: err, ErrorCode: http.StatusBadRequest}
			}


		case constant.ContentTypeCBOR:
			x := codec.CborHandle{}
			err := codec.NewDecoderBytes([]byte(msg.Payload), &x).Decode(&target)
			if err != nil {
				message := fmt.Sprintf(unmarshalErrorMessage, "CBOR")
				err = fmt.Errorf("%s : %s", message, err.Error())
				return &types.MessageError{Err: err, ErrorCode: http.StatusBadRequest}
			}

			// Needed for Marking event as handled

		default:
			message := "content type for input data not supported"
			err := fmt.Errorf("'%s' %s", msg.ContentType, message)
			return &types.MessageError{Err: err, ErrorCode: http.StatusBadRequest}
		}
	}

	target = reflect.ValueOf(target).Elem().Interface()

	var result interface{}
	var continuePipeline = true

	for _, trxFunc := range rt.transforms {
		if result != nil {
			continuePipeline, result = trxFunc(result)
		} else {
			continuePipeline, result = trxFunc(target, contentType)
		}
		if continuePipeline != true {
			if result != nil {
				if err, ok := result.(error); ok {
					return &types.MessageError{Err: err, ErrorCode: http.StatusUnprocessableEntity}
				}
			}
			break
		}
	}
	return nil
}
