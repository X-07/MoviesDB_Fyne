package main

import (
	"image/color"
	"io"
	icon "moviesDB/icons"
	"moviesDB/importing"
	"moviesDB/modele"
	"moviesDB/utils"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type appUI struct {
	collec  string
	dbName  string
	dbType  string
	pathImg string

	app  fyne.App
	win  fyne.Window
	list *widget.List

	headerIdent       *widget.Button
	headerTitle       *widget.Entry
	headerDownloadBtn *widget.Button

	ficheOriginalTitle *widget.Entry
	ficheDateSortie    *widget.Entry
	ficheBtnDirectors  *widget.Button
	ficheDirectors     *widget.Entry
	ficheDuration      *widget.Entry
	ficheAgeMini       *widget.Entry
	ficheBtnCountries  *widget.Button
	ficheCountries     *widget.Entry
	ficheBtnGenres     *widget.Button
	ficheGenres        *widget.Entry
	ficheBtnActors     *widget.Button
	ficheActors        *widget.Entry
	fichePicture       *canvas.Image
	ficheChangerBtn    *widget.Button
	ficheSynopsis      *widget.Entry
	ficheDeleteBtn     *widget.Button

	detailsSeen        *widget.Check
	detailsBadMovie    *widget.Check
	detailsRating      *ratingEntry
	detailsStar        []*extIcon
	detailsRatingPress *ratingEntry
	detailsStarPress   []*extIcon
	detailsDateAjout   *widget.Entry
	detailPicture      *canvas.Image
	detailsComment     *widget.Entry

	infosControle          *widget.Check
	infosReplace           *widget.Check
	infosReplaceInProgress *widget.Check
	infosSupprime          *widget.Check
	infosMissing           *widget.Check
	infosToReEncode        *widget.Check
	infosTimeLag           *widget.Check
	infosBADQuality        *widget.Check
	infosTS                *widget.Check
	infosMD                *widget.Check
	infosSound             *widget.Check
	infosVFQ               *widget.Check
	infosVOSTFR            *widget.Check
	infosOtherPb           *widget.Entry
	infosRIPQuality        *widget.RadioGroup
	infosEncQuality        *widget.RadioGroup
	infosSource            *widget.RadioGroup
	infosFileSize          *widget.Entry
	infosContainer         *widget.SelectEntry
	infosBitRateT          *widget.Entry
	infosCodecV            *widget.SelectEntry
	infos3DTypeLabel       *canvas.Text
	infos3DType            *widget.SelectEntry
	infosFrameRate         *widget.Entry
	infosBitRateV          *widget.Entry
	infosWidth             *widget.Entry
	infosHeight            *widget.Entry
	infosCodecA            *widget.SelectEntry
	infosAudio             *widget.SelectEntry
	infosEchantillonnage   *widget.Entry
	infosBitRateA          *widget.Entry
	infosBtnSubtitles      *widget.Button
	infosSubtitles         *widget.Entry

	popupContainer *container.Split
	popupTitle     *widget.Label
	popupBtnClose  *clickIcon
	popupSlice     []string
	popupList      *widget.List
	popupSaisie    *widget.Entry
	popupSlice2    []string
	popupList2     *widget.List
}

var moviesList []modele.MovieList

var movie modele.Movie
var movieUpdated bool

var space *widget.Label = widget.NewLabel("")

// const emptySelectEntry4 string = "Saisir au moins 4 caractères SVP."
// const emptySelectEntry2 string = "Saisir au moins 2 caractères SVP."

/*****************************************************************************/
/*                         Create UI                                         */
/*****************************************************************************/
func (ui *appUI) makeUI() fyne.CanvasObject {
	ui.createPopupSaisie()

	boxHead := ui.createHeader()
	boxFiche := ui.createFicheTab()
	boxDetails := ui.createDetailsTab()
	boxInfos := ui.createInfosTab()

	navigateur := ui.createNavigateur()

	/* Tabs */
	tabsFiche := container.NewTabItem("Fiche", boxFiche)
	tabsDetails := container.NewTabItem("Details", boxDetails)
	tabsInfos := container.NewTabItem("Infos", boxInfos)
	tabsBar := container.NewAppTabs(tabsFiche, tabsDetails, tabsInfos)
	tabsBar.SetTabLocation(container.TabLocationTop)
	//tabsBarScroll := container.NewVScroll(tabsBar)
	content := container.NewBorder(boxHead, nil, nil, nil, tabsBar)

	split := container.NewHSplit(navigateur, content)
	split.Offset = 0.3

	return split
}

/*****************************************************************************/
/*                         Create Menu                                       */
/*****************************************************************************/
func (ui *appUI) makeMenu() *fyne.MainMenu {
	var collecMenu *fyne.Menu
	emptyItem := fyne.NewMenuItem("Vide", func() {})

	importItem := fyne.NewMenuItem("Import", func() {})
	importItemGCStar := fyne.NewMenuItem("GCStar (*.gcs)", func() {
		ui.menuImportGCStar(collecMenu)
	})
	importItem.ChildMenu = fyne.NewMenu("",
		importItemGCStar,
	)

	exportItem := fyne.NewMenuItem("Export", func() {})
	exportItemXML := fyne.NewMenuItem("XML (*.xml)", func() {})
	exportItem.ChildMenu = fyne.NewMenu("",
		exportItemXML,
	)

	darkItem := fyne.NewMenuItem("Dark Theme", func() {
		//ui.app.Settings().SetTheme(theme.DarkTheme())
	})

	lightItem := fyne.NewMenuItem("Light Theme", func() {
		//ui.app.Settings().SetTheme(theme.LightTheme())
	})

	collecManageItem := fyne.NewMenuItem("Gérer", func() {})
	collecMenu = fyne.NewMenu("Collection", collecManageItem, fyne.NewMenuItemSeparator())
	//var collecItem []*fyne.MenuItem
	for _, collec := range modele.GetCollecList() {
		name := collec.Name
		id := collec.ID
		//collecItem[idx] := fyne.NewMenuItem(collec.Name, func() { ui.menuDispCollec(collec.Name) })
		collecMenu.Items = append(collecMenu.Items, fyne.NewMenuItem(name, func() { ui.menuDispCollec(id) }))
	}

	aboutItem := fyne.NewMenuItem("About", func() {
		dialog.ShowInformation("About", about_text, ui.win)
	})

	helpItem := fyne.NewMenuItem("Help", func() {
		helpWindow := ui.app.NewWindow("Help")
		helpWindow.SetContent(menuHelp())
		helpWindow.Resize(fyne.NewSize(480, 720))
		helpWindow.Show()
	})

	return fyne.NewMainMenu(
		fyne.NewMenu("Fichier", importItem, exportItem),
		fyne.NewMenu("Edition", emptyItem),
		fyne.NewMenu("Filtre", emptyItem),
		fyne.NewMenu("Configuration", lightItem, darkItem, fyne.NewMenuItemSeparator(), emptyItem),
		collecMenu,
		fyne.NewMenu("Aide", aboutItem, helpItem),
	)
}

func (ui *appUI) menuImportGCStar(collecMenu *fyne.Menu) {
	var collecModal *widget.PopUp
	var collecType *widget.SelectEntry
	var collecGCStarSelect *widget.Button
	var collecProgress *widget.ProgressBarInfinite
	var collecError *widget.Label

	title := newBoldLeftLabel("---  Import GCStar  ---")
	btnClose := newClickIcon()
	btnClose.Resource = theme.CancelIcon()
	btnClose.OnTapped = func() {
		collecModal.Hide()
	}
	ligTitle := container.NewHBox(&layout.Spacer{}, title, &layout.Spacer{}, btnClose)

	collecName := widget.NewEntry()
	collecName.OnChanged = func(saisie string) {
		if saisie != "" {
			if collecType.Text != "" {
				collecGCStarSelect.Enable()
			}
		} else {
			collecGCStarSelect.Disable()
		}
	}

	options := []string{"Movies", "Movies3D", "Séries"}
	collecType = widget.NewSelectEntry(options)
	collecType.OnChanged = func(saisie string) {
		trouve := false
		for _, val := range options {
			if saisie == val {
				trouve = true
				if collecName.Text != "" {
					collecGCStarSelect.Enable()
				}
			}
		}
		if !trouve {
			collecType.SetText("")
			collecGCStarSelect.Disable()
		}
	}

	collecGCStarSelect = widget.NewButton("Sélectionner", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, ui.win)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()
			collecProgress.Start()
			collec := collecName.Text
			fic := reader.URI().Path()
			switch collecType.Text {
			case "Movies":
				importing.LoadMovies(fic, filepath.Join(moviesDBDir, collec+".sqlite"))
			case "Movies3D":
				importing.LoadMovies(fic, filepath.Join(moviesDBDir, collec+".sqlite"))

			}
			dirName := filepath.Dir(fic)
			fileName := filepath.Base(fic)
			fileExt := filepath.Ext(fic)
			fileName = "." + strings.TrimSuffix(fileName, fileExt) + "_pictures"
			filePath := filepath.Join(moviesDBDir, collec)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				if err := os.MkdirAll(filePath, 0755); err != nil {
					collecError.SetText(err.Error())
					collecProgress.Stop()
					return
				}
			}

			for _, movieList := range modele.GetMoviesList() {
				picture := modele.GetMoviesByID(movieList.ID).Picture
				if picture != "" {
					sourcePath := filepath.Join(dirName, fileName, picture)
					destPath := filepath.Join(moviesDBDir, collec, picture)
					out, err := os.Create(destPath)
					if err != nil {
						collecError.SetText(err.Error())
						collecProgress.Stop()
						return
					}

					defer out.Close()

					in, err := os.Open(sourcePath)
					if err != nil {
						collecError.SetText(err.Error())
						collecProgress.Stop()
						return
					}
					defer in.Close()

					_, err = io.Copy(out, in)
					if err != nil {
						collecError.SetText(err.Error())
						collecProgress.Stop()
						return
					}
				}
			}
			collecError.SetText("Import OK")
			collection := modele.Collections{
				Name: collec,
				Type: collecType.Text,
			}
			modele.InsertCollec(&collection)

			collecItem := fyne.NewMenuItem(collec, func() { ui.menuDispCollec(collection.ID) })
			collecMenu.Items = append(collecMenu.Items, collecItem)
			collecMenu.Refresh()

			ui.menuDispCollec(int64(collection.ID))

			collecModal.Hide()
		}, ui.win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".gcs"}))
		baseDir, err := storage.ListerForURI(storage.NewFileURI(importing.BddGCStarDir))
		if err == nil {
			fd.SetLocation(baseDir)
		}
		fd.Resize(fyne.NewSize(800, 600))
		fd.Show()
	})
	collecGCStarSelect.Disable()

	collectionForm := container.NewMax(widget.NewForm(
		&widget.FormItem{Text: "Nom de la nouvelle collection :", Widget: collecName},
		&widget.FormItem{Text: "Type de la collection :", Widget: collecType},
		&widget.FormItem{Text: "Choisir la collection GCStar à importer :", Widget: collecGCStarSelect},
	))

	collecProgress = widget.NewProgressBarInfinite()
	collecProgress.Stop()

	collecError = widget.NewLabel("")

	box := container.NewVBox(ligTitle, widget.NewLabel(" "), collectionForm, collecProgress, collecError)
	collecModal = widget.NewModalPopUp(box, ui.win.Canvas())
	//collecModal.Resize(fyne.NewSize(800, 600))
	collecModal.Show()
}

