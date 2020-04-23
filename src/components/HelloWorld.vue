<template>
  <div class="text-center">
    <h1>{{ msg }}</h1>
    <h2>The dead simple interface for gohugo.io (and git)</h2>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
