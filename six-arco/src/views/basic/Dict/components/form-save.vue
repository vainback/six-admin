<template>
    <a-modal
            :title="props.isEdit ? '编辑' : '添加'"
            width="500px"
            v-model:visible="props.visible"
            @cancel="emit('close')"
            @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-form-item field="label" label="标题" validate-trigger="blur" :rules="[{required: true, message: '标题必须填写'}]">
                <a-input v-model.trim="form.label" placeholder="标题"></a-input>
            </a-form-item>
            <a-form-item field="value" label="值" validate-trigger="blur"  :rules="[{required: true, message: '字典值必须填写，且必须在同类型字典中唯一'}]">
                <a-input v-model.trim="form.value" placeholder="字典值，在同类型字典中唯一"></a-input>
            </a-form-item>
            <a-form-item v-if="form.type != 'root'" field="color" label="颜色">
                <a-color-picker v-model="form.color" showText disabledAlpha/>
            </a-form-item>
            <a-form-item field="is_sync" label="同步" v-if="isTenantRoot()">
                <a-space>
                    <a-switch
                            type="round"
                            v-model="form.is_sync"
                            checked-text=" 同 步 "
                            unchecked-text=" 不 同 步 "
                            :checked-value="1"
                            :unchecked-value="-1"
                    >
                    </a-switch>
                    <a-typography-text type="danger">选择同步将同步到全部租户的字典信息中</a-typography-text>
                </a-space>
            </a-form-item>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import {Dict, reqDict} from "@/api/basic/dict";
import {isTenantRoot} from "@/utils/auth";

const size = ref(import.meta.env.VITE_STYLE_SIZE);
const props = defineProps<{
    visible: boolean;
    isEdit: boolean;
    formData: Dict;
}>();

const form = ref<Dict>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
const action = ref<string>(props.isEdit ? 'edit' : 'add');
const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

const submit = async () => {
    try {
        await reqDict(action.value, { model: form.value });
        emit('closeAndRefresh');
    } catch (e) {
        console.log(e);
    }
};

watchEffect(() => {
    form.value = { ...props.formData };
    action.value = props.isEdit ? 'edit' : 'add';
});
</script>
<style scoped lang="less"></style>
