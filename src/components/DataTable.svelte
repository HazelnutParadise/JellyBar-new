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

    // 初始化表格
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
                searching: false,
                pageLength: 10,
                dom: 'rtip',
                destroy: true,
                responsive: true,
                autoWidth: false,
                pagingType: 'simple_numbers',
                rowCallback: function(row, data, index) {
                    jQuery(row).off('mouseenter mouseleave')
                },
                createdRow: function(row, data, dataIndex) {
                    jQuery(row).css('background-color', 'white')
                    jQuery(row).hover(
                        function() {
                            jQuery(this).css('background-color', 'white')
                        },
                        function() {
                            jQuery(this).css('background-color', 'white')
                        }
                    )
                },
                drawCallback: function() {
                    jQuery('.paginate_button').addClass('button is-small')
                    jQuery('.paginate_button.current').addClass('is-primary')
                    jQuery('.dataTables_info').addClass('has-text-grey is-size-7')
                    jQuery(this)
                        .find('tbody tr')
                        .css('background-color', 'white')
                }
            }

            const tableElement = document.getElementById(id)
            if (tableElement) {
                table = jQuery(tableElement).DataTable({
                    ...commonConfig,
                    ...config
                })
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
    /* DataTable 相關樣式 */
    :global(.dataTables_wrapper) {
        padding: 1rem 0;
    }

    :global(.dataTables_info) {
        padding: 0.5rem 0;
        color: #666;
        font-size: 0.875rem;
    }

    :global(.dataTables_paginate) {
        padding: 0.5rem 0;
    }

    :global(.paginate_button) {
        margin: 0 0.25rem;
        padding: 0.25rem 0.75rem;
        border-radius: 4px;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    :global(.paginate_button.current) {
        background-color: var(--theme-primary);
        color: white;
        border: none;
    }

    :global(.paginate_button:not(.current):hover) {
        background-color: #f5f5f5;
    }

    :global(.dataTables_empty) {
        padding: 2rem !important;
        text-align: center;
        color: #666;
    }

    :global(.dataTables_wrapper .dataTables_paginate) {
        margin-top: 1rem;
        padding-top: 0.5rem;
        border-top: 1px solid #dee2e6;
    }

    :global(.dataTables_wrapper .dataTables_info) {
        margin-top: 1rem;
        padding-top: 0.5rem;
    }
</style>

<slot /> 