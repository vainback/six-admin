<template>
    <a-modal
        title="设置权限"
        v-model:visible="props.visible"
        @before-close="() => checkedKeys = []"
        @cancel="emit('close')"
        @ok="submit"
    >
        <div style="width: 100%; height: 500px; overflow-y: auto">
            <a-spin
                :size="20"
                :loading="loading"
                tip="加载中"
                style="width: 100%"
            >
                <a-tree
                    ref="menutree"
                    v-if="!loading"
                    checkable
                    check-strictly
                    default-expand-all
                    @expand="es => expandedKeys = es.map(item => Number(item))"
                    :data="treeData"
                    :checkedKeys="checkedKeys"
                    @check="onCheck"
                />
            </a-spin>
        </div>
    </a-modal>
</template>
<script setup lang="ts">
import {ref, watch} from 'vue';
    import {
        AuthRule,
        reqAuthRule,
        TreeAuthRule,
    } from '@/api/system/auth-rule';
    import { RequestParam } from '@/api/base';
    import {
        AuthRelation,
        getCheckedRule,
        set,
    } from '@/api/system/auth-relation';
    import { TreeNodeData } from '@arco-design/web-vue';

    const props = defineProps<{
        visible: boolean;
        id: number;
    }>();
    const loading = ref(false);
    const treeData = ref([]);
    const checkedKeys = ref<number[]>([]);
    // const expandedKeys = ref<number[]>([1,2])
    const menutree = ref(null);
    const emit = defineEmits(['close', 'refresh', 'closeAndRefresh']);

    const submit = async () => {
        if (menutree.value) {
            const arr = checkedKeys.value.concat(menutree.value.getHalfCheckedNodes().map(item => item.key))
            checkedKeys.value = Array.from(new Set(arr)).sort()    
        }
        try {
            await set(props.id, checkedKeys.value);
        } finally {
            emit('close');
        }
    };

    const getData = async () => {
        loading.value = true;
        try {
            const keys = await getCheckedRule({
                model: { role_id: props.id },
            } as RequestParam<AuthRelation>);
            checkedKeys.value = keys.data || [];
            const res = await reqAuthRule('select', {
                model: { status: 1 },
            } as RequestParam<AuthRule>);
            treeData.value = replaceKeys(res.data || []);
        } finally {
            expandedKeys.value = [1,2]
            loading.value = false;
        }
    };

    const replaceKeys = (list: TreeAuthRule[]): TreeNodeData[] => {
        const tree: TreeNodeData[] = [];
        list.forEach((item) => {
            let row: TreeNodeData = {
                key: item.id,
                title: item.locale,
            };
            if (item.children && item.children.length > 0) {
                row.children = replaceKeys(item.children);
            }
            tree.push(row);
        });
        return tree;
    };

    const onCheck = (cs) => {
        checkedKeys.value = cs.map(item => Number(item))
    }

    watch(
        () => props.visible,
        (newVal, oldVal) => {
            if (props.visible) {
                getData()
            }
        }
    )

</script>
<style scoped lang="less"></style>