func (ui *appUI) menuDispCollec(id int64) {
	collection := modele.GetCollecByID(id)
	ui.collec = collection.Name
	ui.dbType = collection.Type
	ui.dbName = filepath.Join(moviesDBDir, collection.Name+".sqlite")
	ui.pathImg = filepath.Join(moviesDBDir, collection.Name)
	modele.Db.Close()
	modele.OpenDB(ui.dbName)
	moviesList = modele.GetMoviesList()
	ui.list.Refresh()
	ui.list.UnselectAll()
	ui.list.Select(0)
}

func menuHelp() fyne.CanvasObject {
	e := widget.NewMultiLineEntry()
	e.Text = help_text
	//e.Disable()
	e.Wrapping = fyne.TextWrapWord
	return e
}

/*****************************************************************************/
/*                         Create Navigateur                                 */
/*****************************************************************************/
func (ui *appUI) createNavigateur() fyne.CanvasObject {
	moviesList = modele.GetMoviesList()

	ui.list = widget.NewList(
		func() int {
			return len(moviesList)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(nil), widget.NewLabel(""))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			icon := obj.(*fyne.Container).Objects[0].(*widget.Icon)
			if moviesList[id].Seen {
				icon.Resource = theme.CheckButtonCheckedIcon()
			} else {
				icon.Resource = theme.CheckButtonIcon()
			}
			icon.Refresh()
			obj.(*fyne.Container).Objects[1].(*widget.Label).SetText(moviesList[id].Title)
		},
	)
	ui.list.OnSelected = func(id widget.ListItemID) {
		movie = modele.GetMoviesByID(int64(moviesList[id].ID))
		movieUpdated = false
		ui.loadHeader(id, &movie)
		ui.loadFicheTab(id, &movie)
		ui.loadDetailsTab(id, &movie)
		ui.loadInfosTab(id, &movie)
	}
	ui.list.OnUnselected = func(id widget.ListItemID) {
		ui.detailsSeen.OnChanged = func(checked bool) {}
		if movieUpdated {
			modele.UpdateMovies(&movie)
			dialog.ShowInformation("Update", "Base mise à jour : "+movie.Title, ui.win)
		}
		ui.razHeader()
		ui.razFicheTab()
		ui.razDetailsTab()
		ui.razInfosTab()
	}
	ui.list.Select(0)
	//list.SetItemHeight(5, 50)

	return ui.list
}

/*****************************************************************************/
/*                         Create Head Page                                  */
/*****************************************************************************/
func (ui *appUI) createHeader() fyne.CanvasObject {
	ui.headerIdent = widget.NewButton("", nil)
	ui.headerTitle = widget.NewEntry()
	ui.headerTitle.TextStyle = fyne.TextStyle{Bold: true}
	ui.headerDownloadBtn = widget.NewButton("Télécharger", nil)
	head := container.NewBorder(nil, nil, ui.headerIdent, ui.headerDownloadBtn, ui.headerTitle)

	return head
}

func (ui *appUI) loadHeader(id int, movie *modele.Movie) {
	ui.headerIdent.SetText(strconv.FormatInt(int64(movie.ID), 10))
	ui.headerTitle.SetText(movie.Title)
	ui.headerTitle.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Title = saisie
		moviesList[id].Title = saisie
		ui.list.Refresh()
	}
	ui.headerDownloadBtn.OnTapped = func() {
		ui.downloadFiche()
	}
}

func (ui *appUI) razHeader() {
	ui.headerTitle.OnChanged = func(saisie string) {}
}

