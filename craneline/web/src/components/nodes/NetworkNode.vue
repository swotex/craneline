<script setup lang="ts">
import { Handle, Position } from '@vue-flow/core'
import { NodeToolbar } from '@vue-flow/node-toolbar'
import type { NodeProps } from '@vue-flow/core'

const props = defineProps<NodeProps>()
</script>

<template>
  <!-- <NodeToolbar :is-visible="data.toolbarVisible" :position="data.toolbarPosition">
    <button>delete</button>
    <button>copy</button>
    <button>expand</button>
  </NodeToolbar> -->

    <div class="rounded px-4 py-2 text-white hover:bg-sky-500" :class="props.selected 
      ? 'bg-sky-500 border border-amber-400 ' 
      : 'bg-sky-400 border border-sky-500'">
        <div>
            Network
        </div>
        <div>
            {{ data.label ?? 'Custom Node' }}
        </div>
    </div>

 <!-- Plusieurs handles en entrée, à gauche -->
  <!-- <Handle id="input-1" type="target" :position="Position.Left" :style="{ top: '70%', width: '0.5em', height: '0.5em', left: '0.05em' }" />
  <Handle id="input-2" type="target" :position="Position.Right" :style="{ top: '70%', width: '0.5em', height: '0.5em', right: '0.05em' }" />
  <Handle id="input-3" type="target" :position="Position.Bottom" :style="{ left: '25%', width: '0.5em', height: '0.5em', bottom: '0.05em' }" />
  <Handle id="input-4" type="target" :position="Position.Bottom" :style="{ left: '75%', width: '0.5em', height: '0.5em', bottom: '0.05em' }" /> -->

  <Handle id="input-1" type="target" :position="Position.Left" :style="{ width: '0.5em', height: '0.5em', left: '0.05em' }" />
  <Handle id="input-2" type="target" :position="Position.Right" :style="{ width: '0.5em', height: '0.5em', right: '0.05em' }" />
  <Handle id="input-3" type="target" :position="Position.Bottom" :style="{ width: '0.5em', height: '0.5em', bottom: '0.05em' }" />
  <Handle id="input-4" type="target" :position="Position.Top" :style="{ width: '0.5em', height: '0.5em', top: '0.05em' }" />

  <!-- Plusieurs handles en sortie, à droite -->
  <!-- <Handle id="output-1" type="target" :position="Position.Right" :style="{ top: '30%', width: '0.5em', height: '0.5em', right: '0.05em' }" /> -->
  <!-- <Handle id="output-2" type="target" :position="Position.Right" :style="{ top: '70%', width: '0.5em', height: '0.5em', right: '0.05em' }" /> -->
</template>