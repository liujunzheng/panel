<script setup lang="ts">
import Editor from '@guolao/vue-monaco-editor'
import type { MessageReactive } from 'naive-ui'
import { NButton, NDataTable, NFlex, NPopconfirm, NSpace, NSwitch, NTag } from 'naive-ui'

import cert from '@/api/panel/cert'
import { formatDateTime } from '@/utils'
import ObtainModal from '@/views/cert/ObtainModal.vue'

const props = defineProps({
  algorithms: {
    type: Array<any>,
    required: true
  },
  websites: {
    type: Array<any>,
    required: true
  },
  accounts: {
    type: Array<any>,
    required: true
  },
  dns: {
    type: Array<any>,
    required: true
  }
})

const { algorithms, websites, accounts, dns } = toRefs(props)

let messageReactive: MessageReactive | null = null

const updateModel = ref<any>({
  domains: [],
  type: 'P256',
  dns_id: null,
  account_id: null,
  website_id: null,
  auto_renew: true,
  cert: '',
  key: '',
  script: ''
})
const updateModal = ref(false)
const updateCert = ref<any>()
const showModal = ref(false)
const showModel = ref<any>({
  cert: '',
  key: ''
})
const deployModal = ref(false)
const deployModel = ref<any>({
  id: null,
  websites: []
})
const obtain = ref(false)
const obtainCert = ref(0)

const columns: any = [
  {
    title: '域名',
    key: 'domains',
    minWidth: 200,
    resizable: true,
    render(row: any) {
      if (row.domains == null || row.domains.length == 0) {
        return h(NTag, null, { default: () => '无' })
      }
      return h(NFlex, null, {
        default: () =>
          row.domains.map((domain: any) =>
            h(
              NTag,
              { type: 'primary' },
              {
                default: () => domain
              }
            )
          )
      })
    }
  },
  {
    title: '类型',
    key: 'type',
    width: 100,
    render(row: any) {
      return h(
        NTag,
        {
          type: 'info',
          bordered: false
        },
        {
          default: () => {
            switch (row.type) {
              case 'P256':
                return 'EC 256'
              case 'P384':
                return 'EC 384'
              case '2048':
                return 'RSA 2048'
              case '4096':
                return 'RSA 4096'
              default:
                return '上传'
            }
          }
        }
      )
    }
  },
  {
    title: '关联账号',
    key: 'account_id',
    minWidth: 200,
    resizable: true,
    ellipsis: { tooltip: true },
    render(row: any) {
      if (row.account_id == 0) {
        return '无'
      }
      return accounts.value?.find((item: any) => item.value === row.account_id)?.label
    }
  },
  {
    title: '颁发者',
    key: 'issuer',
    width: 150,
    ellipsis: { tooltip: true },
    render(row: any) {
      return row.issuer == '' ? '无' : row.issuer
    }
  },
  {
    title: '过期时间',
    key: 'not_after',
    width: 200,
    ellipsis: { tooltip: true },
    render(row: any) {
      return formatDateTime(row.not_after)
    }
  },
  {
    title: 'OCSP',
    key: 'ocsp_server',
    minWidth: 200,
    resizable: true,
    render(row: any) {
      if (row.ocsp_server == null || row.ocsp_server.length == 0) {
        return h(NTag, null, { default: () => '无' })
      }
      return h(NFlex, null, {
        default: () =>
          row.ocsp_server.map((server: any) =>
            h(NTag, null, {
              default: () => server
            })
          )
      })
    }
  },
  {
    title: '自动续签',
    key: 'auto_renew',
    width: 100,
    align: 'center',
    resizable: true,
    render(row: any) {
      return h(NSwitch, {
        size: 'small',
        rubberBand: false,
        value: row.auto_renew,
        onUpdateValue: () => handleAutoRenewUpdate(row)
      })
    }
  },
  {
    title: '操作',
    key: 'actions',
    width: 350,
    align: 'center',
    hideInExcel: true,
    render(row: any) {
      return [
        row.type != 'upload' && row.cert == '' && row.key == ''
          ? h(
              NButton,
              {
                size: 'small',
                type: 'info',
                style: 'margin-left: 15px;',
                onClick: async () => {
                  obtain.value = true
                  obtainCert.value = row.id
                }
              },
              {
                default: () => '签发'
              }
            )
          : null,
        row.cert != '' && row.key != ''
          ? h(
              NButton,
              {
                size: 'small',
                type: 'info',
                onClick: () => {
                  deployModel.value.id = row.id
                  if (row.website_id != 0) {
                    deployModel.value.websites.push(row.website_id)
                  }
                  deployModal.value = true
                }
              },
              {
                default: () => '部署'
              }
            )
          : null,
        row.cert_url != '' && row.type != 'upload'
          ? h(
              NButton,
              {
                size: 'small',
                type: 'success',
                style: 'margin-left: 15px;',
                onClick: async () => {
                  messageReactive = window.$message.loading('请稍后...', {
                    duration: 0
                  })
                  useRequest(cert.renew(row.id))
                    .onSuccess(() => {
                      refresh()
                      window.$message.success('续签成功')
                    })
                    .onComplete(() => {
                      messageReactive?.destroy()
                    })
                }
              },
              {
                default: () => '续签'
              }
            )
          : null,
        row.cert != '' && row.key != ''
          ? h(
              NButton,
              {
                size: 'small',
                type: 'tertiary',
                style: 'margin-left: 15px;',
                onClick: () => {
                  showModel.value.cert = row.cert
                  showModel.value.key = row.key
                  showModal.value = true
                }
              },
              {
                default: () => '查看'
              }
            )
          : null,
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 15px;',
            onClick: () => {
              updateCert.value = row.id
              updateModel.value.domains = row.domains
              updateModel.value.type = row.type
              updateModel.value.dns_id = row.dns_id == 0 ? null : row.dns_id
              updateModel.value.account_id = row.account_id == 0 ? null : row.account_id
              updateModel.value.website_id = row.website_id == 0 ? null : row.website_id
              updateModel.value.auto_renew = row.auto_renew
              updateModel.value.cert = row.cert
              updateModel.value.key = row.key
              updateModel.value.script = row.script
              updateModal.value = true
            }
          },
          {
            default: () => '修改'
          }
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: async () => {
              useRequest(cert.certDelete(row.id)).onSuccess(() => {
                refresh()
                window.$message.success('删除成功')
              })
            }
          },
          {
            default: () => {
              return '确定删除证书吗？'
            },
            trigger: () => {
              return h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                  style: 'margin-left: 15px;'
                },
                {
                  default: () => '删除'
                }
              )
            }
          }
        )
      ]
    }
  }
]

