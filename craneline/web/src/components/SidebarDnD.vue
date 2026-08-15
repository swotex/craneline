<script setup lang="ts">
// import useDragAndDrop from '@/assets/useDnD.js'

interface DockerImage {
  name: string
  tag: string
  ports: number[]
  category: 'network' | 'database' | 'app'
  description: string
}

// const { onDragStart } = useDragAndDrop()
const props = defineProps<{ images: DockerImage[] }>()


function onDragStart(event: DragEvent, image: DockerImage) {
  event.dataTransfer?.setData('application/vueflow', JSON.stringify(image))
  event.dataTransfer.effectAllowed = 'move'
}
</script>

<template>
  <aside class="bg-zinc-900 border-r border-zinc-700">
    <div class="text-white">You can drag these nodes to the pane.</div>

    <!-- <div class="nodes"> -->
    <div class="flex flex-col items-center gap-3">
      
      <div
      v-for="image in images"
      :key="image.name"
      class="bg-emerald-400 border border-emerald-500 text-white w-38 h-10 rounded flex justify-center items-center"
      :draggable="true"
      @dragstart="onDragStart($event, image)"
    >
      {{ image.name }}:{{ image.tag }}
    </div>

      <!-- <div class="bg-sky-400 border border-sky-500 text-white w-38 h-10 rounded flex justify-center items-center" :draggable="true" @dragstart="onDragStart($event, 'network')">Network</div> -->

      <!-- <div class="bg-emerald-400 border border-emerald-500 text-white w-38 h-10 rounded flex justify-center items-center" :draggable="true" @dragstart="onDragStart($event, 'custom')">Output Node</div> -->
    </div>
  </aside>
</template>
