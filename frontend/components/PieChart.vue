<script lang="ts" setup>
const props = defineProps({
  names: Array<String>,
  scores: Array<Number>,
  type: String,
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
  <div><img class="w-72" :src="url" /></div>
</template>

<style scoped></style>
