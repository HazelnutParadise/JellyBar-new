<script>
    import { onMount, createEventDispatcher } from 'svelte'

    export let content = ''
    export let height = 500
    export let placeholder = '請輸入文章內容...'

    const dispatch = createEventDispatcher()
    let vditor
    let editorElement

    // 添加載入資源的函數
    const loadResources = () => {
        return new Promise((resolve) => {
            // 檢查是否已載入
            if (
                document.querySelector('#vditor-css') &&
                document.querySelector('#vditor-js')
            ) {
                resolve()
                return
            }

            // 載入 CSS
            const css = document.createElement('link')
            css.id = 'vditor-css'
            css.rel = 'stylesheet'
            css.href =
                'https://cdn.jsdelivr.net/npm/vditor@3.9.6/dist/index.css'
            document.head.appendChild(css)

            // 載入 JS
            const script = document.createElement('script')
            script.id = 'vditor-js'
            script.src =
                'https://cdn.jsdelivr.net/npm/vditor@3.9.6/dist/index.min.js'
            script.onload = () => resolve()
            document.head.appendChild(script)
        })
    }

    onMount(async () => {
        // 確保資源載入
        await loadResources()

        const initEditor = () => {
            if (typeof window.Vditor === 'undefined') {
                setTimeout(initEditor, 100)
                return
            }

            vditor = new window.Vditor('vditor', {
                height,
                mode: 'wysiwyg',
                toolbar: [
                    'emoji',
                    'headings',
                    'bold',
                    'italic',
                    'strike',
                    '|',
                    'line',
                    'quote',
                    'list',
                    'ordered-list',
                    'check',
                    'code',
                    'inline-code',
                    '|',
                    'upload',
                    'link',
                    'table',
                    '|',
                    'preview',
                    'fullscreen',
                    'outline',
                    'help',
                ],
                placeholder,
                lang: 'zh_TW',
                cache: {
                    enable: false,
                },
                preview: {
                    theme: {
                        current: 'light',
                    },
                    hljs: {
                        enable: false,
                    },
                },
                counter: {
                    enable: true,
                    type: 'text',
                },
                input: (value) => {
                    dispatch('change', value)
                },
                after: () => {
                    if (content) {
                        vditor.setValue(content)
                    }
                },
            })
        }

        initEditor()

        return () => {
            if (vditor) {
                vditor.destroy()
            }
        }
    })

    export function getValue() {
        return vditor ? vditor.getValue() : ''
    }

    export function setValue(value) {
        if (vditor) {
            vditor.setValue(value)
        }
    }
</script>

<style>
    :global(.vditor) {
        border: 1px solid #ddd;
        border-radius: 4px;
    }

    :global(.vditor-toolbar) {
        background-color: #f6f8fa;
        border-bottom: 1px solid #ddd;
        padding: 5px;
    }

    :global(.vditor-toolbar button) {
        all: unset;
        cursor: pointer;
        padding: 6px !important;
        margin: 0 2px !important;
        border-radius: 3px !important;
        display: inline-flex !important;
        align-items: center !important;
        justify-content: center !important;
        width: 28px !important;
        height: 28px !important;
        box-sizing: border-box !important;
        position: relative !important;
    }

    :global(.vditor-toolbar button:hover) {
        background-color: #e1e4e8 !important;
    }

    :global(.vditor-toolbar button.vditor-menu--current) {
        background-color: #e1e4e8 !important;
    }

    :global(.vditor-toolbar svg) {
        width: 16px !important;
        height: 16px !important;
        fill: currentColor !important;
        display: inline-block !important;
        vertical-align: middle !important;
    }

    :global(.vditor-toolbar .vditor-icon) {
        color: #666 !important;
        display: inline-block !important;
        height: 16px !important;
        width: 16px !important;
        text-align: center !important;
        line-height: 1 !important;
    }

    :global(.vditor-toolbar .vditor-icon:before) {
        font-size: 16px !important;
    }

    :global(.vditor-reset) {
        font-size: 16px !important;
    }

    :global(.vditor-wysiwyg) {
        padding: 10px 15px !important;
    }

    /* 修復分隔線 */
    :global(.vditor-toolbar__divider) {
        height: 20px !important;
        width: 1px !important;
        margin: 0 5px !important;
        background-color: #ddd !important;
        display: inline-block !important;
        vertical-align: middle !important;
    }
</style>

<div id="vditor" class="vditor" bind:this={editorElement}></div>
