<script lang="ts">
    import { onMount, onDestroy } from 'svelte'

    export let id: string
    export let config: any = {}
    
    let table: any
    let scriptsLoaded = false

    // 載入必要的腳本和樣式
    const loadDependencies = async () => {
        if (scriptsLoaded) return

        // 載入 jQuery
        if (!(window as any).jQuery) {
            const jqueryScript = document.createElement('script')
            jqueryScript.src = 'https://code.jquery.com/jquery-3.7.0.min.js'
            document.head.appendChild(jqueryScript)
            await new Promise((resolve) => jqueryScript.onload = resolve)
        }

        // 載入 DataTables CSS
        if (!document.querySelector('link[href*="datatables"]')) {
            const datatablesCss = document.createElement('link')
            datatablesCss.rel = 'stylesheet'
            datatablesCss.type = 'text/css'
            datatablesCss.href = 'https://cdn.datatables.net/1.13.7/css/jquery.dataTables.css'
            document.head.appendChild(datatablesCss)

            const bulmaDataTablesCss = document.createElement('link')
            bulmaDataTablesCss.rel = 'stylesheet'
            bulmaDataTablesCss.type = 'text/css'
            bulmaDataTablesCss.href = 'https://cdn.datatables.net/1.13.7/css/dataTables.bulma.min.css'
            document.head.appendChild(bulmaDataTablesCss)
        }

        // 載入 DataTables JS
        if (!(window as any).jQuery?.fn?.DataTable) {
            const dataTablesScript = document.createElement('script')
            dataTablesScript.src = 'https://cdn.datatables.net/1.13.7/js/jquery.dataTables.min.js'
            document.head.appendChild(dataTablesScript)
            await new Promise((resolve) => dataTablesScript.onload = resolve)

            const bulmaDataTablesScript = document.createElement('script')
            bulmaDataTablesScript.src = 'https://cdn.datatables.net/1.13.7/js/dataTables.bulma.min.js'
            document.head.appendChild(bulmaDataTablesScript)
            await new Promise((resolve) => bulmaDataTablesScript.onload = resolve)
        }

        scriptsLoaded = true
    }

    // 修改初始化表格的函數
    const initTable = async () => {
        await loadDependencies()

        const jQuery = (window as any).jQuery
        if (!jQuery?.fn?.DataTable) {
            console.error('DataTables 未載入')
            return
        }

        try {
            if (table) {
                table.destroy()
                table = null
            }

            await new Promise((resolve) => setTimeout(resolve, 0))

            // 統一的 DataTable 配置
            const commonConfig = {
                language: {
                    url: '//cdn.datatables.net/plug-ins/1.13.7/i18n/zh-HANT.json',
                },
                searching: true,
                pageLength: 10,
                dom: '<"table-controls"lf>rtip',
                destroy: true,
                responsive: true,
                autoWidth: false,
                pagingType: 'simple_numbers',
                createdRow: null,
                rowCallback: null,
                drawCallback: function(settings) {
                    const api = this.api();
                    if (config.drawCallback) {
                        config.drawCallback.call(this, settings);
                    }
                }
            }

            const tableElement = document.getElementById(id)
            if (tableElement) {
                table = jQuery(tableElement).DataTable({
                    ...commonConfig,
                    ...config
                })

                // 發出事件通知父組件 DataTable 實例已創建
                window.dispatchEvent(new CustomEvent('datatableCreated', {
                    detail: {
                        id: id,
                        instance: table
                    }
                }));
            }
        } catch (error) {
            console.error('初始化表格時發生錯誤:', error)
        }
    }

    onMount(async () => {
        await new Promise((resolve) => setTimeout(resolve, 500))
        await initTable()
    })

    onDestroy(() => {
        if (table) {
            table.destroy()
            table = null
        }
    })
</script>

