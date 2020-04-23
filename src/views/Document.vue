<template>
  <div class="flex flex-col py-2">
    <div class="mb-5  px-5 flex flex-col items-start flex-grow justify-between">
      <div class="flex items-center pt-3 w-full justify-between">
        <Breadcrumb v-if="documentPath" :youarehere="documentPath" />
        <button class="button ml-auto mr-2" @click.prevent.prevent="saveFile">save</button>
        <button
            type="button"
            @click.prevent.prevent="closeFile"
            class="inline-flex items-center px-2.5 py-1.5 border border-transparent text-xs leading-4 font-medium rounded text-gray-700 bg-gray-100 hover:bg-gray-50 focus:outline-none focus:border-gray-300 focus:shadow-outline-gray active:bg-gray-200 transition ease-in-out duration-150"
        >
          close
        </button>
      </div>
      <div class="flex items-center py-2 w-full justify-between">
          <transition name="fade" mode="out-in">
            <div class="ml-auto text-xs text-gray-400" :key="saved_message">{{saved_message}}</div>
          </transition>
      </div>
    </div>

    <!--    <FrontMatter :frontmatter="currentFile.data"></FrontMatter>-->

    <div class="editor w-full max-w-3xl mx-auto" v-if="currentFile">
      <editor-content class="editor__content text-3xl mb-5" :editor="title_editor" />
    </div>
    <div class="editor w-full max-w-3xl mx-auto" v-if="currentFile">
      <editor-menu-bar :editor="editor" v-slot="{ commands, isActive }">
        <div class="menubar">
          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.bold() }"
            @click.prevent="commands.bold"
          >
            <Icon name="bold" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.italic() }"
            @click.prevent="commands.italic"
          >
            <Icon name="italic" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.strike() }"
            @click.prevent="commands.strike"
          >
            <Icon name="strike" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.underline() }"
            @click.prevent="commands.underline"
          >
            <Icon name="underline" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.code() }"
            @click.prevent="commands.code"
          >
            <Icon name="code" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.paragraph() }"
            @click.prevent="commands.paragraph"
          >
            <Icon name="paragraph" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.heading({ level: 1 }) }"
            @click.prevent="commands.heading({ level: 1 })"
          >H1</button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.heading({ level: 2 }) }"
            @click.prevent="commands.heading({ level: 2 })"
          >H2</button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.heading({ level: 3 }) }"
            @click.prevent="commands.heading({ level: 3 })"
          >H3</button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.bullet_list() }"
            @click.prevent="commands.bullet_list"
          >
            <Icon name="ul" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.ordered_list() }"
            @click.prevent="commands.ordered_list"
          >
            <Icon name="ol" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.blockquote() }"
            @click.prevent="commands.blockquote"
          >
            <Icon name="quote" />
          </button>

          <button
            class="menubar__button"
            :class="{ 'is-active': isActive.code_block() }"
            @click.prevent="commands.code_block"
          >
            <Icon name="code" />
          </button>

          <button class="menubar__button" @click.prevent="commands.horizontal_rule">
            <Icon name="hr" />
          </button>

          <button class="menubar__button" @click.prevent="commands.undo">
            <Icon name="undo" />
          </button>

          <button class="menubar__button" @click.prevent="commands.redo">
            <Icon name="redo" />
          </button>
        </div>
      </editor-menu-bar>
      <editor-content class="editor__content" :editor="editor" />
    </div>
  </div>
</template>
<script>
import fm from "gray-matter";
import fs from "fs-extra";
import * as path from "path";

import Breadcrumb from '@/components/Breadcrumb'
import {
  Blockquote,
  CodeBlock,
  HardBreak,
  Heading,
  OrderedList,
  BulletList,
  ListItem,
  TodoItem,
  TodoList,
  Bold,
  Code,
  Italic,
  Link,
  History,
  Strike,
  Underline,
  HorizontalRule
} from "tiptap-extensions";
import { Editor, EditorContent, EditorMenuBar } from "tiptap";
import Icon from "@/components/Icon";
import { Placeholder } from "tiptap-extensions";
import { md2html, html2md } from "@/tools";
import { mapMutations, mapState } from "vuex";
import NavBar from "@/components/NavBar.vue";

var vfile = require("vfile");
import Doc from "./Doc";

