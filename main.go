package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pscohn/gonote/models"
)

var (
	list        = flag.Bool("l", false, "list all categories")
	newNote     = flag.Bool("n", false, "save a new note")
	newCategory = flag.Bool("c", false, "create a category")
	get         = flag.String("g", "", "get notes from category")
	dest        = flag.String("d", "", "destination for note")
	db          models.Database
)

func listCategories() {
	allCategories := []models.Category{}
	db.DB.Find(&allCategories)
	if len(allCategories) > 0 {
		for _, c := range allCategories {
			fmt.Println(c.Name)
		}
	} else {
		fmt.Println("no categories. create one with `gonote -c [category]`")
	}
}

func createCategory(arg string) {
	category := models.Category{Name: arg}
	db.DB.Create(&category)
	fmt.Printf("created category \"%v\"\n", arg)
}

func createNote(arg string) {
	if *dest == "" {
		//	TODO: handle error better
		return
	}
	destCategory := models.Category{}
	db.DB.Where("Name = ?", *dest).First(&destCategory)
	note := models.Note{CategoryID: destCategory.Id, Note: arg}
	db.DB.Save(&note)
	fmt.Printf("added to %v: \"%v\"\n", *dest, arg)
}

func getNotes() {
	if *get == "" {
		fmt.Println("no category specified")
		return
	}
	destCategory := models.Category{}
	db.DB.Where("Name = ?", *dest).First(&destCategory)
	notes := []models.Note{}
	db.DB.Where("category_id = ?", destCategory.Id).Find(&notes)
	if len(notes) > 0 {
		fmt.Println("----------------")
		fmt.Println(*dest)
		fmt.Println("----------------")
		for _, n := range notes {
			fmt.Println(n.Note)
		}
	} else {
		fmt.Println("no notes")
	}
}

func runCmd(arg string) {

	if *list {
		listCategories()
		return
	}

	if *newCategory {
		createCategory(arg)
		return
	}

	if *newNote {
		createNote(arg)
		return
	}

	getNotes()
	return
}

func main() {
	db.Connect()
	_ = db.Setup()
	defer db.Close()
	flag.Parse()
	arg := strings.Join(flag.Args(), " ")
	runCmd(arg)
}
