<template>
    <div class="login-form-wrapper">
        <div class="login-form-title">登录 Six-Admin</div>
        <!--        <div class="login-form-sub-title">{{ $t('login.form.title') }}</div>-->
        <div class="login-form-error-msg">{{ errorMessage }}</div>
        <a-form
            ref="loginForm"
            :model="userInfo"
            class="login-form"
            layout="vertical"
            @submit="handleSubmit"
        >
            <a-form-item
                field="tenant_id"
                :rules="[{ required: true, message: '必须选择租户' }]"
                :validate-trigger="['change', 'blur']"
                hide-label
            >
                <a-select v-model="userInfo.tenant_id">
                    <template #prefix>
                        <icon-user-group />
                    </template>
                    <a-option
                        v-for="item in tenant_list"
                        :label="item.name"
                        :value="item.id"
                    ></a-option>
                </a-select>
            </a-form-item>
            <a-form-item
                field="username"
                :rules="[
                    {
                        required: true,
                        message: $t('login.form.userName.errMsg'),
                    },
                ]"
                :validate-trigger="['change', 'blur']"
                hide-label
            >
                <a-input
                    v-model="userInfo.username"
                    :placeholder="$t('login.form.userName.placeholder')"
                >
                    <template #prefix>
                        <icon-user />
                    </template>
                </a-input>
            </a-form-item>
            <a-form-item
                field="password"
                :rules="[
                    {
                        required: true,
                        message: $t('login.form.password.errMsg'),
                    },
                ]"
                :validate-trigger="['change', 'blur']"
                hide-label
            >
                <a-input-password
                    v-model="userInfo.password"
                    :placeholder="$t('login.form.password.placeholder')"
                    allow-clear
                >
                    <template #prefix>
                        <icon-lock />
                    </template>
                </a-input-password>
            </a-form-item>
            <a-space :size="16" direction="vertical">
                <div class="login-form-password-actions">
                    <a-checkbox
                        checked="rememberPassword"
                        :model-value="rememberPassword"
                        @change="setRememberPassword as any"
                    >
                        {{ $t('login.form.rememberPassword') }}
                    </a-checkbox>
                    <a-link>{{ $t('login.form.forgetPassword') }}</a-link>
                </div>
                <a-button
                    type="primary"
                    html-type="submit"
                    long
                    :loading="loading"
                >
                    {{ $t('login.form.login') }}
                </a-button>
                <div style="height: 50px"></div>
            </a-space>
        </a-form>
    </div>
</template>

<script lang="ts" setup>
    import { ref, reactive, onMounted } from 'vue';
    import { useRouter } from 'vue-router';
    import { Message } from '@arco-design/web-vue';
    import { ValidatedError } from '@arco-design/web-vue/es/form/interface';
    import { useI18n } from 'vue-i18n';
    import { useStorage } from '@vueuse/core';
    import { useUserStore } from '@/store';
    import useLoading from '@/hooks/loading';
    import { login, LoginData, tenant } from '@/api/system/user';
    import { Tenant } from '@/api/system/tenant';
    import { setToken } from '@/utils/auth';
import { RoleType, UserState } from '@/store/modules/user/types';

    const router = useRouter();
    const { t } = useI18n();
    const errorMessage = ref('');
    const { loading, setLoading } = useLoading();
    const userStore = useUserStore();

    const rememberPassword = ref(
        !!localStorage.getItem('rememberPassword') || true
    );
    const userInfo = reactive({
        tenant_id: Number(localStorage.getItem('login.tenant_id')) || 1,
        username: localStorage.getItem('login.username') || 'admin',
        password: localStorage.getItem('login.password') || '123456',
    });

    const tenant_list = ref<Tenant[]>([]);
    onMounted(() => {
        getTenant();
    });
    const getTenant = async () => {
        try {
            const res = await tenant();
            tenant_list.value = res.data;
        } catch (err) {
            console.log(err);
        } finally {
        }
    };

    const handleSubmit = async ({
        errors,
        values,
    }: {
        errors: Record<string, ValidatedError> | undefined;
        values: Record<string, any>;
    }) => {
        if (loading.value) return;
        if (!errors) {
            setLoading(true);
            try {
                const res = await login({model: values as LoginData});
                setToken(res.data.token);

                localStorage.setItem("userinfo", JSON.stringify(res.data.userinfo));
                localStorage.setItem("userRole", res.data.userinfo.role_ids.map(item => item.toString()));
                Message.success(t('login.form.login.success'));
                const { redirect, ...othersQuery } = router.currentRoute.value.query;
                await router.push({
                    name: (redirect as string) || 'Workplace',
                    query: {
                        ...othersQuery,
                    },
                });
                const { tenant_id, username, password } = values;
                // 实际生产环境需要进行加密存储。
                if (rememberPassword.value) {
                    localStorage.setItem('login.tenant_id', tenant_id);
                    localStorage.setItem('login.username', username);
                    localStorage.setItem('login.password', password);
                }
            } catch (err) {
                errorMessage.value = (err as Error).message;
            } finally {
                setLoading(false);
            }
        }
    };
    const setRememberPassword = (value: boolean) => {
        // loginConfig.value.rememberPassword = value;
        rememberPassword.value = value;
        localStorage.setItem('rememberPassword', value + '');
    };
</script>

<style lang="less" scoped>
    .login-form {
        &-wrapper {
            width: 320px;
        }

        &-title {
            color: var(--color-text-1);
            font-weight: 500;
            font-size: 24px;
            line-height: 32px;
        }

        &-sub-title {
            color: var(--color-text-3);
            font-size: 16px;
            line-height: 24px;
        }

        &-error-msg {
            height: 32px;
            color: rgb(var(--red-6));
            line-height: 32px;
        }

        &-password-actions {
            display: flex;
            justify-content: space-between;
        }

        &-register-btn {
            color: var(--color-text-3) !important;
        }
    }
</style>
