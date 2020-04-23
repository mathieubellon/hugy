<template>
<div class="GlobalSettings">
    <NavBar></NavBar>
    <div class="container mx-auto mt-10 px-10">
        <div class="my-10">
            <div class="font-bold">Git path</div>
            <input type="text" v-model="workspacePath">
        </div>
        <div class="my-10">
            <div class="font-bold">Workspace path</div>
            <input type="text" v-model="workspacePath">
            <a @click.prevent="openWorkspaceDir">open workspacedir</a>
        </div>
        <div class="my-10">
            <div class="font-bold">Factory default</div>
            <div class="text-sm mb-2 text-gray-600">This will reset Hugy configuration to factory default but will **NOT** delete your projects managed by Hugy</div>
            <a @click.prevent="resetStore()" class="button is-default my-5">Click here to reset factory default</a>
        </div>
        <div class="my-10">
            <div class="font-bold">Install Git utility on windows</div>
            <div class="text-sm mb-2 text-gray-600">This will download git binary</div>
            <a @click.prevent="installGit()" class="button is-default my-5">Click here install Git on your computer</a>
        </div>
    </div>
</div>
</template>

<script>
import { mapMutations, mapState } from 'vuex'
import NavBar from '@/components/NavBar'
import {resourcesPath} from '@/resources'; // Path to resources.ts
const { execFile } = require('child_process');
import * as path from "path";
import {shell} from 'electron';
export default {
    name: "GlobalSettings",
    components:{
        NavBar
    },
    computed:{
        ...mapState(['workspacePath'])
    },
    methods: {
        ...mapMutations(["initialiseStore", "resetStore"]),
        installGit(){
            const portablegitpath = path.join(resourcesPath, 'portablegit')
            const child = execFile(portablegitpath, [], (error, stdout, stderr) => {
                if (error) {
                    throw error;
                }
                console.log(stdout);
            });
        },
        openWorkspaceDir(){
            shell.openItem(this.workspacePath)
        },
    },

}
</script>

<style lang="scss">

</style>
