package mvc

import "github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"

func GetMVCFolders() []types.LayeredFolderStruct {
	return []types.LayeredFolderStruct{
		{FolderName: "bin"},
		{FolderName: "model"},
		{FolderName: "controller"},
		{FolderName: "routes"},
		{FolderName: "middleware"},
		{FolderName: "utilities"},
	}
}
