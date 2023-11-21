<script lang="ts" setup>
  //
</script>

<template>
  <v-container class="fill-height">
    <v-responsive class="text-center fill-height">
      <v-img height="50" src="@/assets/logo.svg" class="logo mt-2" />
      <v-main class="mt-n2">
        <router-view />
      </v-main>
    </v-responsive>
  </v-container>
</template>

<style>
  .logo img {
    width: initial !important;
  }
</style>