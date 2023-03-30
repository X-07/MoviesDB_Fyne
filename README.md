# MoviesDB_Fyne
Test Fyne library for GO 

Je développe en **GO** un gestionnaire de collection de films (façon GCStar pour les connaisseurs).
Il n'est pas question, ici, de visualiser les films, d'autres le font très bien (**VLC**...) mais juste de gérer une collection.
Ce logiciel sera composé à 90% d'une interface utilisateur (**UI**) et de 10% d'une BDD (**SQLite** ici).

Pour l'**UI** j'ai donc exploré ce que propose **Fyne** : https://github.com/fyne-io/fyne

Voici le résultat obtenu ...

### Onglet 'Fiche'
![Image 1](/ScreenShots/MoviesDB-1.jpg)

### Onglet 'Fiche' + saisie réalisateurs
![Image 1a](/ScreenShots/MoviesDB-1a.jpg)

### Onglet 'Détails'
![Image 2](/ScreenShots/MoviesDB-2.jpg)
- Possibilité de créer des Widgets personnalisés (ici au survol de la souris l'étoile s'allume puis se mémorise au click)

### Onglet 'Infos'
![Image 3](/ScreenShots/MoviesDB-3.jpg)

- Assez facile à prendre en main
- Assez verbeux
- Look! Chacun se fera son avis
- IMPOSSIBLE de positionner finement les Widgets avec le concept d'occuper la place miniamale sauf pour 1 qui prend toute la place restante!
Ce qui peut donner des listes à 1 seul élément visible (+ ascenseur!) 

## Pas fan du look
j'ai donc abandonné ce projet (que je laisse ici comme example) et le l'ai continué en GOTK3 : https://github.com/gotk3/gotk3
- Plus complexe
- Beaucoup plus verbeux

### Interface principale Gtk3 à un stade plus avancé!
![Image Gtk3](/ScreenShots/MoviesDB-Gotk3.jpg)
