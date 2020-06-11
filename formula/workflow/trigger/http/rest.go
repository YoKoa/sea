package http

import (
	"fmt"
	"github.com/YoKoa/sea/formula/workflow/constant"
	"github.com/YoKoa/sea/formula/workflow/rt"
	"github.com/YoKoa/sea/formula/workflow/types"
	"github.com/YoKoa/sea/formula/workflow/webserver"
	"io/ioutil"
	"net/http"
)

type Trigger struct {
	Runtime       *rt.FlowRuntime
	Webserver     *webserver.WebServer
}

func (trigger *Trigger) Initialize() error {
	trigger.Webserver.SetupTriggerRoute(trigger.requestHandler)
	return nil
}

func (trigger *Trigger) requestHandler(writer http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	contentType := r.Header.Get(constant.ContentType)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("Error reading HTTP Body: %s", err.Error())))
		return
	}

	msg := types.Message{
		ContentType: contentType,
		Payload:     data,
	}

	messageError := trigger.Runtime.ProcessMessage(msg)
	if messageError != nil {
		// ProcessMessage logs the error, so no need to log it here.
		writer.WriteHeader(messageError.ErrorCode)
		writer.Write([]byte(messageError.Err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("ok"))
}


