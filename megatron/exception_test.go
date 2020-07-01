package exception

import (
	"log"
	"testing"
)

func TestTry(t *testing.T) {
	Try(func() {
		log.Println("try...")
		Throw(2,"error2")
	}).Catch(1, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(2, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Finally(func() {
		log.Println("finally")
	})
}