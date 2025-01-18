import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Codegen {
    id: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    table: string;
    title: string;
    fields: any[] | null;
    parent_module: number | null;
    is_soft_delete: boolean;
    is_tenant: boolean;
}

export const EmptyCodegen: Codegen = {
    id: 0,
    table: '',
    title: '',
    fields: [],
    parent_module: null,
    is_soft_delete: true,
    is_tenant: true,
}

const root = 'admin/codegen'
export function reqCodegen(action: string, data: RequestParam<Codegen> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}