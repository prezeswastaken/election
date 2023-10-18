<script lang="ts" setup>
import { Bar } from "vue-chartjs";
import { Doughnut } from "vue-chartjs";
const props = defineProps({
  names: Array<String>,
  scores: Array<Number>,
  type: String,
});

const chartData = ref({
  labels: props.names,
  datasets: [
    {
      label: "Score (%)",
      data: props.scores,
      backgroundColor: [
        "#c6a0f6",
        "#ed8796",
        "#f5a97f",
        "#eed49f",
        "#a6da95",
        "#7dc4e4",
        "#f5bde6",
      ],
      color: "#cad3f5",
      hoverOffset: 50,
    },
  ],
});
const chartOptions = {
  plugins: {
    legend: {
      labels: {
        color: "#FFFFFF",
      },
    },
  },
};
</script>

<template>
  <div class="flex gap-32 justify-center items-center">
    <div>
      <Bar class="w-96" :data="chartData" :options="chartOptions" />
    </div>
    <div>
      <Doughnut class="w-72" :data="chartData" :options="chartOptions" />
    </div>
  </div>
</template>

<style scoped></style>
