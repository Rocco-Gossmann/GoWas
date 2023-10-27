package core

import (
	"syscall/js"
)

type RessourceType string
type RessourceHandle int

const (
	RESTYPE_AUDIO  RessourceType = "audio"
	RESTYPE_BINARY RessourceType = "binary"
)

type Ressource struct {
	handle        RessourceHandle
	ressourceType RessourceType
}

func _RequestRessource(resourceType RessourceType, fileName string) Ressource {
	ressource := Ressource{}
	ressource.ressourceType = resourceType

	response := js.Global().Get("requestRessource").Invoke(string(resourceType), fileName)
	ressourceHandle := response.Int()

	if ressourceHandle == -1 {
		panic("tried to request a ressource of unknown type")
	}

	ressource.handle = RessourceHandle(ressourceHandle)

	return ressource
}
