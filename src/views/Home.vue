<template>
  <!-- Select existing folder or create new one -->
  <div class="flex flex-col items-center justify-center h-full">
    <Logo class="w-20 h-20" />
    <div class="text-2xl font-black">👋 Hi, I am Hugy</div>
    <div class="text-base max-w-2xl text-center">
      I am the `dead simple`
      <a href="https://gohugo.io/">gohugo.io</a> companion for content editors.
      <div>Edit content and instantly preview and publish your changes</div>
    </div>

    <div class="bg-gray-50 p-2 rounded my-2" v-if="projectPath">
      <div class="text-xs text-gray-500 font-semibold">Recently opened</div>
      <router-link
        :to="{name:'projectfiles', params: {currentDir:projectPath}}"
        class="text-left no-underline justify-center text-lg my-2 hover:text-blue-700 underline"
      >{{projectPath}}</router-link>
    </div>

    <!-- Display probable existing folder path previously used -->
    <!-- Display "create new folder" or "Open existing folder" buttons -->
    <button class="button text-lg" @click.prevent="selectDirectory">Open an existing Hugo project</button>
    <router-link class="text-gray-500 text-sm mt-10" to="globalsettings">Application settings</router-link>
  </div>
</template>

<script>
const { remote } = require("electron");
import { homedir } from "os";
import { mapMutations, mapState } from "vuex";

import Logo from "@/assets/logo/logo.svg";

export default {
  name: "Home",
  components: {
    Logo
  },
  computed: {
    ...mapState(["projectPath"])
  },
  methods: {
    ...mapMutations(["setProjectPath"]),
    selectDirectory() {
      remote.dialog.showOpenDialog(
        {
          properties: ["openDirectory"],
          defaultPath: homedir()
        },
        names => {
          this.setProjectPath(names[0]);
          this.$router.push({
            name: "projectfiles",
            params: { currentDir: names[0] }
          });
        }
      );
    }
  }
};
</script>
