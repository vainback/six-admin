<template>
    <div class="six-container" style="height: 100%">
        <a-card style="width: 100%; height: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd(0)"
                        v-permission="'system:rule:add'"
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
            <!-- row-key必须设置 不设置的话 展开不受控制 点击一个会全部展开 -->
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
                        title="菜单"
                        data-index="locale"
                    ></a-table-column>
                    <a-table-column
                        title="图标"
                        data-index="icon"
                    ></a-table-column>
                    <a-table-column
                        title="路由"
                        data-index="path"
                    ></a-table-column>
                    <a-table-column
                        title="权限"
                        data-index="auth"
                    ></a-table-column>
                    <a-table-column
                        title="后端路由"
                        data-index="api_route"
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
                                @click="toAdd(record.id, record.auth)"
                                v-permission="'system:rule:add'"
                                >添加
                            </a-button>
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'system:rule:save'"
                                >编辑
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
                                    v-permission="'system:rule:del'"
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
    </div>
</template>
<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';

    import {
        AuthRule,
        EmptyRule,
        reqAuthRule,
        TreeAuthRule,
    } from '@/api/system/auth-rule';
    import FormSave from '@/views/system/AuthRule/components/form-save.vue';
    import { Message } from '@arco-design/web-vue';
    import { cloneDeep } from 'lodash';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyRule });
        query.value.order_by = [{field: '`order`', is_desc: false}]
        queryList();
        window.onresize = () =>
            (tableScroll.value.maxHeight = window.innerHeight - 170);
    });
    // table start --
    const tableScroll = ref({
        maxHeight: window.innerHeight - 170,
    });
    const expandedKeys = ref([1, 2]);
    const queryLoading = ref(false);
    const timePicker = ref<any[]>([]);
    const query = ref<RequestParam<AuthRule>>({
        is_delete: false,
        order_by: [{ field: '`order`', is_desc: false }],
        model: { ...EmptyRule },
    });

    const list = ref<TreeAuthRule[]>([]);
    const queryList = async () => {
        try {
            queryLoading.value = true;
                       query.value.start_time = timePicker.value[0]
                ? dayjs(timePicker.value[0]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            query.value.end_time = timePicker.value[1]
                ? dayjs(timePicker.value[1]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            const res = await reqAuthRule('tree-select', query.value);
            list.value = res.data || [];

            // listTotal.value = res.total || 0
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

    const toAdd = (pid: number, prefix?: string) => {
        formModal.value.data = {
            id: null,
            parent_id: pid,
            type: 'menu',
            component: pid == 0 ? 'Layout' : '',
            path: '',
            auth: '',
            locale: '',
            icon: '',
            order: 98,
            status: 1,
            authPrefix: prefix ? prefix + ':' : '',
        };
        formModal.value.isEdit = false;
        formModal.value.show = true;
    };

    const toEdit = (row: AuthRule) => {
        let rowData = cloneDeep(row);
        let auths = rowData.auth.split(':');
        formModal.value.data = rowData;
        formModal.value.data.auth = auths.pop();
        let prefix = auths.join(':');
        formModal.value.data.authPrefix = prefix ? prefix + ':' : '';
        formModal.value.isEdit = true;
        formModal.value.show = true;
    };

    const toDel = async (id: number) => {
        try {
            await reqAuthRule('del', { model: { id: id } as AuthRule });
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
        if (isRefresh) {
            setTimeout(() => {
                queryList();
            }, 150);
        }
    };
</script>
<style scoped lang="less"></style>