/*****************************************************************************/
/*                         Create Fiche Tab                                  */
/*****************************************************************************/
func (ui *appUI) createFicheTab() fyne.CanvasObject {
	ficheOriginalTitleLabel := widget.NewLabelWithStyle("Titre original :", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
	ui.ficheOriginalTitle = widget.NewEntry()

	ui.ficheDateSortie = widget.NewEntry()
	ui.ficheBtnDirectors = widget.NewButtonWithIcon("Réalisateur :", theme.MenuDropDownIcon(), nil)
	ui.ficheDirectors = widget.NewEntry()
	ui.ficheDuration = widget.NewEntry()
	ui.ficheAgeMini = widget.NewEntry()
	ui.ficheBtnCountries = widget.NewButtonWithIcon("Nationalités :", theme.MenuDropDownIcon(), nil)
	ui.ficheCountries = widget.NewEntry()
	ui.ficheBtnGenres = widget.NewButtonWithIcon("Genres :", theme.MenuDropDownIcon(), nil)
	ui.ficheGenres = widget.NewEntry()
	ui.ficheBtnActors = widget.NewButtonWithIcon("Acteurs :", theme.MenuDropDownIcon(), nil)
	//ui.ficheActors = widget.NewMultiLineEntry()
	ui.ficheActors = widget.NewEntry()

	ficheForm := container.NewVBox(
		container.NewGridWithColumns(2,
			newBoldRightLabel("Date de sortie :"), ui.ficheDateSortie,
			container.NewHBox(layout.NewSpacer(), ui.ficheBtnDirectors), ui.ficheDirectors,
			newBoldRightLabel("Durée"), ui.ficheDuration,
			newBoldRightLabel("Age minimum :"), ui.ficheAgeMini,
			container.NewHBox(layout.NewSpacer(), ui.ficheBtnCountries), ui.ficheCountries,
			container.NewHBox(layout.NewSpacer(), ui.ficheBtnGenres), ui.ficheGenres,
			container.NewHBox(layout.NewSpacer(), ui.ficheBtnActors), ui.ficheActors,
		),
	)

	img := &canvas.Image{}
	img.SetMinSize(fyne.NewSize(480, 640))
	img.FillMode = canvas.ImageFillContain
	ui.fichePicture = img
	ui.ficheChangerBtn = widget.NewButton("Changer", nil)

	ui.ficheSynopsis = widget.NewMultiLineEntry()
	//ui.synopsis.Resize(fyne.NewSize(20, 200))

	ui.ficheDeleteBtn = widget.NewButton("Supprimer", nil)

	ficheL1 := container.NewBorder(nil, nil, ficheOriginalTitleLabel, nil, ui.ficheOriginalTitle)
	ficheL2_Form := container.NewVBox(ficheForm, layout.NewSpacer(), container.NewHBox(ui.ficheChangerBtn, layout.NewSpacer()))
	ficheL2 := container.NewBorder(ficheL1, widget.NewLabelWithStyle("Synopsis :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), container.NewPadded(ui.fichePicture), nil, ficheL2_Form)
	ficheL3 := container.NewHBox(layout.NewSpacer(), ui.ficheDeleteBtn)
	fiche := container.NewBorder(ficheL2, ficheL3, nil, nil, ui.ficheSynopsis)

	return fiche
}

func (ui *appUI) loadFicheTab(id int, movie *modele.Movie) {
	ui.ficheOriginalTitle.SetText(movie.OriginalTitle)
	ui.ficheOriginalTitle.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.OriginalTitle = saisie
	}
	ui.ficheDateSortie.SetText(movie.DateSortie)
	ui.ficheDateSortie.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.DateSortie = saisie
	}

	//ui.ficheDirectors.SetText(strings.Replace(movie.Directors, ",", "\n", -1))
	ui.ficheDirectors.SetText(movie.Directors)
	ui.ficheDirectors.OnChanged = func(saisie string) {
		if movie.Directors != saisie {
			movieUpdated = true
			movie.Directors = saisie
		}
	}
	ui.ficheBtnDirectors.OnTapped = func() {
		ui.popupTitle.SetText("Réalisateurs")
		ui.ficheDirectors.OnChanged = func(saisie string) {}
		if movie.Directors != "" {
			ui.popupSlice = strings.Split(movie.Directors, ", ")
		}
		ui.popupList.Refresh()
		ui.popupSaisie.SetText("")
		ui.popupSaisie.OnChanged = func(saisie string) {
			if len(saisie) > 2 {
				ui.popupSlice2 = modele.GetDirectorsList(saisie)
			} else {
				ui.popupSlice2 = []string{}
			}
			ui.popupList2.Refresh()
		}
		ui.ficheBtnDirectors.SetIcon(theme.MenuDropUpIcon())
		ui.ficheDirectors.SetText("")
		modal := widget.NewModalPopUp(ui.popupContainer, ui.win.Canvas())
		modal.Resize(fyne.NewSize(300, 400))
		ui.popupBtnClose.OnTapped = func() {
			ui.ficheDirectors.SetText(strings.Join(ui.popupSlice, ", "))
			ui.ficheBtnDirectors.SetIcon(theme.MenuDropDownIcon())
			ui.ficheDirectors.Show()
			if movie.Directors != ui.ficheDirectors.Text {
				movieUpdated = true
				movie.Directors = ui.ficheDirectors.Text
			}
			ui.ficheDirectors.OnChanged = func(saisie string) {
				if movie.Directors != saisie {
					movieUpdated = true
					movie.Directors = saisie
				}
			}
			modal.Hide()
		}
		modal.Show()
	}

	ui.ficheDuration.SetText(strconv.FormatInt(int64(movie.Duration), 10))
	ui.ficheDuration.OnChanged = func(saisie string) {
		movieUpdated = true
		value, err := strconv.Atoi(saisie)
		if err != nil {
			value = 0
		}
		movie.Duration = value
	}

	ui.ficheAgeMini.SetText(movie.AgeMini)
	ui.ficheAgeMini.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.AgeMini = saisie
	}

	//ui.ficheCountries.SetText(strings.Replace(movie.Countries, ",", "\n", -1))
	ui.ficheCountries.SetText(movie.Countries)
	ui.ficheCountries.OnChanged = func(saisie string) {
		if movie.Countries != saisie {
			movieUpdated = true
			movie.Countries = saisie
		}
	}
	ui.ficheBtnCountries.OnTapped = func() {
		ui.popupTitle.SetText("Nationalité")
		ui.ficheCountries.OnChanged = func(saisie string) {}
		if movie.Countries != "" {
			ui.popupSlice = strings.Split(movie.Countries, ", ")
		}
		ui.popupList.Refresh()
		ui.popupSaisie.SetText("")
		ui.popupSaisie.OnChanged = func(saisie string) {
			if len(saisie) > 2 {
				ui.popupSlice2 = modele.GetCountriesList(saisie)
			} else {
				ui.popupSlice2 = []string{}
			}
			ui.popupList2.Refresh()
		}
		ui.ficheBtnCountries.SetIcon(theme.MenuDropUpIcon())
		ui.ficheCountries.SetText("")
		modal := widget.NewModalPopUp(ui.popupContainer, ui.win.Canvas())
		modal.Resize(fyne.NewSize(280, 400))
		ui.popupBtnClose.OnTapped = func() {
			ui.ficheCountries.SetText(strings.Join(ui.popupSlice, ", "))
			ui.ficheBtnCountries.SetIcon(theme.MenuDropDownIcon())
			ui.ficheCountries.Show()
			if movie.Countries != ui.ficheCountries.Text {
				movieUpdated = true
				movie.Countries = ui.ficheCountries.Text
			}
			ui.ficheCountries.OnChanged = func(saisie string) {
				if movie.Countries != saisie {
					movieUpdated = true
					movie.Countries = saisie
				}
			}
			modal.Hide()
		}
		modal.Show()
	}

	//ui.ficheGenres.SetText(strings.Replace(movie.Genres, ",", "\n", -1))
	ui.ficheGenres.SetText(movie.Genres)
	ui.ficheGenres.OnChanged = func(saisie string) {
		if movie.Genres != saisie {
			movieUpdated = true
			movie.Genres = saisie
		}
	}
	ui.ficheBtnGenres.OnTapped = func() {
		ui.popupTitle.SetText("Genres")
		ui.ficheGenres.OnChanged = func(saisie string) {}
		if movie.Genres != "" {
			ui.popupSlice = strings.Split(movie.Genres, ", ")
		}
		ui.popupList.Refresh()
		ui.popupSaisie.SetText("")
		ui.popupSaisie.OnChanged = func(saisie string) {
			if len(saisie) > 2 {
				ui.popupSlice2 = modele.GetGenresList(saisie)
			} else {
				ui.popupSlice2 = []string{}
			}
			ui.popupList2.Refresh()
		}
		ui.ficheBtnDirectors.SetIcon(theme.MenuDropUpIcon())
		ui.ficheBtnGenres.SetIcon(theme.MenuDropUpIcon())
		ui.ficheGenres.SetText("")
		modal := widget.NewModalPopUp(ui.popupContainer, ui.win.Canvas())
		modal.Resize(fyne.NewSize(280, 400))
		ui.popupBtnClose.OnTapped = func() {
			ui.ficheGenres.SetText(strings.Join(ui.popupSlice, ", "))
			ui.ficheBtnGenres.SetIcon(theme.MenuDropDownIcon())
			ui.ficheGenres.Show()
			if movie.Genres != ui.ficheGenres.Text {
				movieUpdated = true
				movie.Genres = ui.ficheGenres.Text
			}
			ui.ficheGenres.OnChanged = func(saisie string) {
				if movie.Genres != saisie {
					movieUpdated = true
					movie.Genres = saisie
				}
			}
			modal.Hide()
		}
		modal.Show()
	}

	//ui.ficheActors.SetText(strings.Replace(movie.Actors, ",", "\n", -1))
	ui.ficheActors.SetText(movie.Actors)
	ui.ficheActors.OnChanged = func(saisie string) {
		if movie.Actors != saisie {
			movieUpdated = true
			movie.Actors = saisie
		}
	}
	ui.ficheBtnActors.OnTapped = func() {
		ui.popupTitle.SetText("Acteurs")
		ui.ficheActors.OnChanged = func(saisie string) {}
		if movie.Actors != "" {
			ui.popupSlice = strings.Split(movie.Actors, ", ")
		}
		ui.popupList.Refresh()
		//ui.popupSaisie.SetOptions(modele.GetActorsList())
		ui.popupSaisie.SetText("")
		ui.popupSaisie.OnChanged = func(saisie string) {
			if len(saisie) > 2 {
				ui.popupSlice2 = modele.GetActorsList(saisie)
			} else {
				ui.popupSlice2 = []string{}
			}
			ui.popupList2.Refresh()
		}
		ui.ficheBtnDirectors.SetIcon(theme.MenuDropUpIcon())
		ui.ficheBtnActors.SetIcon(theme.MenuDropUpIcon())
		ui.ficheActors.SetText("")
		modal := widget.NewModalPopUp(ui.popupContainer, ui.win.Canvas())
		modal.Resize(fyne.NewSize(300, 400))
		ui.popupBtnClose.OnTapped = func() {
			ui.ficheActors.SetText(strings.Join(ui.popupSlice, ", "))
			ui.ficheBtnActors.SetIcon(theme.MenuDropDownIcon())
			ui.ficheActors.Show()
			if movie.Actors != ui.ficheActors.Text {
				movieUpdated = true
				movie.Actors = ui.ficheActors.Text
			}
			ui.ficheActors.OnChanged = func(saisie string) {
				if movie.Actors != saisie {
					movieUpdated = true
					movie.Actors = saisie
				}
			}
			modal.Hide()
		}
		modal.Show()
	}

	ui.loadFichePicture(movie)

	ui.ficheSynopsis.SetText(movie.Synopsis)
	ui.ficheSynopsis.Wrapping = fyne.TextWrapWord
	ui.ficheSynopsis.Refresh()
	ui.ficheSynopsis.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Synopsis = saisie
	}

	// ui.ficheChangerBtn.OnTapped = func() {
	// 	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
	// 		if err != nil {
	// 			dialog.ShowError(err, ui.win)
	// 			return
	// 		}
	// 		if reader == nil {
	// 			return
	// 		}
	// 		defer reader.Close()
	// 		movieUpdated = true
	// 		resImg, _ := fyne.LoadResourceFromPath(reader.URI().Path())
	// 		movie.Picture = resImg.Name()
	//
	// 		ui.loadFichePicture(movie)
	// 		ui.loadDetailsPicture(movie)
	// 	}, ui.win)
	// 	fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
	// 	baseDir, err := storage.ListerForURI(storage.NewFileURI(pathImg))
	// 	if err == nil {
	// 		fd.SetLocation(baseDir)
	// 	}
	// 	fd.Resize(fyne.NewSize(800, 600))
	// 	fd.Show()
	// }
	ui.ficheChangerBtn.OnTapped = func() {
		ui.createPictureSelect(movie)
	}

	ui.ficheDeleteBtn.OnTapped = func() {
		dialog.ShowInformation("Delete", ui.headerTitle.Text, ui.win)
	}
}

