<template>
    <div class="six-container" style="height: 100%">
        <a-card style="width: 100%; height: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd(0)"
                        v-permission="'system:job:add'"
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
                row-key="id"
                :scroll="tableScroll"
                :style="{ height: tableScroll.maxHeight + 'px' }"
                :data="list"
            >
                <template #columns>
                    <a-table-column
                        title="名称"
                        data-index="name"
                    ></a-table-column>
                    <a-table-column
                        title="介绍"
                        data-index="description"
                    ></a-table-column>
                    <a-table-column title="操作">
                        <template #cell="{ record }">
                            <a-button
                                type="text"
                                status="normal"
                                :size="size"
                                @click="toAdd(record.id)"
                                >添加
                            </a-button>
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'system:job:save'"
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
                                    v-permission="'system:job:del'"
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

    import FormSave from '@/views/system/Job/components/form-save.vue';
    import { Message } from '@arco-design/web-vue';
    import { EmptyJob, Job, reqJobs, TreeJob } from '@/api/system/job';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyJob });
        query.value.order_by = [{field: 'id', is_desc: false}]
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
    const query = ref<RequestParam<Job>>({
        is_delete: false,
        order_by: [{ field: '`id`', is_desc: false }],
        model: { ...EmptyJob },
    });
    const list = ref<TreeJob[]>([]);
    const queryList = async () => {
        try {
            queryLoading.value = true;
                       query.value.start_time = timePicker.value[0]
                ? dayjs(timePicker.value[0]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            query.value.end_time = timePicker.value[1]
                ? dayjs(timePicker.value[1]).format('YYYY-MM-DD HH:mm:ss')
                : '';
            const res = await reqJobs('tree-select', query.value);
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
        let form = { ...EmptyJob };
        form.parent_id = pid;
        formModal.value.data = form;
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
            await reqJobs('del', { model: { id: id } as Job });
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
