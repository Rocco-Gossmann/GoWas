package core

import (
	"fmt"
	"syscall/js"
)

type RessourceType string
type RessourceState string
type RessourceHandle int

const (
	RESTYPE_AUDIO  RessourceType = "audio"
	RESTYPE_BINARY RessourceType = "binary"
)

const (
	RESSTATE_LOADING    RessourceState = "loading"
	RESSTATE_PROCESSING RessourceState = "processing"
	RESSTATE_READY      RessourceState = "ready"
	RESSTATE_NOTFOUND   RessourceState = "notfound"
)

type Ressource struct {
	handle        RessourceHandle
	ressourceType RessourceType
	state         RessourceState
	jsData        js.Value
	binData       []byte
}

func (me *Ressource) _Process() {
	if me.state != RESSTATE_PROCESSING {
		return
	}

	switch me.ressourceType {
	case RESTYPE_BINARY:
		me.binData = make([]byte, me.jsData.Get("length").Int())
		js.CopyBytesToGo(me.binData, me.jsData)
		me.jsData = js.ValueOf(nil)
		me.state = RESSTATE_READY

	default:
		panic(fmt.Sprintf("unprocessable Ressource type: %v", me.ressourceType))
	}
}

func _RequestRessource(resourceType RessourceType, fileName string) Ressource {
	ressource := Ressource{}
	ressource.state = RESSTATE_LOADING
	ressource.ressourceType = resourceType
	ressource.jsData = js.ValueOf(nil)

	response := js.Global().Get("requestRessource").Invoke(string(resourceType), fileName)
	ressourceHandle := response.Int()

	if ressourceHandle == -1 {
		panic("tried to request a ressource of unknown type")
	}

	ressource.handle = RessourceHandle(ressourceHandle)

	return ressource
}
