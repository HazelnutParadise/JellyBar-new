<script>
    import '../app.css'
    import { onMount, onDestroy } from 'svelte'
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    import setTitle from '../js/setTitle'
    import VditorEditor from '../components/VditorEditor.svelte'

    export let title
    export let siteName
    setTitle(title, siteName)

    let newCategory = ''

    // 假資料
    let articles = [
        {
            id: 1,
            title: 'Rust繁中簡學！',
            created_at: '2024-03-20',
            updated_at: '2024-05-21',
            category: 'Rust',
        },
        {
            id: 2,
            title: 'Web開發教學',
            created_at: '2024-03-21',
            updated_at: '2024-03-21',
            category: 'Web',
        },
    ]

    let categories = [
        { id: 1, name: 'Rust', articleCount: 1 },
        { id: 2, name: 'Web', articleCount: 1 },
        { id: 3, name: 'Python', articleCount: 0 },
    ]

    let editingCategory = null
    let editingCategoryName = ''

    const handleEditCategory = async (category) => {
        editingCategory = category
        editingCategoryName = category.name
    }

    const handleSaveCategory = async () => {
        if (!editingCategoryName.trim()) {
            alert('類別名稱不能為空！')
            return
        }

        if (
            categories.some(
                (c) =>
                    c.name === editingCategoryName.trim() &&
                    c.id !== editingCategory.id,
            )
        ) {
            alert('類別名稱已存在！')
            return
        }

        try {
            // 發送更新請求到後端 API
            const response = await fetch(
                `/api/categories/${editingCategory.id}`,
                {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ name: editingCategoryName.trim() }),
                },
            )

            if (!response.ok) {
                throw new Error('更新類別失敗')
            }

            categories = categories.map((c) =>
                c.id === editingCategory.id
                    ? { ...c, name: editingCategoryName.trim() }
                    : c,
            )

            // 更新相關文章的類別名稱
            articles = articles.map((article) =>
                article.category === editingCategory.name
                    ? { ...article, category: editingCategoryName.trim() }
                    : article,
            )

            editingCategory = null
            editingCategoryName = ''
            alert('類別更新成功！')
        } catch (error) {
            console.error('更新類別時發生錯誤:', error)
            alert('更新類別失敗：' + error.message)
        }
    }

    const handleDeleteCategory = async (id) => {
        const category = categories.find((c) => c.id === id)

        if (category.articleCount > 0) {
            const confirmMessage = `警告：「${category.name}」類別下還有 ${category.articleCount} 篇文章！\n確定要刪除嗎？`
            if (!confirm(confirmMessage)) return

            const secondConfirm = `最後確認：刪除「${category.name}」類別將會影響 ${category.articleCount} 篇文章的分類。\n此操作無法復原，確定要繼續嗎？`
            if (!confirm(secondConfirm)) return
        } else {
            if (!confirm(`確定要刪除「${category.name}」類別嗎？`)) return
        }

        try {
            // 發送刪除請求到後端 API
            const response = await fetch(`/api/categories/${id}`, {
                method: 'DELETE',
            })

            if (!response.ok) {
                throw new Error('刪除類別失敗')
            }

            categories = categories.filter((category) => category.id !== id)
            alert('類別已成功刪除！')
        } catch (error) {
            console.error('刪除類別時發生錯誤:', error)
            alert('刪除類別失敗：' + error.message)
        }
    }

    let selectedCategory = null

    const viewCategoryArticles = (category) => {
        showArticles = true
        selectedCategory = category
    }

    let showArticles = true

    // 修改日期篩選相關變數
    let dateFilter = {
        created: {
            from: null,
            to: null,
        },
        updated: {
            from: null,
            to: null,
        },
    }

    // 新增搜尋關鍵字
    let searchKeyword = ''

    // 修改文章排序相關變量
    let articleSort = null // 當前排序的欄位，null 表示未排序
    let articleSortDirection = 'asc'

    // 修改文章排序相關函數
    const toggleArticleSort = (field) => {
        if (articleSort === field) {
            articleSortDirection =
                articleSortDirection === 'asc' ? 'desc' : 'asc'
        } else {
            articleSort = field
            articleSortDirection = 'asc'
        }
    }

    // 更新篩選和排序邏輯
    $: filteredArticles = articles
        .filter((article) => {
            let matchCategory = selectedCategory
                ? article.category === selectedCategory.name
                : true

            let matchDate = true
            // 檢查發布日期篩選
            if (dateFilter.created.from || dateFilter.created.to) {
                if (dateFilter.created.from) {
                    matchDate =
                        matchDate &&
                        new Date(article.created_at) >=
                            new Date(dateFilter.created.from)
                }
                if (dateFilter.created.to) {
                    matchDate =
                        matchDate &&
                        new Date(article.created_at) <=
                            new Date(dateFilter.created.to)
                }
            }

            // 檢查修改日期篩選
            if (dateFilter.updated.from || dateFilter.updated.to) {
                if (dateFilter.updated.from) {
                    matchDate =
                        matchDate &&
                        new Date(article.updated_at) >=
                            new Date(dateFilter.updated.from)
                }
                if (dateFilter.updated.to) {
                    matchDate =
                        matchDate &&
                        new Date(article.updated_at) <=
                            new Date(dateFilter.updated.to)
                }
            }

            let matchSearch = true
            if (searchKeyword.trim()) {
                matchSearch = article.title
                    .toLowerCase()
                    .includes(searchKeyword.toLowerCase())
            }

            return matchCategory && matchDate && matchSearch
        })
        .sort((a, b) => {
            if (!articleSort) return 0 // 如果沒有選擇排序欄位，保持原順序

            const direction = articleSortDirection === 'asc' ? 1 : -1
            switch (articleSort) {
                case 'title':
                    return direction * a.title.localeCompare(b.title, 'zh-TW')
                case 'created':
                    return (
                        direction *
                        (new Date(a.created_at) - new Date(b.created_at))
                    )
                case 'updated':
                    return (
                        direction *
                        (new Date(a.updated_at) - new Date(b.updated_at))
                    )
                case 'category':
                    return (
                        direction *
                        a.category.localeCompare(b.category, 'zh-TW')
                    )
                default:
                    return 0
            }
        })

    // 清除所有篩選
    const clearAllFilters = () => {
        selectedCategory = null
        dateFilter = {
            created: {
                from: null,
                to: null,
            },
            updated: {
                from: null,
                to: null,
            },
        }
        searchKeyword = ''
    }

    // 更新類別文章數量
    const updateCategoryCount = () => {
        categories = categories.map((category) => ({
            ...category,
            articleCount: articles.filter(
                (article) => article.category === category.name,
            ).length,
        }))
    }

    // 修改日期驗證
    $: {
        if (dateFilter.created.from && dateFilter.created.to) {
            if (
                new Date(dateFilter.created.to) <
                new Date(dateFilter.created.from)
            ) {
                dateFilter.created.to = dateFilter.created.from
            }
        }
        if (dateFilter.updated.from && dateFilter.updated.to) {
            if (
                new Date(dateFilter.updated.to) <
                new Date(dateFilter.updated.from)
            ) {
                dateFilter.updated.to = dateFilter.updated.from
            }
        }
    }

    // 添加新的狀態變量
    let editingCategoryPage = null
    let categoryPageContent = ''

    // 修改 script 部分，添加或更新 cleanupEditor 函數
    const cleanupEditor = async () => {
        // 如果內容有變更，則顯示確認對話框
        if (categoryPageContent.trim()) {
            const confirmed = confirm('確定要離開嗎？未儲存的變更將會遺失。')
            if (!confirmed) {
                return
            }
        }
        editingCategoryPage = null
        categoryPageContent = ''
    }

    // 添加類別篩選相關變量
    let categorySearchKeyword = ''
    let categorySort = null // 當前排序的欄位，null 表示未排序
    let categorySortDirection = 'asc' // 'asc' | 'desc'

    // 修改類別排序相關函數
    const toggleSort = (field) => {
        if (categorySort === field) {
            categorySortDirection =
                categorySortDirection === 'asc' ? 'desc' : 'asc'
        } else {
            categorySort = field
            categorySortDirection = 'asc'
        }
    }

    // 更新類別篩選和排序邏輯
    $: filteredCategories = categories
        .filter((category) => {
            if (!categorySearchKeyword.trim()) return true
            return category.name
                .toLowerCase()
                .includes(categorySearchKeyword.toLowerCase())
        })
        .sort((a, b) => {
            if (!categorySort) return 0 // 如果沒有選擇排序欄位，保持原順序

            const direction = categorySortDirection === 'asc' ? 1 : -1
            if (categorySort === 'name') {
                return direction * a.name.localeCompare(b.name, 'zh-TW')
            } else if (categorySort === 'count') {
                return direction * (a.articleCount - b.articleCount)
            }
            return 0
        })

    let articleTable
    let categoryTable

    // 修改初始化表格的函數
    const initTables = async () => {
        if (!window.jQuery?.fn?.DataTable) {
            console.error('DataTables 未載入')
            return
        }

        try {
            if (articleTable) {
                articleTable.destroy()
                articleTable = null
            }
            if (categoryTable) {
                categoryTable.destroy()
                categoryTable = null
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
                // 添加以下設置來禁用懸停效果
                rowCallback: function (row, data, index) {
                    jQuery(row).off('mouseenter mouseleave')
                },
                createdRow: function (row, data, dataIndex) {
                    jQuery(row).css('background-color', 'white')
                    jQuery(row).hover(
                        function () {
                            jQuery(this).css('background-color', 'white')
                        },
                        function () {
                            jQuery(this).css('background-color', 'white')
                        },
                    )
                },
                drawCallback: function () {
                    // 原有的代碼
                    jQuery('.paginate_button').addClass('button is-small')
                    jQuery('.paginate_button.current').addClass('is-primary')
                    jQuery('.dataTables_info').addClass(
                        'has-text-grey is-size-7',
                    )

                    // 添加以下代碼來確保所有行的背景色
                    jQuery(this)
                        .find('tbody tr')
                        .css('background-color', 'white')
                },
            }

            if (showArticles) {
                const table = document.getElementById('articleTable')
                if (table) {
                    articleTable = window.jQuery(table).DataTable({
                        ...commonConfig,
                        order: [[1, 'asc']],
                        columnDefs: [
                            { orderable: false, targets: [0, 5] },
                            { width: '50px', targets: 0 },
                            { width: '120px', targets: [3, 4] },
                            { width: '100px', targets: 2 },
                            { width: '150px', targets: 5 },
                        ],
                    })
                }
            } else {
                const table = document.getElementById('categoryTable')
                if (table) {
                    categoryTable = window.jQuery(table).DataTable({
                        ...commonConfig,
                        order: [[0, 'asc']],
                        columnDefs: [
                            { orderable: false, targets: 2 },
                            { width: '60%', targets: 0 },
                            { width: '20%', targets: 1 },
                            { width: '20%', targets: 2 },
                        ],
                    })
                }
            }
        } catch (error) {
            console.error('初始化表格時發生錯誤:', error)
        }
    }

    // 修改監聽方式
    $: {
        if (typeof window !== 'undefined' && showArticles !== undefined) {
            // 使用 RAF 確保在瀏覽器重繪後執行
            requestAnimationFrame(() => {
                initTables()
            })
        }
    }

    // ���改切換函數
    const handleTabChange = async (isArticles) => {
        showArticles = isArticles
        // 使用 RAF 確保在瀏覽器重繪後執行
        requestAnimationFrame(() => {
            initTables()
        })
    }

    // 修改 onMount
    onMount(async () => {
        // 等待 scripts 載入完成
        await new Promise((resolve) => setTimeout(resolve, 500))
        await initTables()
        updateCategoryCount()
    })

    // 在組件銷毀時清理表格
    onDestroy(() => {
        if (articleTable) {
            articleTable.destroy()
            articleTable = null
        }
        if (categoryTable) {
            categoryTable.destroy()
            categoryTable = null
        }
    })

    const handleCategoryClick = (category) => {
        viewCategoryArticles(category)
    }

    // 文章相關函數
    const handleDelete = async (articleId) => {
        if (!confirm('確定要刪除這篇文章嗎？此操作無法復原。')) {
            return
        }

        try {
            // 發送刪除請求到後端 API
            const response = await fetch(`/api/articles/${articleId}`, {
                method: 'DELETE',
            })

            if (!response.ok) {
                throw new Error('刪除文章失敗')
            }

            // 從前端狀態中移除文章
            articles = articles.filter((article) => article.id !== articleId)
            updateCategoryCount() // 更新類別文章數
            alert('文章已成功刪除！')
        } catch (error) {
            console.error('刪除文章時發生錯誤:', error)
            alert('刪除文章失敗：' + error.message)
        }
    }

    const handleAddCategory = async () => {
        if (!newCategory.trim()) {
            alert('類別名稱不能為空！')
            return
        }

        if (categories.some((c) => c.name === newCategory.trim())) {
            alert('類別名稱已存在！')
            return
        }

        try {
            // 發送新增請求到後端 API
            const response = await fetch('/api/categories', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name: newCategory.trim() }),
            })

            if (!response.ok) {
                throw new Error('新增類別失敗')
            }

            const newCategoryData = await response.json()
            categories = [
                ...categories,
                {
                    id: newCategoryData.id,
                    name: newCategory.trim(),
                    articleCount: 0,
                },
            ]
            newCategory = ''
            alert('類別新增成功！')
        } catch (error) {
            console.error('新增類別時發生錯誤:', error)
            alert('新增類別失敗：' + error.message)
        }
    }

    const handleEditCategoryPage = async (category) => {
        try {
            // 從後端獲取類別頁面內容
            const response = await fetch(`/api/categories/${category.id}/page`)
            if (!response.ok) {
                throw new Error('獲取類別頁面內容失敗')
            }
            const data = await response.json()
            editingCategoryPage = category
            categoryPageContent = data.content || ''
        } catch (error) {
            console.error('獲取類別頁面內容時發生錯誤:', error)
            alert('獲取類別頁面內容失敗：' + error.message)
        }
    }

    const handleSaveCategoryPage = async () => {
        if (!editingCategoryPage) return

        try {
            // 發送請求到後端保存類別頁面內容
            const response = await fetch(
                `/api/categories/${editingCategoryPage.id}/page`,
                {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ content: categoryPageContent }),
                },
            )

            if (!response.ok) {
                throw new Error('保存類別頁面失敗')
            }

            alert('類別頁面保存成功！')
            editingCategoryPage = null
            categoryPageContent = ''
        } catch (error) {
            console.error('保存類別頁面時發生錯誤:', error)
            alert('保存失敗：' + error.message)
        }
    }

    // 初始化數據
    onMount(async () => {
        try {
            // 獲取文章列表
            const articlesResponse = await fetch('/api/articles')
            if (!articlesResponse.ok) throw new Error('獲取文章列表失敗')
            articles = await articlesResponse.json()

            // 獲取類別列表
            const categoriesResponse = await fetch('/api/categories')
            if (!categoriesResponse.ok) throw new Error('獲取類別列表失敗')
            categories = await categoriesResponse.json()

            updateCategoryCount()
            await initTables()
        } catch (error) {
            console.error('初始化數據時發生錯誤:', error)
            alert('載入數據失敗：' + error.message)
        }
    })
