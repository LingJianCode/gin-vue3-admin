<template>
  <el-config-provider :locale="locale" :size="size">
    <!-- 开启水印 -->
    <!-- <el-watermark
      :font="{ color: fontColor }"
      :content="watermarkEnabled ? defaultSettings.watermarkContent : ''"
      :z-index="9999"
      class="wh-full"
    > -->

    <el-watermark
      :font="{ color: fontColor }"
      :content="watermarkEnabled ? watermarkContent : ''"
      :z-index="9999"
      class="wh-full"
    >
      <router-view />
    </el-watermark>
  </el-config-provider>
</template>

<script setup>
import { useAppStore, useSettingsStore, useUserStore } from "@/store";
import defaultSettings from "@/settings";
import { ThemeMode } from "@/enums/settings/theme.enum";



const appStore = useAppStore();
const settingsStore = useSettingsStore();
const userStore = useUserStore();

// 下面是箭头函数的简写，不用return的情况
const locale = computed(() => appStore.locale);
const size = computed(() => appStore.size);
const watermarkEnabled = computed(() => settingsStore.watermarkEnabled);
// 箭头函数使用return的情况
const watermarkContent = computed(() => {
  return userStore.userInfo.username ? userStore.userInfo.username : defaultSettings.watermarkContent
})

// 明亮/暗黑主题水印字体颜色适配
const fontColor = computed(() => {
  return settingsStore.theme === ThemeMode.DARK ? "rgba(255, 255, 255, .15)" : "rgba(0, 0, 0, .15)";
});
</script>