const { loading, data, page, total, pageSize, pageCount, refresh } = usePagination(
  (page, pageSize) => cert.certs(page, pageSize),
  {
    initialData: { total: 0, list: [] },
    initialPageSize: 20,
    total: (res: any) => res.total,
    data: (res: any) => res.items
  }
)

const handleUpdateCert = () => {
  useRequest(cert.certUpdate(updateCert.value, updateModel.value)).onSuccess(() => {
    refresh()
    updateModal.value = false
    updateModel.value.domains = []
    updateModel.value.type = 'P256'
    updateModel.value.dns_id = null
    updateModel.value.account_id = null
    updateModel.value.website_id = null
    updateModel.value.auto_renew = true
    updateModel.value.cert = ''
    updateModel.value.key = ''
    updateModel.value.script = ''
    window.$message.success('更新成功')
  })
}

const handleAutoRenewUpdate = (row: any) => {
  updateModel.value.domains = row.domains
  updateModel.value.type = row.type
  updateModel.value.dns_id = row.dns_id == 0 ? null : row.dns_id
  updateModel.value.account_id = row.account_id == 0 ? null : row.account_id
  updateModel.value.website_id = row.website_id == 0 ? null : row.website_id
  updateModel.value.auto_renew = !row.auto_renew
  updateModel.value.cert = row.cert
  updateModel.value.key = row.key
  updateModel.value.script = row.script
  useRequest(cert.certUpdate(row.id, updateModel.value))
    .onSuccess(() => {
      refresh()
      window.$message.success('更新成功')
    })
    .onComplete(() => {
      updateModel.value.domains = []
      updateModel.value.type = 'P256'
      updateModel.value.dns_id = null
      updateModel.value.account_id = null
      updateModel.value.website_id = null
      updateModel.value.auto_renew = true
      updateModel.value.cert = ''
      updateModel.value.key = ''
      updateModel.value.script = ''
    })
}

const handleDeployCert = async () => {
  const promises = deployModel.value.websites.map((website: any) =>
    cert.deploy(deployModel.value.id, website)
  )
  await Promise.all(promises)

  deployModal.value = false
  deployModel.value.id = null
  deployModel.value.websites = []
  window.$message.success('部署成功')
}

const handleShowModalClose = () => {
  showModel.value.cert = ''
  showModel.value.key = ''
}

onMounted(() => {
  refresh()
  window.$bus.on('cert:refresh-cert', () => {
    refresh()
  })
})

onUnmounted(() => {
  window.$bus.off('cert:refresh-cert')
})
</script>

