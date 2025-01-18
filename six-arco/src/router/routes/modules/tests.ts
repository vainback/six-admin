import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const BASIC: AppRouteRecordRaw = {
    path: '/testss',
    name: 'Testss',
    component: DEFAULT_LAYOUT,
    meta: {
        locale: '基础设置',
        requiresAuth: true,
        icon: 'icon-desktop',
        order: 98,
    },
    children: [
        {
            path: 'tests',
            name: 'Tests',
            component: () => import('@/views/codegen/tests/index.vue'),
            meta: {
                locale: '字典管理',
                requiresAuth: true,
            },
        },
    ],
};

export default BASIC;
