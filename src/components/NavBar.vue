<template>
  <div class="bg-gray-50 px-5">
    <div class="flex items-center  py-2">
      <router-link to="/" >
        <Logo />
      </router-link>
      <router-link
          v-if="projectBasename"
          :to="{name:'projectfiles', params: {currentDir:projectPath}}"
          class="text-left ml-2 font-medium text-gray-900 hover:text-blue-700 hover:underline"
      >
        {{ projectBasename }}
      </router-link>
      <div v-else class="ml-2 font-medium">Hugy</div>
      <div class="ml-auto">
         <transition name="fade" mode="out-in">
        <button v-if="serverIsLaunching" class="button is-orange">
          <Info class="w-5 stroke-current text-white mr-2"/>
          <BounceLoader :loading="true" :size="14" color="#fff" class="mr-2"></BounceLoader>Launching
        </button>
        <button
            v-else-if="hugoServerPort"
            type="button"
            @click.prevent="hugoStop()"
            class="button is-red"
        >
          <Stop class="w-5 stroke-current text-white mr-2"/>Stop server
        </button>
        <button v-else @click.prevent="hugoRun()" class="button">
          <Play class="w-5 stroke-current text-white mr-2"/>Show website preview
        </button>
      </transition>
      </div>
     
    </div>
  </div>

  <!--

  -->
</template>

<script>
import * as hugo from "@/hugo";
import { mapState, mapMutations } from "vuex";
import Logo from "@/assets/logo/logo.svg";
import Play from "@/assets/svg/play-circle.svg";
import Stop from "@/assets/svg/stop-circle.svg";
import Info from "@/assets/svg/alert-circle.svg";
import * as path from "path";
import Breadcrumb from '@/components/Breadcrumb'

export default {
  name: "NavBar",
  components: {
    Logo,
    Play,
    Stop,
    Info,
    Breadcrumb
  },
  computed: {
    ...mapState(["projectPath", "hugoServerPort", "serverIsLaunching"]),
    projectBasename(){
      if (this.projectPath){
        return path.basename(this.projectPath)
      }
    }
  },
  methods: {
    ...mapMutations(["toggleServerIsLaunching"]),
    hugoStop() {
      hugo.stopserver();
    },
    hugoRun() {
      this.toggleServerIsLaunching(true);
      const siteDir = path.join(this.projectPath);
      hugo.runserver(siteDir);
    }
  }
};
</script>


<style lang="scss">
.tab {
  @apply flex-1 py-4 px-1 text-center border-b border-gray-300 font-medium text-sm leading-5 text-gray-500;
  &:focus {
    @apply outline-none;
    @apply text-blue-800;
    @apply border-blue-700;
  }
  &.router-link-active {
    @apply text-blue-600;
    @apply border-blue-700;
  }
}
</style>