<style>
    /* 基礎表格樣式 */
    :global(table.display) {
        width: 100% !important;
        border-collapse: separate !important;
        border-spacing: 0;
        background-color: white;
    }

    /* 表格控制區域樣式 */
    :global(.table-controls) {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 1rem;
    }

    /* 表頭樣式 */
    :global(table.display thead th) {
        background-color: #f8fafc !important;
        color: #1e293b !important;
        font-weight: 600 !important;
        padding: 1rem !important;
        text-align: left !important;
        border-bottom: 2px solid #e2e8f0 !important;
        white-space: nowrap;
    }

    /* 表格內容樣式 */
    :global(table.display tbody td) {
        padding: 1rem !important;
        border-bottom: 1px solid #e2e8f0 !important;
        color: #475569 !important;
        background-color: white !important;
    }

    /* 表格行樣式 */
    :global(table.display tbody tr) {
        background-color: white !important;
    }

    :global(table.display tbody tr:hover) {
        background-color: #f8fafc !important;
    }

    :global(table.display tbody tr:hover td) {
        background-color: #f8fafc !important;
    }

    /* 分頁控制樣式 */
    :global(.dataTables_wrapper .dataTables_paginate) {
        margin-top: 1rem !important;
        padding: 1rem 0 !important;
    }

    :global(.dataTables_wrapper .dataTables_paginate .paginate_button) {
        padding: 0.5rem 1rem !important;
        margin: 0 0.25rem !important;
        border-radius: 0.375rem !important;
        border: 1px solid #e2e8f0 !important;
        background: white !important;
        color: #475569 !important;
        box-shadow: none !important;
    }

    :global(.dataTables_wrapper .dataTables_paginate .paginate_button:hover) {
        background: #f1f5f9 !important;
        color: #1e293b !important;
        border-color: #e2e8f0 !important;
    }

    :global(.dataTables_wrapper .dataTables_paginate .paginate_button.current) {
        background: #3b82f6 !important;
        border-color: #3b82f6 !important;
        color: white !important;
    }

    :global(.dataTables_wrapper .dataTables_paginate .paginate_button.current:hover) {
        background: #2563eb !important;
        color: white !important;
    }

    :global(.dataTables_wrapper .dataTables_paginate .paginate_button.disabled) {
        opacity: 0.5 !important;
        cursor: not-allowed !important;
        background: #f1f5f9 !important;
    }

    /* 搜尋框樣式 */
    :global(.dataTables_wrapper .dataTables_filter) {
        margin-bottom: 0 !important;
    }

    :global(.dataTables_wrapper .dataTables_filter input) {
        padding: 0.5rem !important;
        border: 1px solid #e2e8f0 !important;
        border-radius: 0.375rem !important;
        margin-left: 0.5rem !important;
        outline: none !important;
        min-width: 200px !important;
    }

    :global(.dataTables_wrapper .dataTables_filter input:focus) {
        border-color: #3b82f6 !important;
        box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1) !important;
    }

    /* 表格資訊樣式 */
    :global(.dataTables_wrapper .dataTables_info) {
        padding: 1rem 0 !important;
        color: #64748b !important;
        clear: both !important;
    }

    /* 每頁顯示數量選擇器樣式 */
    :global(.dataTables_wrapper .dataTables_length) {
        margin-bottom: 0 !important;
        white-space: nowrap !important;
    }

    :global(.dataTables_wrapper .dataTables_length select) {
        padding: 0.5rem !important;
        border: 1px solid #e2e8f0 !important;
        border-radius: 0.375rem !important;
        margin: 0 0.5rem !important;
        outline: none !important;
        background-color: white !important;
        min-width: 70px !important;
    }

    /* 按鈕樣式覆蓋 */
    :global(.edit-role-btn),
    :global(.toggle-status-btn),
    :global(.delete-user-btn) {
        background: transparent !important;
        border: none !important;
        padding: 0.5rem !important;
        cursor: pointer !important;
        display: inline-flex !important;
        align-items: center !important;
        justify-content: center !important;
    }

    :global(.edit-role-btn:hover),
    :global(.toggle-status-btn:hover),
    :global(.delete-user-btn:hover) {
        background: #f8fafc !important;
        border-radius: 0.375rem !important;
    }

    /* 確保操作列的按鈕容器樣式 */
    :global(.action-buttons) {
        display: flex !important;
        gap: 0.5rem !important;
        justify-content: center !important;
        align-items: center !important;
    }

    /* SVG 圖標樣式 */
    :global(button svg) {
        width: 20px !important;
        height: 20px !important;
    }
</style>

<slot /> 