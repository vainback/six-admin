import {RequestParam} from "@/api/base";
import axios from "axios";

export interface AuthRelation {
    id: number;
    rule_id: number;
    role_id: number;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}
const root = '/auth/relation'
export function getCheckedRule(data: RequestParam<AuthRelation>) {
    return axios.post(`${root}/select/rule`, data);
}

export function set(role_id: number, rule_ids: number[]) {
    return axios.post(`${root}/set`, {role_id: role_id, rule_ids: rule_ids});
}