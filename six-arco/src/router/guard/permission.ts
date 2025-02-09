import type { Router, RouteRecordNormalized } from 'vue-router';
import NProgress from 'nprogress'; // progress bar

import usePermission from '@/hooks/permission';
import { useUserStore, useAppStore } from '@/store';
import { appRoutes } from '../routes';
import { WHITE_LIST, NOT_FOUND } from '../constants';
import {getPermissionBtns} from "@/api/user";

export default function setupPermissionGuard(router: Router) {
    router.beforeEach(async (to, from, next) => {
        const appStore = useAppStore();
        const userStore = useUserStore();
        const Permission = usePermission();
        const permissionsAllow = Permission.accessRouter(to);
        if (appStore.menuFromServer) {
            // if (
            //     !appStore.appAsyncMenus.length &&
            //     !WHITE_LIST.find((el) => el.name === to.name)
            // ) {
            //     await appStore.fetchServerMenuConfig();
            // }


            if (!WHITE_LIST.find((el) => el.name === to.name)) {

                const {data: btns} = await getPermissionBtns();
                localStorage.setItem("permission:btns", JSON.stringify(btns));

                await appStore.fetchServerMenuConfig();
            }
            const serverMenuConfig = [...appStore.appAsyncMenus, ...WHITE_LIST];
            let exist = false;
            while (serverMenuConfig.length && !exist) {
                // 拿出第一个元素的值
                const element = serverMenuConfig.shift();
                if (element?.name === to.name) exist = true;

                if (element?.children) {
                    serverMenuConfig.push(
                        ...(element.children as unknown as RouteRecordNormalized[])
                    );
                }
            }
            if (exist && permissionsAllow) {
                next();
            } else {
                next(NOT_FOUND)
            }
        } else {
            if (permissionsAllow) next();
            else {
                const destination =
                    Permission.findFirstPermissionRoute(
                        appRoutes,
                        ""
                    ) || NOT_FOUND;
                next(destination);
            }
        }
        NProgress.done();
    });
}
