package appcontext

type Instance struct {
}

var AppDependencies *Instance

func LoadDependencies() {
	AppDependencies = &Instance{}
}
