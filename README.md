# 🚨 Warning ️🚨

This is a POC stage project. Test it but only on versionned Hugo projet.

# hugy

Hugy is a desktop (Electron) app aimed at content editors onboarding in projects using Hugo.io. 
Static sites are great and Hugo is a great static site generator but due to its specifity (a command line tool) it may be hard
to say to content editors to use it in projects.

Hugy assumes that your workflow is git-centered as it will allow the content editor to push changes to your git repo.

# Roadmap

In v1 Hugy will allow a content editor (on Windows and Mac) to clone (or create) a Hugo project, edit some content and push all the changes to Github or Gitlab.
The UI will also serve as an "advanced tutorial" explaining Hugo related mechanics as the editor work on content (eg. if the copywriter add a new file we will explain that a new url will exists for this content because Hugo maps it ti files hierarchy)

We also plan to allow a developper to configure Hugy UI with a special config .hugy file  (for example to disabled config.toml edition of certain fields) 


# Development

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