</script>

<style>
    h1 {
        padding: 20px;
    }

    .admin-container {
        min-height: 100vh;
        padding: 2rem;
        background-color: #f7f9fc;
    }

    .admin-actions {
        display: flex;
        justify-content: flex-end;
    }

    .table-container {
        box-shadow: 0 2px 3px rgba(10, 10, 10, 0.1);
        border-radius: 6px;
        padding: 0 20px;
        overflow: hidden;
    }

    @media screen and (max-width: 768px) {
        .admin-container {
            padding: 2rem 1rem;
        }
    }

    .article-title {
        max-width: 500px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .article-title a {
        color: var(--theme-primary);
        text-decoration: none;
    }

    .article-title a:hover {
        text-decoration: underline;
    }

    .table td {
        vertical-align: middle;
    }

    .buttons.is-centered {
        justify-content: center;
        margin-bottom: 0;
    }

    .section-tabs {
        display: flex;
        justify-content: center;
        gap: 1.5rem;
        margin-top: -2rem;
        margin-bottom: 3rem;
    }

    .section-tab {
        box-sizing: border-box;
        padding: 1rem 2rem;
        border: 1px solid var(--neutral-medium);
        border-radius: 8px;
        background: var(--neutral-light);
        color: var(--neutral-dark);
        font-weight: 600;
        transition: all 0.3s ease;
        display: flex;
        align-items: center;
        gap: 0.5rem;
        cursor: pointer;
    }

    .section-tab:hover {
        border: none;
        background: var(--theme-secondary);
        transform: translateY(-2px);
        color: var(--theme-primary);
    }

    .section-tab.active {
        background: var(--theme-primary);
        color: white;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    }

    .section-tab.active:hover {
        background: var(--theme-primary);
        opacity: 0.9;
    }

    .filter-section {
        display: flex;
        flex-direction: column;
    }

    .filter-controls {
        display: flex;
        gap: 1rem;
        align-items: center;
    }

    .filter-controls-right {
        display: flex;
        gap: 1rem;
        align-items: center;
        margin-left: auto;
        flex-shrink: 0; /* 防止壓縮 */
    }

    .filter-tags {
        min-height: 2.5rem;
        display: flex;
        align-items: center;
    }

    .filter-tags .tags {
        margin-bottom: 0;
    }

    .box {
        background: white;
        border: 1px solid #eee;
        box-shadow: 0 2px 15px rgba(0, 0, 0, 0.05);
    }

    .admin-container {
        background-color: #f7f9fc;
    }

    .field-group {
        position: relative;
        display: flex;
        align-items: flex-end;
        gap: 0.5rem;
    }

    .filter-controls {
        display: flex;
        gap: 1rem;
        align-items: flex-end;
    }

    .filter-controls-right {
        display: flex;
        gap: 1rem;
        align-items: flex-end;
        margin-left: auto;
    }

    .modal-card {
        width: 90%;
        max-width: 1200px;
        height: 90vh;
    }

    .modal-card-body {
        padding: 1rem;
        display: flex;
        flex-direction: column;
    }

    .field {
        height: 100%;
        margin: 0;
    }

    /* 更模態相關���式 */
    :global(.modal) {
        z-index: 1000; /* 確保模態框在導航欄上方 */
    }

    :global(.modal-background) {
        z-index: 1001;
    }

    :global(.modal-card) {
        z-index: 1002;
        width: 90%;
        max-width: 1200px;
        height: 90vh;
    }

    .modal-card-body {
        padding: 1rem;
        display: flex;
        flex-direction: column;
        overflow: hidden; /* 防止內容溢出 */
    }

    .field {
        height: 100%;
        margin: 0;
        overflow: hidden; /* 止內容溢出 */
    }

    /* 加排序標題樣式 */
    .sort-header {
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        color: inherit;
        cursor: pointer;
        user-select: none;
    }

    .sort-header:hover {
        color: var(--theme-primary);
    }

    .sort-header .icon {
        font-size: 0.8em;
        opacity: 0.7;
        transition: transform 0.2s ease;
    }

    .sort-header:hover .icon {
        opacity: 1;
    }

    /* 調整表格標題樣式 */
    .table th {
        white-space: nowrap;
    }

    .table th a {
        text-decoration: none;
    }

    /* 更新浮動標籤樣式 */
    .field.has-floating-label {
        position: relative;
        margin-top: 1.5em;
    }

    .field.has-floating-label .floating-label {
        position: absolute;
        top: -1.5em;
        left: 0;
        font-size: 0.75em;
        color: #666;
        pointer-events: none;
        line-height: 1;
    }

    /* 調整日期輸入框的樣式 */
    .field.has-floating-label input[type='date'] {
        width: 150px; /* 調整寬度以確保日期完整顯示 */
        padding: 0.5em;
    }

    /* 確保標籤不會被其他元素覆蓋 */
    .field.has-floating-label {
        z-index: 1;
    }

    .field.has-floating-label .floating-label {
        z-index: 2;
    }

    .date-field-group {
        display: flex;
        gap: 1rem;
        align-items: flex-end;
    }

    .date-field {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .date-field label {
        font-size: 0.75rem;
        color: #666;
    }

    .date-field input {
        width: 150px;
        height: 2.5em;
    }

    /* 修改類別選擇器樣式 */
    .select-field {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .select-field label {
        font-size: 0.75rem;
        color: #666;
    }

    .select-field .select {
        min-width: 150px;
        position: relative;
    }

    .select-field .select select {
        height: 2.5em;
        padding-right: 2.5em;
        width: 100%;
    }

    .select-field .select::after {
        border-color: #666;
        right: 0.75em;
        z-index: 4;
        border-width: 2px;
        margin-top: -0.375em;
    }

    .date-column {
        min-width: 120px;
        white-space: nowrap;
    }

    .date-field-group {
        display: flex;
        gap: 1rem;
        align-items: flex-end;
    }

    .date-type-select {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .date-type-select label {
        font-size: 0.75rem;
        color: #666;
    }

    .date-type-select .select {
        min-width: 120px;
    }

    .date-field {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .date-field label {
        font-size: 0.75rem;
        color: #666;
    }

    .date-field input {
        width: 160px;
        height: 2.5em;
    }

    .date-group {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        padding: 0.5rem;
        border: 1px solid #eee;
        border-radius: 4px;
        background: #fafafa;
    }

    .date-label {
        font-size: 0.75rem;
        font-weight: 600;
        color: #666;
        margin-bottom: 0.25rem;
    }

    .date-inputs {
        display: flex;
        gap: 0.5rem;
    }

    .date-field {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .date-field label {
        font-size: 0.75rem;
        color: #666;
    }

    .date-field input {
        width: 140px;
        height: 2.25em;
    }

    /* 調整日期欄位寬度 */
    .date-column {
        min-width: 130px;
        white-space: nowrap;
    }

    /* 修改排序圖標相關樣式 */
    .sort-header {
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        color: inherit;
        cursor: pointer;
        user-select: none;
    }

    .sort-header:hover {
        color: var(--theme-primary);
    }

    .sort-header .icon {
        font-size: 0.8em;
        opacity: 0.7;
        transition:
            opacity 0.2s ease,
            transform 0.2s ease;
    }

    .sort-header.active .icon {
        opacity: 1;
    }

    .sort-header .icon i {
        display: none;
    }

    .sort-header.active .icon i.fa-sort-up,
    .sort-header.active .icon i.fa-sort-down {
        display: inline-block;
    }

    .sort-header:not(.active) .icon i.fa-sort {
        display: inline-block;
    }

    /* 調整表格容器樣式 */
    .table-container {
        margin-bottom: 2rem;
    }

    /* 確保操作按鈕列的寬度合適 */
    .table td:last-child {
        min-width: 300px;
    }

    .category-link {
        background: none;
        border: none;
        padding: 0;
        color: var(--theme-primary);
        cursor: pointer;
        font-size: 1em;
        text-align: left;
        transition: color 0.2s ease;
        text-decoration: none;
    }

    .category-link:hover {
        text-decoration: underline;
    }

    /* 確保表格行的背景色始終保持白色 */
    .table tbody tr {
        background-color: white !important;
    }

    /* 修改文章標題連結樣式 */
    .article-title a {
        color: var(--theme-primary);
        transition: color 0.2s ease;
        text-decoration: none;
    }

    .article-title a:hover {
        text-decoration: underline;
    }

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

    /* 表格樣式統一 */
    .table {
        width: 100%;
        margin-bottom: 0;
        background-color: white;
    }

    .table th {
        background-color: #f8f9fa;
        border-bottom: 2px solid #dee2e6;
        color: #495057;
        font-weight: 600;
        padding: 0.75rem;
    }

    .table td {
        padding: 0.75rem;
        vertical-align: middle;
        border-bottom: 1px solid #dee2e6;
        background-color: white;
    }

    .table tr:last-child td {
        border-bottom: none;
    }

    /* 確保表格容器樣式一致 */
    .table-container {
        border: 1px solid #dee2e6;
        border-radius: 6px;
        overflow: hidden;
        background: white;
        box-shadow: 0 2px 3px rgba(10, 10, 10, 0.1);
    }

    /* 調整分頁控件的間距 */
    :global(.dataTables_wrapper .dataTables_paginate) {
        margin-top: 1rem;
        padding-top: 0.5rem;
        border-top: 1px solid #dee2e6;
    }

    /* 調整信息顯示的間距 */
    :global(.dataTables_wrapper .dataTables_info) {
        margin-top: 1rem;
        padding-top: 0.5rem;
    }
</style>

<svelte:head>
    <link
        rel="stylesheet"
        type="text/css"
        href="https://cdn.datatables.net/1.13.7/css/jquery.dataTables.css"
    />
    <link
        rel="stylesheet"
        type="text/css"
        href="https://cdn.datatables.net/1.13.7/css/dataTables.bulma.min.css"
    />
    <script src="https://code.jquery.com/jquery-3.7.0.min.js"></script>
    <script
        src="https://cdn.datatables.net/1.13.7/js/jquery.dataTables.min.js"
    ></script>
    <script
        src="https://cdn.datatables.net/1.13.7/js/dataTables.bulma.min.js"
    ></script>
</svelte:head>

<AdminNavbar {siteName} />

<div class="admin-container">
    <div class="container">
        <h1 class="title is-2 has-text-centered mb-6 mt-6">{title}</h1>

        <div class="section-tabs mb-6">
            <button
                class="section-tab"
                class:active={showArticles}
                on:click={() => handleTabChange(true)}
            >
                <span class="icon">
                    <i class="fas fa-file-alt"></i>
                </span>
                <span>文章管理</span>
            </button>
            <button
                class="section-tab"
                class:active={!showArticles}
                on:click={() => handleTabChange(false)}
            >
                <span class="icon">
                    <i class="fas fa-tags"></i>
                </span>
                <span>類別管理</span>
            </button>
        </div>

        {#if showArticles}
            <div class="box">
                <div class="filter-section mb-5">
                    <div class="filter-controls mb-3">
                        <div class="field-group is-expanded">
                            <div class="control has-icons-left">
                                <input
                                    type="text"
                                    class="input"
                                    placeholder="搜尋文章標題..."
                                    bind:value={searchKeyword}
                                />
                                <span class="icon is-left">
                                    <i class="fas fa-search"></i>
                                </span>
                            </div>
                        </div>

                        <div class="filter-controls-right">
                            <div class="select-field">
                                <label>類別</label>
                                <div class="select">
                                    <select bind:value={selectedCategory}>
                                        <option
                                            value={null}
                                            disabled
                                            selected={!selectedCategory}
                                        >
                                            選擇類別
                                        </option>
                                        {#each categories as category}
                                            <option value={category}>
                                                {category.name} ({category.articleCount})
                                            </option>
                                        {/each}
                                    </select>
                                </div>
                            </div>

                            <div class="date-field-group">
                                <div class="date-group">
                                    <div class="date-label">發布日期</div>
                                    <div class="date-inputs">
                                        <div class="date-field">
                                            <label for="createdDateFrom"
                                                >從</label
                                            >
                                            <input
                                                id="createdDateFrom"
                                                type="date"
                                                class="input"
                                                bind:value={dateFilter.created
                                                    .from}
                                                max={dateFilter.created.to ||
                                                    undefined}
                                            />
                                        </div>
                                        <div class="date-field">
                                            <label for="createdDateTo">至</label
                                            >
                                            <input
                                                id="createdDateTo"
                                                type="date"
                                                class="input"
                                                bind:value={dateFilter.created
                                                    .to}
                                                min={dateFilter.created.from ||
                                                    undefined}
                                            />
                                        </div>
                                    </div>
                                </div>

                                <div class="date-group">
                                    <div class="date-label">修改日期</div>
                                    <div class="date-inputs">
                                        <div class="date-field">
                                            <label for="updatedDateFrom"
                                                >從</label
                                            >
                                            <input
                                                id="updatedDateFrom"
                                                type="date"
                                                class="input"
                                                bind:value={dateFilter.updated
                                                    .from}
                                                max={dateFilter.updated.to ||
                                                    undefined}
                                            />
                                        </div>
                                        <div class="date-field">
                                            <label for="updatedDateTo">至</label
                                            >
                                            <input
                                                id="updatedDateTo"
                                                type="date"
                                                class="input"
                                                bind:value={dateFilter.updated
                                                    .to}
                                                min={dateFilter.updated.from ||
                                                    undefined}
                                            />
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <div class="field-group">
                                <a
                                    href="/admin/article/new"
                                    class="button is-success"
                                >
                                    <span class="icon">
                                        <i class="fas fa-plus"></i>
                                    </span>
                                    <span>新增文章</span>
                                </a>
                            </div>
                        </div>
                    </div>

                    <div class="filter-tags">
                        {#if selectedCategory || dateFilter.created.from || dateFilter.created.to || dateFilter.updated.from || dateFilter.updated.to || searchKeyword}
                            <div class="tags">
                                {#if searchKeyword}
                                    <span class="tag is-medium is-warning">
                                        搜尋：{searchKeyword}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (searchKeyword = '')}
                                        ></button>
                                    </span>
                                {/if}
                                {#if selectedCategory}
                                    <span class="tag is-medium is-info">
                                        類別：{selectedCategory.name}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (selectedCategory = null)}
                                        ></button>
                                    </span>
                                {/if}
                                {#if dateFilter.created.from}
                                    <span class="tag is-medium is-success">
                                        發布日期從：{dateFilter.created.from}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (dateFilter.created.from =
                                                    null)}
                                        ></button>
                                    </span>
                                {/if}
                                {#if dateFilter.created.to}
                                    <span class="tag is-medium is-success">
                                        發布日期至：{dateFilter.created.to}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (dateFilter.created.to = null)}
                                        ></button>
                                    </span>
                                {/if}
                                {#if dateFilter.updated.from}
                                    <span class="tag is-medium is-success">
                                        修改日期從：{dateFilter.updated.from}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (dateFilter.updated.from =
                                                    null)}
                                        ></button>
                                    </span>
                                {/if}
                                {#if dateFilter.updated.to}
                                    <span class="tag is-medium is-success">
                                        修改日期至：{dateFilter.updated.to}
                                        <button
                                            class="delete"
                                            on:click={() =>
                                                (dateFilter.updated.to = null)}
                                        ></button>
                                    </span>
                                {/if}
                                <button
                                    class="button is-small is-light"
                                    on:click={clearAllFilters}
                                >
                                    清除所有篩選
                                </button>
                            </div>
                        {:else}
                            <span class="has-text-grey">尚未設定篩選條件</span>
                        {/if}
                    </div>
                </div>

                <div class="table-container">
                    <table id="articleTable" class="table is-fullwidth">
                        <thead>
                            <tr>
                                <th class="is-narrow">#</th>
                                <th>標題</th>
                                <th class="is-narrow">分類</th>
                                <th class="date-column">發布日期</th>
                                <th class="date-column">修改日期</th>
                                <th class="is-narrow has-text-centered">操作</th
                                >
                            </tr>
                        </thead>
                        <tbody>
                            {#each filteredArticles as article, index}
                                <tr>
                                    <td class="has-text-grey">{index + 1}</td>
                                    <td>
                                        <div class="article-title">
                                            <a
                                                href={`/article/${article.id}`}
                                                target="_blank"
                                                class="has-text-dark"
                                            >
                                                {article.title}
                                            </a>
                                        </div>
                                    </td>
                                    <td>
                                        <span class="tag is-info is-light">
                                            {article.category}
                                        </span>
                                    </td>
                                    <td class="has-text-grey">
                                        {article.created_at}
                                    </td>
                                    <td class="has-text-grey">
                                        {article.updated_at}
                                    </td>
                                    <td>
                                        <div
                                            class="buttons is-centered are-small"
                                        >
                                            <a
                                                href={`/admin/article/edit/${article.id}`}
                                                class="button is-info is-outlined"
                                                title="編輯文章"
                                            >
                                                <span class="icon">
                                                    <i class="fas fa-edit"></i>
                                                </span>
                                            </a>
                                            <button
                                                class="button is-danger is-outlined"
                                                title="刪除文章"
                                                on:click={() =>
                                                    handleDelete(article.id)}
                                            >
                                                <span class="icon">
                                                    <i class="fas fa-trash-alt"
                                                    ></i>
                                                </span>
                                            </button>
                                        </div>
                                    </td>
                                </tr>
                            {/each}
                            {#if filteredArticles.length === 0}
                                <tr>
                                    <td
                                        colspan="6"
                                        class="has-text-centered py-6"
                                    >
                                        <p class="has-text-grey">
                                            {selectedCategory
                                                ? `「${selectedCategory.name}」類別目前沒有文章`
                                                : '前沒有任何文章'}
                                        </p>
                                    </td>
                                </tr>
                            {/if}
                        </tbody>
                    </table>
                </div>
            </div>
        {:else}
            <div class="box">
                <div class="field">
                    <label class="label" for="newCategory">新增類別</label>
                    <div class="field has-addons">
                        <div class="control is-expanded">
                            <input
                                class="input"
                                type="text"
                                placeholder="輸入新類別名稱"
                                bind:value={newCategory}
                            />
                        </div>
                        <div class="control">
                            <button
                                class="button is-primary"
                                on:click={handleAddCategory}
                            >
                                新增
                            </button>
                        </div>
                    </div>
                </div>

                <!-- 添加類別篩選區域 -->
                <div class="filter-section mt-5 mb-4">
                    <div class="filter-controls">
                        <div class="field-group is-expanded">
                            <div class="control has-icons-left">
                                <input
                                    type="text"
                                    class="input"
                                    placeholder="搜尋類別名稱..."
                                    bind:value={categorySearchKeyword}
                                />
                                <span class="icon is-left">
                                    <i class="fas fa-search"></i>
                                </span>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="table-container">
                    <table id="categoryTable" class="table is-fullwidth">
                        <thead>
                            <tr>
                                <th>類別名稱</th>
                                <th>文章數量</th>
                                <th class="is-narrow">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each filteredCategories as category}
                                <tr>
                                    <td>
                                        {#if editingCategory?.id === category.id}
                                            <div class="field has-addons">
                                                <div class="control">
                                                    <input
                                                        class="input is-small"
                                                        type="text"
                                                        bind:value={editingCategoryName}
                                                    />
                                                </div>
                                                <div class="control">
                                                    <button
                                                        class="button is-small is-success"
                                                        on:click={handleSaveCategory}
                                                    >
                                                        儲存
                                                    </button>
                                                </div>
                                            </div>
                                        {:else}
                                            <button
                                                class="category-link"
                                                on:click={() =>
                                                    handleCategoryClick(
                                                        category,
                                                    )}
                                            >
                                                {category.name}
                                            </button>
                                        {/if}
                                    </td>
                                    <td>
                                        <span class="tag is-info is-light">
                                            {category.articleCount} 篇
                                        </span>
                                    </td>
                                    <td>
                                        <div class="buttons are-small">
                                            {#if editingCategory?.id !== category.id}
                                                <button
                                                    class="button is-info is-small"
                                                    on:click={() =>
                                                        handleEditCategory(
                                                            category,
                                                        )}
                                                >
                                                    編輯類別名稱
                                                </button>
                                            {/if}
                                            <button
                                                class="button is-primary is-small"
                                                on:click={() =>
                                                    handleEditCategoryPage(
                                                        category,
                                                    )}
                                            >
                                                編輯類別頁面
                                            </button>
                                            <button
                                                class="button is-danger is-small"
                                                class:is-light={category.articleCount >
                                                    0}
                                                on:click={() =>
                                                    handleDeleteCategory(
                                                        category.id,
                                                    )}
                                            >
                                                刪除
                                            </button>
                                        </div>
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        {/if}
    </div>
</div>
<Footer {siteName} />

<!-- 修改模態的部分 -->
{#if editingCategoryPage}
    <div class="modal is-active">
        <div class="modal-background"></div>
        <div class="modal-card">
            <header class="modal-card-head">
                <p class="modal-card-title">
                    編輯「{editingCategoryPage.name}」類別頁面
                </p>
                <button
                    class="delete"
                    aria-label="close"
                    on:click={cleanupEditor}
                ></button>
            </header>
            <section class="modal-card-body">
                <div class="field">
                    <VditorEditor
                        bind:content={categoryPageContent}
                        height="70vh"
                        placeholder="請輸入類別頁面內容..."
                        on:change={(e) => (categoryPageContent = e.detail)}
                    />
                </div>
            </section>
            <footer class="modal-card-foot">
                <button
                    class="button is-success"
                    on:click={handleSaveCategoryPage}
                >
                    儲存變更
                </button>
                <button class="button" on:click={cleanupEditor}> 取消 </button>
            </footer>
        </div>
    </div>
{/if}
