<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-primary text-white shadow-6">
      <q-toolbar>
        <q-toolbar-title>
          <q-btn flat color="white" label="Todo-vue" @click="goHome" />
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container class="q-pa-md">
      <router-view />
    </q-page-container>

    <q-footer elevated class="bg-green-6 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <a href="https://github.com/sebastianaldi17/DockerVue-Todo"
            style="text-decoration: none; color: aliceblue;">Source code</a>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script>
export default {
  methods: {
    goHome() {
      this.$router.push("/")
    }
  },
  name: 'LayoutDefault',
}
</script>
