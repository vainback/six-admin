<template>
    <div class="six-container" style="height: 100%">
        <a-card style="height: 100%; width: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd"
                        v-permission="'system:tenant:add'"
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
                :pagination="{
                    total: total,
                    current: query.page || 1,
                    pageSize: query.limit || 10,
                    size: size,
                }"
                :data="list"
                @page-change="pageChange"
            >
                <template #columns>
                    <a-table-column title="ID" data-index="id"></a-table-column>
                    <a-table-column
                        title="租户名"
                        data-index="name"
                    ></a-table-column>
                    <a-table-column
                        title="签名"
                        data-index="sign"
                    ></a-table-column>
                    <a-table-column
                        title="域名"
                        data-index="domain"
                    ></a-table-column>
                    <a-table-column title="状态" data-index="status">
                        <template #cell="{ record }">
                            <a-tag v-if="record.status" color="blue"
                                >正常
                            </a-tag>
                            <a-tag v-else>禁用</a-tag>
                        </template>
                    </a-table-column>
                    <a-table-column title="操作">
                        <template #cell="{ record }">
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'system:tenant:save'"
                                >编辑
                            </a-button>
                            <a-popconfirm
                                content="确认删除？"
                                okText="删除"
                                cancelText="取消"
                                type="error"
                                :ok-button-props="{
                                    status: 'danger',
                                }"
                                @ok="toDel(record.id)"
                            >
                                <a-button
                                    type="text"
                                    status="danger"
                                    :size="size"
                                    v-permission="'system:tenant:del'"
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
            :form-data="formModal.data"
            :is-edit="formModal.isEdit"
            @close="closeModal"
            @close-and-refresh="closeModal(true)"
        ></form-save>
    </div>
</template>

<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import { Message } from '@arco-design/web-vue';
    import { EmptyTenant, reqTenant, Tenant } from '@/api/system/tenant';
    import FormSave from '@/views/system/Tenant/component/form-save.vue';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyTenant });
        queryList();
        window.onresize = () =>
            (tableScroll.value.maxHeight = window.innerHeight - 170);
    });

    // table start --
    const tableScroll = ref({
        maxHeight: window.innerHeight - 170,
    });
    const queryLoading = ref(false);
    const timePicker = ref<any[]>([]);
    const query = ref<RequestParam<Tenant>>({
        is_delete: false,
        keyword: '',
        limit: 10,
        page: 1,
        model: { ...EmptyTenant },
    });

    const list = ref<Tenant[]>([]);
    const total = ref<number>(0);
    const queryList = async () => {
        try {
            queryLoading.value = true;
                       query.value.start_time = timePicker.value[0]
                ? dayjs(timePicker.value[0]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            query.value.end_time = timePicker.value[1]
                ? dayjs(timePicker.value[1]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            const { data } = await reqTenant('list', query.value);
            list.value = data.list || [];
            total.value = data.total || 0;
        } finally {
            queryLoading.value = false;
        }
    };

    const pageChange = (page) => {
        query.value.page = page;
        queryList();
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

    const toAdd = () => {
        formModal.value.data = { ...EmptyTenant };
        formModal.value.data.status = 1;
        formModal.value.isEdit = false;
        formModal.value.show = true;
    };

    const toEdit = (row) => {
        formModal.value.data = { ...row };
        formModal.value.isEdit = true;
        formModal.value.show = true;
    };

    const toDel = async (id) => {
        try {
            await reqTenant('del', { model: { id: id } as Tenant });
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
</script>

<style scoped lang="less"></style>
