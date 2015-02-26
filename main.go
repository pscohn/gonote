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
	get         = flag.String("g", "", "get recent notes")
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

	if *newNote {
		if *dest == "" {
			//	TODO: handle error better
			return
		}
		destCategory := models.Category{}
		db.DB.Where("Name = ?", *dest).First(&destCategory)
		note := models.Note{CategoryID: destCategory.Id, Note: arg}
		db.DB.Save(&note)
		fmt.Printf("added to %v: \"%v\"\n", *dest, arg)
	} else {
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
}

func main() {
	db.Connect()
	_ = db.Setup()
	defer db.Close()
	flag.Parse()
	arg := strings.Join(flag.Args(), " ")
	runCmd(arg)
}
