<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
    <div style="max-height: calc(100vh - 300px)" >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <Tpl-From-Items>
        </a-form>
    </div>
    </a-modal>
</template>
<script setup lang="ts">
import {onBeforeUnmount, ref, shallowRef, watchEffect} from 'vue';
import {Message, Modal} from '@arco-design/web-vue';
import {req<Tpl-Module>, <Tpl-Module>} from "@/api/codegen/<Tpl-Module-Lower>";
import { uploadHeaders, uploadUrl } from '@/api/user-single';
<Tpl - Editor - Import>

const urlPrefix = ref(import.meta.env.VITE_API_BASE_URL);
const size = ref(import.meta.env.VITE_STYLE_SIZE);
const props = defineProps<{
    visible: boolean;
    isEdit: boolean;
    <Tpl - Dict>
    formData: <Tpl-Module>;
}>();

const form = ref<<Tpl-Module>>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

const submit = async () => {
    try {
        await req<Tpl-Module>(props.isEdit ? 'save' : 'add', { model: form.value });
        emit('closeAndRefresh');
        Message.success(props.isEdit ? '修改成功' : '添加成功');
    } catch (e) {
    }
};

watchEffect(() => {
    form.value = { ...props.formData };
});

<Tpl - ImageList - Success - Error>



</script>
<style scoped lang="less"></style>
