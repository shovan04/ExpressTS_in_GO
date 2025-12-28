package types

type FilePath struct {
	DirPath    string
	FileName   string
	FolderName *string // optional
}

type WriteFileStruct struct {
	Path    FilePath
	Content []byte
}
