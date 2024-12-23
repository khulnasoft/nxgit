---
date: "2017-08-23T09:00:00+02:00"
title: "Installation avec le binaire pré-compilé"
slug: "install-from-binary"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "Binaire pré-compilé"
    weight: 20
    identifier: "install-from-binary"
---

# Installation avec le binaire pré-compilé

Tous les binaires sont livrés avec le support de SQLite, MySQL et PostgreSQL, et sont construits avec les ressources incorporées. Gardez à l'esprit que cela peut être différent pour les versions antérieures. L'installation basée sur nos binaires est assez simple, il suffit de choisir le fichier correspondant à votre plateforme à partir de la [page de téléchargement](https://dl.nxgit.io/nxgit). Copiez l'URL et remplacer l'URL dans les commandes suivantes par la nouvelle:

```
wget -O nxgit https://dl.nxgit.io/nxgit/1.3.2/nxgit-1.3.2-linux-amd64
chmod +x nxgit
```

## Test

Après avoir suivi les étapes ci-dessus, vous aurez un binaire `nxgit` dans votre répertoire de travail. En premier lieu, vous pouvez tester si le binaire fonctionne comme prévu et ensuite vous pouvez le copier vers la destination où vous souhaitez le stocker. Lorsque vous lancez Nxgit manuellement à partir de votre CLI, vous pouvez toujours le tuer en appuyant sur `Ctrl + C`.

```
./nxgit web
```

## Dépannage

### Anciennes version de glibc

Les anciennes distributions Linux (comme Debian 7 ou CentOS 6) peuvent ne pas être capable d'exécuter le binaire Nxgit, résultant généralement une erreur du type ```./nxgit: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.14' not found (required by ./nxgit)```. Cette erreur est due au driver SQLite que nous intégrons dans le binaire Nxgit. Dans le futur, nous fournirons des binaires sans la dépendance pour la bibliothèque glibc. En attendant, vous pouvez mettre à jour votre distribution ou installer Nxgit depuis le [code source]({{< relref "from-source.fr-fr.md" >}}).

### Exécuter Nxgit avec un autre port

Si vous obtenez l'erreur `702 runWeb()] [E] Failed to start server: listen tcp 0.0.0.0:3000: bind: address already in use`, Nxgit à besoin d'utiliser un autre port. Vous pouvez changer le port par défaut en utilisant `./nxgit web -p $PORT`.

## Il manque quelque chose ?

Est-ce que nous avons oublié quelque chose sur cette page ? N'hésitez pas à nous contacter sur notre [serveur Discord](https://discord.gg/NsatcWJ), vous obtiendrez des réponses à toute vos questions assez rapidement.
