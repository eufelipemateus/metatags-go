package metatags

import (
	"log"
	"reflect"
	"testing"
)

func TestGetMetatags(t *testing.T){
	metas :=	  GetMetatagsFromURL("https://felipemateus.com")

	if reflect.TypeOf(metas).String() != "map[string]string"{
		log.Panic("Não foi possivel retornar metatags")
	}
}