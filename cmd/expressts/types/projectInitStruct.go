package types

type ProjectInitOptions struct {
	ConfigType  string
	ProjectArch *string
	Confirm     *bool
}

type ProjectInitStruct struct {
	ProjectName        string
	ProjectDescription string
	Options            ProjectInitOptions
}
