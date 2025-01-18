<template>
    <div class="six-container" style="height: 100%">
        <a-card>
            <a-layout>
                <a-layout-header>
                    <a-tabs
                        v-model="defaultCheckedKey"
                        position="top"
                        @tab-click="tabClick"
                    >
                        <a-tab-pane key="root" title="字典类型" />
                        <a-tab-pane
                            v-for="item in listRoot"
                            :key="item.value"
                            :title="item.label"
                        />
                    </a-tabs>
                </a-layout-header>
                <a-layout-content>
                    <div class="six-flex-between">
                        <a-space>
                            <a-button
                                type="primary"
                                :size="size"
                                @click="toAdd"
                                v-permission="'basic:dict:add'"
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
                        :pagination="false"
                        :style="{ height: tableScroll.maxHeight + 'px' }"
                        row-key="id"
                        :data="
                            (() => {
                                if (query.model.type == 'root') {
                                    return listRoot;
                                }
                                return listOther;
                            })()
                        "
                    >
                        <template #columns>
                            <a-table-column
                                title="标题"
                                data-index="label"
                            ></a-table-column>
                            <a-table-column
                                title="值"
                                data-index="value"
                            ></a-table-column>
                            <a-table-column
                                title="颜色"
                                data-index="color"
                            >
                                <template #cell="{ record }">
                                    <a-color-picker v-if="record.type != 'root'" v-model="record.color" showText  disabled/>
                                </template>
                            </a-table-column>
                            <a-table-column
                                v-if="isTenantRoot()"
                                title="是否同步"
                                data-index="is_sync"
                            >
                                <template #cell="{ record }">
                                    <a-tag v-if="record.is_sync" color="blue"
                                        >同步
                                    </a-tag>
                                    <a-tag v-else>不同步</a-tag>
                                </template>
                            </a-table-column>
                            <a-table-column title="操作">
                                <template #cell="{ record }">
                                    <a-button
                                        type="text"
                                        status="warning"
                                        :size="size"
                                        @click="toEdit(record)"
                                        v-permission="'basic:dict:save'"
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
                                            type="text"
                                            status="danger"
                                            :size="size"
                                            v-permission="'basic:dict:del'"
                                            >删除
                                        </a-button>
                                    </a-popconfirm>
                                </template>
                            </a-table-column>
                        </template>
                    </a-table>
                </a-layout-content>
            </a-layout>
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
    import { Dict, EmptyDict, reqDict } from '@/api/basic/dict';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import FormSave from '@/views/basic/Dict/components/form-save.vue';
    import { Message } from '@arco-design/web-vue';
    import { cloneDeep } from 'lodash';
    import { isTenantRoot } from '@/utils/auth';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);

    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyDict });
        query.value.model.type = 'root';
        queryList();
        window.onresize = () =>
            (tableScroll.value.maxHeight = window.innerHeight - 170 - 55);
    });

    const tableScroll = ref({
        maxHeight: window.innerHeight - 170 - 55,
    });
    const queryLoading = ref(false);
    const timePicker = ref<any[]>([]);
    const query = ref<RequestParam<Dict>>({
        is_delete: false,
        keyword: '',
        model: { ...EmptyDict },
    });
    const defaultCheckedKey = ref<string>('root');

    const listRoot = ref<Dict[]>([]);
    const listOther = ref<Dict[]>([]);

    const queryList = async () => {
        try {
            queryLoading.value = true;
            query.value.model.type = defaultCheckedKey.value;
            const { data } = await reqDict('select', query.value);
            if (defaultCheckedKey.value === 'root') {
                listRoot.value = data || [];
            } else {
                listOther.value = data || [];
            }
        } finally {
            queryLoading.value = false;
        }
    };
    const refresh = async () => {
        query.value.keyword = '';
        timePicker.value = [];
        await queryList();
    };

    const tabClick = (e) => {
        defaultCheckedKey.value = e;
        queryList();
    };

    const formModal = ref({
        show: false,
        isEdit: false,
        data: {},
    });

    const toAdd = () => {
        formModal.value.data = { ...EmptyDict };
        formModal.value.data.type = cloneDeep(query.value.model.type);
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
            await reqDict('del', { model: { id: id } as Dict });
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
