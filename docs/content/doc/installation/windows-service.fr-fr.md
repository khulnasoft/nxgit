---
date: "2017-08-23T09:00:00+02:00"
title: "Démarrer en tant que service Windows"
slug: "windows-service"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "Service Windows"
    weight: 30
    identifier: "windows-service"
---

# Activer un service Windows

Pour activer le service Windows Nxgit, ouvrez une `cmd` en tant qu'Administrateur puis utilisez la commande suivante :

```
sc create nxgit start= auto binPath= ""C:\nxgit\nxgit.exe" web --config "C:\nxgit\custom\conf\app.ini""
```

N'oubliez pas de remplacer `C:\nxgit` par le chemin que vous avez utilisez pour votre installation.

Ensuite, ouvrez "Services Windows", puis recherchez le service `nxgit`, faites un clique droit et selectionnez "Run". Si tout fonctionne, vous devriez être capable d'accèder à Nxgit à l'URL `http://localhost:3000` (ou sur le port configuré si différent de 3000).

## Désactiver un service Windows

Pour désactiver le service Windows Nxgit, ouvrez une `cmd` en tant qu'Administrateur puis utilisez la commande suivante :

```
sc delete nxgit
```
