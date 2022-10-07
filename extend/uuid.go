package extend

import "github.com/satori/go.uuid"

func Uuid() string {
	id := uuid.NewV4()
   	ids := id.String()
	return ids
}