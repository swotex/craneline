<script setup lang="ts">
import { ref } from 'vue'
import { VueFlow, useVueFlow, ConnectionLineType } from '@vue-flow/core'
import { Background } from '@vue-flow/background'
import { MiniMap } from '@vue-flow/minimap'

import CustomNode from '@/components/nodes/CustomNode.vue'
import NetworkNode from '@/components/nodes/NetworkNode.vue'
import useDragAndDrop from '@/assets/useDnD.js'
import DropzoneBackground from '@/components/DropzoneBackground.vue'
import Sidebar from '@/components/SidebarDnD.vue'
import type { NodeMouseEvent } from '@vue-flow/core'
import DockerImageDrawer from '@/components/DockerImageConfigDrawer.vue'
import PortNode from '@/components/nodes/PortNode.vue'
import VolumeNode from '@/components/nodes/VolumeNode.vue'
import EnvFileNode from '@/components/nodes/EnvFileNode.vue'


import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from '@/components/ui/drawer'


import { mockDockerImages } from './mockDockerImages'

const { onDragOver, onDrop, onDragLeave, isDragOver } = useDragAndDrop()
const { onConnect, addEdges, onNodeDoubleClick, findNode } = useVueFlow()


import '@vue-flow/minimap/dist/style.css'

const nodes = ref([])


const drawerOpen = ref(false)
const selectedNode = ref()

onNodeDoubleClick((event: NodeMouseEvent) => {
  selectedNode.value = event.node
  drawerOpen.value = true
})

function handleOpenConfig(nodeId: string) {
  const node = findNode(nodeId)
  console.log('node trouvé:', node) // <-- vérifie ça
  selectedNode.value = node
  drawerOpen.value = true
  console.log('drawerOpen:', drawerOpen.value) // <-- et ça
}

onConnect(addEdges)


// Helper lines

import { onMounted, onUnmounted } from 'vue'
import type { NodeChange } from '@vue-flow/core'
import { getHelperLines } from '@/assets/helperLines.ts'
import HelperLines from '@/components/HelperLines.vue'

const { onNodesChange, applyNodeChanges, getNodes } = useVueFlow()

const isShiftPressed = ref(false)
const helperLineHorizontal = ref<number | undefined>(undefined)
const helperLineVertical = ref<number | undefined>(undefined)

function onKeyDown(e: KeyboardEvent) {
  if (e.key === 'Shift') isShiftPressed.value = true
}
function onKeyUp(e: KeyboardEvent) {
  if (e.key === 'Shift') isShiftPressed.value = false
}

onMounted(() => {
  window.addEventListener('keydown', onKeyDown)
  window.addEventListener('keyup', onKeyUp)
})
onUnmounted(() => {
  window.removeEventListener('keydown', onKeyDown)
  window.removeEventListener('keyup', onKeyUp)
})

// on intercepte les changements AVANT qu'ils soient appliqués
onNodesChange((changes: NodeChange[]) => {
  // reset par défaut
  helperLineHorizontal.value = undefined
  helperLineVertical.value = undefined

  // on ne modifie le comportement que si :
  // - shift est appuyé
  // - un seul node est en train d'être déplacé (dragging)
  // - c'est bien un changement de type 'position'
  if (
    isShiftPressed.value &&
    changes.length === 1 &&
    changes[0].type === 'position' &&
    changes[0].dragging &&
    changes[0].position
  ) {
    const result = getHelperLines(changes[0], getNodes.value)

    helperLineHorizontal.value = result.horizontal
    helperLineVertical.value = result.vertical

    // on écrase la position du node par la position "snappée"
    changes[0].position.x = result.snapPosition.x ?? changes[0].position.x
    changes[0].position.y = result.snapPosition.y ?? changes[0].position.y
  }

  applyNodeChanges(changes)
})
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/components/ui/resizable'
</script>

<template>

<ResizablePanelGroup
    direction="horizontal"
    class="w-full h-full"
    @drop="onDrop"
  >
    <ResizablePanel class="min-w-45 max-w-[80%]" :default-size="15">
        <Sidebar class="w-full h-full" :images="mockDockerImages" />
    </ResizablePanel>

    <ResizableHandle />

    <ResizablePanel :default-size="85">
      <VueFlow class="bg-zinc-800" :nodes="nodes" @dragover="onDragOver" @dragleave="onDragLeave" :delete-key-code="['Backspace', 'Delete']" :default-edge-options="{ type: 'smoothstep' }" :connection-line-type="ConnectionLineType.SmoothStep">
        <DropzoneBackground
          :style="{
            backgroundColor: isDragOver ? '#52525b' : 'transparent',
            transition: 'background-color 0.2s ease',
          }"
        >
          <p class="text-white" v-if="isDragOver">Drop here</p>
        </DropzoneBackground>
        <Background :gap="50" patternColor="#d4d4d8" variant="lines" />
        <MiniMap maskColor="#3f3f46" pannable zoomable />
        <template #node-custom="nodeProps">
          <CustomNode v-bind="nodeProps" @open-config="handleOpenConfig" />
        </template>
        <template #node-network="nodeProps">
          <NetworkNode v-bind="nodeProps" />
        </template>
        <template #node-dockerImage="nodeProps">
          <CustomNode v-bind="nodeProps"  @open-config="handleOpenConfig" />
        </template>
        <template #node-ports="nodeProps">
          <PortNode v-bind="nodeProps" />
        </template>
        <template #node-volume="nodeProps">
          <VolumeNode v-bind="nodeProps" />
        </template>
        <template #node-envFile="nodeProps">
          <EnvFileNode v-bind="nodeProps" />
        </template>


        
        <HelperLines
          :horizontal="helperLineHorizontal"
          :vertical="helperLineVertical"
        />
      </VueFlow>
    </ResizablePanel>
  </ResizablePanelGroup>
  <DockerImageDrawer v-model:open="drawerOpen" :node="selectedNode" />
</template>

<style>
/* these are necessary styles for vue flow */
@import '@vue-flow/core/dist/style.css';

/* this contains the default theme, these are optional styles */
@import '@vue-flow/core/dist/theme-default.css';
html, body, #app {
  height: 100%;
  margin: 0;
}

.app-wrapper {
  width: 100%;
  height: 100vh;
}
</style>
