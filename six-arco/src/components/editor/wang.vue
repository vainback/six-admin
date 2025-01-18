<template>
    <div>
        <div style="border: 1px solid #ccc; margin-top: 10px">
            <Toolbar
                :editor="editorRef"
                :defaultConfig="toolbarConfig"
                mode="default"
                style="border-bottom: 1px solid #ccc"
            />
            <Editor
                :defaultConfig="editorConfig"
                mode="default"
                v-model="valueHtml"
                style="height: 400px; overflow-y: hidden"
                @onCreated="handleCreated"
                @onChange="handleChange"
                @onDestroyed="handleDestroyed"
            />
        </div>
    </div>
</template>
<script setup>
    import '@wangeditor/editor/dist/css/style.css';
    import { onBeforeUnmount, ref, shallowRef, onMounted } from 'vue';
    import { Editor, Toolbar } from '@wangeditor/editor-for-vue';
    import { defineEmits } from 'vue';

    const emit = defineEmits(['change']);
    // 编辑器实例，必须用 shallowRef，重要！
    const editorRef = shallowRef();

    // 内容 HTML
    const valueHtml = ref('');

    // 模拟 ajax 异步获取内容
    onMounted(() => {
        valueHtml.value = '';
    });

    const toolbarConfig = {};
    const editorConfig = { placeholder: '请输入内容...' };

    // 组件销毁时，也及时销毁编辑器，重要！
    onBeforeUnmount(() => {
        const editor = editorRef.value;
        if (editor == null) return;

        editor.destroy();
    });

    // 编辑器回调函数
    const handleCreated = (editor) => {
        console.log('created', editor);
        editorRef.value = editor; // 记录 editor 实例，重要！
    };
    const handleChange = (editor) => {
        emit('change', editor.getHtml());
    };
    const handleDestroyed = (editor) => {
        console.log('destroyed', editor);
    };
</script>
