package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pscohn/gonote/models"
)

var (
	list        = flag.Bool("l", false, "list all categories")
	newNote     = flag.Bool("n", false, "save a new note")
	newCategory = flag.Bool("c", false, "create a category")
	get         = flag.Bool("g", false, "get notes from category")
	execute     = flag.Bool("e", false, "execute command")
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

func printNotes(notes *[]models.Note, heading string) {
	if len(*notes) > 0 {
		fmt.Println("----------------")
		fmt.Println(heading)
		fmt.Println("----------------")
		for _, n := range *notes {
			fmt.Println(n.Id, n.Note)
		}
		fmt.Println("")
	} else {
		fmt.Println("no notes")
	}
}

func getAllNotes() {
	notes := []models.Note{}
	db.DB.Find(&notes)
	printNotes(&notes, "all notes")
}

func getNotesForCategory(arg string) {
	destCategory := models.Category{}
	db.DB.Where("Name = ?", arg).First(&destCategory)
	notes := []models.Note{}
	db.DB.Where("category_id = ?", destCategory.Id).Find(&notes)
	printNotes(&notes, arg)
}

func executeCommand(arg string) {
	findNote := models.Note{}
	db.DB.First(&findNote, arg)

	split := strings.Split(findNote.Note, " ")
	prog := split[0]
	rest := strings.Join(split[1:], " ")

	fmt.Println(prog, rest)
	var e bytes.Buffer
	cmd := exec.Command(prog, split[1:]...)
	cmd.Stderr = &e
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(string(e.Bytes()))
		return
	}
	fmt.Println(string(stdout))
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

	if *get {
		if arg == "" {
			getAllNotes()
			return
		}
		getNotesForCategory(arg)
		return
	}

	if *execute {
		executeCommand(arg)
		return
	}
}

func main() {
	db.Setup()
	db.Connect()
	defer db.Close()
	flag.Parse()
	arg := strings.Join(flag.Args(), " ")
	runCmd(arg)
}