func (ui *appUI) razFicheTab() {
	ui.ficheOriginalTitle.OnChanged = func(saisie string) {}
	ui.ficheDateSortie.OnChanged = func(saisie string) {}
	ui.ficheDirectors.OnChanged = func(saisie string) {}
	ui.ficheDuration.OnChanged = func(saisie string) {}
	ui.ficheAgeMini.OnChanged = func(saisie string) {}
	ui.ficheCountries.OnChanged = func(saisie string) {}
	ui.ficheGenres.OnChanged = func(saisie string) {}
	ui.ficheActors.OnChanged = func(saisie string) {}
	ui.ficheSynopsis.OnChanged = func(saisie string) {}
	ui.ficheChangerBtn.OnTapped = func() {}
	ui.ficheDeleteBtn.OnTapped = func() {}
}

func (ui *appUI) loadFichePicture(movie *modele.Movie) {
	ui.fichePicture.Resource, _ = fyne.LoadResourceFromPath(filepath.Join(ui.pathImg, movie.Picture))
	ui.fichePicture.Show()
	ui.fichePicture.Refresh()
}

/*****************************************************************************/
/*                         Create Details Tab                                */
/*****************************************************************************/
func (ui *appUI) createDetailsTab() fyne.CanvasObject {
	img := &canvas.Image{}
	img.SetMinSize(fyne.NewSize(240, 320))
	img.FillMode = canvas.ImageFillContain
	ui.detailPicture = img
	detailsCol1 := container.NewVBox(
		container.NewPadded(ui.detailPicture),
		newBoldLeftLabel("Commentaires :"),
	)

	ui.detailsSeen = widget.NewCheck("Vu", func(checked bool) {})
	ui.detailsBadMovie = widget.NewCheck("Nul", func(checked bool) {})
	ui.detailsRating = newRatingEntry()
	//ui.detailsRating.SetPlaceHolder("De 0 à 10")
	ui.detailsStar = make([]*extIcon, 10)
	for idx := 0; idx < 10; idx++ {
		ui.detailsStar[idx] = newExtIcon(icon.StarOnIco, icon.StarOnHoverIco, icon.StarOffIco, icon.StarOffHoverIco)
		ui.detailsStar[idx].indx = idx
		ui.detailsStar[idx].OnMouseIn = func(id int) {
			for i := 0; i < id; i++ {
				ui.detailsStar[i].MouseHoverIn()
			}
		}
		ui.detailsStar[idx].OnMouseOut = func(id int) {
			for i := 0; i < id; i++ {
				ui.detailsStar[i].MouseHoverOut()
			}
		}
		ui.detailsStar[idx].OnTapped = func(id int) {
			ui.detailsStar[id].tapped = true
			ui.detailsRating.SetText(utils.ItoA(id + 1))
			ui.loadDetailsNoteStar(id + 1)
		}
	}
	detailsNoteAndStars := container.NewHBox(ui.detailsRating, layout.NewSpacer(), widget.NewLabel("  "), layout.NewSpacer(), ui.detailsStar[0], ui.detailsStar[1], ui.detailsStar[2], ui.detailsStar[3], ui.detailsStar[4], ui.detailsStar[5], ui.detailsStar[6], ui.detailsStar[7], ui.detailsStar[8], ui.detailsStar[9])

	ui.detailsRatingPress = newRatingEntry()
	//ui.detailsRatingPress.SetPlaceHolder("De 0 à 10")
	ui.detailsStarPress = make([]*extIcon, 10)
	for idx := 0; idx < 10; idx++ {
		ui.detailsStarPress[idx] = newExtIcon(icon.StarOnIco, icon.StarOnHoverIco, icon.StarOffIco, icon.StarOffHoverIco)
		ui.detailsStarPress[idx].indx = idx
		ui.detailsStarPress[idx].OnMouseIn = func(id int) {
			//ui.detailsRatingPress.SetText(ItoA(id + 1))
			for i := 0; i < id; i++ {
				ui.detailsStarPress[i].MouseHoverIn()
			}
		}
		ui.detailsStarPress[idx].OnMouseOut = func(id int) {
			//ui.detailsRatingPress.SetText("  ")
			for i := 0; i < id; i++ {
				ui.detailsStarPress[i].MouseHoverOut()
			}
		}
		ui.detailsStarPress[idx].OnTapped = func(id int) {
			ui.detailsStarPress[id].tapped = true
			ui.detailsRatingPress.SetText(utils.ItoA(id + 1))
			ui.loadDetailsNotePresseStar(id + 1)
		}
	}
	detailsNotePresseAndStars := container.NewHBox(ui.detailsRatingPress, layout.NewSpacer(), widget.NewLabel("  "), layout.NewSpacer(), ui.detailsStarPress[0], ui.detailsStarPress[1], ui.detailsStarPress[2], ui.detailsStarPress[3], ui.detailsStarPress[4], ui.detailsStarPress[5], ui.detailsStarPress[6], ui.detailsStarPress[7], ui.detailsStarPress[8], ui.detailsStarPress[9])

	ui.detailsDateAjout = widget.NewEntry()
	detailsDateAjout := container.NewBorder(nil, nil, nil, widget.NewLabel("                                                            "), ui.detailsDateAjout)
	detailsCol2 := container.NewVBox(
		ui.detailsSeen,
		space,
		newBoldRightLabel("Note :"),
		newBoldRightLabel("Note Presse :"),
		space,
		newBoldRightLabel("Date d'ajout :"),
	)
	detailsCol3 := container.NewVBox(
		ui.detailsBadMovie,
		space,
		detailsNoteAndStars,
		detailsNotePresseAndStars,
		space,
		detailsDateAjout,
	)

	detailsL1 := container.NewHBox(detailsCol1, detailsCol2, detailsCol3, layout.NewSpacer())

	ui.detailsComment = widget.NewMultiLineEntry()
	details := container.NewBorder(detailsL1, nil, nil, nil, ui.detailsComment)

	return details
}

