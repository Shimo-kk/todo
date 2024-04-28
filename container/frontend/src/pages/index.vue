<script setup lang="ts">
import { useCsrfStore } from '@/store/csrfStore'

const csrf = useCsrfStore()
</script>

<template>
  <div>
    <h1 class="text-3xl font-bold underline">Top</h1>
    <h2>{{ csrf.token }}</h2>
  </div>
</template>
