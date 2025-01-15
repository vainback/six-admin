import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Tenant {
    id: number;
    name: string;
    sign: string;
    domain: string;
    status: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string | null;
}

export const EmptyTenant: Tenant = {
    id: 0,
    name: '',
    sign: '',
    domain: '',
    status: 0,
}

const root = '/tenant'
export function reqTenant(action: string, data: RequestParam<Tenant> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}