func (ui *appUI) loadDetailsTab(id int, movie *modele.Movie) {
	ui.detailsSeen.SetChecked(movie.Seen)
	ui.detailsSeen.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Seen = checked
		moviesList[id].Seen = checked
		ui.list.Refresh()
		if checked {
			ui.detailsRating.Enable()
			for idx := 0; idx < 10; idx++ {
				ui.detailsStar[idx].Enable()
			}
			value, err := strconv.Atoi(ui.detailsRating.Text)
			if err != nil {
				value = 0
			}
			ui.loadDetailsNoteStar(value)
		} else {
			ui.detailsRating.Disable()
			for idx := 0; idx < 10; idx++ {
				ui.detailsStar[idx].Disable()
			}
			ui.loadDetailsNoteStar(0)
		}
	}
	if movie.Seen {
		ui.detailsRating.Enable()
		for idx := 0; idx < 10; idx++ {
			ui.detailsStar[idx].Enable()
		}
	} else {
		ui.detailsRating.Disable()
		for idx := 0; idx < 10; idx++ {
			ui.detailsStar[idx].disabled = true
		}
	}

	ui.detailsBadMovie.SetChecked(movie.BadMovie)
	ui.detailsBadMovie.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.BadMovie = checked
		if checked {
			ui.detailsSeen.SetChecked(true)
			ui.detailsSeen.Disable()
			ui.detailsRating.Disable()
			for idx := 0; idx < 10; idx++ {
				ui.detailsStar[idx].Disable()
			}
			ui.loadDetailsNoteStar(0)
		} else {
			ui.detailsSeen.Enable()
			ui.detailsSeen.SetChecked(false)
		}
	}

	rating := "  "
	if movie.Rating > 0 {
		rating = utils.ItoA(movie.Rating)
	}
	ui.detailsRating.SetText(rating)
	ui.detailsRating.OnChanged = func(saisie string) {
		if saisie != "  " {
			ui.detailsRating.SetText(strings.Trim(saisie, " "))
		}
		value := utils.AtoI(strings.Trim(saisie, " "))
		if value > 10 {
			ui.detailsRating.SetText("10")
			value = 10
		}
		if movie.Rating != value {
			movieUpdated = true
			movie.Rating = value
		}
		ui.loadDetailsNoteStar(value)
	}
	if movie.Seen {
		ui.detailsRating.Enable()
		ui.loadDetailsNoteStar(movie.Rating)
	} else {
		ui.detailsRating.Disable()
		ui.loadDetailsNoteStar(0)
	}

	ratingPress := "  "
	if movie.RatingPress > 0 {
		ratingPress = utils.ItoA(movie.RatingPress)
	}
	ui.detailsRatingPress.SetText(ratingPress)
	ui.detailsRatingPress.OnChanged = func(saisie string) {
		if saisie != "  " {
			ui.detailsRatingPress.SetText(strings.Trim(saisie, " "))
		}
		value := utils.AtoI(strings.Trim(saisie, " "))
		if value > 10 {
			ui.detailsRatingPress.SetText("10")
			value = 10
		}
		if movie.RatingPress != value {
			movieUpdated = true
			movie.RatingPress = value
		}
		ui.loadDetailsNotePresseStar(value)
	}
	ui.loadDetailsNotePresseStar(movie.RatingPress)

	ui.detailsDateAjout.SetText(movie.DateAjout)
	ui.detailsDateAjout.OnChanged = func(saisie string) {
		//movieUpdated = true
		//movie.DateAjout = saisie
	}

	ui.loadDetailsPicture(movie)

	ui.detailsComment.SetText(movie.Comment)
	ui.detailsComment.Wrapping = fyne.TextWrapWord
	ui.detailsComment.Refresh()
	ui.detailsComment.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Comment = saisie
	}
}

func (ui *appUI) razDetailsTab() {
	ui.detailsSeen.OnChanged = func(checked bool) {}
	ui.detailsBadMovie.OnChanged = func(checked bool) {}
	ui.detailsRating.OnChanged = func(saisie string) {}
	ui.detailsRatingPress.OnChanged = func(saisie string) {}
	ui.detailsDateAjout.OnChanged = func(saisie string) {}
	ui.detailsComment.OnChanged = func(saisie string) {}
}

func (ui *appUI) loadDetailsPicture(movie *modele.Movie) {
	ui.detailPicture.Resource, _ = fyne.LoadResourceFromPath(filepath.Join(ui.pathImg + movie.Picture))
	ui.detailPicture.Show()
	ui.detailPicture.Refresh()
}

func (ui *appUI) loadDetailsNoteStar(value int) {
	for i := 0; i < 10; i++ {
		if i < value {
			ui.detailsStar[i].SetResource(icon.StarOnIco)
		} else {
			ui.detailsStar[i].SetResource(icon.StarOffIco)
		}
	}
}

func (ui *appUI) loadDetailsNotePresseStar(value int) {
	for i := 0; i < 10; i++ {
		if i < value {
			ui.detailsStarPress[i].SetResource(icon.StarOnIco)
		} else {
			ui.detailsStarPress[i].SetResource(icon.StarOffIco)
		}
	}
}

