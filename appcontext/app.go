package appcontext

func Init() {
	Load()
	LoadDependencies()

	SetupLogger()
}
