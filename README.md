<h1 align="center"> Hugy </h1> 
<br>

<p align="center">
<a href="https://github.com/hbyio/hugy"><img src="https://raw.githubusercontent.com/hbyio/hugy/master/src/assets/logo/hugy/icon.svg" alt="icon-128x128-2x" height="200px" border="0"></a>
</p>
<h2 align="center">
The <strong>dead simple</strong> Hugo companion.
</h2>

Hugy is an Electron desktop app acting as a GUI for the [Hugo](https://gohugo.io) static site generator.

Designed for copywriters onboarding Hugo projects without command line knowledge needed.

## 🚨 *Warning - not production ready - pre-alpha stage*
This is a proof of concept at the moment, almost pre-alpha. Test it but only on versioned Hugo projects.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Screenshots](#screenshots)
- [Feedback](#feedback)
- [Development](#development)


## Introduction

Hugy is a desktop (Electron) app aimed at content editors onboarding projects using the [Hugo](https://gohugo.io) static site generator. 
Static sites are great and Hugo is a great static site generator, but due to its technical nature (a command line tool) it may be hard to ask copywriters to use it in projects.

Hugy assumes that your workflow is git-centered as it will allow the content editor to push changes to a git repo.

## Roadmap

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
* Configure Hugy with a versioned config file
* Embedded documentation (Both for Hugy and Hugo)

## Screenshots

### Live edit
<img src="https://user-images.githubusercontent.com/1445897/80204154-a13a7c80-8628-11ea-8b4f-44c0c6133c37.png" alt="live_edit" width="100%" border="0">

### Manage Hugo
<img src="https://user-images.githubusercontent.com/1445897/80204164-a5ff3080-8628-11ea-9fb8-546ec964c7bf.png" alt="console" width="100%" border="0">

### Instant real preview
<img src="https://user-images.githubusercontent.com/1445897/80203780-fde96780-8627-11ea-9f9d-d55fca7bd096.png" alt="live_edit" width="100%" border="0">


## Feedback
Use GitHub issues. We probably will create a discourse/Gitter space in the near future.


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
Not working yet
