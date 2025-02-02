<script setup lang="ts">
import {
  ElContainer,
  ElHeader,
  ElAside,
  ElMain,
  ElFooter,
  ElButton,
  ElMenu,
  ElMenuItem,
  ElSwitch,
  ElSubMenu,
  ElAvatar
} from 'element-plus';
import { Moon, Sunny, UserFilled } from '@element-plus/icons-vue';
import { useConfigStore } from '@/store'
const config = useConfigStore()
</script>

<template>
  <div class="common-layout">
    <ElContainer id="main-top" style="height: 100vh;">
      <ElHeader id="main-header">
        <div class="left">SDCRAFT</div>
        <div class="right">
          <ElMenu :default-active="$route.path" router mode="horizontal" :ellipsis="false" :collapse="false">
            <ElMenuItem index="/">Index</ElMenuItem>
            <ElSubMenu index="/account">
              <template #title>Account</template>
              <ElMenuItem index="/signin">Login</ElMenuItem>
              <ElMenuItem index="2-2">Register</ElMenuItem>
            </ElSubMenu>
          </ElMenu>
          <ElSwitch
          @change="config.toggleDark()"
          v-model=" config.dark"
          inline-prompt
          >
            <template #active-action>
              <Moon class="light-icon"></Moon>
            </template>
            <template #inactive-action>
              <Sunny class="light-icon"></Sunny>
            </template>
          </ElSwitch>
        </div>
      </ElHeader>
      <ElContainer>
        <ElMain>
          <RouterView/>
        </ElMain>
      </ElContainer>
      <ElFooter>Copyright 2025-Now SDCRAFT.</ElFooter>
    </ElContainer>
  </div>
</template>
<style scoped>
</style>
