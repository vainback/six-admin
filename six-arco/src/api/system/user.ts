import axios from "axios";
import user from "@/store/modules/user";
import {Tenant} from "@/api/system/tenant";
import {RequestParam} from "@/api/base";
import {AuthRule} from "@/api/system/auth-rule";

export interface LoginData {
    tenant_id: number;
    username: string;
    password: string;
}

export interface LoginRes {
    token: string;
    userinfo: Userinfo;
}

export interface Userinfo {
    id: number;
    username: string;
    nickname: string;
    password: string;
    status: number;
    role_ids: number[] | null;
    is_root: boolean | null;
    organization_id: number | null;
    job_id: number | null;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}

export const EmptyUser: Userinfo = {
    id: 0,
    username: '',
    nickname: '',
    password: '',
    status: 0,
    role_ids: null,
    is_root: null,
    organization_id: null,
    job_id: null,
    tenant_id: 0,
}

const userSingle = '/user/single'

export function login(data: RequestParam<LoginData>) {
    return axios.post<LoginRes>(`${userSingle}/login`, data);
}

export function tenant() {
    return axios.post(`${userSingle}/tenant`);
}

const root = '/auth/user'
export function reqUser(action: string, data: RequestParam<Userinfo> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}
