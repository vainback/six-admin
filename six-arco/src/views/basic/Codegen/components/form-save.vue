<template>
    <a-modal
        fullscreen
        :title="props.isEdit ? '编辑' : '添加'"
        v-model:visible="props.visible"
        @cancel="emit('close')"
        @ok="submit"
    >
        <a-form :model="form" :size="size" label-align="right" auto-label-width>
            <a-space wrap>
                <a-form-item label="模块名称" required>
                    <a-input
                        v-model="form.title"
                        placeholder="例如：用户管理"
                    ></a-input>
                </a-form-item>
                <a-form-item label="数据表名" required>
                    <a-input
                        v-model="form.table"
                        placeholder="必须是小写加下划线格式"
                    ></a-input>
                </a-form-item>
                <a-form-item label="上级菜单" required>
                    <a-tree-select
                        allow-search
                        :data="menu"
                        v-model="form.parent_module"
                        style="width: 200px"
                        placeholder="请选择上级菜单"
                    ></a-tree-select>
                </a-form-item>
                <a-form-item label="软删除" required>
                    <a-switch
                        style="width: 65px"
                        type="round"
                        v-model="form.is_soft_delete"
                        checked-text="ON"
                        unchecked-text="OFF"
                    >
                    </a-switch>
                </a-form-item>
                <a-form-item label="多租户" required>
                    <a-switch
                        style="width: 65px"
                        type="round"
                        v-model="form.is_tenant"
                        checked-text="ON"
                        unchecked-text="OFF"
                    >
                    </a-switch>
                </a-form-item>
                <a-form-item label="默认字段">
                    <a-input-tag
                        :default-value="['id', 'create_time', 'update_time']"
                        style="width: 250px"
                        readonly
                    />
                </a-form-item>
                <a-form-item>
                    <a-button
                        v-if="!form.fields || (form.fields as any[]).length === 0"
                        type="primary"
                        @click="pushField"
                        >添加字段
                    </a-button>
                </a-form-item>
            </a-space>
            <a-divider />
            <a-space direction="vertical" fill>
                <a-space
                    wrap
                    v-for="(item, index) in form.fields as any[]"
                    :key="index"
                >
                    <a-form-item label="字段名" required>
                        <a-input
                            v-model="item.name"
                            placeholder="必须是小写加下划线格式"
                        ></a-input>
                    </a-form-item>
                    <a-form-item label="字段标题" required>
                        <a-input
                            v-model="item.title"
                            placeholder="在列表和表单展示的标题"
                        ></a-input>
                    </a-form-item>
                    <a-form-item label="字段类型" required>
                        <a-select
                            v-model="item.type"
                            placeholder="请选择"
                            allow-search
                            style="width: 135px"
                        >
                            <a-option value="int(11)">整形</a-option>
                            <a-option value="bigint(20)">大整形</a-option>
                            <a-option value="varchar(255)">字符串</a-option>
                            <a-option value="float">浮点数</a-option>
                            <a-option value="decimal(10,2)"
                                >高精度浮点数
                            </a-option>
                            <a-option value="下拉框">下拉框</a-option>
                            <a-option value="下拉框多选">下拉框多选</a-option>
                            <a-option value="单选框">单选框</a-option>
                            <a-option value="复选框">复选框</a-option>
                            <a-option value="开关">开关</a-option>
                            <a-option value="时间选择器">时间选择器</a-option>
                            <a-option value="日期选择器">日期选择器</a-option>
                            <a-option value="日期时间选择器"
                                >日期时间选择器
                            </a-option>
                            <a-option value="上传图片">上传图片</a-option>
                            <a-option value="上传图片（多图）"
                                >上传图片-多图
                            </a-option>
                            <a-option value="上传视频">上传视频</a-option>
                            <a-option value="上传文件">上传文件</a-option>
                            <a-option value="text">文本域</a-option>
                            <a-option value="富文本">富文本</a-option>
                        </a-select>
                    </a-form-item>
                    <a-form-item
                        label="字典类型"
                        v-if="
                            [
                                '下拉框',
                                '下拉框多选',
                                '单选框',
                                '复选框',
                            ].includes(item.type)
                        "
                        required
                    >
                        <a-select
                            v-model="item.dict_type"
                            placeholder="该类型字段需选择字典"
                            allow-search
                        >
                            <a-option
                                v-for="item in dict"
                                :value="item.value"
                                >{{ item.label }}</a-option
                            >
                        </a-select>
                    </a-form-item>
                    <a-form-item label="默认值" required>
                        <a-select
                            v-model="item.default"
                            placeholder="可输入"
                            style="width: 120px"
                            allow-create
                        >
                            <a-option value="null">NULL</a-option>
                            <a-option value="">空字符串</a-option>
                        </a-select>
                    </a-form-item>
                    <a-form-item label="索引" style="margin-right: 30px">
                        <a-select
                            v-model="item.index"
                            :default-value="''"
                            style="width: 100px"
                        >
                            <a-option value="">无索引</a-option>
                            <a-option value="unique">unique</a-option>
                            <a-option value="index">index</a-option>
                        </a-select>
                    </a-form-item>
                    <a-form-item
                        label="模糊查询"
                        v-if="(item.type as string) == 'varchar(255)'"
                    >
                        <a-switch
                            style="width: 65px"
                            type="round"
                            v-model="item.is_keyword"
                            checked-text="是"
                            unchecked-text="否"
                        >
                        </a-switch>
                    </a-form-item>
                    <a-form-item label="是否必填">
                        <a-switch
                            style="width: 65px"
                            type="round"
                            v-model="item.is_required"
                            checked-text="是"
                            unchecked-text="否"
                        >
                        </a-switch>
                    </a-form-item>
                    <a-form-item label="字段注释">
                        <a-input v-model="item.comment"></a-input>
                    </a-form-item>
                    <a-space style="margin-top: -15px">
                        <a-button
                            status="danger"
                            @click="(form.fields as any[]).splice(index, 1)"
                        >
                            <template #icon>
                                <div style="padding: 0 15px">
                                    <icon-delete />
                                </div>
                            </template>
                        </a-button>
                        <a-button
                            v-if="index == (form.fields as any[]).length - 1"
                            type="primary"
                            @click="pushField"
                        >
                            <template #icon>
                                <div style="padding: 0 15px">
                                    <icon-plus />
                                </div>
                            </template>
                        </a-button>
                    </a-space>
                </a-space>
            </a-space>
        </a-form>
    </a-modal>
