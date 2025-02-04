<script setup lang="ts">
import SplitLine from '@/components/SplitLine.vue';
import { Key, MessageBox, User } from '@element-plus/icons-vue';
import { ElContainer, ElInput } from 'element-plus';

</script>

<template>
  <div class="box-container">
    <ElContainer class="box" style="width: 300px;">
      <h1 style="margin-left: 0.5rem;">Sign In</h1>
      <SplitLine />
      <div class="input-field">
        <ElInput placeholder="xxx" :prefix-icon="User"/>
        <ElInput placeholder="xxx@xxx.xxx" :prefix-icon="MessageBox"/>
        <ElInput :show-password="true" type="password" placeholder="Please input password" :prefix-icon="Key"/>
        
      </div>
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
  div.input-field > * {
    margin: 0.15rem 0;
  }
</style>
