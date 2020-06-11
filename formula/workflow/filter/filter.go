package filter


type Filter struct {
	FilterValues []string
}

func (filter *Filter) FilterBySH(filterValue...interface{})(bool, interface{}){
	return true, filterValue[0].(string)
}


func ExampleFilter(f []string) func(...interface{}) (bool, interface{}) {
	transforms := Filter{
		FilterValues: deviceNames,
	}
	return transforms.FilterBySH
}