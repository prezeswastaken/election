<script lang="ts" setup>
const props = defineProps({
  id: Number,
  name: String,
  score: Number,
});

const emit = defineEmits(["fetch-contestants"]);

async function handleSubmit() {
  const response = await useFetch(
    `http://localhost:8000/api/contestants/${props.id}`,
    {
      method: "delete",
    },
  );
  emit("fetch-contestants");
  console.log("Deleting contestant with id", props.id);
  console.log(response.status);
}
</script>

<template>
  <form
    @submit.prevent="handleSubmit"
    class="flex justify-around py-3 px-5 w-96 rounded-lg border border-white"
  >
    <div class="flex flex-col w-full">
      <p>{{ name }}</p>
      <p>{{ score }} (%)</p>
    </div>
    <button
      class="flex items-center text-center text-red-300 duration-300 hover:text-red-500 text-red"
    >
      X
    </button>
  </form>
</template>

<style scoped></style>
