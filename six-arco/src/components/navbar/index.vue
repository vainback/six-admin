<template>
    <div class="navbar">
        <div class="left-side">
            <a-space>
                <img
                    alt="logo"
                    src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
                />
                <a-typography-title
                    :style="{ margin: 0, fontSize: '18px' }"
                    :heading="5"
                >
                    Six Admin
                </a-typography-title>
                <icon-menu-fold
                    v-if="!topMenu && appStore.device === 'mobile'"
                    style="font-size: 22px; cursor: pointer"
                    @click="toggleDrawerMenu"
                />
            </a-space>
        </div>
        <div class="center-side">
            <Menu v-if="topMenu" />
        </div>
        <ul class="right-side">
            <!-- <li>
                <a-tooltip :content="$t('settings.search')">
                    <a-button class="nav-btn" type="outline" :shape="'circle'">
                        <template #icon>
                            <icon-search />
                        </template>
                    </a-button>
                </a-tooltip>
            </li> -->
            <!-- <li>
                <a-tooltip :content="$t('settings.language')">
                    <a-button
                        class="nav-btn"
                        type="outline"
                        :shape="'circle'"
                        @click="setDropDownVisible"
                    >
                        <template #icon>
                            <icon-language />
                        </template>
                    </a-button>
                </a-tooltip>
                <a-dropdown trigger="click" @select="changeLocale as any">
                    <div ref="triggerBtn" class="trigger-btn"></div>
                    <template #content>
                        <a-doption
                            v-for="item in locales"
                            :key="item.value"
                            :value="item.value"
                        >
                            <template #icon>
                                <icon-check
                                    v-show="item.value === currentLocale"
                                />
                            </template>
                            {{ item.label }}
                        </a-doption>
                    </template>
                </a-dropdown>
            </li> -->
            <li>
                <a-tooltip :content="$t('settings.navbar.alerts')">
                    <div class="message-box-trigger">
                        <a-badge :count="9" dot>
                            <a-button
                                class="nav-btn"
                                type="outline"
                                :shape="'circle'"
                                @click="setPopoverVisible"
                            >
                                <icon-notification />
                            </a-button>
                        </a-badge>
                    </div>
                </a-tooltip>
                <a-popover
                    trigger="click"
                    :arrow-style="{ display: 'none' }"
                    :content-style="{ padding: 0, minWidth: '400px' }"
                    content-class="message-popover"
                >
                    <div ref="refBtn" class="ref-btn"></div>
                    <template #content>
                        <message-box />
                    </template>
                </a-popover>
            </li>
            <li>
                <a-tooltip
                    :content="
                        theme === 'light'
                            ? $t('settings.navbar.theme.toDark')
                            : $t('settings.navbar.theme.toLight')
                    "
                >
                    <a-button
                        class="nav-btn"
                        type="outline"
                        :shape="'circle'"
                        @click="handleToggleTheme"
                    >
                        <template #icon>
                            <icon-moon-fill v-if="theme === 'dark'" />
                            <icon-sun-fill v-else />
                        </template>
                    </a-button>
                </a-tooltip>
            </li>

            <li>
                <a-tooltip
                    :content="
                        isFullscreen
                            ? $t('settings.navbar.screen.toExit')
                            : $t('settings.navbar.screen.toFull')
                    "
                >
                    <a-button
                        class="nav-btn"
                        type="outline"
                        :shape="'circle'"
                        @click="toggleFullScreen"
                    >
                        <template #icon>
                            <icon-fullscreen-exit v-if="isFullscreen" />
                            <icon-fullscreen v-else />
                        </template>
                    </a-button>
                </a-tooltip>
            </li>
            <li>
                <a-tooltip :content="$t('settings.title')">
                    <a-button
                        class="nav-btn"
                        type="outline"
                        :shape="'circle'"
                        @click="setVisible"
                    >
                        <template #icon>
                            <icon-settings />
                        </template>
                    </a-button>
                </a-tooltip>
            </li>
            <li>
                <a-dropdown trigger="click">
                    <a-avatar :size="28" style="font-size: 18px;cursor: pointer;background-color: #14a9f8">{{userinfo.nickname.substring(0, 1)}}</a-avatar>
                    <template #content>
                        <!-- <a-doption>
                            <a-space @click="switchRoles">
                                <icon-tag />
                                <span>
                                    {{ $t('messageBox.switchRoles') }}
                                </span>
                            </a-space>
                        </a-doption>
                        <a-doption>
                            <a-space @click="$router.push({ name: 'Info' })">
                                <icon-user />
                                <span>
                                    {{ $t('messageBox.userCenter') }}
                                </span>
                            </a-space>
                        </a-doption>
                        <a-doption>
                            <a-space @click="$router.push({ name: 'Setting' })">
                                <icon-settings />
                                <span>
                                    {{ $t('messageBox.userSettings') }}
                                </span>
                            </a-space>
                        </a-doption> -->
                        <a-doption>
                            <a-space @click="resetPassword.visible = true">
                                <icon-edit />
                                <span>
                                    {{ '修改密码' }}
                                </span>
                            </a-space>
                        </a-doption>
                        <a-doption>
                            <a-space @click="handleLogout">
                                <icon-export />
                                <span>
                                    {{ $t('messageBox.logout') }}
                                </span>
                            </a-space>
                        </a-doption>
                    </template>
                </a-dropdown>
            </li>
        </ul>
        <a-modal
                title="修改密码"
                width="500px"
                v-model:visible="resetPassword.visible"
                @cancel="resetPassword.visible = false"
                @ok="resetPasswordSubmit"
        >
            <a-form :model="resetPassword.form" :size="size" label-align="right" auto-label-width>
                <a-form-item field="old_password" label="原密码" validate-trigger="blur" :rules="[{required: true, message: '原密码必须填写'}]">
                    <a-input-password
                            v-model="resetPassword.form.old_password"
                            placeholder="请填写原密码"
                    />
                </a-form-item>
                <a-form-item field="new_password" label="新密码" validate-trigger="blur" :rules="[{required: true, message: '新密码必须填写'}]">
                    <a-input-password
                            v-model="resetPassword.form.new_password"
                            placeholder="请填写新密码"
                    />
                </a-form-item>
                <a-form-item field="confirm_pass" label="确认密码" validate-trigger="blur" :rules="[{required: true, message: '确认密码必须填写'}]">
                    <a-input-password
                            v-model="resetPassword.form.confirm_pass"
                            placeholder="请填写确认密码且与新密码一致"
                    />
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script lang="ts" setup>
    import { computed, ref, inject } from 'vue';
    import { Message, Modal } from '@arco-design/web-vue';
    import { useDark, useToggle, useFullscreen } from '@vueuse/core';
    import { useAppStore, useUserStore } from '@/store';
    import { LOCALE_OPTIONS } from '@/locale';
    import useLocale from '@/hooks/locale';
    import useUser from '@/hooks/user';
    import Menu from '@/components/menu/index.vue';
    import MessageBox from '../message-box/index.vue';
    import {Userinfo} from "@/api/system/user";
    import user from "../../store/modules/user";
    import {isTenantRoot} from "@/utils/auth";
