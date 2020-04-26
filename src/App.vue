<template>
  <div id="app" class="flex flex-col overflow-y-hidden">
    <NavBar v-if="$route.name != 'home'"></NavBar>
    <main class="flex-grow overflow-y-auto">
      <transition name="fade" mode="out-in" :duration="150">
        <router-view name="navbar"></router-view>
      </transition>
      <transition name="fade" mode="out-in" :duration="150">
        <router-view :key="$route.fullPath"></router-view>
      </transition>
    </main>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import NavBar from "./components/NavBar.vue";
export default Vue.extend({
  name: "app",
  components: {
    NavBar
  },
  data() {
    return {
      transitionName: "slide-left"
    };
  }
});
</script>

<style lang="scss">
main {
  &::-webkit-scrollbar {
    display: none;
  }
}
</style>
