<script setup lang="ts">
import { ref, onMounted } from 'vue'
import useDragAndDrop from '@/assets/useDnD.js'
import { TriangleAlert } from '@lucide/vue';
import { imagesService } from '@/services/imagesService'

interface DockerImage {
  name: string
  tag: string
  ports: number[]
  category: 'network' | 'database' | 'app'
  description: string
}

const { onDragStart } = useDragAndDrop()
// const props = defineProps<{ images: DockerImage[] }>()
const images = ref([])
const error = ref(null)


function onDragStartImage(event: DragEvent, image: DockerImage | null, type: string) {
  if (image != null) {
    event.dataTransfer?.setData('application/vueflow', JSON.stringify(image))
  } 
  
  event.dataTransfer?.setData('type', type)
  event.dataTransfer.effectAllowed = 'move'
  // draggedType.value = type
}

onMounted(async () => {
  try {
    images.value = await imagesService.getAll()
  } catch (e : any) {
    error.value = e.message
  } finally {
    // loading.value = false
  }
})
</script>

<template>
  <aside class="bg-zinc-900 border-r border-zinc-700">
    <div class="text-white">You can drag these nodes to the pane.</div>

    
    <!-- <div class="nodes"> -->
      <div class="flex flex-col items-center gap-3">

      <!-- <div class="bg-sky-400 border border-sky-500 text-white w-38 h-10 rounded flex justify-center items-center" :draggable="true" @dragstart="onDragStartImage($event, null, 'network')">Network</div> -->
      <div
        class="bg-indigo-400 border border-indigo-500 text-white w-38 h-10 rounded flex justify-center items-center cursor-grab hover:bg-indigo-300 transition-colors"
        :draggable="true"
        @dragstart="onDragStartImage($event, null, 'network')"
      >
        Network
      </div>

      <div
        class="bg-amber-700 border border-amber-800 text-white w-38 h-10 rounded flex justify-center items-center cursor-grab hover:bg-amber-600 transition-colors"
        :draggable="true"
        @dragstart="onDragStartImage($event, null, 'volume')"
      >
        Volume
      </div>

      <div
        class="bg-yellow-600 border border-yellow-700 text-white w-38 h-10 rounded flex justify-center items-center cursor-grab hover:bg-yellow-500 transition-colors"
        :draggable="true"
        @dragstart="onDragStartImage($event, null, 'envFile')"
      >
        .env
      </div>


      <div
        class="bg-zinc-900 border-2 border-dashed border-amber-500/50 text-amber-400 w-38 h-10 rounded flex justify-center items-center gap-1.5 cursor-grab hover:border-amber-400 hover:bg-zinc-800 transition-colors"
        :draggable="true"
        @dragstart="onDragStartImage($event, null, 'ports')"
      >
        <TriangleAlert class="size-3" />
        Ports exposés
      </div>
      
      <div
        v-for="image in images"
        :key="image.id"
        class="bg-teal-600 border border-teal-700 text-white w-38 h-10 rounded flex justify-center items-center cursor-grab hover:bg-teal-500 transition-colors"
        :draggable="true"
        @dragstart="onDragStartImage($event, image, 'dockerImage')"
      >
        {{ image.name }}
      </div>

      

      <!-- <div class="bg-emerald-400 border border-emerald-500 text-white w-38 h-10 rounded flex justify-center items-center" :draggable="true" @dragstart="onDragStart($event, 'custom')">Output Node</div> -->
    </div>
  </aside>
</template>
