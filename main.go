package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pscohn/gonote/models"
)

var (
	newNote     = flag.Bool("n", false, "save a new note")
	newCategory = flag.Bool("c", false, "create a category")
	get         = flag.Bool("g", false, "get recent notes")
	dest        = flag.String("d", "", "destination for note")
	db          models.Database
)

func runCmd(arg string) {

	if *newCategory {
		category := models.Category{Name: arg}
		db.DB.Create(&category)
		fmt.Printf("created category \"%v\"\n", arg)
		return
	}

	if *dest == "" {
		//	TODO: handle error better
		return
	}

	fmt.Println("destination:", *dest)
	if *newNote {
		note := models.Note{Id: 0, Note: arg}
		db.DB.Save(&note)
		fmt.Printf("added to %v: \"%v\"\n", *dest, arg)
	} else if *get {
		fmt.Println("get note:", arg)
	}
}

func main() {
	db.Connect()
	_ = db.Setup()
	defer db.Close()
	flag.Parse()
	arg := strings.Join(flag.Args(), " ")
	runCmd(arg)
}
