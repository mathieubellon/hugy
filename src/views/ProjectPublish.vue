<template>
  <div class="max-w-3xl px-5 mx-auto">
    <div class="border border-gray-200  rounded p-5 my-4">
      <div class="flex items-center">
        <img src="@/assets/logo/github_mark.png" alt=""/>
        <img class="w-16" src="@/assets/logo/github_logo.png" alt=""/>
      </div>
      <div class="font-bold">Connect to github</div>
      <Auth></Auth>
    </div>
    <div class="border border-gray-200 rounded p-5 my-4">
      <div class="flex items-center">
        <img class="w-24" src="@/assets/logo/gitlab.png" alt=""/>
      </div>
      <div class="font-bold">Connect to gitlab</div>
      <Auth></Auth>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

import Auth from "@/components/Auth.vue";
import ProjectNav from "@/components/ProjectNav";
export default {
  name: "ProjectPublish",
  components: {
    Auth,
  },
  computed: {
    ...mapState(["projectPath"])
  }
};
</script>

<style lang="scss">
</style>