/*****************************************************************************/
/*                         Create Infos Tab                                  */
/*****************************************************************************/
func (ui *appUI) createInfosTab() fyne.CanvasObject {
	ui.infosControle = widget.NewCheck("Contrôlé", func(checked bool) {})
	problem := widget.NewLabel("Problème :")
	RIPQuality := widget.NewLabel("RIP Qualité :")
	encodageQuality := widget.NewLabel("Encodage Qualité :")
	source := widget.NewLabel("Source :")
	vbox1 := container.NewVBox(ui.infosControle, problem, space, space, space, RIPQuality, encodageQuality, source, space)

	ui.infosReplace = widget.NewCheck("A Remplacer !", func(checked bool) {})
	ui.infosReplaceInProgress = widget.NewCheck("Remplacement en cours", func(checked bool) {})
	ui.infosSupprime = widget.NewCheck("Supprimé : NUL", func(checked bool) {})
	infosL1 := container.NewHBox(ui.infosReplace, ui.infosReplaceInProgress, ui.infosSupprime)

	ui.infosMissing = widget.NewCheck("Absent", func(checked bool) {})
	ui.infosToReEncode = widget.NewCheck("A Réencoder", func(checked bool) {})
	ui.infosTimeLag = widget.NewCheck("Décalage Image/Son", func(checked bool) {})
	ui.infosBADQuality = widget.NewCheck("BAD Qualité", func(checked bool) {})
	infosL2 := container.NewHBox(ui.infosMissing, ui.infosToReEncode, ui.infosTimeLag, ui.infosBADQuality)

	ui.infosTS = widget.NewCheck("TS", func(checked bool) {})
	ui.infosMD = widget.NewCheck("MD", func(checked bool) {})
	ui.infosSound = widget.NewCheck("Son", func(checked bool) {})
	ui.infosVFQ = widget.NewCheck("VFQ", func(checked bool) {})
	ui.infosVOSTFR = widget.NewCheck("VOSTFR", func(checked bool) {})
	infosL3 := container.NewHBox(ui.infosTS, ui.infosMD, ui.infosSound, ui.infosVFQ, ui.infosVOSTFR)

	autrePb := widget.NewLabel("Autre PB :")
	ui.infosOtherPb = widget.NewEntry()
	infosL4 := container.NewBorder(nil, nil, autrePb, layout.NewSpacer(), ui.infosOtherPb)

	ui.infosRIPQuality = widget.NewRadioGroup([]string{"VCD", "VHS", "VGA", "DVD", "sVGA", "XGA", "HD", "FHD", "wqHD", "4K", "8K"}, func(saisie string) {})
	ui.infosRIPQuality.Horizontal = true

	ui.infosEncQuality = widget.NewRadioGroup([]string{"Light", "Good", "Hight"}, func(saisie string) {})
	ui.infosEncQuality.Horizontal = true

	ui.infosSource = widget.NewRadioGroup([]string{"TVRip", "VHS", "VHSRip", "HDRip", "DVD", "DVDRip", "BD", "BDRip", "BR-4K", "WEBRip", "WEB-DL"}, func(saisie string) {})
	ui.infosSource.Horizontal = true

	vbox2 := container.NewVBox(infosL1, infosL2, infosL3, infosL4, space, ui.infosRIPQuality, ui.infosEncQuality, ui.infosSource, space)
	block1 := container.NewHBox(vbox1, vbox2)

	ui.infosFileSize = widget.NewEntry()
	ui.infosCodecV = widget.NewSelectEntry([]string{})
	ui.infosFrameRate = widget.NewEntry()
	ui.infosWidth = widget.NewEntry()
	ui.infosCodecA = widget.NewSelectEntry([]string{})
	ui.infosEchantillonnage = widget.NewEntry()
	ui.infosBtnSubtitles = widget.NewButtonWithIcon("SousTitres :", theme.MenuDropDownIcon(), nil)
	ui.infosSubtitles = widget.NewEntry()

	ui.infosContainer = widget.NewSelectEntry([]string{})
	ui.infosBitRateT = widget.NewEntry()
	ui.infos3DTypeLabel = canvas.NewText("3D Type :  ", color.Black)
	ui.infos3DTypeLabel.Alignment = fyne.TextAlignTrailing
	ui.infos3DTypeLabel.TextStyle = fyne.TextStyle{Bold: true}
	ui.infos3DType = widget.NewSelectEntry([]string{"Side by Side", "Top Bottom"})
	ui.infosBitRateV = widget.NewEntry()
	ui.infosHeight = widget.NewEntry()
	ui.infosAudio = widget.NewSelectEntry([]string{})
	ui.infosBitRateA = widget.NewEntry()

	infosForm := container.NewVBox(
		container.NewGridWithColumns(4,
			newBoldRightLabel("Taille (Go) :"), ui.infosFileSize, newBoldRightLabel("Conteneur :"), ui.infosContainer,
			space, space, newBoldRightLabel("Débit Total (Kbps):"), ui.infosBitRateT,
			space, space, space, space,
			newBoldRightLabel("Codec Vidéo :"), ui.infosCodecV, ui.infos3DTypeLabel, ui.infos3DType,
			newBoldRightLabel("Cadence (fps) :"), ui.infosFrameRate, newBoldRightLabel("Débit Vidéo (Kbps) :"), ui.infosBitRateV,
			newBoldRightLabel("Largeur (px) :"), ui.infosWidth, newBoldRightLabel("Hauteur (px) :"), ui.infosHeight,
			space, space, space, space,
			newBoldRightLabel("Codec Audio :"), ui.infosCodecA, newBoldRightLabel("Audio :"), ui.infosAudio,
			newBoldRightLabel("Echantillonnage (KHz) :"), ui.infosEchantillonnage, newBoldRightLabel("Débit Audio (Kbps) :"), ui.infosBitRateA,
			space, space, space, space,
			container.NewHBox(layout.NewSpacer(), ui.infosBtnSubtitles), ui.infosSubtitles, space, space,
		),
		// container.NewGridWithColumns(4,
		// 	space, container.NewBorder(nil, ui.elmtL1, nil, nil, ui.elmtList), space, space,
		// ),
	)

	// //infos := container.NewBorder(block1, nil, nil, nil, grid)
	infos := container.NewVBox(block1, infosForm)

	return infos
}

