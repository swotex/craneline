<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'

const props = defineProps<NodeProps>()
</script>

<template>
  <div
    class="w-full h-full rounded border p-2 transition-all"
    :class="props.selected
      ? 'bg-yellow-700 border-yellow-500 ring-1 ring-yellow-500 text-white shadow-lg shadow-yellow-600/50'
      : 'bg-yellow-600 border-yellow-700 text-white shadow-sm'"
  >
    <div class="font-semibold text-sm">{{ data.label }}</div>
    <div class="text-xs opacity-80">
      {{ Object.keys(data.vars ?? {}).length }} variable(s)
    </div>
  </div>

  <Handle type="target" :position="Position.Left" />
  <Handle type="source" :position="Position.Right" />
</template>