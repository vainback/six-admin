import axios from "axios";
import {RequestParam} from "@/api/base";

export interface TreeAuthRule extends AuthRule{
    children?: TreeAuthRule[];
}

export interface AuthRule {
    id: number;
    parent_id: number;
    type: string;
    component: string;
    name: string;
    path: string;
    authPrefix?: string; // 兼容前端样式 后端不需要的字段
    auth: string;
    api_route: string;
    locale: string;
    icon: string;
    order: number;
    status: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}


export const EmptyRule: AuthRule = {
    id: 0,
    parent_id: 0,
    type: '',
    component: '',
    name: '',
    path: '',
    auth: '',
    api_route: '',
    locale: '',
    icon: '',
    order: 0,
    status: 0,
}

const root = '/auth/rule'

export function reqAuthRule(action: string, data: RequestParam<AuthRule> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}

