<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'
import { SlidersHorizontal } from '@lucide/vue';
import { Button } from '@/components/ui/button'

const props = defineProps<NodeProps>()
const emit = defineEmits<{ openConfig: [nodeId: string] }>()
</script>

<template>
  <div
    class="w-full h-full rounded border p-2 transition-all"
    :class="props.selected
      ? 'bg-teal-700 border-teal-500 ring-1 ring-teal-500 text-white shadow-lg shadow-teal-600/50'
      : 'bg-teal-600 border-teal-700 text-white shadow-sm'"
  >
    <div class="flex justify-between items-center">
      <div class="font-semibold text-sm">{{ data.label }}:{{ data.tag }}</div>
      <Button variant="outline" size="icon" class="nodrag bg-transparent size-auto p-1" @click="emit('openConfig', props.id)"><SlidersHorizontal /></Button>
    </div>
    <div class="text-xs opacity-80 mb-1">{{ data.description }}</div>
    <div class="text-xs space-y-0.5">
      <div v-for="port in data.ports" :key="port">Port: {{ port }}</div>
    </div>
  </div>

  <Handle id="input-1" type="source" :position="Position.Left" :style="{ top: '30%', width: '0.5em', height: '0.5em', left: '0.08em' }" />
  <Handle id="input-2" type="source" :position="Position.Left" :style="{ top: '70%', width: '0.5em', height: '0.5em', left: '0.08em' }" />

  <Handle id="output-1" type="source" :position="Position.Right" :style="{ top: '30%', width: '0.5em', height: '0.5em', right: '0.08em' }" />
  <Handle id="output-2" type="source" :position="Position.Right" :style="{ top: '70%', width: '0.5em', height: '0.5em', right: '0.08em' }" />
</template>