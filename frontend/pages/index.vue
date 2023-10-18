<script setup lang="ts">
type Contestant = {
  name: string;
  id: number;
  score: number;
};

const contestantsFetched = ref(false);

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
  contestantsFetched.value = true;
}
</script>

<template>
  <div
    class="flex flex-col gap-10 justify-center items-center min-h-screen font-serif text-white uppercase bg-slate-950"
  >
    <h1 class="mb-5 text-7xl">Wyniki wyborów do sejmu 2023</h1>
    <div class="flex gap-5">
      <Chart
        v-if="contestantsFetched"
        :names="contestants.map((contestant) => contestant.name)"
        :scores="contestants.map((contestant) => contestant.score)"
      />
    </div>
  </div>
</template>
