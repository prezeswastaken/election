<script lang="ts" setup>
definePageMeta({
  layout: "admin-layout",
});
type Contestant = {
  id: number;
  name: string;
  score: number;
};

const contestants: Ref<Contestant[]> = ref([]);

onMounted(() => {
  fetchContestants();
});

async function fetchContestants() {
  const { data, error } = await useFetch(
    "http://localhost:8000/api/contestants",
  );
  if (error.value != null) {
    console.log("Error from fetch: ", error.value);
    return;
  }
  contestants.value = data.value as Contestant[];
  console.log("Data from nuxt fetch: ", contestants.value);
  console.log(contestants.value.map((contestant) => contestant.name));
  console.log(contestants.value.map((contestant) => contestant.score));
}
</script>

<template>
  <div v-for="contestant in contestants" :key="contestant.id">
    <Contestant
      :id="contestant.id"
      :name="contestant.name"
      :score="contestant.score"
      @fetch-contestants="fetchContestants"
    />
  </div>
</template>

<style scoped></style>