func (ui *appUI) loadInfosTab(id int, movie *modele.Movie) {
	ui.infosControle.SetChecked(movie.Control)
	ui.infosControle.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Control = checked
	}

	ui.infosReplace.SetChecked(movie.Replace)
	ui.infosReplace.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Replace = checked
	}

	ui.infosReplaceInProgress.SetChecked(movie.ReplaceInProgress)
	ui.infosReplaceInProgress.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.ReplaceInProgress = checked
	}

	ui.infosSupprime.SetChecked(movie.Deleted)
	ui.infosSupprime.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Deleted = checked
	}

	ui.infosMissing.SetChecked(movie.Missing)
	ui.infosMissing.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Missing = checked
	}

	ui.infosToReEncode.SetChecked(movie.ToReEncode)
	ui.infosToReEncode.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.ToReEncode = checked
	}

	ui.infosTimeLag.SetChecked(movie.TimeLag)
	ui.infosTimeLag.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.TimeLag = checked
	}

	ui.infosBADQuality.SetChecked(movie.BADQuality)
	ui.infosBADQuality.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.BADQuality = checked
	}

	ui.infosTS.SetChecked(movie.TS)
	ui.infosTS.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.TS = checked
	}

	ui.infosMD.SetChecked(movie.MD)
	ui.infosMD.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.MD = checked
	}

	ui.infosSound.SetChecked(movie.Sound)
	ui.infosSound.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.Sound = checked
	}

	ui.infosVFQ.SetChecked(movie.VFQ)
	ui.infosVFQ.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.VFQ = checked
	}

	ui.infosVOSTFR.SetChecked(movie.VOSTFR)
	ui.infosVOSTFR.OnChanged = func(checked bool) {
		movieUpdated = true
		movie.VOSTFR = checked
	}

	ui.infosOtherPb.SetText(movie.OtherPb)
	ui.infosOtherPb.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.OtherPb = saisie
	}

	ui.infosRIPQuality.SetSelected(movie.RIPQuality)
	ui.infosRIPQuality.OnChanged = func(saisie string) {
		movieUpdated = true
		switch saisie {
		case "VCD":
			movie.RIPQuality = "VCD"
		case "VHS":
			movie.RIPQuality = "VHS"
		case "VGA":
			movie.RIPQuality = "VGA"
		case "DVD":
			movie.RIPQuality = "DVD"
		case "sVGA":
			movie.RIPQuality = "sVGA"
		case "XGA":
			movie.RIPQuality = "XGA"
		case "HD":
			movie.RIPQuality = "HD"
		case "fHD":
			movie.RIPQuality = "FHD"
		case "wqHD":
			movie.RIPQuality = "wqHD"
		case "4K":
			movie.RIPQuality = "4K"
		case "8K":
			movie.RIPQuality = "8K"
		}
	}

	ui.infosEncQuality.SetSelected(movie.EncQuality)
	ui.infosEncQuality.OnChanged = func(saisie string) {
		movieUpdated = true
		switch saisie {
		case "Light":
			movie.EncQuality = "Light"
		case "Good":
			movie.EncQuality = "VHS"
		case "Hight":
			movie.EncQuality = "Hight"
		case "Web-DL":
			movie.EncQuality = "Web-DL"
		}
	}

	ui.infosSource.SetSelected(movie.Source)
	ui.infosSource.OnChanged = func(saisie string) {
		movieUpdated = true
		switch saisie {
		case "TVRip":
			movie.RIPQuality = "TVRip"
		case "VHSRip":
			movie.RIPQuality = "VHSRip"
		case "HDRip":
			movie.RIPQuality = "HDRip"
		case "DVDRip":
			movie.RIPQuality = "DVDRip"
		case "BDRip":
			movie.RIPQuality = "BDRip"
		case "WEBRip":
			movie.RIPQuality = "WEBRip"
		case "WEB-DL":
			movie.RIPQuality = "WEB-DL"
		}
	}

	ui.infosFileSize.SetText(utils.FtoA(movie.FileSize))
	ui.infosFileSize.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.FileSize = utils.AtoF(saisie)
	}

	ui.infosCodecV.SetOptions(modele.GetCodecVideoList())
	ui.infosCodecV.SetText(movie.CodecV)
	ui.infosCodecV.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.CodecV = saisie
	}

	ui.infosFrameRate.SetText(movie.FrameRate)
	ui.infosFrameRate.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.FrameRate = saisie
	}

	ui.infosWidth.SetText(utils.ItoA(movie.Width))
	ui.infosWidth.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Width = utils.AtoI(saisie)
	}

	ui.infosCodecA.SetOptions(modele.GetCodecAudioList())
	ui.infosCodecA.SetText(movie.CodecA)
	ui.infosCodecA.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.CodecA = saisie
	}

	ui.infosEchantillonnage.SetText(utils.ItoA(movie.Sampling))
	ui.infosEchantillonnage.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Sampling = utils.AtoI(saisie)
	}

	ui.infosSubtitles.SetText(movie.Subtitles)
	ui.infosSubtitles.OnChanged = func(saisie string) {
		if movie.Subtitles != saisie {
			movieUpdated = true
			movie.Subtitles = saisie
		}
	}
	ui.infosBtnSubtitles.OnTapped = func() {
		ui.popupTitle.SetText("Sous-Titres")
		ui.infosSubtitles.OnChanged = func(saisie string) {}
		if movie.Subtitles != "" {
			ui.popupSlice = strings.Split(movie.Subtitles, ", ")
		}
		ui.popupList.Refresh()
		ui.popupSaisie.SetText("")
		ui.popupSlice2 = modele.GetSubtitleList()
		ui.popupList2.Refresh()
		ui.infosBtnSubtitles.SetIcon(theme.MenuDropUpIcon())
		ui.infosSubtitles.SetText("")
		modal := widget.NewModalPopUp(ui.popupContainer, ui.win.Canvas())
		modal.Resize(fyne.NewSize(240, 400))
		ui.popupBtnClose.OnTapped = func() {
			ui.infosSubtitles.SetText(strings.Join(ui.popupSlice, ", "))
			ui.infosBtnSubtitles.SetIcon(theme.MenuDropDownIcon())
			ui.infosSubtitles.Show()
			if movie.Subtitles != ui.infosSubtitles.Text {
				movieUpdated = true
				movie.Subtitles = ui.infosSubtitles.Text
			}
			ui.infosSubtitles.OnChanged = func(saisie string) {
				if movie.Subtitles != saisie {
					movieUpdated = true
					movie.Subtitles = saisie
				}
			}
			modal.Hide()
		}
		modal.Show()
	}

	ui.infosContainer.SetOptions(modele.GetContainerList())
	ui.infosContainer.SetText(movie.Container)
	ui.infosContainer.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Container = saisie
	}

	ui.infosBitRateT.SetText(utils.ItoA(movie.BitRateT))
	ui.infosBitRateT.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.BitRateT = utils.AtoI(saisie)
	}

	options := []string{"Side by Side", "Top Bottom"}
	if ui.dbType != "Movies3D" {
		ui.infos3DType.SetText("")
		ui.infos3DType.Disable()
		ui.infos3DTypeLabel.Color = &color.NRGBA{R: 160, G: 160, B: 160, A: 255}
		ui.infos3DTypeLabel.Refresh()
	} else {
		ui.infos3DType.SetText(movie.Type3D)
		ui.infos3DType.Enable()
		ui.infos3DTypeLabel.Color = color.Black
		ui.infos3DTypeLabel.Refresh()
	}
	ui.infos3DType.OnChanged = func(saisie string) {
		if ui.dbType != "Movies3D" {
			ui.infos3DType.SetText("")
		} else {
			if saisie != options[0] && saisie != options[1] {
				ui.infos3DType.SetText("")
			} else {
				movieUpdated = true
				movie.Type3D = saisie
			}
		}
	}

	ui.infosBitRateV.SetText(utils.ItoA(movie.BitRateV))
	ui.infosBitRateV.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.BitRateV = utils.AtoI(saisie)
	}

	ui.infosHeight.SetText(utils.ItoA(movie.Height))
	ui.infosHeight.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Height = utils.AtoI(saisie)
	}

	ui.infosAudio.SetOptions(modele.GetAudioList())
	ui.infosAudio.SetText(movie.Audio)
	ui.infosAudio.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.Audio = saisie
	}

	ui.infosBitRateA.SetText(utils.ItoA(movie.BitRateA))
	ui.infosBitRateA.OnChanged = func(saisie string) {
		movieUpdated = true
		movie.BitRateA = utils.AtoI(saisie)
	}
}

func (ui *appUI) razInfosTab() {
	ui.infosControle.OnChanged = func(checked bool) {}
	ui.infosReplace.OnChanged = func(checked bool) {}
	ui.infosReplaceInProgress.OnChanged = func(checked bool) {}
	ui.infosSupprime.OnChanged = func(checked bool) {}
	ui.infosMissing.OnChanged = func(checked bool) {}
	ui.infosToReEncode.OnChanged = func(checked bool) {}
	ui.infosTimeLag.OnChanged = func(checked bool) {}
	ui.infosBADQuality.OnChanged = func(checked bool) {}
	ui.infosTS.OnChanged = func(checked bool) {}
	ui.infosMD.OnChanged = func(checked bool) {}
	ui.infosSound.OnChanged = func(checked bool) {}
	ui.infosVFQ.OnChanged = func(checked bool) {}
	ui.infosVOSTFR.OnChanged = func(checked bool) {}
	ui.infosOtherPb.OnChanged = func(saisie string) {}
	ui.infosRIPQuality.OnChanged = func(saisie string) {}
	ui.infosEncQuality.OnChanged = func(saisie string) {}
	ui.infosSource.OnChanged = func(saisie string) {}
	ui.infosFileSize.OnChanged = func(saisie string) {}
	ui.infosCodecV.OnChanged = func(saisie string) {}
	ui.infosFrameRate.OnChanged = func(saisie string) {}
	ui.infosWidth.OnChanged = func(saisie string) {}
	ui.infosCodecA.OnChanged = func(saisie string) {}
	ui.infosEchantillonnage.OnChanged = func(saisie string) {}
	ui.infosSubtitles.OnChanged = func(saisie string) {}
	ui.infosContainer.OnChanged = func(saisie string) {}
	ui.infosBitRateT.OnChanged = func(saisie string) {}
	ui.infos3DType.OnChanged = func(saisie string) {}
	ui.infosBitRateV.OnChanged = func(saisie string) {}
	ui.infosHeight.OnChanged = func(saisie string) {}
	ui.infosAudio.OnChanged = func(saisie string) {}
	ui.infosBitRateA.OnChanged = func(saisie string) {}
}

/*****************************************************************************/
/*                         Create Popup Saisies multiples                    */
/*****************************************************************************/
func (ui *appUI) createPopupSaisie() {
	var popupBtnSupp *widget.Button
	//-------------Bouton close
	ui.popupTitle = newBoldRightLabel("")
	ui.popupBtnClose = newClickIcon()
	ui.popupBtnClose.Resource = theme.CancelIcon()
	popupHaut := container.NewHBox(layout.NewSpacer(), ui.popupTitle, layout.NewSpacer(), ui.popupBtnClose)

	//-------------Liste d'éléments choisis
	ui.popupSlice = []string{}
	ui.popupList = widget.NewList(
		func() int {
			return len(ui.popupSlice)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(ui.popupSlice[id])
		},
	)
	ui.popupList.OnSelected = func(id widget.ListItemID) {
		popupBtnSupp.OnTapped = func() {
			list := []string{}
			for idx, elmt := range ui.popupSlice {
				if idx != id {
					list = append(list, elmt)
				}
			}
			ui.popupSlice = list
			ui.popupList.Refresh()
			ui.popupList.UnselectAll()
		}
	}

	//-------------Ligne de saisie, d'ajout et de suppression
	ui.popupSaisie = widget.NewEntry()
	ui.popupSaisie.SetText("")
	popupBtnAdd := widget.NewButton("Ajouter", func() {
		if ui.popupSaisie.Text != "" {
			ui.popupSlice = append(ui.popupSlice, ui.popupSaisie.Text)
			ui.popupList.Refresh()
			ui.popupSaisie.SetText("")
		}
	})
	popupBtnSupp = widget.NewButton("Enlever", nil)
	popupSaisie := container.NewBorder(nil, nil, nil, container.NewHBox(popupBtnAdd, popupBtnSupp), ui.popupSaisie)

	//-------------Liste d'éléments à choisir
	ui.popupSlice2 = []string{}
	ui.popupList2 = widget.NewList(
		func() int {
			return len(ui.popupSlice2)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel(""))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(ui.popupSlice2[id])
		},
	)
	ui.popupList2.OnSelected = func(id widget.ListItemID) {
		ui.popupSlice = append(ui.popupSlice, ui.popupSlice2[id])
		ui.popupList.Refresh()
	}

	split := container.NewVSplit(
		container.NewBorder(popupHaut, popupSaisie, nil, nil, ui.popupList),
		ui.popupList2,
	)
	split.Offset = 0.6

	ui.popupContainer = split
}

