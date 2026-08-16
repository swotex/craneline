<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'

const props = defineProps<NodeProps>()
</script>

<template>
  <div
    class="w-full h-full rounded border p-2 transition-all"
    :class="props.selected
      ? 'bg-amber-800 border-amber-600 ring-1 ring-amber-600 text-white shadow-lg shadow-amber-700/50'
      : 'bg-amber-700 border-amber-800 text-white shadow-sm'"
  >
    <div class="font-semibold text-sm">{{ data.label }}</div>
    <div class="text-xs opacity-80">{{ data.driver }}</div>
  </div>

  <Handle type="target" :position="Position.Left" />
  <Handle type="source" :position="Position.Right" />
</template>