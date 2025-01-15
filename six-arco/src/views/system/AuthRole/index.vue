<template>
    <div class="six-container">
        <a-card style="width: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd(0)"
                        v-permission="'system:role:add'"
                    >
                        <template #icon>
                            <icon-plus />
                        </template>
                        <template #default>添加</template>
                    </a-button>
                    <a-input
                        v-model="query.keyword"
                        placeholder="输入关键词搜索"
                        :size="size"
                        @input="queryList"
                    ></a-input>
                    <a-range-picker
                        v-model="timePicker"
                        show-time
                        :time-picker-props="{
                            defaultValue: ['00:00:00', '00:00:00'],
                        }"
                        :size="size"
                        @ok="queryList"
                        @clear="queryList"
                    ></a-range-picker>
                </a-space>
                <a-space>
                    <template #split>
                        <a-divider direction="vertical" />
                    </template>
                    <a-tooltip content="刷新" :mini="size === 'mini'">
                        <div style="cursor: pointer" @click="refresh">
                            <icon-refresh />
                        </div>
                    </a-tooltip>
                </a-space>
            </div>
            <div style="height: 10px"></div>
            <a-table
                :loading="queryLoading"
                :scroll="tableScroll"
                :style="{ height: tableScroll.maxHeight + 'px' }"
                row-key="id"
                v-model:expandedKeys="expandedKeys"
                :pagination="false"
                :data="list"
            >
                <template #columns>
                    <a-table-column
                        title="角色标题"
                        data-index="title"
                    ></a-table-column>
                    <a-table-column
                        title="角色签名"
                        data-index="sign"
                    ></a-table-column>
                    <a-table-column title="状态" data-index="status">
                        <template #cell="{ record }">
                            <a-tag v-if="record.status" color="blue"
                                >展示
                            </a-tag>
                            <a-tag v-else>隐藏</a-tag>
                        </template>
                    </a-table-column>
                    <a-table-column title="操作">
                        <template #cell="{ record }">
                            <a-button
                                type="text"
                                status="normal"
                                :size="size"
                                @click="toAdd(record.id)"
                                v-permission="'system:role:add'"
                                >添加
                            </a-button>
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'system:role:save'"
                                >编辑
                            </a-button>
                            <a-button
                                type="text"
                                status="success"
                                :size="size"
                                @click="toSetAuth(record.id)"
                                v-permission="'system:role:select-rule'"
                                >设置权限
                            </a-button>
                            <a-popconfirm
                                content="确认删除？"
                                okText="删除"
                                cancelText="取消"
                                type="error"
                                :ok-button-props="{ status: 'danger' }"
                                @ok="toDel(record.id)"
                            >
                                <a-button
                                    :disabled="
                                        record.children &&
                                        record.children.length > 0
                                    "
                                    type="text"
                                    status="danger"
                                    :size="size"
                                    v-permission="'system:role:del'"
                                    >删除
                                </a-button>
                            </a-popconfirm>
                        </template>
                    </a-table-column>
                </template>
            </a-table>
        </a-card>
        <form-save
            :visible="formModal.show"
            :is-edit="formModal.isEdit"
            :form-data="formModal.data"
            @close="closeModal"
            @close-and-refresh="closeModal(true)"
        ></form-save>
        <set-auth
            :id="setAuthModal.id"
            :visible="setAuthModal.show"
            @close="setAuthModal.show = false"
        ></set-auth>
    </div>
</template>
<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { Message } from '@arco-design/web-vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import {
        AuthRole,
        EmptyRole,
        reqAuthRole,
        TreeAuthRole,
    } from '@/api/system/auth-role';
    import FormSave from '@/views/system/AuthRole/components/form-save.vue';
    import SetAuth from '@/views/system/AuthRole/components/set-auth.vue';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyRole });
        query.value.order_by = [{field: 'id', is_desc: false}]
        queryList();
        window.onresize = () =>
            (tableScroll.value.maxHeight = window.innerHeight - 170);
    });

    // table start --
    const tableScroll = ref({
        maxHeight: window.innerHeight - 170,
    });
    const expandedKeys = ref([]);

    const queryLoading = ref(false);
    const timePicker = ref<any[]>([]);
    const query = ref<RequestParam<AuthRole>>({
        is_delete: false,
        order_by: [{ field: 'create_time', is_desc: false }],
        model: { ...EmptyRole },
    });

    const list = ref<TreeAuthRole[]>([]);
    const queryList = async () => {
        try {
            queryLoading.value = true;
            query.value.start_time = timePicker.value[0]
                ? dayjs(timePicker.value[0]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            query.value.end_time = timePicker.value[1]
                ? dayjs(timePicker.value[1]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            const res = await reqAuthRole('tree-select', query.value);
            list.value = res.data || [];
        } catch (e) {
        } finally {
            queryLoading.value = false;
        }
    };

    const refresh = async () => {
        query.value.keyword = '';
        timePicker.value = [];
        await queryList();
    };
    // table end --
    const formModal = ref({
        show: false,
        isEdit: false,
        data: {},
    });

    const toAdd = (pid) => {
        formModal.value.data = {
            id: 0,
            parent_id: pid,
            title: '',
            sign: '',
            status: 1,
        };
        formModal.value.isEdit = false;
        formModal.value.show = true;
    };

    const toEdit = (row) => {
        formModal.value.data = row;
        formModal.value.isEdit = true;
        formModal.value.show = true;
    };

    const toDel = async (id: number) => {
        try {
            await reqAuthRole('del', { model: { id: id } as AuthRole });
            Message.success('删除成功');
            await queryList();
        } catch (e) {
            console.log(e);
        }
    };

    const closeModal = (isRefresh: boolean) => {
        formModal.value = {
            show: false,
            isEdit: false,
            data: {},
        };
        if (isRefresh) queryList();
    };

    const setAuthModal = ref({
        show: false,
        id: 0,
    });
    const toSetAuth = (id: number) => {
        setAuthModal.value.id = id;
        setAuthModal.value.show = true;
    };
</script>
<style scoped lang="less"></style>
