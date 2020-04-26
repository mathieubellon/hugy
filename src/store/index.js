import Vue from 'vue'
import Vuex from 'vuex';

Vue.use(Vuex)

const getDefaultState = () => {
  return {
    serverIsLaunching: false,
    projectPath: null,
    currentDir: null,
    currentFile: null,

    hugoServerPID: null,
    hugoServerPort: null,
    console: [],
  }
}

const store = new Vuex.Store({
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

    setCurrentFile(state, file) {
      state.currentFile = file
    },
    setCurrentDir(state, dir) {
      state.currentDir = dir
    },

    setHugoServer(state, pid) {
      state.hugoServerPID = pid
    },
    setHugoServerPort(state, port) {
      state.hugoServerPort = port
    },


    toggleServerIsLaunching(state, islaunching) {
      state.serverIsLaunching = islaunching
    },
    logToConsole(state, msg) {
      state.console.push(msg)
    },
  },
  actions: {},
  modules: {}
})

/*store.subscribe((mutation, state) => {
  // Store the state object as a JSON string
  localStorage.setItem('store', JSON.stringify(state))
})*/

export default store
