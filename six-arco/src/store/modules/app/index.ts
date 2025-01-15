import { defineStore } from 'pinia';
import { Notification } from '@arco-design/web-vue';
import type { NotificationReturn } from '@arco-design/web-vue/es/notification/interface';
import type { RouteRecordNormalized } from 'vue-router';
import defaultSettings from '@/config/settings.json';
import {getMenuList, getPermissionBtns} from '@/api/user';
import { AppState } from './types';

const useAppStore = defineStore('app', {
    state: (): AppState => ({ ...defaultSettings }),

    getters: {
        appCurrentSetting(state: AppState): AppState {
            return { ...state };
        },
        appDevice(state: AppState) {
            return state.device;
        },
        appAsyncMenus(state: AppState): RouteRecordNormalized[] {
            if (!state.serverMenu) {
                return localMenu as RouteRecordNormalized[];
            }
            return state.serverMenu as unknown as RouteRecordNormalized[];
        },
    },

    actions: {
        // Update app settings
        updateSettings(partial: Partial<AppState>) {
            // @ts-ignore-next-line
            this.$patch(partial);
        },

        // Change theme color
        toggleTheme(dark: boolean) {
            if (dark) {
                this.theme = 'dark';
                document.body.setAttribute('arco-theme', 'dark');
            } else {
                this.theme = 'light';
                document.body.removeAttribute('arco-theme');
            }
        },
        toggleDevice(device: string) {
            this.device = device;
        },
        toggleMenu(value: boolean) {
            this.hideMenu = value;
        },
        async fetchServerMenuConfig() {
            try {
                const { data } = await getMenuList();
                this.serverMenu = [...localMenu, ...data]
            } catch (error) {
            }
        },
        clearServerMenu() {
            this.serverMenu = [];
        },
    },
});

export default useAppStore;

const localMenu = [
    {
        path: '/dashboard',
        name: 'dashboard',
        meta: {
            locale: '仪表盘',
            requiresAuth: true,
            icon: 'icon-dashboard',
            order: 0,
        },
        children: [
            {
                path: 'workplace',
                name: 'Workplace',
                meta: {
                    locale: '工作台',
                    requiresAuth: true,
                },
            }
        ],
    },
];
