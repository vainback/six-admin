<template>
    <a-modal
        :title="props.isEdit ? '编辑' : '添加'"
        width="600px"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-row class="six-flex-center">
                <a-col :span="11">
                    <a-form-item label="账号" required>
                        <a-input
                            v-model="form.username"
                            placeholder="请输入账号"
                        ></a-input>
                    </a-form-item>
                </a-col>
                <a-col :span="1"></a-col>
                <a-col :span="11">
                    <a-form-item label="密码" :required="!isEdit">
                        <a-input
                            type="password"
                            v-model="form.password"
                            :placeholder="props.isEdit ? '输入后将进行修改，留空则不修改':'请输入密码'"
                        ></a-input>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row class="six-flex-center">
                <a-col :span="11">
                    <a-form-item label="昵称" required>
                        <a-input
                            v-model="form.nickname"
                            placeholder="请输入昵称 / 姓名"
                        ></a-input>
                    </a-form-item>
                </a-col>
                <a-col :span="1"></a-col>
                <a-col :span="11">
                    <a-form-item label="角色" required>
                        <a-tree-select
                            :max-tag-count="1"
                            tree-checkable
                            tree-check-strictly
                            v-model="form.role_ids"
                            :data="roles"
                            placeholder="请选择权限角色"
                            :fieldNames="{
                                key: 'id',
                                title: 'title',
                                children: 'children',
                            }"
                        ></a-tree-select>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row class="six-flex-center">
                <a-col :span="11">
                    <a-form-item label="部门" required>
                        <a-tree-select
                            v-model="form.organization_id"
                            :data="orgs"
                            placeholder="请选择所属部门"
                            :fieldNames="{
                                key: 'id',
                                title: 'name',
                                children: 'children',
                            }"
                        ></a-tree-select>
                    </a-form-item>
                </a-col>
                <a-col :span="1"></a-col>
                <a-col :span="11">
                    <a-form-item label="职位" required>
                        <a-tree-select
                            v-model="form.job_id"
                            :data="jobs"
                            placeholder="请选择职位"
                            :fieldNames="{
                                key: 'id',
                                title: 'name',
                                children: 'children',
                            }"
                        ></a-tree-select>
                    </a-form-item>
                </a-col>
            </a-row>
            <a-row class="six-flex-center">
                <a-col :span="11">
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
                </a-col>
                <a-col :span="1"></a-col>
                <a-col :span="11"></a-col>
            </a-row>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
    import { ref, watch, watchEffect } from 'vue';
    import { Message } from '@arco-design/web-vue';
    import { reqUser, Userinfo } from '@/api/system/user';
    import { reqAuthRole, TreeAuthRole } from '@/api/system/auth-role';
    import { reqOrgs } from '@/api/system/organization';
    import { reqJobs } from '@/api/system/job';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    const props = defineProps<{
        visible: boolean;
        isEdit: boolean;
        formData: Userinfo;
    }>();

    const form = ref<Userinfo>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
    const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

    const submit = async () => {
        try {
            await reqUser(props.isEdit ? 'edit' : 'add', { model: form.value });
            emit('closeAndRefresh');
            Message.success(props.isEdit ? '修改成功' : '添加成功');
        } catch (e) {}
    };

    const loading = ref(false);
    const roles = ref<TreeAuthRole[]>([]);
    const orgs = ref<TreeAuthRole[]>([]);
    const jobs = ref<TreeAuthRole[]>([]);
    const getData = async () => {
        loading.value = true;
        try {
            const { data } = await reqAuthRole('tree-select', null);
            if ((form.value.role_ids || []).includes(-20240929)) {
                roles.value = [{id: -20240929, title: 'ROOT', disabled: true},...(data || [])];
            } else {
                roles.value = data || []
            }
            const { data: orgRes } = await reqOrgs('tree-select', null);
            orgs.value = orgRes || [];
            const { data: jobRes } = await reqJobs('tree-select', null);
            jobs.value = jobRes || [];
        } finally {
            loading.value = false;
        }
    };
    const isEdit = ref(false);

    watch(
        () => props.visible,
        (newVal, oldVal) => {
            if (props.visible) {
                getData();
            }
        }
    );

    watchEffect(() => {
        form.value = { ...props.formData };
    });
</script>
<style scoped lang="less"></style>