import FrontMatter from "@/components/FrontMatter.vue";
export default {
  name: "Document",
  data() {
    return {
      documentPath:'',
      virtualFile: {},
      editor: {},
      title_editor: {},
      saved_message:'hit cmd+s to save file',
    };
  },
  components: {
    EditorContent,
    EditorMenuBar,
    Icon,
    FrontMatter,
    NavBar,
    Breadcrumb
  },
  computed: {
    ...mapState(["currentFile", "projectPath"]),
  },
  created() {
    this.documentPath = this.$route.params.documentPath
    this.fetchFile();
    console.log("content detail created");
    

    
  },
  mounted() {
    this.$mousetrap.bind("command+s", this.saveFile);
  },
  watch: {
    $route: "fetchFile"
  },
  beforeDestroy() {
    // Always destroy your editor instance when it's no longer needed
    this.editor.destroy();
    this.setCurrentFile({});
    console.log("destroyed content detail");
  },
  methods: {
    ...mapMutations(["setCurrentFile"]),

    fetchFile() {
      // manipulate document in the vfile spec https://github.com/vfile/vfile#use
      this.virtualFile = vfile({path: this.documentPath});
      let buffer = fs.readFileSync(
        path.join(this.documentPath),
        "utf8"
      );
      this.virtualFile.data.frontmatter = fm(buffer).data;
      this.virtualFile.data.contentmd = fm(buffer).content;
      this.virtualFile.data.contenthtml = md2html(this.virtualFile.data.contentmd);
      
      // Throw the vfile in vuex
      this.setCurrentFile(this.virtualFile);
      
      // setup 2 editors, one for title, one for content
      this.title_editor = new Editor({
        autoFocus: true,
        extensions:[
          new Doc(),
        ],
        content: this.virtualFile.data.frontmatter.title,
      });
      this.editor = new Editor({
        extensions: [
          new Blockquote(),
          new BulletList(),
          new CodeBlock(),
          new HardBreak(),
          new Heading({ levels: [1, 2, 3] }),
          new ListItem(),
          new OrderedList(),
          new TodoItem(),
          new TodoList(),
          new Link(),
          new Bold(),
          new Code(),
          new Italic(),
          new Strike(),
          new Underline(),
          new History(),
          new HorizontalRule()
        ],
        content: this.virtualFile.data.contenthtml
      });

    },
    saveFile() {
      this.currentFile.data.title = this.title_editor.getJSON().content[0].content[0].text;
      try {
        fs.writeFileSync(
          path.join(this.currentFile.path),
          fm.stringify(html2md(this.editor.getHTML()), this.currentFile.data),
          "utf8"
        );
        console.log(fm.stringify(html2md(this.editor.getHTML()), this.currentFile.data))
        this.saved_message = `Saved at ${Date()}`
        console.log("file saved");
      } catch (e) {
        console.error("file not saved", e);
        return;
      }
    },
    closeFile() {
      this.setCurrentFile({});
      this.$router.go(-1);
    }
  }
};
</script>

<style lang="scss">
$color-black: #000000;
$color-white: #ffffff;
$color-grey: #dddddd;

.editor {
  position: relative;
  @apply outline-none max-w-2xl mx-auto p-4;
}
.editor__content .ProseMirror {
  overflow-wrap: break-word;
  word-wrap: break-word;
  word-break: break-word;
  @apply bg-gray-50 rounded outline-none p-2;
  * {
    caret-color: currentColor;
  }

  h1 {
    @apply leading-none text-red-700;
    font-size: 2em;
    margin: 25px 0;
  }

  h2 {
    @apply leading-none text-red-700;
    font-size: 1.6em;
    margin: 25px 0;
  }

  h2 {
    @apply leading-none text-red-700;
    font-size: 1.2em;
    margin: 25px 0;
  }

  pre {
    padding: 0.7rem 1rem;
    border-radius: 5px;
    background: $color-black;
    color: $color-white;
    font-size: 0.8rem;
    overflow-x: auto;

    code {
      display: block;
    }
  }

  p code {
    padding: 0.2rem 0.4rem;
    border-radius: 5px;
    font-size: 0.8rem;
    font-weight: bold;
    background: rgba($color-black, 0.1);
    color: rgba($color-black, 0.8);
  }

  ul,
  ol {
    padding-left: 1rem;
    list-style: disc;
  }

  li > p,
  li > ol,
  li > ul {
    margin: 0;
    
  }

  a {
    color: inherit;
  }

  blockquote {
    border-left: 3px solid rgba($color-black, 0.1);
    color: rgba($color-black, 0.8);
    padding-left: 0.8rem;
    font-style: italic;

    p {
      margin: 0;
    }
  }

  img {
    max-width: 100%;
    border-radius: 3px;
  }

  table {
    border-collapse: collapse;
    table-layout: fixed;
    width: 100%;
    margin: 0;
    overflow: hidden;

    td,
    th {
      min-width: 1em;
      border: 2px solid $color-grey;
      padding: 3px 5px;
      vertical-align: top;
      box-sizing: border-box;
      position: relative;

      > * {
        margin-bottom: 0;
      }
    }

    th {
      font-weight: bold;
      text-align: left;
    }

    .selectedCell:after {
      z-index: 2;
      position: absolute;
      content: "";
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      background: rgba(200, 200, 255, 0.4);
      pointer-events: none;
    }

    .column-resize-handle {
      position: absolute;
      right: -2px;
      top: 0;
      bottom: 0;
      width: 4px;
      z-index: 20;
      background-color: #adf;
      pointer-events: none;
    }
  }

  .tableWrapper {
    margin: 1em 0;
    overflow-x: auto;
  }

  .resize-cursor {
    cursor: ew-resize;
    cursor: col-resize;
  }
}

.menubar {
  @apply bg-gray-50;
  transition: visibility 0.2s 0.4s, opacity 0.2s 0.4s;

  &.is-hidden {
    visibility: hidden;
    opacity: 0;
  }

  &.is-focused {
    visibility: visible;
    opacity: 1;
    transition: visibility 0.2s, opacity 0.2s;
  }

  &__button {
    font-weight: bold;
    display: inline-flex;
    background: transparent;
    border: 0;
    color: $color-black;
    padding: 0.2rem 0.5rem;
    margin-right: 0.2rem;
    border-radius: 3px;
    cursor: pointer;

    &:hover {
      background-color: rgba($color-black, 0.05);
    }

    &.is-active {
      background-color: rgba($color-black, 0.1);
    }
  }

  span#{&}__button {
    font-size: 13.3333px;
  }
}
</style>
