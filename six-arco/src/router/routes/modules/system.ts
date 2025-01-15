import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const SYSTEM: AppRouteRecordRaw = {
    path: '/system',
    name: 'System',
    component: DEFAULT_LAYOUT,
    meta: {
        locale: '系统管理',
        requiresAuth: true,
        icon: 'icon-settings',
        order: 98,
    },
    children: [
        {
            path: 'authrule',
            name: 'AuthRule',
            component: () => import('@/views/system/AuthRule/index.vue'),
            meta: {
                locale: '菜单管理',
                requiresAuth: true,
            },
        },
        {
            path: 'authrole',
            name: 'AuthRole',
            component: () => import('@/views/system/AuthRole/index.vue'),
            meta: {
                locale: '角色管理',
                requiresAuth: true,
            },
        },
        {
            path: 'authuser',
            name: 'AuthUser',
            component: () => import('@/views/system/AuthUser/index.vue'),
            meta: {
                locale: '用户管理',
                requiresAuth: true,
            },
        },
        {
            path: 'tenant',
            name: 'Tenant',
            component: () => import('@/views/system/Tenant/index.vue'),
            meta: {
                locale: '租户管理',
                requiresAuth: true,
            },
        },
        {
            path: 'organization',
            name: 'Organization',
            component: () => import('@/views/system/Organization/index.vue'),
            meta: {
                locale: '组织架构',
                requiresAuth: true,
            },
        },
        {
            path: 'job',
            name: 'Job',
            component: () => import('@/views/system/Job/index.vue'),
            meta: {
                locale: '职位管理',
                requiresAuth: true,
            },
        },
        {
            path: 'logs',
            name: 'Logs',
            component: () => import('@/views/system/Logs/index.vue'),
            meta: {
                locale: '日志管理',
                requiresAuth: true,
            },
        },
    ],
};

export default SYSTEM;
