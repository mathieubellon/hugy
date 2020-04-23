<template>
<form class="max-w-3xl  mx-auto p-5">
  <div>
     <div>
      <div class="text-gray-900 tex-sm my-5 bg-yellow-300 rounded p-5">
        <p class="font-medium leading-5">
          Warning. This form is purely  for demonstration purposes right now. Modifications won't be saved
        </p>
      </div>
    </div>
    <div>
      <div class="text-gray-700 tex-sm my-5 bg-gray-50 rounded p-5">
        <p class="font-medium leading-5">
          Project configuration.
        </p>
        <p class="text-sm text-gray-600 leading-normal ">
           Edit your project configuration. All changes are saved in <code>config.toml</code>, a file you will find at the root of your project files.
        </p>
      </div>
    </div>

    <div class="mt-8 pt-8">
      <div>
        <h3 class="text-lg leading-6 font-medium text-gray-900">
          Required informations
        </h3>
      </div>
      <div class="mt-6 ">
        <div class="my-3">
          <label for="title" class="block text-sm font-medium leading-5 text-gray-700">
           Website title (this value appears in browser title bar for example)
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <input
                id="title"
                class="form-input block w-full transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            />
          </div>
        </div>

        <div class="my-3">
          <label for="base_url" class="block text-sm font-medium leading-5 text-gray-700">
            Base URL
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <input
                id="base_url"
                class="form-input block w-full transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            />
          </div>
        </div>


        <div class="my-3">
          <label for="country" class="block text-sm font-medium leading-5 text-gray-700">
            Content Language (if we can't find a translation for a content, default to this lang)
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <select
                id="language"
                class="form-select block w-full transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            >
              <option>EN (us)</option>
              <option>French</option>
              <option>Spanish</option>
            </select>
          </div>
        </div>

        <div class="my-3">
          <label for="defaultlanguage" class="block text-sm font-medium leading-5 text-gray-700">
            Default language (multilingual websites must have a default.)
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <select
                id="defaultlanguage"
                class="form-select block w-full transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            >
              <option>EN (us)</option>
              <option>French</option>
              <option>Spanish</option>
            </select>
          </div>
        </div>

      </div>
    </div>
    <div class="pt-4">
      <div class="mt-6">
        <fieldset>
          <legend class="text-base font-medium text-gray-900">
            Translation
          </legend>
          <div class="mt-4">
            <div class="relative flex items-start">
              <div class="absolute flex items-center h-5">
                <input
                    id="comments"
                    type="checkbox"
                    class="form-checkbox h-4 w-4 text-indigo-600 transition duration-150 ease-in-out"
                />
              </div>
              <div class="pl-7 text-sm leading-5">
                <label for="comments" class="font-medium text-gray-700">Enable Missing Translation Placeholders</label>
                <p class="text-gray-500">If some piece of content goes missing while translating the website show a placeholder</p>
              </div>
            </div>
          </div>
        </fieldset>

      </div>
    </div>
  </div>
  <div class="mt-8 pt-5">
    <div class="flex justify-end">
      <span class="inline-flex rounded-md shadow-sm">
        <button
            type="button"
            class="py-2 px-4 border border-gray-300 rounded-md text-sm leading-5 font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:bg-gray-50 active:text-gray-800 transition duration-150 ease-in-out"
        >
          Cancel
        </button>
      </span>
      <span class="ml-3 inline-flex rounded-md shadow-sm">
        <button
            type="submit"
            class="inline-flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out"
        >
          Save
        </button>
      </span>
    </div>
  </div>
</form>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "ProjectSettings",

  computed: {
    ...mapState(["projectPath"])
  }
};
</script>

<style lang="scss">
</style>
