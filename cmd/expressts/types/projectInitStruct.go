package types


type ProjectInitOptions struct {
	ConfigType string
}

type ProjectInitStruct struct {
	ProjectName string
	ProjectDescription string
	Options     ProjectInitOptions
}