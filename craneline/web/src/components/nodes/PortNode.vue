<!-- PortsInfoNode.vue -->
<script setup lang="ts">
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import type { NodeProps } from '@vue-flow/core'
import { TriangleAlert } from '@lucide/vue';

const props = defineProps<NodeProps>()
</script>

<template>
  <Card class="w-64 border-dashed border-2 border-amber-500/50 bg-zinc-900/95 shadow-none">
    <CardHeader class="pb-2 pt-3 px-3">
      <CardTitle class="text-sm flex items-center gap-2 text-amber-400">
        <TriangleAlert class="size-4" />
        Exposed ports
      </CardTitle>
    </CardHeader>

    <Separator class="bg-zinc-700" />

    <CardContent class="px-3 py-2 space-y-1.5">
      <div
        v-for="portInfo in data.ports"
        :key="portInfo.port"
        class="flex items-center justify-between text-xs"
      >
        <span class="text-zinc-300">{{ portInfo.service }}</span>
        <div class="flex items-center gap-1.5">
          <Badge variant="outline" class="text-zinc-400 border-zinc-600 font-mono text-[10px] px-1.5">
            {{ portInfo.port }}
          </Badge>
          <Badge
            :variant="portInfo.public ? 'destructive' : 'secondary'"
            class="text-[10px] px-1.5"
          >
            {{ portInfo.public ? 'Public' : 'Interne' }}
          </Badge>
        </div>
      </div>
    </CardContent>
  </Card>
</template>