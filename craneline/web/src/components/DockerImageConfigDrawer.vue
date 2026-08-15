<script setup lang="ts">
import type { GraphNode } from '@vue-flow/core'
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle } from '@/components/ui/drawer'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'

defineProps<{
  node: GraphNode | null
}>()

const open = defineModel<boolean>('open', { required: true })
</script>

<template>
  <Drawer v-model:open="open" swipe-direction="right">
    <DrawerContent v-if="node">
      <DrawerHeader>
        <DrawerTitle>{{ node.data.label }}</DrawerTitle>
      </DrawerHeader>

      <div class="p-4 space-y-3">
        <div>
          <Label>Nom</Label>
          <Input v-model="node.data.label" />
        </div>
        <div>
          <Label>IP</Label>
          <Input v-model="node.data.ip" />
        </div>
        <div>
          <Label>Port</Label>
          <Input v-model="node.data.port" type="number" />
        </div>
      </div>
    </DrawerContent>
  </Drawer>
</template>