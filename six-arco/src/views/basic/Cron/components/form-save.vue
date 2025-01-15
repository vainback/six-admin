<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        width="500px"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-form-item label="标识" required>
                <a-input v-model="form.name" placeholder="请输入标识【服务端定时执行方法储存的Map.Key】"></a-input>
            </a-form-item>
            <a-form-item label="标题" required>
                <a-input v-model="form.title" placeholder="请输入标题"></a-input>
            </a-form-item>
            <a-form-item label="定时表达式" required tooltip="默认支持秒级定时，需要六位字段以空格分隔">
                <a-input v-model="form.times" placeholder="请输入定时表达式">
                    <template #append>
                        <a-link href="https://pkg.go.dev/github.com/robfig/cron">参考文档</a-link>
                    </template>
                </a-input>
            </a-form-item>
            <a-form-item label="状态">
                <a-switch
                    type="round"
                    v-model="form.status"
                    checked-text=" 正 常 "
                    unchecked-text=" 禁 用 "
                    :checked-value="1"
                    :unchecked-value="-1"
                >
                </a-switch>
            </a-form-item>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import { Message } from '@arco-design/web-vue';
import {reqCron, Cron} from "@/api/basic/cron";

const size = ref(import.meta.env.VITE_STYLE_SIZE);
const props = defineProps<{
    visible: boolean;
    isEdit: boolean;
    formData: Cron;
}>();

const form = ref<Cron>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

const submit = async () => {
    try {
        await reqCron(props.isEdit ? 'save' : 'add', { model: form.value });
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
