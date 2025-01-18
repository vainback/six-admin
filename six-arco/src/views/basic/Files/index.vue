<template>
    <div class="six-container" style="height: 100%">
        <a-card style="height: 100%; width: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-upload
                        :action="uploadUrl"
                        :headers="uploadHeaders"
                        :show-file-list="false"
                        @success="uploadSuccess"
                        @error="uploadError"
                        v-permission="'basic:files:add'"
                    >
                        <template #upload-button>
                            <a-button type="primary" :size="size">
                                <template #icon>
                                    <icon-upload />
                                </template>
                                <template #default>上传文件</template>
                            </a-button>
                        </template>
                    </a-upload>
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
                        title="文件类型"
                        data-index="mime"
                        width="200"
                    ></a-table-column>
                    <a-table-column
                        title="存储方式"
                        data-index="type"
                        width="100"
                    >
                        <template #cell="{record}">
                            <span>{{ {local: '本地', cos: '腾讯云', oss: '阿里云', qiniu: '七牛云'}[record.type] || record.type }}</span>
                        </template>
                    </a-table-column>
                    <a-table-column title="上传时间" width="180">
                        <template #cell="{ record }">
                            {{ new Date(record.create_time).toLocaleString() }}
                        </template>
                    </a-table-column>
                    <a-table-column title="文件" data-index="url">
                        <template #cell="{ record }">
                            <div v-if="record.mime.includes('image')">
                                <a-image
                                    v-if="record.type == 'local'"
                                    width="50"
                                    :src="urlPrefix + record.url"
                                ></a-image>
                                <a-image
                                    v-else
                                    width="50"
                                    :src="record.url"
                                ></a-image>
                            </div>
                            <div v-else-if="record.mime.includes('video')">
                                <a-link
                                    v-if="record.type == 'local'"
                                    :href="urlPrefix+ record.url"
                                    >查看视频
                                </a-link>
                                <a-link v-else :href="record.url"
                                    >查看视频
                                </a-link>
                            </div>
                            <div v-else>
                                <a-typography-text copyable ellipsis>
                                    {{
                                        record.type == 'local'
                                                ? urlPrefix  + record.url
                                                : record.url
                                    }}
                                </a-typography-text>
                            </div>
                        </template>
                    </a-table-column>
                    <a-table-column title="操作" width="180">
                        <template #cell="{ record }">
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
                                    v-permission="'basic:files:del'"
                                    >删除
                                </a-button>
                            </a-popconfirm>
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
    import { Message } from '@arco-design/web-vue';
    import { EmptyFile, Files, reqFiles } from '@/api/basic/files';
    import { uploadHeaders, uploadUrl } from '@/api/user-single';
    import dayjs from 'dayjs';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    const urlPrefix = ref(import.meta.env.VITE_API_BASE_URL);
    onMounted(() => {
        query.value = defaultRequestParam({ ...EmptyFile });
        query.value.limit = 10;
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
    const query = ref<RequestParam<Files>>({
        is_delete: false,
        keyword: '',
        limit: 10,
        page: 1,
        model: { ...EmptyFile },
    });

    const list = ref<Files[]>([]);
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
            const { data } = await reqFiles('list', query.value);
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

    const toDel = async (id) => {
        try {
            await reqFiles('del', { model: { id: id } as Files, is_delete: true });
            Message.success('删除成功');
            await queryList();
        } catch (e) {
            console.log(e);
        }
    };

    const uploadSuccess = async (e) => {
        console.log(e)
        Message.success('上传成功');
        await queryList();
    };

    const uploadError = async (e) => {
        Message.error('上传成功');
    };
</script>

<style scoped lang="less">
div.arco-typography {
    margin: 0 !important;
}
</style>
