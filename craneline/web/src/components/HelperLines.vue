<!-- HelperLines.vue -->
<script setup lang="ts">
import { useVueFlow } from '@vue-flow/core'

defineProps<{
  horizontal?: number
  vertical?: number
}>()

const { viewport } = useVueFlow()
</script>

<template>
  <svg
    class="pointer-events-none absolute inset-0 h-full w-full"
    style="z-index: 10;"
  >
    <line
      v-if="horizontal !== undefined"
      :x1="0" :x2="'100%'"
      :y1="horizontal * viewport.zoom + viewport.y"
      :y2="horizontal * viewport.zoom + viewport.y"
      stroke="#0ea5e9" stroke-width="1" stroke-dasharray="4"
    />
    <line
      v-if="vertical !== undefined"
      :y1="0" :y2="'100%'"
      :x1="vertical * viewport.zoom + viewport.x"
      :x2="vertical * viewport.zoom + viewport.x"
      stroke="#0ea5e9" stroke-width="1" stroke-dasharray="4"
    />
  </svg>
</template>