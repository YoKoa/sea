package main

import (
	"github.com/YoKoa/sea/formula/workflow"
	"github.com/YoKoa/sea/formula/workflow/conversion"
	"github.com/YoKoa/sea/formula/workflow/filter"
)

func main() {

	// 1) First thing to do is to create an workflow and initialize it.
	wf := workflow.Flow{
		ServiceKey: "example",
		TargetType: []byte{},
	}
	wf.Initialize()

	// 2) Since our Filter Function requires the list of fields we would
	fields := []string{"a", "b", "c"}

	// 3) This is our pipeline configuration, the collection of functions to
	//    Filter -> ...Conversion -> Output
	//    You need to customize all the pipeline functions for your data format
	wf.SetFunctionsPipeline(
		filter.ExampleFilter(fields),
		conversion.NewConversion().TransformToJSON,
		// output
	)
	// 4) TODO shows how to access the application's specific configuration settings.

	// 5) Lastly, we'll go ahead and tell the flow to "start" and begin listening for stream data
	wf.Run()
}