</template>
<script setup lang="ts">
    import { ref, watchEffect } from 'vue';
    import { Message } from '@arco-design/web-vue';
    import { reqCodegen, Codegen } from '@/api/basic/code';
    import { Dict } from '@/api/basic/dict';

    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    // const size = 'small';
    const props = defineProps<{
        visible: boolean;
        isEdit: boolean;
        formData: Codegen;
        menu: any[];
        dict: Dict[];
    }>();

    const form = ref<Codegen>({ ...props.formData }); // 深拷贝 防止与列表数据双向绑定
    const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

    const submit = async () => {
        console.log('xxxxxxxxxxxxxxxxx');
        try {
            // form.value.fields = (form.value.fields as any[]).map(
            //     (item: any) => {
            //         if ((item.default as string | null) == 'null') {
            //             item.default = null;
            //         }
            //     }
            // );
            console.log(form.value);
            console.log(JSON.stringify(form.value));

            await reqCodegen(props.isEdit ? 'save' : 'add', {
                model: form.value,
            });
            // emit('closeAndRefresh');
            Message.success(props.isEdit ? '修改成功' : '添加成功');
        } catch (e) {}
    };

    const pushField = () => {
        if (
            form.value.fields == null ||
            (form.value.fields as any[]).length === 0
        ) {
            form.value.fields = [{}];
        } else {
            console.log(form.value.fields, '可以进来？');
            (form.value.fields as any[]).push({});
        }
    };

    watchEffect(() => {
        form.value = { ...props.formData };

        let fields = [];
        fields.push({
            title: '测试普通字符串',
            type: 'varchar(255)',
            default: '',
            index: 'index',
            comment: '测试普通字符串',
            is_keyword: true,
            is_required: true,
            dict_type: '',
            name: 'te_str'
        });
        fields.push({
            title: '测试普通整形',
            type: 'int(11)',
            default: '',
            index: 'index',
            comment: '测试普通整形',
            is_keyword: true,
            is_required: true,
            dict_type: '',
            name: 'te_int'
        });
        fields.push({
            title: '测试大整形',
            type: 'bigint(20)',
            default: '',
            index: 'index',
            comment: '测试大整形',
            is_keyword: true,
            is_required: true,
            dict_type: '',
            name: 'te_bigint'
        });
        fields.push({
            title: '测试浮点数',
            type: 'float',
            default: '0',
            index: '',
            comment: '测试浮点数',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_float'
        });
        fields.push({
            title: '测试高精度浮点',
            type: 'decimal(10,2)',
            default: '0',
            index: '',
            comment: '测试高精度浮点',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_decimal'
        });
        
        
        
        fields.push({
            title: '测试单选下拉框',
            type: '下拉框',
            default: '',
            index: '',
            comment: '测试单选下拉框',
            is_keyword: false,
            is_required: true,
            dict_type: 'test',
            name: 'te_select'
        });
        
        fields.push({
            title: '测试多选下拉框',
            type: '下拉框多选',
            default: '',
            index: '',
            comment: '测试多选下拉框',
            is_keyword: false,
            is_required: true,
            dict_type: 'test2',
            name: 'te_select_many'
        });
        
        fields.push({
            title: '单选框',
            type: '单选框',
            default: '',
            index: '',
            comment: '单选框',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_radio'
        });
        
        fields.push({
            title: '复选框',
            type: '复选框',
            default: '',
            index: '',
            comment: '复选框',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_checkbox'
        });
        
        fields.push({
            title: '开关',
            type: '开关',
            default: '0',
            index: '',
            comment: '开关',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_switch'
        });
        
        fields.push({
            title: '时间选择器',
            type: '时间选择器',
            default: '',
            index: '',
            comment: '时间选择器',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_timepicker'
        });
        fields.push({
            title: '日期选择器',
            type: '日期选择器',
            default: '',
            index: '',
            comment: '日期选择器',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_datepicker'
        });
        fields.push({
            title: '日期时间选择器',
            type: '日期时间选择器',
            default: '',
            index: '',
            comment: '日期时间选择器',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_datetimepicker'
        });
        
        fields.push({
            title: '上传图片',
            type: '上传图片',
            default: '',
            index: '',
            comment: '上传图片',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_image_one'
        });
        fields.push({
            title: '上传图片（多图）',
            type: '上传图片（多图）',
            default: '',
            index: '',
            comment: '上传图片（多图）',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_image_many'
        });
        
        fields.push({
            title: '上传视频',
            type: '上传视频',
            default: '',
            index: '',
            comment: '上传视频',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_video'
        });
        fields.push({
            title: '上传文件',
            type: '上传文件',
            default: '',
            index: '',
            comment: '上传文件',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_file'
        });
        
        fields.push({
            title: '文本域',
            type: 'text',
            default: '',
            index: '',
            comment: '文本域',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_text'
        });
        
        fields.push({
            title: '富文本',
            type: '富文本',
            default: '',
            index: '',
            comment: '富文本',
            is_keyword: false,
            is_required: true,
            dict_type: '',
            name: 'te_editor'
        });
        
        form.value = {
            table: 'tests',
            title: '测试模块',
            fields: fields,
            parent_module: 99,
            is_soft_delete: true,
            is_tenant: true,
        } as Codegen;
    });
</script>
<style scoped lang="less"></style>
