<template>
<div class="flex flex-row max-w-3xl mx-auto px-5">
    <div class="sm:hidden">
      <select class="form-select block w-full">
        <option>Files</option>
        <option>Configuration</option>
        <option>Publish</option>
        <option>Console</option>
      </select>
    </div>
    <div class="hidden w-full sm:block">
        <nav class="flex justify-between w-full">
          <router-link :to="{name:'projectfiles', params: {currentDir:projectPath}}" class="tab">Files</router-link>
          <router-link :to="{name:'projectsettings'}" class="tab">Configuration</router-link>

          <router-link :to="{name:'projectpublish'}" class="tab">Publish</router-link>
          <router-link :to="{name:'projectconsole'}" class="tab flex items-center justify-center">
            <transition name="fade">
                <LogoHugo class="opacity-25 w-16" :class="{'opacity-100':serverIsLaunching || hugoServerPort}"/>
            </transition>
            <transition name="fade" mode="out-in"></transition>
          </router-link>
        </nav>
    </div>
</div>
</template>

<script>
import {mapState} from 'vuex'
import LogoHugo from '@/assets/logo/hugo.svg'
export default {
    name: "ProjectNav",
    components:{
        LogoHugo
    },
    computed:{
        ...mapState(["projectPath", "serverIsLaunching","hugoServerPort"])
    },
}
</script>

<style scoped>

</style>
