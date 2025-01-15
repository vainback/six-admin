import {AuthRule} from "@/api/system/auth-rule";
import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Organization {
    id: number;
    parent_id: number;
    name: string;
    description: string;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}

 export interface TreeOrganization extends Organization{
    children?: TreeOrganization[];
}
const root = '/organization'

export const EmptyOrganization = {
    id: 0,
    parent_id: 0,
    name: '',
    description:'',
}

export function reqOrgs(action: string, data: RequestParam<Organization> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}