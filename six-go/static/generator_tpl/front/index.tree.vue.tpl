<template>
    <div class="six-container" style="height: 100%">
        <a-card style="width: 100%; height: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd(0)"
                        v-permission="'<Tpl-AuthName>:add'"
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
                    <Tpl-Table-Columns>
                    <a-table-column title="操作">
                        <template #cell="{ record }">
                            <a-button
                                type="text"
                                status="normal"
                                :size="size"
                                @click="toAdd(record.id, record.auth)"
                                v-permission="'<Tpl-AuthName>:add'"
                                >添加
                            </a-button>
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'<Tpl-AuthName>:save'"
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
                                    v-permission="'<Tpl-AuthName>:del'"
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
            <Tpl - Dict>
            @close="closeModal"
            @close-and-refresh="closeModal(true)"
        ></form-save>
    </div>
</template>
<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';

    import {
        <Tpl-Module>,
        Empty<Tpl-Module>,
        req<Tpl-Module>,
        Tree<Tpl-Module>,
    } from '@/api/<Tpl-Module>';
    import FormSave from '@/views/system/<Tpl-Module>/components/form-save.vue';
    import { Message } from '@arco-design/web-vue';
    import { cloneDeep } from 'lodash';
    import dayjs from 'dayjs';
    <Tpl - Import - Dict>

    const urlPrefix = ref(import.meta.env.VITE_API_BASE_URL);
    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        <Tpl - Request - Dict - Object>
        query.value = defaultRequestParam({ ...Empty<Tpl-Module> });
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
    const query = ref<RequestParam<<Tpl-Module>>>({
        is_delete: false,
        order_by: [{ field: '`order`', is_desc: false }],
        model: { ...Empty<Tpl-Module> },
    });

    const list = ref<Tree<Tpl-Module>[]>([]);
    const queryList = async () => {
        try {
            queryLoading.value = true;
                       query.value.start_time = timePicker.value[0]
                ? dayjs(timePicker.value[0]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            query.value.end_time = timePicker.value[1]
                ? dayjs(timePicker.value[1]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            const res = await req<Tpl-Module>('tree-select', query.value);
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

    <Tpl - Request - Dict - Object>
    // table end --
    const formModal = ref({
        show: false,
        isEdit: false,
        data: {},
    });

    const toAdd = (pid: number) => {
        formModal.value.data = { ...Empty<Tpl-Module> }
        formModal.value.data.parent_id = pid
        formModal.value.isEdit = false;
        formModal.value.show = true;
    };

    const toEdit = (row: <Tpl-Module>) => {
        formModal.value.data = row;
        formModal.value.isEdit = true;
        formModal.value.show = true;
    };

    const toDel = async (id: number) => {
        try {
            await req<Tpl-Module>('del', { model: { id: id } as <Tpl-Module> });
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
