<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        width="600px"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-row>
                <a-col :span="11">
                    <a-form-item field="type" label="菜单类型">
                        <a-radio-group type="button" v-model.trim="form.type">
                            <a-radio value="menu">菜单</a-radio>
                            <a-radio value="page">页面</a-radio>
                            <a-radio value="button">按钮</a-radio>
                        </a-radio-group>
                    </a-form-item>
                </a-col>
                <a-col :span="2"></a-col>
                <a-col :span="11">
                    <a-form-item field="type" label="ID">
                        <a-input-number :disabled="props.isEdit" v-model="form.id" placeholder="默认为自增id，可自定义"></a-input-number>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row>
                <a-col :span="11">
                    <a-form-item field="locale" label="菜单名称">
                        <a-input
                            v-model.trim="form.locale"
                            placeholder="请输入菜单名称"
                        ></a-input>
                    </a-form-item>
                </a-col>
                <a-col :span="2"></a-col>
                <a-col :span="11">
                    <a-form-item field="icon" label="菜单图标">
                        <a-input
                                v-model.trim="form.icon"
                                placeholder="请输入菜单图标"
                        ></a-input>
                    </a-form-item>
                </a-col>
                
            </a-row>
            <a-row>
                <a-col :span="11">
                    <a-form-item field="path" label="Name">
                        <a-input
                                v-model.trim="form.name"
                                placeholder="请输入前端路由name"
                        ></a-input>
                    </a-form-item>
                </a-col>
                <a-col :span="2"></a-col>
                <a-col :span="11">
                    <a-form-item field="name" label="Path">
                        <a-input
                                v-model.trim="form.path"
                                placeholder="请输入前端路由path"
                        ></a-input>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row>
                <a-col :span="11">
                    <a-form-item field="api" label="权限标识">
                        <a-input
                            v-model.trim="form.auth"
                            placeholder="请输入权限标识"
                        >
                            <template #prepend v-if="form.authPrefix">
                                {{form.authPrefix}}
                            </template>
                        </a-input>
                    </a-form-item>
                </a-col>
                <a-col :span="2"></a-col>
                <a-col :span="11">
                    <a-form-item field="api_route" label="后端路由">
                        <a-input
                                v-model.trim="form.api_route"
                                placeholder="请输入后端路由路径"
                        >
                        </a-input>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row>
                <a-col :span="11">
                    <a-form-item field="api" label="排序">
                        <a-input-number
                            v-model.number="form.order"
                            placeholder="请输入排序序号"
                        ></a-input-number>
                    </a-form-item>
                </a-col>
                <a-col :span="2"></a-col>
                <a-col :span="11">
                    <a-form-item field="icon" label="状态">
                        <a-switch
                                type="round"
                                v-model="form.status"
                                checked-text=" 展 示 "
                                unchecked-text=" 隐 藏 "
                                :checked-value="1"
                                :unchecked-value="-1"
                        >
                        </a-switch>
                    </a-form-item>
                </a-col>
            </a-row>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
    import { AuthRule, reqAuthRule } from '@/api/system/auth-rule';
    import { nextTick, ref, watchEffect } from 'vue';
    import { cloneDeep } from 'lodash';
    import { Message } from '@arco-design/web-vue';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    const props = defineProps<{
        visible: boolean;
        isEdit: boolean;
        formData: AuthRule;
    }>();

    const form = ref<AuthRule>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
    // const action = ref<string>(props.isEdit ? 'edit' : 'add');
    const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

    const submit = async() => {
        try {
            let values = {...form.value}
            values.auth = form.value.authPrefix + form.value.auth;
            await reqAuthRule(props.isEdit ? 'save' : 'add', { model: values });
            emit('closeAndRefresh');
            Message.success(props.isEdit ? '修改成功' : '添加成功');
        } catch (e) {}
    };

    watchEffect(() => {
        form.value = { ...props.formData };
    });
</script>
<style scoped lang="less"></style>
