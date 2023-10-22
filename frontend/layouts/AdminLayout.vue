<script lang="ts" setup></script>

<template>
  <div
    class="flex flex-col gap-10 items-center p-20 min-h-screen font-serif text-3xl text-white bg-slate-950"
  >
    <h1 class="text-7xl">Admin panel</h1>
    <NavBar />
    <slot />
  </div>
</template>

<style scoped></style>