<template>
  <n-space vertical size="large">
    <n-data-table
      striped
      remote
      :scroll-x="1600"
      :loading="loading"
      :columns="columns"
      :data="data"
      :row-key="(row: any) => row.id"
      v-model:page="page"
      v-model:pageSize="pageSize"
      :pagination="{
        page: page,
        pageCount: pageCount,
        pageSize: pageSize,
        itemCount: total,
        showQuickJumper: true,
        showSizePicker: true,
        pageSizes: [20, 50, 100, 200]
      }"
    />
  </n-space>
  <n-modal
    v-model:show="updateModal"
    preset="card"
    title="修改证书"
    style="width: 60vw"
    size="huge"
    :bordered="false"
    :segmented="false"
  >
    <n-space vertical>
      <n-alert v-if="updateModel.type != 'upload'" type="info">
        可以通过选择网站 / DNS 中的任意一项来自动签发和部署证书，也可以手动输入域名并设置 DNS
        解析来签发证书，还可以填写部署脚本来自动部署证书。
      </n-alert>
      <n-form :model="updateModel">
        <n-form-item v-if="updateModel.type != 'upload'" path="domains" label="域名">
          <n-dynamic-input
            v-model:value="updateModel.domains"
            placeholder="example.com"
            :min="1"
            show-sort-button
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type != 'upload'" path="type" label="密钥类型">
          <n-select
            v-model:value="updateModel.type"
            placeholder="选择密钥类型"
            clearable
            :options="algorithms"
          />
        </n-form-item>
        <n-form-item path="website_id" label="网站">
          <n-select
            v-model:value="updateModel.website_id"
            placeholder="选择用于部署证书的网站"
            clearable
            :options="websites"
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type != 'upload'" path="account_id" label="账号">
          <n-select
            v-model:value="updateModel.account_id"
            placeholder="选择用于签发证书的账号"
            clearable
            :options="accounts"
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type != 'upload'" path="account_id" label="DNS">
          <n-select
            v-model:value="updateModel.dns_id"
            placeholder="选择用于签发证书的DNS"
            clearable
            :options="dns"
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type == 'upload'" path="cert" label="证书">
          <n-input
            v-model:value="updateModel.cert"
            type="textarea"
            placeholder="输入 PEM 证书文件的内容"
            :autosize="{ minRows: 10, maxRows: 15 }"
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type == 'upload'" path="key" label="私钥">
          <n-input
            v-model:value="updateModel.key"
            type="textarea"
            placeholder="输入 KEY 私钥文件的内容"
            :autosize="{ minRows: 10, maxRows: 15 }"
          />
        </n-form-item>
        <n-form-item v-if="updateModel.type != 'upload'" path="key" label="部署脚本">
          <n-input
            v-model:value="updateModel.script"
            type="textarea"
            placeholder="脚本中的 {cert} 和 {key} 会被替换为证书和私钥内容"
            :autosize="{ minRows: 5, maxRows: 10 }"
          />
        </n-form-item>
      </n-form>
      <n-button type="info" block @click="handleUpdateCert">提交</n-button>
    </n-space>
  </n-modal>
  <n-modal
    v-model:show="deployModal"
    preset="card"
    title="部署证书"
    style="width: 60vw"
    size="huge"
    :bordered="false"
    :segmented="false"
  >
    <n-space vertical>
      <n-form :model="deployModel">
        <n-form-item path="website_id" label="网站">
          <n-select
            v-model:value="deployModel.websites"
            placeholder="选择需要部署证书的网站"
            clearable
            multiple
            :options="websites"
          />
        </n-form-item>
      </n-form>
      <n-button type="info" block @click="handleDeployCert">提交</n-button>
    </n-space>
  </n-modal>
  <n-modal
    v-model:show="showModal"
    preset="card"
    title="查看证书"
    style="width: 80vw"
    size="huge"
    :bordered="false"
    :segmented="false"
    @close="handleShowModalClose"
  >
    <n-tabs type="line" animated>
      <n-tab-pane name="cert" tab="证书">
        <Editor
          v-model:value="showModel.cert"
          theme="vs-dark"
          height="60vh"
          mt-8
          :options="{
            readOnly: true,
            automaticLayout: true
          }"
        />
      </n-tab-pane>
      <n-tab-pane name="key" tab="密钥">
        <Editor
          v-model:value="showModel.key"
          theme="vs-dark"
          height="60vh"
          mt-8
          :options="{
            readOnly: true,
            automaticLayout: true
          }"
        />
      </n-tab-pane>
    </n-tabs>
  </n-modal>
  <obtain-modal v-model:id="obtainCert" v-model:show="obtain" />
</template>

<style scoped lang="scss"></style>
