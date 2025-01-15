import axios from "axios";
import {getToken} from "@/utils/auth";
import {ref} from "vue";

const domain = ref(import.meta.env.VITE_API_BASE_URL)

const root = '/user/single/'

export const uploadUrl = `${domain.value}${root}upload`

export const uploadHeaders = {
    "Authorization": `Bearer ${getToken() || ''}`,
    "Six-Token": getToken() || '',
}

export function reqResetPassword(data: any) {
    return axios.post(`${root}reset/password`, data);
}