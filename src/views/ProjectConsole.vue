<template>
  <div class="max-w-3xl p-5 mx-auto">
    <div class="my-2 rounded-md bg-gray-50 px-6 py-5 sm:flex sm:items-start sm:justify-between">
      <div class="sm:flex sm:items-start">
        <Process />
        <div class="mt-3 sm:mt-0 sm:ml-4">
          <div class="text-sm leading-5 font-medium text-gray-900">
            <div>Process id : {{ hugoServerPID }}</div>
            <div>Port : {{ hugoServerPort }}</div>
          </div>
          <div class="mt-1 text-sm leading-5 text-gray-600 sm:flex sm:items-center">
            <div>Server running http://localhost:{{hugoServerPort}}</div>
            <span class="hidden sm:mx-2 sm:inline" aria-hidden="true">&middot;</span>
            <div class="mt-1 sm:mt-0">
              <a class="link" @click.prevent="cleanConsole">Clean console</a>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-4 sm:mt-0 sm:ml-6 sm:flex-shrink-0">
        <span class="inline-flex rounded-md shadow-sm">
          <button
            type="button"
            @click.prevent="openServerUrl"
            v-if="hugoServerPort"
            class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700 bg-white hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:text-gray-800 active:bg-gray-50 transition ease-in-out duration-150"
          >Open in browser</button>
          <button
            type="button"
            v-else
            disabled
            class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-400 disabled:opacity-75 bg-white hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:text-gray-800 active:bg-gray-50 transition ease-in-out duration-150"
          >Open in browser</button>
        </span>
      </div>
    </div>

    <div class="console-screen break-words flex-grow" v-if="console && console.length > 0">
      <pre
        class="my-2 text-xs even:bg-gray-200"
        v-for="output in console"
        v-html="output.msg"
        :key="output.time"
      ></pre>
    </div>
    <div v-else class="console-screen">
      <div>Hugo is not running</div>
    </div>
  </div>
</template>

<script>
import { mapMutations, mapState } from "vuex";
import * as hugo from "@/hugo";
import * as path from "path";
import Process from "@/assets/svg/cpu.svg";
const { shell } = require("electron");

export default {
  name: "Console",
  components: {
    Process
  },
  computed: {
    ...mapState([
      "projectPath",

      "hugoServerPID",
      "hugoServerPort",

      "serverIsLaunching",
      "console"
    ])
  },
  methods: {
    ...mapMutations([
      "setCurrentProject",
      "sethugoServerPID",
      "setSubWindow",
      "cleanConsole"
    ]),
    openServerUrl() {
      shell.openExternal(`http://localhost:${this.hugoServerPort}`);
    },
    hugoStop() {
      hugo.stopserver();
    },
    hugoRun() {
      const siteDir = path.join(this.projectPath);
      hugo.runserver(siteDir);
    }
  }
};
</script>

<style lang="scss">
.console-screen {
  @apply font-normal text-sm bg-gray-900 break-words text-white p-2 rounded overflow-y-auto;
  pre {
    white-space: pre-wrap; /* css-3 */
    white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
    white-space: -pre-wrap; /* Opera 4-6 */
    white-space: -o-pre-wrap; /* Opera 7 */
    word-wrap: break-word; /* Internet Explorer 5.5+ */
  }
}
</style>
