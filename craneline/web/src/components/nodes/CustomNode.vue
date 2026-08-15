<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core'
import type { NodeProps } from '@vue-flow/core'
import { SlidersHorizontal } from '@lucide/vue';
import { NodeResizeControl } from '@vue-flow/node-resizer'
import '@vue-flow/node-resizer/dist/style.css'

import { Button } from '@/components/ui/button'

const props = defineProps<NodeProps>()
const emit = defineEmits<{ openConfig: [nodeId: string] }>()
</script>

<template>

  <div class="w-full h-full rounded px-4 py-2 text-white hover:bg-emerald-400" :class="props.selected 
      ? 'bg-emerald-500 border border-amber-400' 
      : 'bg-emerald-400 border border-emerald-500'">
      <Button variant="outline" size="icon" class="nodrag bg-transparent" @click="emit('openConfig', props.id)"><SlidersHorizontal /></Button>
    {{ data.label ?? 'Custom Node' }}
  </div>

 <!-- Plusieurs handles en entrée, à gauche -->
  <Handle id="input-1" type="source" :position="Position.Left" :style="{ top: '30%', width: '0.5em', height: '0.5em', left: '0.08em' }" />
  <Handle id="input-2" type="source" :position="Position.Left" :style="{ top: '70%', width: '0.5em', height: '0.5em', left: '0.08em' }" />

  <!-- Plusieurs handles en sortie, à droite -->
  <Handle id="output-1" type="source" :position="Position.Right" :style="{ top: '30%', width: '0.5em', height: '0.5em', right: '0.08em' }" />
  <Handle id="output-2" type="source" :position="Position.Right" :style="{ top: '70%', width: '0.5em', height: '0.5em', right: '0.08em' }" />
</template>