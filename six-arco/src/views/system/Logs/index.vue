<template>
    <div class="six-container" style="height: 100%">
        <a-card style="height: 100%; width: 100%">
            <div class="six-flex-between">
                <a-space>
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
                    <a-table-column title="操作人" data-index="user.nickname">
                        <template #cell="{ record }"
                            >{{
                                record.user.nickname +
                                ' (' +
                                record.user.username +
                                ')'
                            }}
                        </template>
                    </a-table-column>
                    <a-table-column
                        title="路由名"
                        data-index="route_name"
                        width="300"
                    >
                        <template #cell="{ record }">
                            <a-typography-text ellipsis>
                                {{ record.route_name }}
                            </a-typography-text>
                        </template>
                    </a-table-column>
                    <a-table-column
                        title="API"
                        data-index="route"
                    ></a-table-column>
                    <a-table-column title="IP" data-index="ip"></a-table-column>
                    <a-table-column
                        title="执行时长"
                        data-index="latency"
                    ></a-table-column>
                    <a-table-column title="记录时间" data-index="create_time">
                        <template #cell="{ record }">
                            {{ new Date(record.create_time).toLocaleString() }}
                        </template>
                    </a-table-column>
                    <a-table-column>
                        <template #cell="{ record }">
                            <a-space>
                                <a-popover position="tr">
                                    <template #content>
                                        <vue-json-pretty
                                            :data="
                                                JSON.parse(record.request_body)
                                            "
                                        ></vue-json-pretty>
                                    </template>
                                    <a-button :size="size">请求参数</a-button>
                                </a-popover>
                                <a-popover position="tr">
                                    <template #content>
                                        <a-link status="warning"
                                            >响应的列表数据过长时，将不会记录在日志中
                                        </a-link>
                                        <vue-json-pretty
                                            :data="
                                                JSON.parse(record.response_body)
                                            "
                                        ></vue-json-pretty>
                                    </template>
                                    <a-button :size="size">响应结果</a-button>
                                </a-popover>
                            </a-space>
                        </template>
                    </a-table-column>
                </template>
            </a-table>
        </a-card>
    </div>
</template>

<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import { EmptyLog, Log, reqLog } from '@/api/system/logs';
    import VueJsonPretty from 'vue-json-pretty';
    import 'vue-json-pretty/lib/styles.css';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyLog });
        queryList();
        window.onresize = () =>
            (tableScroll.value.maxHeight = window.innerHeight - 170);
    });

    // table start --
    const tableScroll = ref({
        maxHeight: window.innerHeight - 170,
    });
    const queryLoading = ref(false);
    const timePicker = ref<number[]>([]);
    const query = ref<RequestParam<Log>>({
        is_delete: false,
        keyword: '',
        limit: 15,
        page: 1,
        model: { ...EmptyLog },
    });

    const list = ref<Log[]>([]);
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
            const { data } = await reqLog('list', query.value);
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
</script>

<style scoped lang="less">
    div.arco-typography {
        margin: 0 !important;
    }
</style>
