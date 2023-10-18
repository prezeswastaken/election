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
    class="flex flex-col justify-center items-center min-h-screen text-white bg-slate-950"
  >
    <div class="flex gap-5">
      <Chart
        v-if="contestantsFetched"
        :names="contestants.map((contestant) => contestant.name)"
        :scores="contestants.map((contestant) => contestant.score)"
        type="pie"
      />
    </div>
  </div>
</template>