/*****************************************************************************/
/*                         Create Picture Select                             */
/*****************************************************************************/
// func (ui *appUI) createPictureSelect() {
// 	title := newBoldRightLabel("Sélectionner une affiche")
// 	btnClose := newClickIcon()
// 	btnClose.Resource = theme.CancelIcon()
// 	popupHaut := container.NewHBox(layout.NewSpacer(), title, layout.NewSpacer(), btnClose)

// 	files, err := ioutil.ReadDir(pathImg)
// 	if err != nil {
// 		return
// 	}
// 	var listFiles []string
// 	sort.Strings(listFiles)
// 	for _, ficInfo := range files {
// 		fic := filepath.Join(pathImg, ficInfo.Name())
// 		fileInfo, err := os.Lstat(fic)
// 		if err != nil {
// 			return
// 		}
// 		if !fileInfo.Mode().IsDir() {
// 			if IsImageFile(filepath.Ext(fic)) {
// 				listFiles = append(listFiles, fic)
// 			}
// 		}
// 	}
// 	grid := container.NewGridWithColumns(7, layout.NewSpacer())
// 	for _, file := range listFiles {
// 		img := &canvas.Image{}
// 		img.SetMinSize(fyne.NewSize(120, 160))
// 		img.FillMode = canvas.ImageFillContain
// 		img.Resource, _ = fyne.LoadResourceFromPath(file)
// 		img.Show()
// 		//img.Refresh()
// 		grid.Add(img)
// 	}

//		modal := widget.NewModalPopUp(container.NewBorder(popupHaut, nil, nil, nil, container.NewVScroll(grid)), ui.win.Canvas())
//		btnClose.OnTapped = func() {
//			modal.Hide()
//		}
//		modal.Resize(fyne.NewSize(800, 600))
//		modal.Show()
//	}

type fileInfoStruct struct {
	name string
	date int64
}

func (ui *appUI) createPictureSelect(movie *modele.Movie) {
	var fileSelectModal *widget.PopUp
	var picture string
	var listFiles []fileInfoStruct
	var list *widget.List
	var selectKey string = ""
	var btnSortState bool = false

	title := newBoldRightLabel("Sélectionner l'affiche de ce film")
	btnSort := newClickIcon()
	btnSort.Resource = theme.MenuDropDownIcon()
	btnSort.OnTapped = func() {
		btnSortState = !btnSortState
		if btnSortState {
			btnSort.Resource = theme.MenuDropUpIcon()
			btnSort.Refresh()
			sort.Slice(listFiles, func(i, j int) bool {
				return listFiles[i].date > listFiles[j].date
			})
		} else {
			btnSort.Resource = theme.MenuDropDownIcon()
			btnSort.Refresh()
			sort.Slice(listFiles, func(i, j int) bool {
				return strings.ToLower(listFiles[i].name) < strings.ToLower(listFiles[j].name)
			})
		}
		createPictureListMoveTo(listFiles, list)
		list.Refresh()
	}
	fileSelectHaut := container.NewHBox(layout.NewSpacer(), title, layout.NewSpacer(), btnSort)

	img := &canvas.Image{}
	img.SetMinSize(fyne.NewSize(240, 320))
	fileSelectPicture := container.NewCenter(img)

	if listFiles = createPictureGetDirSelect(ui.pathImg); listFiles == nil {
		dialog.ShowInformation("Erreur", "Pb avec la lecture du répertoire : "+ui.pathImg, ui.win)
		return
	}

	list = widget.NewList(
		func() int {
			return len(listFiles)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.FileImageIcon()), widget.NewLabel(""))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[1].(*widget.Label).SetText(listFiles[id].name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		picture = listFiles[id].name
		img.FillMode = canvas.ImageFillContain
		img.Resource, _ = fyne.LoadResourceFromPath(filepath.Join(ui.pathImg, picture))
		img.Show()
		img.Refresh()
	}
	createPictureListMoveTo(listFiles, list)

	selectKeyLabel := widget.NewLabel("")
	btnSelect := widget.NewButton("Sélectionner", func() {
		movieUpdated = true
		movie.Picture = picture
		ui.loadFichePicture(movie)
		ui.loadDetailsPicture(movie)
		fileSelectModal.Hide()
	})
	btnCancel := widget.NewButton("Annuler", func() { fileSelectModal.Hide() })
	fileSelectBas := container.NewHBox(selectKeyLabel, layout.NewSpacer(), btnCancel, btnSelect)

	fileSelectModal = widget.NewModalPopUp(container.NewBorder(fileSelectHaut, fileSelectBas, nil, fileSelectPicture, list), ui.win.Canvas())
	fileSelectModal.Resize(fyne.NewSize(800, 600))
	fileSelectModal.Canvas.SetOnTypedKey(func(k *fyne.KeyEvent) {
		time.AfterFunc(5*time.Second, func() {
			selectKey = ""
			selectKeyLabel.SetText("")
		})
		selectKey += string(k.Name)
		selectKeyLabel.SetText(selectKey)
		for id, file := range listFiles {
			if strings.HasPrefix(strings.ToLower(file.name), strings.ToLower(selectKey)) {
				list.Select(id)
				if id > 0 {
					list.ScrollTo(id - 1)
				} else {
					list.ScrollTo(0)
				}
				break
			}
		}
	})
	fileSelectModal.Show()
}

func createPictureGetDirSelect(pathImg string) []fileInfoStruct {
	var listFiles []fileInfoStruct
	if files, err := os.ReadDir(pathImg); err != nil {
		return nil
	} else {
		for _, dirEntry := range files {
			fic := filepath.Join(pathImg, dirEntry.Name())
			fileInfo, err := os.Lstat(fic)
			if err != nil {
				return nil
			}
			if !fileInfo.Mode().IsDir() {
				if IsImageFile(filepath.Ext(fic)) {
					listFiles = append(listFiles, fileInfoStruct{name: fileInfo.Name(), date: fileInfo.ModTime().Unix()})
				}
			}
		}
		sort.Slice(listFiles, func(i, j int) bool {
			return strings.ToLower(listFiles[i].name) < strings.ToLower(listFiles[j].name)
		})
	}
	return listFiles
}

func createPictureListMoveTo(listFiles []fileInfoStruct, list *widget.List) {
	for id, file := range listFiles {
		if file.name == movie.Picture {
			list.Select(id)
			if id > 0 {
				list.ScrollTo(id - 1)
			} else {
				list.ScrollTo(0)
			}
			break
		}
	}
}

/*****************************************************************************/
/*                         Utils                                             */
/*****************************************************************************/
func newBoldRightLabel(text string) *widget.Label {
	//return widget.NewLabelWithStyle("Titre original :", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
	return &widget.Label{Text: text, Alignment: fyne.TextAlignTrailing, TextStyle: fyne.TextStyle{Bold: true}}
}

func newBoldLeftLabel(text string) *widget.Label {
	//return widget.NewLabelWithStyle("Titre original :", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	return &widget.Label{Text: text, Alignment: fyne.TextAlignLeading, TextStyle: fyne.TextStyle{Bold: true}}
}

//	func newBoldCenterLabel(text string) *widget.Label {
//		//return widget.NewLabelWithStyle("Titre original :", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
//		return &widget.Label{Text: text, Alignment: fyne.TextAlignCenter, TextStyle: fyne.TextStyle{Bold: true}}
//	}
//
// liste des extensions contenant des médias

var imgContainers = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}

func IsImageFile(ext string) bool {
	result := false
	if _, ok := imgContainers[strings.ToLower(ext)]; ok {
		result = true
	}

	return result
}
