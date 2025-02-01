<script setup lang="ts">
import { ElContainer,ElHeader,ElAside,ElMain,ElFooter } from 'element-plus';
</script>

<template >
  <div class="common-layout">
    <ElContainer>
      <ElHeader>Header</ElHeader>
      <ElContainer>
        <ElAside width="200px">Aside</ElAside>
        <ElContainer>
          <ElMain>Main</ElMain>
          <ElFooter>Copyright 2025-Now SDCRAFT.</ElFooter>
        </ElContainer>
      </ElContainer>
    </ElContainer>
  </div>
</template>
<style scoped>
</style>
