package main

import (
	"fmt"
	icon "moviesDB/icons"
	"moviesDB/modele"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

// click droit puis open in Integrated Terminal
// pour générer les ressources statiques 	: go generate
// pour executer en dev. 					: go run .*  ( go run *.go  |  go run . )
// pour compiler 							: go build -o moviesDB
//
//	: go build -o /home/jluc20mx/moviesDB/moviesDB
//
//	: make run
//	: go mod init moviesDB
//	: go mod tidy

const appDevPath = "/media/jluc20mx/WorkSpace/GO/moviesDB_Fyne/"

var moviesDBDir string
var dbCollecsName string
var appRep string
var err error

var about_text = `
Auteur: J-L THOMAS
Version: 1

Cette application permet de gérer des collections de films ou de séries.

`

var help_text = `
This app can be used to manage tasks.

Each task has parameters (which can be displayed in the Details tab), at least a name, a due date and an "ahead" number, which defaults to 30 days.
Different colors are used to indicate the status of a task.
- Green:  due date is more than 
          "ahead" days away 
          (states future and soon)
- Yellow: due date is less than 
          "ahead" days away (state now)
- Red:    due date has passed already
          (state past)
- Grey:   task is marked as "done"

(Deleted tasks are not displayed, but kept in memory and synchronised with the server)


----------- Menus -----------

** File-Menu **
Export: export your tasks to a file (e.g. for a backup)
Import: import your tasks from e.g. a backup file

** Settings-Menu **
Choose light or dark theme and set connections parameters (IP4 Address and Port) of your sync-server. 
The Apply Button saves the connection parameters in a local file.

** Help-Menu **
Show help text and app info.
  

----------- Tabs ------------
  
** Tasks-Tab **
Tasks are displayed in a scrolled list. The button "Show all" resp. "Apply Filter" is used to show tasks, either all, or only tasks that match the filter criteria.
Icons can be used to add a new task (+), delete the currently selected task (-), copy and edit(modify) the currently selected task. 
The button "Save" is used to store all tasks (in an app-internal file). Note that without saving, all changes (add, delete, ...) are lost when the app is closed.
  
** Details-Tab **
Display and optionally change the details of the selected (or new) task. Use the icons to accept the changes, or display details of the previous/next task of the currently displayed list.
The "Done" checkbox can be used to mark/unmark a task as done.
Mandatory fields are pre-filled when the "add task" icon is used.
New owners and new categories can used, but these are only visible in the filter tab after restart of the app.

** Filter-Tab **
Set filter criteria, which apply if the "Apply Filter" button of the tasks-tab is used.
Status "soon"/"future":  is less/more than 1 month away.
Filter criteria are: state, priority, category, owner.

** Sync-Tab **
(Usable only if a sync-server is running)
Use the "Start" button to sync the tasks with an external server: all tasks are sent to the server (updated in the server) and a new task list is received. This new list is automatically stored on the internal file and then displayed.

`

func main() {
	//test()

	//	defer db.Close()
	appRep, err = getAppPath()
	if err != nil {
		panic(fmt.Sprint("  init:getAppPath : ", err))
	}

	moviesDBDir = filepath.Join(appRep, "Collections")
	dbCollecsName = filepath.Join(moviesDBDir, "Collections.sqlite")

	myApp := app.New()
	myApp.SetIcon(icon.LogoIco)
	myWindow := myApp.NewWindow("moviesDB")
	myUI := &appUI{
		app: myApp,
		win: myWindow,
	}

	if _, err := os.Stat(dbCollecsName); os.IsNotExist(err) {
		modele.OpenDBCollec(dbCollecsName)
		modele.CreateCollecTable()
	} else {
		modele.OpenDBCollec(dbCollecsName)
		collection := modele.GetCollecByID(1)
		myUI.collec = collection.Name
		myUI.dbType = collection.Type
		myUI.dbName = filepath.Join(moviesDBDir, collection.Name+".sqlite")
		myUI.pathImg = filepath.Join(moviesDBDir, collection.Name)
	}
	modele.OpenDB(myUI.dbName)

	myWindow.SetMainMenu(myUI.makeMenu())
	myWindow.SetContent(myUI.makeUI())
	myWindow.Resize(fyne.NewSize(1280, 1100))
	myApp.Settings().SetTheme(theme.DefaultTheme())
	myWindow.ShowAndRun()
}

// récupère le répertoire de l'application
func getAppPath() (string, error) {
	appPath, err := os.Executable()
	if err == nil {
		appPath = filepath.Dir(appPath)
		fmt.Println("appPath - os.Executable()           : " + appPath)

		appPath, err = filepath.EvalSymlinks(appPath)
		if err == nil {
			fmt.Println("appPath - filepath.EvalSymlinks(...): " + appPath)

			parts := strings.Split(appPath, string(os.PathSeparator))
			fmt.Println(parts)
			fmt.Println(filepath.Join(parts...))
			if parts[1] == "tmp" && strings.Contains(parts[2], "go-build") {
				fmt.Println("Session de DEV.")
				appPath = appDevPath
			}
			if parts[len(parts)-1] == "__debug_bin" {
				fmt.Println("Session de DEBUG")
				appPath = appDevPath
			}
		}
	}
	return appPath, err
}
