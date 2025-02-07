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
  ElAvatar,
  ElDropdown,
  ElDropdownMenu,
  ElDropdownItem,
  ElInput
} from 'element-plus';
import { Moon, Sunny, UserFilled } from '@element-plus/icons-vue';
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark({disableTransition: false})
const toggleDark = useToggle(isDark)
</script>

<template>
  <div class="common-layout">
    <ElContainer id="main-top" style="height: 100vh;">
      <ElHeader id="main-header">
        <div class="left">
          SDCRAFT
          <ElMenu :default-active="$route.path" router mode="horizontal" :ellipsis="false" :collapse="false">
            <ElMenuItem index="/">Index</ElMenuItem>
            <ElMenuItem index="/signin">Login</ElMenuItem>
            <ElMenuItem index="/signup">Register</ElMenuItem>
          </ElMenu>
        </div>
        <div class="right">
          <ElSwitch
          v-model="isDark"
          inline-prompt
          >
            <template #active-action>
              <Moon class="light-icon"></Moon>
            </template>
            <template #inactive-action>
              <Sunny class="light-icon"></Sunny>
            </template>
          </ElSwitch>
          <ElDropdown>
            <ElAvatar :icon="UserFilled"></ElAvatar>
            <template #dropdown>
              <ElDropdownMenu>
                <ElDropdownItem>Sign in</ElDropdownItem>
                <ElDropdownItem>Sign up</ElDropdownItem>
                <ElDropdownItem>Profile</ElDropdownItem>
              </ElDropdownMenu>
            </template>
          </ElDropdown>
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