import { reqResetPassword } from '@/api/user-single';
    const size = ref(import.meta.env.VITE_STYLE_SIZE);
    const appStore = useAppStore();
    const userStore = useUserStore();
    const { logout } = useUser();
    const { changeLocale, currentLocale } = useLocale();
    const { isFullscreen, toggle: toggleFullScreen } = useFullscreen();
    const locales = [...LOCALE_OPTIONS];
    // const avatar = computed(() => {
    //     return userStore.avatar;
    // });
    const theme = computed(() => {
        return appStore.theme;
    });
    const topMenu = computed(() => appStore.topMenu && appStore.menu);
    const isDark = useDark({
        selector: 'body',
        attribute: 'arco-theme',
        valueDark: 'dark',
        valueLight: 'light',
        storageKey: 'arco-theme',
        onChanged(dark: boolean) {
            // overridden default behavior
            appStore.toggleTheme(dark);
        },
    });
    const toggleTheme = useToggle(isDark);
    const handleToggleTheme = () => {
        toggleTheme();
    };
    const setVisible = () => {
        appStore.updateSettings({ globalSettings: true });
    };
    const refBtn = ref();
    const triggerBtn = ref();
    const setPopoverVisible = () => {
        const event = new MouseEvent('click', {
            view: window,
            bubbles: true,
            cancelable: true,
        });
        refBtn.value.dispatchEvent(event);
    };
    const handleLogout = () => {
        logout();
    };
    const setDropDownVisible = () => {
        const event = new MouseEvent('click', {
            view: window,
            bubbles: true,
            cancelable: true,
        });
        triggerBtn.value.dispatchEvent(event);
    };
    const switchRoles = async () => {
        const res = await userStore.switchRoles();
        Message.success(res as string);
    };
    const toggleDrawerMenu = inject('toggleDrawerMenu') as () => void;

    const userinfo = ref<Userinfo>(JSON.parse(localStorage.getItem('userinfo') || '') as Userinfo);
    
    const resetPassword = ref({
        visible: false,
        form: {
            old_password: '', // 原密码
            new_password: '', // 新密码
            confirm_pass: '' // 确认密码
        }
    })

    const resetPasswordSubmit = async () => {
        try {
            const res = await reqResetPassword(resetPassword.value.form)
            Message.success("修改成功")
            Modal.error({
                    title: '登录状态失效',
                    content: '您的登录状态已失效，请重新登录',
                    okText: '重新登陆',
                    async onOk() {
                        const userStore = useUserStore();
                        await userStore.logout();
                        window.location.reload();
                    },
                });
        } catch(e) {

        } finally {
            resetPassword.value = {
                visible: false,
                form: {
                    old_password: '', // 原密码
                    new_password: '', // 新密码
                    confirm_pass: '' // 确认密码
                } 
            }
        }
    }
</script>

<style scoped lang="less">
    .navbar {
        display: flex;
        justify-content: space-between;
        height: 100%;
        background-color: var(--color-bg-2);
        border-bottom: 1px solid var(--color-border);
    }

    .left-side {
        display: flex;
        align-items: center;
        padding-left: 20px;
    }

    .center-side {
        flex: 1;
    }

    .right-side {
        display: flex;
        padding-right: 20px;
        list-style: none;

        :deep(.locale-select) {
            border-radius: 20px;
        }

        li {
            display: flex;
            align-items: center;
            padding: 0 10px;
        }

        a {
            color: var(--color-text-1);
            text-decoration: none;
        }

        .nav-btn {
            border-color: rgb(var(--gray-2));
            color: rgb(var(--gray-8));
            font-size: 16px;
        }

        .trigger-btn,
        .ref-btn {
            position: absolute;
            bottom: 14px;
        }

        .trigger-btn {
            margin-left: 14px;
        }
    }
</style>

<style lang="less">
    .message-popover {
        .arco-popover-content {
            margin-top: 0;
        }
    }
</style>
