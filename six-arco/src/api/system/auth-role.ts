import { RequestParam } from '@/api/base';
import axios from 'axios';
import {EmptyJob} from "@/api/system/job";

export interface AuthRole {
    id: number;
    parent_id: number;
    title: string;
    sign: string;
    status: number;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}

export interface TreeAuthRole extends AuthRole {
    children?: TreeAuthRole[];
}

export const EmptyRole : AuthRole = {
    id: 0,
    parent_id: 0,
    title: '',
    sign: '',
    status: 0,
}

const root = 'admin/auth/role';

export function reqAuthRole(action: string, data: RequestParam<AuthRole> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}
