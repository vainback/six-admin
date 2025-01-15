export interface RequestParam<T> {
    page?: number;
    limit?: number;
    start_time?: string;
    end_time?: string;
    keyword?: string;
    order_by?: OrderBy[];
    is_delete?: boolean; // 真实删除传true 回收站查询传true
    model: T;
}

export interface OrderBy {
    field: string;
    is_desc: boolean;
}

// export const defaultRequestParam =

export function defaultRequestParam<T>(model: T): RequestParam<T> {
    return {
        page: 1,
        limit: 15,
        start_time: '',
        end_time: '',
        keyword: '',
        order_by: [{ field: 'id', is_desc: true }],
        is_delete: false,
        model: model,
    };
}
