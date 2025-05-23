<!-- 字典项 -->
<template>
  <div class="app-container">
    <div class="search-bar mt-5">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="关键字" prop="keywords">
          <el-input
            v-model="queryParams.keywords"
            placeholder="字典标签/字典值"
            clearable
            @keyup.enter="handleQuery"
          />
        </el-form-item>
        <el-form-item >
          <el-button v-hasPerm="['sys:dict-item:query']" type="primary" icon="search" @click="handleQuery()">搜索</el-button>
          <el-button icon="refresh" @click="handleResetQuery()">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <div class="mb-[10px]">
        <el-button v-hasPerm="['sys:dict-item:add']" type="success" icon="plus" @click="handleOpenDialog()">新增</el-button>
        <!-- <el-button v-hasPerm="['sys:dict-item:delete']" type="danger" :disabled="ids.length === 0" icon="delete" v-hasPerm="['sys:dict:delete']" @click="handleDelete()">
          删除
        </el-button> -->
      </div>

      <el-table
        v-loading="loading"
        highlight-current-row
        :data="tableData"
        border
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="字典项标签" prop="label" />
        <el-table-column label="字典项值" prop="value" />
        <el-table-column label="排序" prop="sort" />
        <el-table-column label="状态">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? "启用" : "禁用" }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column fixed="right" label="操作" align="center" width="220">
          <template #default="scope">
            <el-button
              v-hasPerm="['sys:dict-item:edit']"
              type="primary"
              link
              size="small"
              icon="edit"
              @click.stop="handleOpenDialog(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              v-hasPerm="['sys:dict-item:delete']"
              type="danger"
              link
              size="small"
              icon="delete"
              @click.stop="handleDelete(scope.row.id)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination
        v-if="total > 0"
        v-model:total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="handleQuery"
      />
    </el-card>

    <!--字典项弹窗-->
    <el-dialog
      v-model="dialog.visible"
      :title="dialog.title"
      width="820px"
      @close="handleCloseDialog"
    >
      <el-form ref="dataFormRef" :model="formData" :rules="computedRules" label-width="100px">
        <el-card shadow="never">
          <el-form-item label="字典项标签" prop="label">
            <el-input v-model="formData.label" placeholder="请输入字典标签" />
          </el-form-item>
          <el-form-item label="字典项值" prop="value">
            <el-input v-model="formData.value" placeholder="请输入字典值" />
          </el-form-item>
          <el-form-item label="状态">
            <el-radio-group v-model="formData.status">
              <el-radio :value="1">启用</el-radio>
              <el-radio :value="0">禁用</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="排序">
            <el-input-number v-model="formData.sort" controls-position="right" />
          </el-form-item>
          <el-form-item label="标签类型">
            <el-tag v-if="formData.tagType" :type="formData.tagType" class="mr-2">
              {{ formData.label }}
            </el-tag>
            <el-radio-group v-model="formData.tagType">
              <el-radio value="success" border size="small">success</el-radio>
              <el-radio value="warning" border size="small">warning</el-radio>
              <el-radio value="info" border size="small">info</el-radio>
              <el-radio value="primary" border size="small">primary</el-radio>
              <el-radio value="danger" border size="small">danger</el-radio>
              <el-radio value="" border size="small">清空</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-card>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmitClick">确 定</el-button>
          <el-button @click="handleCloseDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import DictAPI from "@/api/system/dict.api";

const route = useRoute();

const dictCode = ref(route.query.dictCode);

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const ids = ref([]);
const total = ref(0);

const queryParams = reactive({
  pageNum: 1,
  pageSize: 10,
});

const tableData = ref();

const dialog = reactive({
  title: "",
  visible: false,
});

const formData = reactive({});

const computedRules = computed(() => {
  const rules = {
    value: [{ required: true, message: "请输入字典值", trigger: "blur" }],
    label: [{ required: true, message: "请输入字典标签", trigger: "blur" }],
  };
  return rules;
});

// 查询
function handleQuery() {
  loading.value = true;
  DictAPI.getDictItemPage(dictCode.value, queryParams)
    .then((data) => {
      tableData.value = data.list;
      total.value = data.total;
    })
    .finally(() => {
      loading.value = false;
    });
}

// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  handleQuery();
}

// 行选择
function handleSelectionChange(selection) {
  ids.value = selection.map((item) => item.id);
}

// 打开弹窗
function handleOpenDialog(row) {
  dialog.visible = true;
  dialog.title = row ? "编辑字典项" : "新增字典项";

  if (row?.id) {
    DictAPI.getDictItemFormData(dictCode.value, row.id).then((data) => {
      Object.assign(formData, data);
    });
  } else {
    formData.id = undefined;
    formData.status = 1;
    formData.sort = 0;
  }
}

// 提交表单
function handleSubmitClick() {
  dataFormRef.value.validate((valid) => {
    if (valid) {
      const id = formData.id;
      if (id) {
        DictAPI.updateDictItem(dictCode.value, id, formData).then(() => {
          ElMessage.success("修改成功");
          handleCloseDialog();
          handleResetQuery();
        });
      } else {
        DictAPI.createDictItem(dictCode.value, formData).then(() => {
          ElMessage.success("新增成功");
          handleCloseDialog();
          handleResetQuery();
        });
      }
    }
  });
}

// 关闭弹窗
function handleCloseDialog() {
  dialog.visible = false;
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();
  formData.id = undefined;
}

// 删除字典项
function handleDelete(id) {
  const itemIds = [id || ids.value].join(",");
  if (!itemIds) {
    ElMessage.warning("请勾选删除项");
    return;
  }
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    () => {
      DictAPI.deleteDictItems(dictCode.value, itemIds).then(() => {
        ElMessage.success("删除成功");
        handleResetQuery();
      });
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
}

onMounted(() => {
  handleQuery();
});
</script>
