<template>
    <div class="six-container" style="height: 100%">
        <a-card style="height: 100%; width: 100%">
            <div class="six-flex-between">
                <a-space>
                    <a-button
                        type="primary"
                        :size="size"
                        @click="toAdd"
                        v-permission="'tests:tests:add'"
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
                    <a-table-column title="测试普通字符串" data-index="te_str"></a-table-column>
                    <a-table-column title="测试普通整形" data-index="te_int"></a-table-column>
                    <a-table-column title="测试大整形" data-index="te_bigint"></a-table-column>
                    <a-table-column title="测试浮点数" data-index="te_float"></a-table-column>
                    <a-table-column title="测试高精度浮点" data-index="te_decimal"></a-table-column>

<a-table-column title="测试单选下拉框" data-index="te_select">
	<template #cell="{ record }">
		<a-tag :color="dictMap['test'][record.te_select].color">{{dictMap['test'][record.te_select].label}}</a-tag>
	</template>
</a-table-column>
<a-table-column title="测试多选下拉框" data-index="te_select_many">
	<template #cell="{ record }">
		<a-space wrap><a-tag v-for="item in record.te_select_many" :color="dictMap['test2'][item].color">{{dictMap['test2'][item].label}}</a-tag></a-space>
	</template>
</a-table-column>
<a-table-column title="单选框" data-index="te_radio">
	<template #cell="{ record }">
		<a-tag :color="dictMap['test'][record.te_radio].color">{{dictMap['test'][record.te_radio].label}}</a-tag>
	</template>
</a-table-column>
<a-table-column title="复选框" data-index="te_checkbox">
	<template #cell="{ record }">
		<a-space wrap><a-tag v-for="item in record.te_checkbox" :color="dictMap['test2'][item].color">{{dictMap['test2'][item].label}}</a-tag></a-space>
	</template>
</a-table-column>
<a-table-column title="开关" data-index="te_switch">
	<template #cell="{ record }">
		<a-tag v-if="record.status" color="blue"
			>正常
		</a-tag>
		<a-tag v-else>禁用</a-tag>
	</template>
</a-table-column>                    <a-table-column title="时间选择器" data-index="te_timepicker"></a-table-column>
                    <a-table-column title="日期选择器" data-index="te_datepicker"></a-table-column>
                    <a-table-column title="日期时间选择器" data-index="te_datetimepicker"></a-table-column>

<a-table-column title="上传图片" data-index="te_image_one">
	<template #cell="{ record }">
		<a-space wrap><a-image v-for="url in record.te_image_one" width="50" :src="String(url).includes('http') ? url : urlPrefix+ 'admin/' + url"></a-image></a-space>
	</template>
</a-table-column>
<a-table-column title="上传图片（多图）" data-index="te_image_many">
	<template #cell="{ record }">
		<a-space wrap><a-image v-for="url in record.te_image_many" width="50" :src="String(url).includes('http') ? url : urlPrefix+ 'admin/' + url"></a-image></a-space>
	</template>
</a-table-column>
<a-table-column title="上传视频" data-index="te_video">
	<template #cell="{ record }">
		<a-space wrap><a-link v-for="(url, index) in record.te_video" :href="String(url).includes('http') ? url : urlPrefix+ 'admin/' + url">链接 - {{ index }}</a-link></a-space>
	</template>
</a-table-column>
<a-table-column title="上传文件" data-index="te_file">
	<template #cell="{ record }">
		<a-space wrap><a-link v-for="(url, index) in record.te_file" :href="String(url).includes('http') ? url : urlPrefix+ 'admin/' + url">链接 - {{ index }}</a-link></a-space>
	</template>
</a-table-column>                    <a-table-column title="文本域" data-index="te_text"><template #cell="{ record }"><a-typography-text ellipsis>{{ record.te_text }}</a-typography-text></template></a-table-column>
                    <a-table-column title="操作">
                        <template #cell="{ record }">
                            <a-button
                                type="text"
                                status="warning"
                                :size="size"
                                @click="toEdit(record)"
                                v-permission="'tests:tests:save'"
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
                                    v-permission="'tests:tests:del'"
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
            :dict="dictList"
            @close="closeModal"
            @close-and-refresh="closeModal(true)"
        ></form-save>
    </div>
</template>

<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import { defaultRequestParam, RequestParam } from '@/api/base';
    import { Message } from '@arco-design/web-vue';
    import { EmptyTests, reqTests, Tests } from '@/api/codegen/tests';
    import FormSave from '@/views/codegen/Tests/components/form-save.vue';
    import dayjs from 'dayjs';
    import { Dict, reqDict, EmptyDict } from '@/api/basic/dict';

    const urlPrefix = ref(import.meta.env.VITE_API_BASE_URL);
    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    onMounted(() => {
        requestDict()
        query.value = defaultRequestParam({ ...EmptyTests });
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
    const query = ref<RequestParam<Tests>>({
        is_delete: false,
        keyword: '',
        limit: 10,
        page: 1,
        model: { ...EmptyTests },
    });

    const list = ref<Tests[]>([]);
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
            const { data } = await reqTests('list', query.value);
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

    
    const dictMap = ref({})
	const dictList = ref({})
    const requestDict = async () => {
        const { data } = await reqDict('select', defaultRequestParam({...EmptyDict}));
        let m = {}
		let l = {}
        for (const key in data) {
            let item = data[key]
            if (!m[item.type]) {
                m[item.type] = {}
            }
			if (!l[item.type]) {
				l[item.type] = []
			}
            m[item.type][item.value] = item
			l[item.type].push(item)
        }
        dictMap.value = {...m}
		dictList.value = {...l}
    }

    // table end --
    const formModal = ref({
        show: false,
        isEdit: false,
        data: {},
    });

    const toAdd = () => {
        formModal.value.data = { ...EmptyTests };
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
            await reqTests('del', { model: { id: id } as Tests });
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
