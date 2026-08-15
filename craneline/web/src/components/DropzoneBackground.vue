<script lang="ts" setup>
import { Background } from '@vue-flow/background'
</script>

<template>
  <div class="w-full h-full">
    <Background :size="2" :gap="20" patternColor="#d4d4d8" />

    <div class="w-full h-full pointer-events-none z-1 flex justify-center items-center">
      <slot />
    </div>
  </div>
</template>
