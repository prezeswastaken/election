<script lang="ts" setup>
definePageMeta({
  layout: "admin-layout",
});
const form = ref({
  name: "",
  score: "",
});

async function handleSubmit() {
  const response = await useFetch(
    "http://localhost:8000/api/contestants/create",
    {
      method: "POST",
      body: { ...form.value, score: parseFloat(form.value.score) },
    },
  );
  form.value = {
    name: "",
    score: "",
  };

  console.log(response);
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="flex flex-col gap-5">
    <label class="flex flex-col gap-5">
      Name
      <input
        v-model="form.name"
        class="py-3 px-5 rounded-lg border border-white bg-slate-950"
      />
    </label>
    <label class="flex flex-col gap-3">
      Score (%)
      <input
        v-model="form.score"
        class="py-3 px-5 rounded-lg border border-white bg-slate-950"
      />
    </label>
    <button class="duration-300 hover:text-green-500" type="submit">
      SUBMIT
    </button>
  </form>
</template>

<style scoped></style>
