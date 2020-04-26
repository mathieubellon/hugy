<template>
  <div class="max-w-3xl mx-auto px-5">
    <!--     <BarLoader v-if="loading" :size="100" widthUnit="%" color="#000"></BarLoader>-->
    <Breadcrumb class="my-1" v-if="openedDir" :youarehere="openedDir" />
    <div class="text-lg flex items-center py-2 justify-end">
      <input
        id="search"
        class="form-input rounded bg-gray-50 block w-full sm:text-sm sm:leading-5"
        placeholder="search files"
      />
      <span class="flex-shrink-0 relative z-0 inline-flex">
        <a
          @click.prevent="newFile()"
          class="text-gray-600 hover:text-blue-600 flex items-center text-sm font-medium mx-4"
        >
          <FileText class="w-4 mr-1" />New file
        </a>
        <a
          @click.prevent="newFolder()"
          class="text-gray-600 hover:text-blue-600 flex items-center text-sm font-medium"
        >
          <Folder class="w-4 mr-1" />New folder
        </a>
      </span>
    </div>

    <div class="text-sm w-full mx-auto py-2 rounded">
      <a
        v-show="projectPath != openedDir"
        @click="$router.go(-1)"
        class="flex items-center font-medium hover:bg-blue-50"
      >
        <Folder class="w-4 mr-2" />..
      </a>

      <router-link
        :to="{name:'projectfiles', params: {'currentDir':node.path}}"
        v-for="node in filesList.folders"
        :key="node.path"
        class="flex items-center font-medium hover:bg-blue-50"
      >
        <Folder class="w-4 mr-2" />
        {{ node.name }}
      </router-link>

      <router-link
        :to="{name:'document', params: {'documentPath':node.path}}"
        class="flex items-center font-medium hover:bg-blue-50"
        v-for="node in filesList.files"
        :key="node.path"
      >
        <FileText class="w-4 mr-2" />
        {{ node.name }}
      </router-link>
      <div v-if="filesList && !filesList.children">No files in this folder</div>
    </div>
  </div>
</template>
<script>
const dree = require("dree");
import { mapMutations, mapState } from "vuex";
import FileText from "@/assets/svg/file-text.svg";
import Folder from "@/assets/svg/folder.svg";
import Breadcrumb from "@/components/Breadcrumb";
export default {
  name: "ProjectFiles",
  data: () => {
    return {
      loading: false,
      error: null,
      openedDir: "",
      filesList: {}
    };
  },
  components: {
    FileText,
    Folder,
    Breadcrumb
  },
  computed: {
    ...mapState(["projectPath", "readDirPath"])
  },
  created() {
    this.openedDir = this.$route.params.currentDir;
    this.fetchFilesList();
  },
  watch: {
    $route: "fetchFilesList"
  },
  methods: {
    ...mapMutations(["setCurrentFile", "setReadDirPath"]),
    fetchFilesList() {
      this.error = this.filesList = null;
      this.loading = true;
      const options = {
        stat: false,
        normalize: true,
        followLinks: true,
        size: true,
        hash: true,
        depth: 1,
        exclude: /.git|DS_Store|netlify/
      };
      this.filesList = dree.scan(this.openedDir, options);

      if (Object.prototype.hasOwnProperty.call(this.filesList, "children")) {
        // files and directories are in the 'children' key : filter them out from there
        this.filesList.folders = dree
          .scan(this.openedDir, options)
          .children.filter(function(obj) {
            return obj.type !== "file";
          });
        this.filesList.files = dree
          .scan(this.openedDir, options)
          .children.filter(function(obj) {
            return obj.type !== "directory";
          });
      }
    },
    newFile() {
      /*const file = path.join(this.readDirPath, "neofile.md");
      try {
        fs.ensureFileSync(file);
        this.$toasted.success("file created");
      } catch (e) {
        this.$toasted.error(err);
      }*/
    },
    newFolder() {
      /*const file = path.join(this.readDirPath, "installation");
      try {
        fs.ensureDirSync(file);
        this.$toasted.success("folder created");
      } catch (e) {
        this.$toasted.error(err);
      }*/
    }
  }
};
</script>
