import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const BASIC: AppRouteRecordRaw = {
    path: '/basic',
    name: 'Basic',
    component: DEFAULT_LAYOUT,
    meta: {
        locale: '基础设置',
        requiresAuth: true,
        icon: 'icon-desktop',
        order: 98,
    },
    children: [
        {
            path: 'dict',
            name: 'Dict',
            component: () => import('@/views/basic/Dict/index.vue'),
            meta: {
                locale: '字典管理',
                requiresAuth: true,
            },
        },
        {
            path: 'files',
            name: 'Files',
            component: () => import('@/views/basic/Files/index.vue'),
            meta: {
                locale: '文件管理',
                requiresAuth: true,
            },
        },
        {
            path: 'cron',
            name: 'Cron',
            component: () => import('@/views/basic/Cron/index.vue'),
            meta: {
                locale: '定时任务',
                requiresAuth: true,
            },
        },
        {
            path: 'code',
            name: 'Code',
            component: () => import('@/views/basic/CodeGenerator/index.vue'),
            meta: {
                locale: '代码生成',
                requiresAuth: true,
            },
        },
    ],
};

export default BASIC;
