import { DirectiveBinding } from 'vue';
import { useUserStore } from '@/store';

function checkPermission(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding;
    if ((value as String).trim() != '') {
        const hasBtns: string[] = JSON.parse(localStorage.getItem('permission:btns')) || [];
        if (!hasBtns.includes(value) && el.parentNode) {
            el.parentNode.removeChild(el);
        }
    }
}

export default {
    mounted(el: HTMLElement, binding: DirectiveBinding) {
        checkPermission(el, binding);
    },
    updated(el: HTMLElement, binding: DirectiveBinding) {
        checkPermission(el, binding);
    },
};
