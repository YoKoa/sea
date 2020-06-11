package workflow

import (
	"errors"
	"github.com/YoKoa/sea/formula/workflow/rt"
)

type Flow struct {
	ServiceKey  string
	configDir   string
	TargetType  interface{}
	useRegistry bool
	transforms  []rt.AppFunction
	runtime     *rt.FlowRuntime
}

func (flow *Flow) Initialize() {

}
func (flow *Flow) SetFunctionsPipeline(transforms ...rt.AppFunction) error {
	if len(transforms) == 0 {
		return errors.New("No transforms provided to pipeline")
	}
	return nil
}

func (flow *Flow) Run() {
	flow.runtime = &rt.FlowRuntime{TargetType: flow.TargetType}
	flow.runtime.SetTransforms(flow.transforms)

}
