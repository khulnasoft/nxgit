---
date: "2017-08-23T09:00:00+02:00"
title: "Mise à jour depuis Gogs"
slug: "upgrade-from-gogs"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "upgrade"
    name: "Depuis Gogs"
    weight: 10
    identifier: "upgrade-from-gogs"
---

# Mise à jour depuis Gogs

À partir de la version 0.9.146 (schéma de la base de données : version 15) de Gogs, Il est possible de migrer vers Nxgit simplement et sans encombre.

Veuillez suivre les étapes ci-dessous. Sur Unix, toute les commandes s'exécutent en tant que l'utilisateur utilisé pour votre installation de Gogs :

* Crééer une sauvegarde de Gogs avec la commande `gogs dump`. Le fichier nouvellement créé `gogs-dump-[timestamp].zip` contient toutes les données de votre instance de Gogs.
* Téléchargez le fichier correspondant à votre plateforme à partir de la [page de téléchargements](https://dl.nxgit.io/nxgit).
* Mettez la binaire dans le répertoire d'installation souhaité.
* Copiez le fichier `gogs/custom/conf/app.ini` vers `nxgit/custom/conf/app.ini`.
* Si vous avez personnalisé les répertoires `templates, public` dans `gogs/custom/`, copiez-les vers `nxgit/custom/`.
* Si vous avez d'autres répertoires personnalisés comme `gitignore, label, license, locale, readme` dans `gogs/custom/conf` copiez-les vers `nxgit/custom/options`.
* Copiez le répertoire `gogs/data/` vers `nxgit/data/`.
* Vérifiez votre installation en exécutant Nxgit avec la commande `nxgit web`.
* Lancez le binaire de version majeure en version majeure ( `1.1.4` → `1.2.3` → `1.3.4` → `1.4.2` →  etc ) afin de récupérer les migrations de base de données.
* Connectez vous au panel d'administration de Nxgit et exécutez l'action `Rewrite '.ssh/authorized_keys' file`, puis l'action `Rewrite all update hook of repositories` (obligatoire si le chemin menant à votre configuration personnalisée à changé).

## Modifier les informations spécifiques de gogs

* Renommez `gogs-repositories/` vers `nxgit-repositories/`
* Renommez `gogs-data/` to `nxgit-data/`
* Dans votre fichier `nxgit/custom/conf/app.ini`, modifiez les éléments suivants:

  DE :

  ```ini
  [database]
  PATH = /home/:USER/gogs/data/:DATABASE.db
  [attachment]
  PATH = /home/:USER/gogs-data/attachments
  [picture]
  AVATAR_UPLOAD_PATH = /home/:USER/gogs-data/avatars
  [log]
  ROOT_PATH = /home/:USER/gogs/log
  ```

  VERS :

  ```ini
  [database]
  PATH = /home/:USER/nxgit/data/:DATABASE.db
  [attachment]
  PATH = /home/:USER/nxgit-data/attachments
  [picture]
  AVATAR_UPLOAD_PATH = /home/:USER/nxgit-data/avatars
  [log]
  ROOT_PATH = /home/:USER/nxgit/log
  ```

* Vérifiez votre installation en exécutant Nxgit avec la commande `nxgit web`.

## Dépannage

* Si vous rencontrez des erreurs relatives à des modèles personnalisés dans le dossier `nxgit/custom/templates`, essayez de déplacer un par un les modèles provoquant les erreurs. Il est possible qu'ils ne soient pas compatibles avec Nxgit.

## Démarrer automatiquement Nxgit (Unix)

Distributions utilisant systemd:

* Copiez le script mis à jour vers `/etc/systemd/system/nxgit.service`
* Ajoutez le service avec la commande `sudo systemctl enable nxgit`
* Désactivez Gogs avec la commande `sudo systemctl disable gogs`

Distributions utilisant SysVinit:

* Copiez le script mis à jour vers `/etc/init.d/nxgit`
* Ajoutez le service avec la commande `sudo rc-update add nxgit`
* Désactivez Gogs avec la commande `sudo rc-update del gogs`
