<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        width="500px"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-form-item label="名称" required>
                <a-input v-model="form.name" placeholder="请输入名称"></a-input>
            </a-form-item>
            <a-form-item label="介绍">
                <a-input v-model="form.description" placeholder="请输入简介"></a-input>
            </a-form-item>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import { Message } from '@arco-design/web-vue';
import {Job, reqJobs} from "@/api/system/job";

const size = ref(import.meta.env.VITE_STYLE_SIZE);
const props = defineProps<{
    visible: boolean;
    isEdit: boolean;
    formData: Job;
}>();

const form = ref<Job>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
// const action = ref<string>(props.isEdit ? 'edit' : 'add');
const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

const submit = async () => {
    try {
        await reqJobs(props.isEdit ? 'save' : 'add', { model: form.value });
        emit('closeAndRefresh');
        Message.success(props.isEdit ? '修改成功' : '添加成功');
    } catch (e) {
    }
};

watchEffect(() => {
    form.value = { ...props.formData };
});
</script>
<style scoped lang="less"></style>
