<script setup lang="ts">
import SplitLine from '@/components/SplitLine.vue';
import { ArrowDown, ArrowUp, Key, MessageBox, User } from '@element-plus/icons-vue';
import { ElButton, ElContainer, ElForm, ElFormItem, ElIcon, ElImage, ElInput, ElInputNumber, ElLink } from 'element-plus';
import { Picture as IconPicture } from '@element-plus/icons-vue'
import { reactive } from 'vue';
import type { CaptchaResult } from '@/models/captcha';
const signinInfo = reactive({
  username: '',
  password: '',
  captcha: {} as CaptchaResult<any>,
})

const captchaInfo = reactive({
  src: '',
  schduleID: -1
})
function getCaptcha() {
  fetch('/api/v1/captcha')
    .then(res => res.json())
    .then(data => {
      captchaInfo.src = data.data.base64;
      signinInfo.captcha.challengeID = data.data.id;
      if (captchaInfo.schduleID !== -1) {
        clearTimeout(captchaInfo.schduleID);
      }
      captchaInfo.schduleID = setTimeout(() => {
        captchaInfo.src = '';
      }, 120 * 1000);
    })
}
</script>

<template>
  <div class="box-container">
    <ElContainer class="box" style="width: 300px;">
      <h1 style="margin-left: 0.5rem;">Sign In</h1>
      <SplitLine />
      <ElForm @submit.prevent :model="signinInfo" class="input-field">
        <ElFormItem>
          <ElInput placeholder="Your Username / Email" :prefix-icon="User" v-model="signinInfo.username" clearable />
        </ElFormItem>
        <ElFormItem>
          <ElInput :show-password="true" type="password" placeholder="Your Password" :prefix-icon="Key"
            v-model="signinInfo.password" clearable />
        </ElFormItem>
        <ElFormItem id="captcha">
          <template #label>
            <ElImage @click="getCaptcha" fit="scale-down" :src="captchaInfo.src" style="height: 32px;width: 96px;">
              <template #error>
                <div class="image-slot">
                  Click To Refresh
                </div>
              </template>
            </ElImage>
          </template>
          <ElInputNumber :min="0" controls-position="right" v-model="signinInfo.captcha.result" placeholder="Calculate">
            <template #decrease-icon>
              <ElIcon>
                <ArrowDown />
              </ElIcon>
            </template>
            <template #increase-icon>
              <ElIcon>
                <ArrowUp />
              </ElIcon>
            </template>
          </ElInputNumber>
        </ElFormItem>
        <ElFormItem>
          <ElButton type="primary" style="width: 100%;">Sign in</ElButton>
        </ElFormItem>
        <ElFormItem style="margin: 0.25rem 0;">
          <ElLink style="font-size: 12px;height: 12px;" :underline="false" type="primary">Forgot Password?</ElLink>
        </ElFormItem>
        <SplitLine style="height: 1px;" />
        <ElFormItem>
          <ElButton plain style="width: 100%;">Sign up</ElButton>
        </ElFormItem>
      </ElForm>
    </ElContainer>
  </div>
</template>

<style lang="css" scoped>
html.dark section.box {
  box-shadow: 3px 3px 7px 0px rgba(255, 255, 255, 0.1);
}

html section.box {
  box-shadow: 5px 5px 7px 0px rgba(0, 0, 0, 0.1);
}

section.box {
  border: 2px solid var(--el-border-color);
  border-radius: 8px;
  padding: 0 0.5rem 0.5rem 0.5rem;
  margin-right: 0;
  float: right;
  display: block;
}

.input-field>* {
  margin: 0.5rem 0;
}

.input-field>#captcha {
  .el-image {
    border-radius: var(--el-border-radius-base);

    .image-slot {
      background-color: rgb(103, 184, 255);
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      background: var(--el-fill-color-light);
      color: var(--el-text-color-secondary);
      font-size: 12px;
    }
  }

  .el-input-number {
    width: 100%;
  }
}
</style>
