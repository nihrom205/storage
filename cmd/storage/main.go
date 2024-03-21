package main

import (
	"fmt"
	"log"
)
import "github.com/nihrom205/ozonCurses/05-folders/internal/storage"

func main() {

	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoreFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it restore: ", restoreFile)
}
