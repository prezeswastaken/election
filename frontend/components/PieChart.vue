<script lang="ts" setup>
import { Bar } from "vue-chartjs";
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
      backgroundColor: "#8bd5ca",
      data: props.scores,
    },
  ],
});
// Define your chart configuration
let chartConfig = {
  type: props.type,
  data: {
    labels: props.names,
    datasets: [
      {
        label: "Score (%)",
        data: props.scores,
      },
    ],
  },
};

// Convert the chart configuration to a JSON string
let chartConfigJson = JSON.stringify(chartConfig);

// Define the URL for the QuickChart.io API
let url = `https://quickchart.io/chart?c=${chartConfigJson}`;
</script>

<template>
  <div>
    <Bar :data="chartData" />
  </div>
</template>

<style scoped></style>
