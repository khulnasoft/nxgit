---
date: "2017-04-15T14:56:00+02:00"
title: "Customizing Nxgit"
slug: "customizing-nxgit"
weight: 9
toc: false
draft: false
menu:
  sidebar:
    parent: "advanced"
    name: "Customizing Nxgit"
    weight: 9
    identifier: "customizing-nxgit"
---

# Customizing Nxgit

Customizing Nxgit is typically done using the `custom` folder. This is the central
place to override configuration settings, templates, etc.

If Nxgit is deployed from binary, all default paths will be relative to the Nxgit
binary. If installed from a distribution, these paths will likely be modified to
the Linux Filesystem Standard. Nxgit will create required folders, including `custom/`.
Application settings are configured in `custom/conf/app.ini`. Distributions may
provide a symlink for `custom` using `/etc/nxgit/`.

- [Quick Cheat Sheet](https://docs.nxgit.io/en-us/config-cheat-sheet/)
- [Complete List](https://github.com/khulnasoft/nxgit/blob/master/custom/conf/app.ini.sample)

If the `custom` folder can't be found next to the binary, check the `NXGIT_CUSTOM`
environment variable; this can be used to override the default path to something else.
`NXGIT_CUSTOM` might, for example, be set by an init script.

- [List of Environment Variables](https://docs.nxgit.io/en-us/specific-variables/)

**Note:** Nxgit must perform a full restart to see configuration changes.

## Customizing /robots.txt

To make Nxgit serve a custom `/robots.txt` (default: empty 404), create a file called
`robots.txt` in the `custom` folder with [expected contents](http://www.robotstxt.org/).

## Serving custom public files

To make Nxgit serve custom public files (like pages and images), use the folder
`custom/public/` as the webroot. Symbolic links will be followed.

For example, a file `image.png` stored in `custom/public/`, can be accessed with
the url `http://nxgit.domain.tld/image.png`.

## Changing the default avatar

Place the png image at the following path: `custom/public/img/avatar\_default.png`

## Customizing Nxgit pages

The `custom/templates` folder allows changing every single page of Nxgit. Templates
to override can be found in the [`templates`](https://github.com/khulnasoft/nxgit/tree/master/templates) directory of Nxgit source. Override by
making a copy of the file under `custom/templates` using a full path structure
matching source.

Any statement contained inside `{{` and `}}` are Nxgit's template syntax and
shouldn't be touched without fully understanding these components.

### Adding links and tabs

If all you want is to add extra links to the top navigation bar, or extra tabs to the repository view, you can put them in `extra_links.tmpl` and `extra_tabs.tmpl` inside your `custom/templates/custom/` directory.

For instance, let's say you are in Germany and must add the famously legally-required "Impressum"/about page, listing who is responsible for the site's content:
just place it under your "custom/public/" directory (for instance `custom/public/impressum.html`) and put a link to it in `custom/templates/custom/extra_links.tmpl`.

To match the current style, the link should have the class name "item", and you can use `{{AppSubUrl}}` to get the base URL:
`<a class="item" href="{{AppSubUrl}}/impressum.html">Impressum</a>`

You can add new tabs in the same way, putting them in `extra_tabs.tmpl`.
The exact HTML needed to match the style of other tabs is in the file
`templates/repo/header.tmpl`
([source in GitHub](https://github.com/khulnasoft/nxgit/blob/master/templates/repo/header.tmpl))

### Other additions to the page

Apart from `extra_links.tmpl` and `extra_tabs.tmpl`, there are other useful templates you can put in your `custom/templates/custom/` directory:

- `header.tmpl`, just before the end of the `<head>` tag where you can add custom CSS files for instance.
- `body_outer_pre.tmpl`, right after the start of `<body>`.
- `body_inner_pre.tmpl`, before the top navigation bar, but already inside the main container `<div class="full height">`.
- `body_inner_post.tmpl`, before the end of the main container.
- `body_outer_post.tmpl`, before the bottom `<footer>` element.
- `footer.tmpl`, right before the end of the `<body>` tag, a good place for additional Javascript.

## Adding Analytics to Nxgit

Google Analytics, Matomo (previously Piwik), and other analytics services can be added to Nxgit. To add the tracking code, refer to the `Other additions to the page` section of this document, and add the JavaScript to the `custom/templates/custom/header.tmpl` file.

## Customizing gitignores, labels, licenses, locales, and readmes.

Place custom files in corresponding sub-folder under `custom/options`.

**NOTE:** The files should not have a file extension, e.g. `Labels` rather than `Labels.txt`

### gitignores

To add custom .gitignore, add a file with existing [.gitignore rules](https://git-scm.com/docs/gitignore) in it to `custom/options/gitignore`

### Labels

To add a custom label set, add a file that follows the [label format](https://github.com/khulnasoft/nxgit/blob/master/options/label/Default) to `custom/options/label`  
`#hex-color label name ; label description`

### Licenses

To add a custom license, add a file with the license text to `custom/options/license`

### Locales

Locales are managed via our [crowdin](https://crowdin.com/project/nxgit).  
You can override a locale by placing an altered locale file in `custom/options/locale`.  
Nxgit's default locale files can be found in  the [`options/locale`](https://github.com/khulnasoft/nxgit/tree/master/options/locale) source folder and these should be used as examples for your changes.  
  
To add a completely new locale, as well as placing the file in the above location, you will need to add the new lang and name to the `[i18n]` section in your `app.ini`. Keep in mind that Nxgit will use those settings as **overrides**, so if you want to keep the other languages as well you will need to copy/paste the default values and add your own to them.

```
[i18n]
LANGS = en-US,foo-BAR
NAMES = English,FooBar
```

Locales may change between versions, so keeping track of your customized locales is highly encouraged.

### Readmes

To add a custom Readme, add a markdown formatted file (without an `.md` extension) to `custom/options/readme`

## Customizing the look of Nxgit

As of version 1.6.0 Nxgit has built-in themes. The two built-in themes are, the default theme `nxgit`, and a dark theme `arc-green`. To change the look of your Nxgit install change the value of `DEFAULT_THEME` in the [ui](https://docs.nxgit.io/en-us/config-cheat-sheet/#ui-ui) section of `app.ini` to another one of the available options.  
As of version 1.8.0 Nxgit also has per-user themes. The list of themes a user can choose from can be configured with the `THEMES` value in the [ui](https://docs.nxgit.io/en-us/config-cheat-sheet/#ui-ui) section of `app.ini` (defaults to `nxgit` and `arc-green`, light and dark respectively)
