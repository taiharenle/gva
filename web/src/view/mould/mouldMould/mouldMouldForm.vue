<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="年份:" prop="mouldTime">
          <el-date-picker v-model="formData.mouldTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
       </el-form-item>
        <el-form-item label="门幅:" prop="size">
          <el-input v-model="formData.size" :clearable="true"  placeholder="请输入门幅" />
       </el-form-item>
        <el-form-item label="冲床(只/冲):" prop="machinesNumber">
          <el-input v-model.number="formData.machinesNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="个数(只/吸):" prop="paperNumber">
          <el-input v-model.number="formData.paperNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单价(元/只):" prop="unitPrice">
          <el-input v-model.number="formData.unitPrice" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数量(袋):" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="长度(毫米mm):" prop="length">
          <el-input v-model.number="formData.length" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="重量:" prop="weight">
          <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="颜色:" prop="color">
          <el-input v-model="formData.color" :clearable="true"  placeholder="请输入颜色" />
       </el-form-item>
        <el-form-item label="原料:" prop="material">
          <el-input v-model="formData.material" :clearable="true"  placeholder="请输入原料" />
       </el-form-item>
        <el-form-item label="图片:" prop="mouldPicture">
           <SelectImage v-model="formData.mouldPicture" multiple file-type="image"/>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMouldMould,
  updateMouldMould,
  findMouldMould
} from '@/api/mould/mouldMould'

defineOptions({
    name: 'MouldMouldForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unitPrice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMouldMould({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remouldMould
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
