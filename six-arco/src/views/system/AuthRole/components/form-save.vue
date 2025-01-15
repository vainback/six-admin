<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        width="500px"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-form-item field="title" label="角色标题" validate-trigger="blur" :rules="[{required: true, message: '角色名必须填写'}]">
                <a-input v-model.trim="form.title" placeholder="角色标题"></a-input>
            </a-form-item>
            <a-form-item field="sign" label="角色签名" validate-trigger="blur"  :rules="[{required: true, message: '角色值必须填写，且必须在系统中唯一'}]">
                <a-input v-model.trim="form.sign" placeholder="角色签名，在系统中必须唯一"></a-input>
            </a-form-item>
            <a-form-item field="icon" label="状态">
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
    import { AuthRole, reqAuthRole } from '@/api/system/auth-role';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    const props = defineProps<{
        visible: boolean;
        isEdit: boolean;
        formData: AuthRole;
    }>();

    const form = ref<AuthRole>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
    const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

    const submit = async () => {
        try {
            await reqAuthRole(props.isEdit ? 'save' : 'add', { model: form.value });
            emit('closeAndRefresh');
        } catch (e) {
            console.log(e);
        }
    };

    watchEffect(() => {
        form.value = { ...props.formData };
    });
</script>
<style scoped lang="less"></style>
