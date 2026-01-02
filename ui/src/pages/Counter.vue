<script setup>
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>

<template>
  <div>Current Count: {{ counter.count }}</div>
  <button class="btn" v-on:click="counter.increment()">increment</button>
  <button class="btn" v-on:click="counter.decrement()">decrement</button>
</template>