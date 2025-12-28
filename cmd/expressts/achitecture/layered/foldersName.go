package layered

import "github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"

func GetLayeredFolders() []types.LayeredFolderStruct {
	return []types.LayeredFolderStruct{
		{FolderName: "bin"},
		{FolderName: "config"},
		{FolderName: "constants"},
		{FolderName: "controllers"},
		{FolderName: "services"},
		{FolderName: "repositories"},
		{FolderName: "DTO"},
		{FolderName: "exceptions"},
		{FolderName: "middlewares"},
		{FolderName: "routes"},
		{FolderName: "utils"},
		{FolderName: "interfaces"},
		{FolderName: "mappers"},
	}
}
