import { RouteLocationNormalized, RouteRecordRaw } from 'vue-router';
import { useUserStore } from '@/store';

export default function usePermission() {
    const userStore = useUserStore();
    return {
        accessRouter(route: RouteLocationNormalized | RouteRecordRaw) {
            let ok: boolean = false
            ok = !route.meta?.requiresAuth || !route.meta?.roles || route.meta?.roles?.includes('*')
            if (ok) return ok;

            // const userRoles = userStore.role.split(',')
            const userRoles = localStorage.getItem("userRole")?.split(',') || []
            for (const role of userRoles) {
                if (route.meta?.roles?.includes(role)) {
                    return true;
                }
            }
            return false;
        },
        findFirstPermissionRoute(_routers: any, role = 'admin') {
            const cloneRouters = [..._routers];
            while (cloneRouters.length) {
                const firstElement = cloneRouters.shift();
                if (
                    firstElement?.meta?.roles?.find((el: string[]) => {
                        return el.includes('*') || el.includes(role);
                    })
                )
                    return { name: firstElement.name };
                if (firstElement?.children) {
                    cloneRouters.push(...firstElement.children);
                }
            }
            return null;
        },
        // You can add any rules you want
    };
}
