<template>
<div v-if="statusSummary" class="bg-purple-200 p-4 rounded border border-purple-300 mb-5">
    <button class="button bg-purple-600 text-white" @click.prevent="publish">
        Publish changes
    </button>
    <div v-if="statusSummary.renamed.length > 0" class="mb-3">
        <div class="text-xs text-purple-900 font-semibold">Renamed</div>
        <div v-for="file in statusSummary.renamed">
            {{file}}
        </div>
    </div>
    <div v-if="statusSummary.modified.length > 0" class="mb-3">
        <span class="font-bold text-purple-700 text-xs mr-2">modified</span>
        <div v-for="file in statusSummary.modified">
            <span class="text-purple-700">{{file}}</span>
        </div>
    </div>
    <div v-if="statusSummary.deleted.length > 0" class="mb-3">
        <span class="font-bold text-red-700 text-xs mr-2">deleted</span>
        <div v-for="file in statusSummary.deleted">
            <span class="text-red-700">{{file}}</span>
        </div>
    </div>
    <div v-if="statusSummary.created.length > 0" class="mb-3">
        <span class="font-bold text-green-700 text-xs mr-2">created</span>
        <div v-for="file in statusSummary.created">
            <span class="text-green-700">{{file}}</span>
        </div>
    </div>
    <div v-if="statusSummary.renamed.length > 0" class="mb-3">
        <span class="font-bold text-green-700 text-xs mr-2">renamed</span>
        <div v-for="file in statusSummary.renamed">
            <span class="text-green-700">{{file}}</span>
        </div>
    </div>
    <div v-if="statusSummary.not_added.length > 0" class="mb-3">
        <span class="font-bold text-gray-700 text-xs mr-2">not_added</span>
        <div v-for="file in statusSummary.not_added">
            <span class="text-gray-700">{{file}}</span>
        </div>
    </div>
    <div v-if="statusSummary.staged.length > 0" class="mb-3">
        <span class="font-bold text-gray-700 text-xs mr-2">staged</span>
        <div v-for="file in statusSummary.staged">
            <span class="text-gray-700">{{file}}</span>
        </div>
    </div>
</div>
</template>
