<script lang="ts" setup>
const props = defineProps({
  link: String,
  text: String,
});
</script>

<template>
  <NuxtLink
    class="capitalize duration-300 hover:text-3xl hover:text-green-400"
    :to="link"
    >{{ text }}
  </NuxtLink>
</template>

<style scoped>
.router-link-active {
  @apply text-teal-500;
}
</style>
