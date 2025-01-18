import {AuthRule} from "@/api/system/auth-rule";
import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Job {
    id: number;
    parent_id: number;
    name: string;
    description: string;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}

export interface TreeJob extends Job{
    children?: TreeJob[];
}
const root = 'admin/job'

export const EmptyJob = {
    id: 0,
    parent_id: 0,
    name: '',
    description:'',
}

export function reqJobs(action: string, data: RequestParam<Job> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}