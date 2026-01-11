package ddd

import "github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"

func GetDDDFolders() []types.LayeredFolderStruct {
	return []types.LayeredFolderStruct{
		{FolderName: "domain"},
		{FolderName: "domain/entities"},
		{FolderName: "domain/repositories"},
		{FolderName: "domain/value-objects"},
		{FolderName: "application"},
		{FolderName: "application/use-cases"},
		{FolderName: "application/dtos"},
		{FolderName: "application/interfaces"},
		{FolderName: "infrastructure"},
		{FolderName: "infrastructure/persistence"},
		{FolderName: "infrastructure/external"},
		{FolderName: "interfaces"},
		{FolderName: "interfaces/http"},
		{FolderName: "interfaces/http/controllers"},
		{FolderName: "interfaces/http/routes"},
		{FolderName: "interfaces/http/middlewares"},
	}
}
