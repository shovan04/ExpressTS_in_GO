package layered

import "github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"

func GetLayeredFolders() []types.LayeredFolderStruct {
	return []types.LayeredFolderStruct{
		{FolderName: "bin"},
		{FolderName: "configs"},
		{FolderName: "constants"},
		{FolderName: "controllers"},
		{FolderName: "services"},
		{FolderName: "repositories"},
		{FolderName: "DTOClasses"},
		{FolderName: "exceptions"},
		{FolderName: "middlewares"},
		{FolderName: "routes"},
		{FolderName: "utilities"},
		{FolderName: "interfaces"},
		{FolderName: "mappers"},
	}
}
