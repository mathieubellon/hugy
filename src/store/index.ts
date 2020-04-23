import Vue from 'vue'
import Vuex from 'vuex';
import { Project, RootState } from '@/types'
import path from 'path'
import VuexPersistence from 'vuex-persist'
import localforage from 'localforage'
const uuidv4 = require('uuid/v4');

import { filter } from "lodash"

Vue.use(Vuex)

const vuexLocal = new VuexPersistence<RootState>({
  storage: localforage,
  asyncStorage: true
})

const getDefaultState = (): RootState => {
  return {
    serverIsLaunching: false,
    projectPath: null,
    currentDir: null,
    currentFile: null,
  
    hugoServerPID: null,
    hugoServerPort: null,
    subwindow: null,
    console: [],
    
    currentProject: null,
    currentProjectFiles: null,
    readDirPath: null,


    user: null,
  }
}

const store = new Vuex.Store<RootState>({
  state: getDefaultState(),
  mutations: {
    resetStore(state) {
      Object.assign(state, getDefaultState())
    },

    cleanConsole(state) {
      state.console = []
    },

    setProjectPath(state, projectPath) {
      state.projectPath = projectPath
    },

    setCurrentProject(state, project) {
      state.currentProject = project
    },
    setCurrentProjectFiles(state, files) {
      state.currentProjectFiles = files
    },
    setReadDirPath(state, currentPath) {
      state.readDirPath = currentPath
    },
    setCurrentFile(state, file) {
      state.currentFile = file
    },
    setCurrentDir(state, dir){
      state.currentDir = dir
    },

    setHugoServer(state, pid) {
      state.hugoServerPID = pid
    },
    setHugoServerPort(state, port) {
      state.hugoServerPort = port
    },
    setSubWindow(state, subwindow) {
      state.subwindow = subwindow
    },

    toggleServerIsLaunching(state, islaunching) {
      state.serverIsLaunching = islaunching
    },
    logToConsole(state, msg) {
      state.console.push(msg)
    },
    setUser(state, user) {
      state.user = user
    }
  },
  actions: {},
  modules: {},
  plugins: [vuexLocal.plugin]
})

/*store.subscribe((mutation, state) => {
  // Store the state object as a JSON string
  localStorage.setItem('store', JSON.stringify(state))
})*/

export default store
