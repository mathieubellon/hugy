<h1 align="center"> Hugy </h1> 
<br>
<a href="https://github.com/hbyio/hugy"><img src="https://raw.githubusercontent.com/hbyio/hugy/master/src/assets/logo/hugy/icon.svg" alt="icon-128x128-2x" height="200px" border="0"></a>

The *dead simple* Hugo companion

Hugy is an Electron desktop app acting as a GUI for the [Hugo](https://gohugo.io) static site generator.

Designed for copywriters onboarding Hugo projects without command line knowledge needed.

## 🚨 *Warning - not production ready*
This is a proof of concept at the moment. Test it but only on versionned Hugo projet.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Screenshots](#screenshots)
- [Feedback](#feedback)
- [Development](#development)


## Introduction

Hugy is a desktop (Electron) app aimed at content editors onboarding projects using the [Hugo](https://gohugo.io) static site genertor. 
Static sites are great and Hugo is a great static site generator but due to its specifity (a command line tool) it may be hard
to ask copywriters to use it in projects.

Hugy assumes that your workflow is git-centered as it will allow the content editor to push changes to a git repo.

## Features

### V1.0
* Download available for Mac and Windows
* Create or clone Hugo project
* Navigate content folder, add, edit content with a nice WYSIWYG tool
* Manage project configuration (config.toml)
* Live preview your website as you edit content
* Manage multiples local versions of your website. Rollback changes.
* Publish (push) your changes

### V1.x
* Create new content based on archetypes
* Manage taxonomies
* Manage menus
* Configure Hugy with versionned config file
* Embedded documentation (Both for Hugy and Hugo)

## Screenshots

### Live edit
<img src="https://raw.githubusercontent.com/hbyio/hugy/master/documentation/images/live_edit" alt="live_edit" width="100%" border="0">

### Manage Hugo
<img src="https://raw.githubusercontent.com/hbyio/hugy/master/documentation/images/console" alt="console" width="100%" border="0">

### Preview
<img src="https://raw.githubusercontent.com/hbyio/hugy/master/documentation/images/preview" alt="live_edit" width="100%" border="0">


## Feedback
Use github issues. We probably will create a discourse/Gitter space a the near future.


## Development

Clone the repo and ..

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn electron:serve
```

### Compiles and minifies for production
```
yarn electron:build
```
