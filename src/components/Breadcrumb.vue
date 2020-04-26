import * as path from "path";
<template>
  <div class="breadcrumb flex items-center w-full">
    <div class="flex items-center" v-for="(dirpath, index) in breadcrumb" :key="dirpath.path">
      <router-link
        :to="{name:'projectfiles', params: {currentDir:dirpath.path}}"
        :key="dirpath.path"
        class="text-gray-600 hover:text-blue-500 text-sm"
      >{{ dirpath.name }}</router-link>
      <ChevronRight
        v-if="index != breadcrumb.length - 1"
        class="stroke-current text-gray-400 w-4 h-4"
      />
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import * as path from "path";
import ChevronRight from "@/assets/icons/chevron-right.svg";
export default {
  name: "Breadcrumb",
  props: {
    youarehere: {
      type: String,
      required: true
    }
  },
  components: {
    ChevronRight
  },
  computed: {
    ...mapState(["projectPath"]),
    breadcrumb() {
      // We don not want a nav breadcrumb aboc this point
      let topfolder = path.basename(this.projectPath);
      let dirname = path.dirname(this.youarehere);
      let splitter = dirname.split(path.sep);
      //let relativeBase = path.relative ('', this.projectPath)

      let navigo = [];
      let l = splitter.length;
      let recurnav = this.youarehere;
      for (let i = 0; i < splitter.length; i++) {
        if (!recurnav.includes(topfolder)) break;
        if (i > 0) recurnav = path.dirname(recurnav);
        navigo.push({
          path: recurnav,
          name: path.basename(recurnav)
        });
      }
      return navigo
        .filter(function(navobj) {
          return navobj.path.includes(topfolder);
        })
        .reverse();
    }
  }
};
</script>

<style scoped>
</style>
