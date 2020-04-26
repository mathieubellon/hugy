import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'

import ProjectConsole from '@/views/ProjectConsole.vue'
import Document from '@/views/Document.vue'
import AppSettings from '@/views/AppSettings.vue'
import ProjectSettings from '@/views/ProjectSettings.vue'
import ProjectPublish from '@/views/ProjectPublish.vue'
import ProjectFiles from '@/views/ProjectFiles.vue'
import ProjectNav from "@/components/ProjectNav.vue";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/projectfiles/:currentDir',
    name: 'projectfiles',
    components: { default: ProjectFiles, navbar: ProjectNav },
  },
  {
    path: '/projectsettings',
    name: 'projectsettings',
    components: { default: ProjectSettings, navbar: ProjectNav },
    props: { default: true, navbar: false }
  },
  {
    path: '/projectconsole',
    name: 'projectconsole',
    components: { default: ProjectConsole, navbar: ProjectNav },
    props: { default: true, navbar: false }
  },
  {
    path: '/projectpublish',
    name: 'projectpublish',
    components: { default: ProjectPublish, navbar: ProjectNav },
    props: { default: true, navbar: false }
  },

  {
    path: '/document/:documentPath',
    name: 'document',
    component: Document,
  },
  {
    path: '/settings',
    name: 'settings',
    component: AppSettings,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
console.log(router)
export default router
