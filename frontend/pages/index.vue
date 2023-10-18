<script setup lang="ts">
type Contestant = {
  name: string;
  id: number;
  score: number;
};

const contestantsFetched = ref(false);

const contestants: Ref<Contestant[]> = ref([]);

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
    <button
      class="p-3 m-10 text-white bg-black rounded-3xl duration-300 hover:text-green-500"
      @click="fetchContestants"
    >
      SHOW RESULTS
    </button>
    <div class="flex gap-5">
      <PieChart
        v-if="contestantsFetched"
        :names="contestants.map((contestant) => contestant.name)"
        :scores="contestants.map((contestant) => contestant.score)"
        type="pie"
      />
      <PieChart
        v-if="contestantsFetched"
        :names="contestants.map((contestant) => contestant.name)"
        :scores="contestants.map((contestant) => contestant.score)"
        type="line"
      />
    </div>
    <div v-for="contestant in contestants" :key="contestant.id">
      {{ contestant.name }}
      {{ contestant.score }}
    </div>
  </div>
</template>
