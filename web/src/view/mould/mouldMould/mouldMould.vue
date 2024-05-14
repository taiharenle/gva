<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
         <el-table-column align="left" label="年份" width="180">
            <template #default="scope">{{ formatDate(scope.row.mouldTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="门幅" prop="size" width="120" />
        <el-table-column align="left" label="冲床(只/冲)" prop="machinesNumber" width="120" />
        <el-table-column align="left" label="个数(只/吸)" prop="paperNumber" width="120" />
        <el-table-column align="left" label="单价(元/只)" prop="unitPrice" width="120" />
        <el-table-column align="left" label="数量(袋)" prop="quantity" width="120" />
        <el-table-column align="left" label="长度(毫米mm)" prop="length" width="120" />
        <el-table-column align="left" label="重量" prop="weight" width="120" />
        <el-table-column align="left" label="颜色" prop="color" width="120" />
        <el-table-column align="left" label="原料" prop="material" width="120" />
           <el-table-column label="图片" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.mouldPicture" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMouldMouldFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="年份:"  prop="mouldTime" >
              <el-date-picker v-model="formData.mouldTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="门幅:"  prop="size" >
              <el-input v-model="formData.size" :clearable="true"  placeholder="请输入门幅" />
            </el-form-item>
            <el-form-item label="冲床(只/冲):"  prop="machinesNumber" >
              <el-input v-model.number="formData.machinesNumber" :clearable="true" placeholder="请输入冲床(只/冲)" />
            </el-form-item>
            <el-form-item label="个数(只/吸):"  prop="paperNumber" >
              <el-input v-model.number="formData.paperNumber" :clearable="true" placeholder="请输入个数(只/吸)" />
            </el-form-item>
            <el-form-item label="单价(元/只):"  prop="unitPrice" >
              <el-input v-model.number="formData.unitPrice" :clearable="true" placeholder="请输入单价(元/只)" />
            </el-form-item>
            <el-form-item label="数量(袋):"  prop="quantity" >
              <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入数量(袋)" />
            </el-form-item>
            <el-form-item label="长度(毫米mm):"  prop="length" >
              <el-input v-model.number="formData.length" :clearable="true" placeholder="请输入长度(毫米mm)" />
            </el-form-item>
            <el-form-item label="重量:"  prop="weight" >
              <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入重量" />
            </el-form-item>
            <el-form-item label="颜色:"  prop="color" >
              <el-input v-model="formData.color" :clearable="true"  placeholder="请输入颜色" />
            </el-form-item>
            <el-form-item label="原料:"  prop="material" >
              <el-input v-model="formData.material" :clearable="true"  placeholder="请输入原料" />
            </el-form-item>
            <el-form-item label="图片:"  prop="mouldPicture" >
                <SelectImage
                 multiple
                 v-model="formData.mouldPicture"
                 file-type="image"
                 />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createMouldMould,
  deleteMouldMould,
  deleteMouldMouldByIds,
  updateMouldMould,
  findMouldMould,
  getMouldMouldList
} from '@/api/mould/mouldMould'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MouldMould'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        mouldTime: new Date(),
        name: '',
        size: '',
        machinesNumber: 0,
        paperNumber: 0,
        unitPrice: 0,
        quantity: 0,
        length: 0,
        weight: 0,
        color: '',
        material: '',
        mouldPicture: [],
        })



// 验证规则
const rule = reactive({
               mouldTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               unitPrice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMouldMouldList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMouldMouldFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMouldMouldByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMouldMouldFunc = async(row) => {
    const res = await findMouldMould({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remouldMould
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMouldMouldFunc = async (row) => {
    const res = await deleteMouldMould({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        mouldTime: new Date(),
        name: '',
        size: '',
        machinesNumber: 0,
        paperNumber: 0,
        unitPrice: 0,
        quantity: 0,
        length: 0,
        weight: 0,
        color: '',
        material: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createMouldMould(formData.value)
                  break
                case 'update':
                  res = await updateMouldMould(formData.value)
                  break
                default:
                  res = await createMouldMould(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
