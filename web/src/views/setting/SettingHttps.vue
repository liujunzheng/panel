<script setup lang="ts">
import { useI18n } from 'vue-i18n'

import setting from '@/api/panel/setting'

const { t } = useI18n()

const { data: model } = useRequest(setting.list, {
  initialData: {
    name: '',
    locale: '',
    username: '',
    password: '',
    email: '',
    port: 8888,
    entrance: '',
    offline_mode: false,
    website_path: '',
    backup_path: '',
    https: false,
    cert: '',
    key: ''
  }
})

const handleSave = () => {
  useRequest(setting.update(model.value)).onSuccess(() => {
    window.$message.success(t('settingIndex.edit.toasts.success'))
  })
}
</script>

<template>
  <n-space vertical>
    <n-alert type="warning"> 错误的证书可能导致面板无法访问，请谨慎操作！</n-alert>
    <n-form>
      <n-form-item :label="$t('settingIndex.edit.fields.https.label')">
        <n-switch v-model:value="model.https" />
      </n-form-item>
      <n-form-item v-if="model.https" :label="$t('settingIndex.edit.fields.cert.label')">
        <n-input
          v-model:value="model.cert"
          type="textarea"
          :autosize="{ minRows: 10, maxRows: 15 }"
        />
      </n-form-item>
      <n-form-item v-if="model.https" :label="$t('settingIndex.edit.fields.key.label')">
        <n-input
          v-model:value="model.key"
          type="textarea"
          :autosize="{ minRows: 10, maxRows: 15 }"
        />
      </n-form-item>
    </n-form>
  </n-space>
  <n-button type="primary" @click="handleSave">
    {{ $t('settingIndex.edit.actions.submit') }}
  </n-button>
</template>

<style scoped lang="scss"></style>
