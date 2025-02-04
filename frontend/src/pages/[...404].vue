<script setup lang="ts">
import { ElButton, ElProgress } from 'element-plus';
import { onMounted, ref } from 'vue';

var percentage = ref(0)

function backToHome() {
   window.location.href = '/'
}

onMounted(() => {
  setInterval(() => {
    if (percentage.value >= 100) {
      backToHome()
    } else {
      percentage.value += 1
    }
  }, 50)
})
</script>

<template>
  <div class="container">
    <div>
      <h1>404</h1>
      <h1>Not Found</h1>
      <ElProgress type="dashboard" :percentage="percentage" :width=300 :stroke-width=20>
        <template #default="{ percentage }">
          <span class="percentage-value">Redirecting...</span>
          <span class="percentage-label">{{ Math.round((100 - Math.min(percentage, 100)) / 20) }} s</span>
          <ElButton type="primary" round @click="backToHome()" style="margin-top: 0.25rem;">Tap to go</ElButton>
        </template>
      </ElProgress>
    </div>
  </div>
</template>

<style lang="css" scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  > div {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-self: center;
  }
  .percentage-value {
    display: block;
    margin-top: 10px;
    font-size: 28px;
  }
  .percentage-label {
    display: block;
    margin-top: 10px;
    font-size: 20px;
  }
  h1 {
    margin-block-start: 0;
    margin-block-end: 0.5rem;
  }
  svg {
    transition: 0s !important;
  }
}
</style>
