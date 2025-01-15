<template>
    <div class="six-container" style="height: 100%">
        <a-layout>
            <a-layout-sider style="width: 250px">
                <a-card style="height: 100%; overflow-y: auto">
                    <a-button
                        long
                        :size="size"
                        :type="
                            selectedKeys.length == 0 ? 'primary' : 'secondary'
                        "
                        @click="clearSelectKeys"
                        >全部组织
                    </a-button>
                    <a-tree
                        ref="tree"
                        :fieldNames="{
                            key: 'id',
                            title: 'name',
                            children: 'children',
                        }"
                        show-line
                        :data="organizations"
                        :expanded-keys="expandKeys"
                        :selected-keys="selectedKeys"
                        @select="onSelect"
                    ></a-tree>
                </a-card>
            </a-layout-sider>
            <a-layout-content>
                <a-card style="height: 100%">
                    <div class="six-flex-between">
                        <a-space>
                            <a-button
                                type="primary"
                                :size="size"
                                @click="toAdd"
                                v-permission="'system:user:add'"
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
                            <a-table-column
                                title="ID"
                                data-index="id"
                            ></a-table-column>
                            <a-table-column
                                title="账号"
                                data-index="username"
                            ></a-table-column>
                            <a-table-column
                                title="昵称"
                                data-index="nickname"
                            ></a-table-column>
                            <a-table-column
                                title="职位"
                                data-index="job.name"
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
                                        v-permission="'system:user:edit'"
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
                                        v-permission="'system:user:del'"
                                    >
                                        <a-button
                                            type="text"
                                            status="danger"
                                            :size="size"
                                            >删除
                                        </a-button>
                                    </a-popconfirm>
                                </template>
                            </a-table-column>
                        </template>
                    </a-table>
                </a-card>
            </a-layout-content>
        </a-layout>
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
    import {EmptyOrganization, reqOrgs, TreeOrganization} from '@/api/system/organization';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import { EmptyUser, reqUser, Userinfo } from '@/api/system/user';
    import FormSave from '@/views/system/AuthUser/components/form-save.vue';
    import { Message } from '@arco-design/web-vue';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyUser });
        queryOrganizations();
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
    const query = ref<RequestParam<Userinfo>>({
        is_delete: false,
        keyword: '',
        limit: 10,
        page: 1,
        start_time: '',
        end_time: '',
        model: { ...EmptyUser },
    });

    const list = ref<Userinfo[]>([]);
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
            const { data } = await reqUser('list', query.value);
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

    // organization tree start --
    const expandKeys = ref<number[]>([]);
    const organizations = ref<TreeOrganization[]>([]);
    const selectedKeys = ref<number[]>([]);
    const queryOrganizations = async () => {
        try {
            let tsQuery = defaultRequestParam({ ...EmptyOrganization });
            tsQuery.order_by = [{field: 'id', is_desc: false}]
            const { data } = await reqOrgs('tree-select', tsQuery);
            organizations.value = data;
            await getExpandKes(data);
        } finally {
        }
    };
    const getExpandKes = async (data: TreeOrganization[]) => {
        data.forEach((item) => {
            expandKeys.value.push(item.id);
            if (item.children && item.children.length > 0) {
                getExpandKes(item.children);
            }
        });
    };

    const onSelect = (ss) => {
        selectedKeys.value = ss.map((item) => Number(item));
        query.value.model.organization_id = ss[0] || 0;
        queryList();
    };

    const clearSelectKeys = () => {
        selectedKeys.value = [];
        queryList();
    };
    // organization tree end --

    const formModal = ref({
        show: false,
        isEdit: false,
        data: {},
    });

    const toAdd = () => {
        formModal.value.data = { ...EmptyUser };
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
            await reqUser('del', { model: { id: id } as Userinfo });
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
