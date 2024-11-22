<script>
    import '../app.css'
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    import VditorEditor from '../components/VditorEditor.svelte'
    import { onMount } from 'svelte'

    export let siteName
    export let title
    export let id = null

    // 文章狀態選項
    const statusOptions = [
        { value: 'draft', label: '草稿', icon: 'fa-file' },
        { value: 'review', label: '待審核', icon: 'fa-clock' },
        { value: 'publish', label: '已發布', icon: 'fa-check-circle' },
        { value: 'private', label: '私人', icon: 'fa-lock' },
    ]

    let article = {
        title: '',
        content: '# 歡迎使用編輯器\n\n這裡可以開始編寫您的文章...',
        category: '',
        description: '',
        status: 'draft',
        tags: [],
    }

    // 預設分類列表
    let categories = [
        '前端開發',
        '後端技術',
        'JavaScript',
        'Python',
        'Go語言',
        'DevOps',
        '資料庫',
        '系統架構',
        '程式設計',
        '演算法',
        '資訊安全',
        '雲端服務',
        '容器技術',
        '人工智慧',
        '機器學習',
        '心得分享',
        '專案管理',
    ]

    let searchTerm = ''
    let showDropdown = false

    // 根據搜尋詞過濾分類
    $: filteredCategories = searchTerm
        ? categories.filter((cat) =>
              cat.toLowerCase().includes(searchTerm.toLowerCase()),
          )
        : categories

    // 選擇分類
    function selectCategory(category) {
        article.category = category
        searchTerm = category
        showDropdown = false
    }

    // 新增分類
    function addNewCategory() {
        const newCategory = searchTerm.trim()
        if (newCategory && !categories.includes(newCategory)) {
            categories = [...categories, newCategory]
            selectCategory(newCategory)
        }
    }

    // 清除選擇的分類
    function clearCategory() {
        article.category = ''
        searchTerm = ''
        showDropdown = false
    }

    // 點擊外部時關閉下拉選單
    function handleClickOutside(event) {
        const searchContainer = document.querySelector('.search-container')
        if (searchContainer && !searchContainer.contains(event.target)) {
            showDropdown = false
        }
    }

    onMount(() => {
        // 如果有 id，載入文章資料
        if (id) {
            // TODO: 從後端載入文章資料
            console.log('載入文章:', id)
        }

        document.addEventListener('click', handleClickOutside)
        return () => {
            document.removeEventListener('click', handleClickOutside)
        }
    })

    // 取得對應狀態的按鈕文字和圖標
    $: submitButtonConfig = {
        draft: { icon: 'fa-save', text: '儲存草稿' },
        review: { icon: 'fa-paper-plane', text: '送出審核' },
        publish: { icon: 'fa-paper-plane', text: '發布文章' },
        private: { icon: 'fa-lock', text: '儲存私人文章' },
    }[article.status]

    let editor

    onMount(() => {
        if (window.hljs) {
            window.hljs.configure({ ignoreUnescapedHTML: true })
            window.hljs.highlightAll = () => {}
        }
    })

    function handleEditorChange(event) {
        article.content = event.detail
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        console.log('儲存文章:', article)
    }
</script>

<style>
    .page-wrapper {
        padding-top: 80px;
        padding-bottom: 20px;
        min-height: 100vh;
        background-color: var(--neutral-light);
    }

    .admin-layout {
        display: grid;
        grid-template-columns: minmax(0, 1fr) 320px;
        gap: 24px;
        padding: 24px;
        max-width: 1400px;
        margin: 0 auto;
    }

    .main-content {
        background: #fff;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
        min-height: calc(100vh - 120px);
    }

    .editor-container {
        padding: 24px;
    }

    .title-input {
        width: 100%;
        font-size: 2em;
        border: none;
        padding: 12px;
        margin-bottom: 24px;
        border-bottom: 2px solid var(--neutral-medium);
        background: transparent;
        font-family: inherit;
        transition: border-color 0.3s ease;
    }

    .title-input:focus {
        outline: none;
        border-bottom-color: var(--theme-accent);
    }

    .sidebar {
        display: flex;
        flex-direction: column;
        gap: 24px;
    }

    .sidebar-box {
        background: #fff;
        padding: 20px;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
    }

    .status-section {
        margin: 15px 0;
    }

    .status-section label {
        display: block;
        margin-bottom: 8px;
        color: var(--neutral-dark);
        font-size: 0.9em;
    }

    .select-wrapper {
        position: relative;
        width: 100%;
    }

    .status-select {
        width: 100%;
        padding: 10px 32px 10px 12px; /* 增加右邊 padding 給箭頭圖標 */
        border-radius: 6px;
        border: 1px solid var(--neutral-medium);
        background-color: white;
        font-family: inherit;
        transition: border-color 0.3s ease;
        appearance: none; /* 移除原生下拉箭頭 */
        -webkit-appearance: none;
        -moz-appearance: none;
        cursor: pointer;
    }

    .select-icon {
        position: absolute;
        right: 12px;
        top: 50%;
        transform: translateY(-50%);
        pointer-events: none;
        color: var(--neutral-dark);
    }

    .status-select:focus {
        outline: none;
        border-color: var(--theme-accent);
    }

    .status-select option {
        padding: 8px;
        font-size: 0.95em;
    }

    .action-buttons {
        display: flex;
        justify-content: space-between;
        gap: 12px;
        margin-top: 20px;
    }

    .button {
        flex: 1;
        padding: 10px;
        border: none;
        border-radius: 6px;
        font-family: inherit;
        font-size: 0.95em;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 8px;
        transition: all 0.3s ease;
    }

    .preview-button {
        background-color: var(--neutral-light);
        color: var(--theme-text);
    }

    .preview-button:hover {
        background-color: var(--neutral-medium);
    }

    .publish-button {
        background-color: var(--theme-accent);
        color: white;
    }

    .publish-button:hover {
        background-color: var(--interactive-dark);
    }

    h3 {
        font-size: 1.1em;
        font-weight: 600;
        margin-bottom: 15px;
        color: var(--theme-text);
    }

    .textarea {
        width: 100%;
        min-height: 120px;
        padding: 12px;
        border: 1px solid var(--neutral-medium);
        border-radius: 6px;
        resize: vertical;
        font-family: inherit;
        transition: border-color 0.3s ease;
    }

    .textarea:focus {
        outline: none;
        border-color: var(--theme-accent);
    }

    .input {
        width: 100%;
        padding: 10px;
        border: 1px solid var(--neutral-medium);
        border-radius: 6px;
        font-family: inherit;
        transition: border-color 0.3s ease;
    }

    .input:focus {
        outline: none;
        border-color: var(--theme-accent);
    }

    @media (max-width: 768px) {
        .admin-layout {
            grid-template-columns: 1fr;
        }

        .main-content {
            min-height: auto;
        }

        .editor-container {
            padding: 16px;
        }
    }

    .category-select {
        width: 100%;
    }

    .select-container {
        position: relative;
        margin-bottom: 4px;
    }

    .search-input {
        width: 100%;
        padding-right: 35px;
    }

    .toggle-button {
        position: absolute;
        right: 0;
        top: 0;
        height: 100%;
        width: 35px;
        background: none;
        border: none;
        color: var(--neutral-dark);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .dropdown-list {
        position: absolute;
        width: 100%;
        max-height: 200px;
        overflow-y: auto;
        background: white;
        border: 1px solid var(--neutral-medium);
        border-radius: 6px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        z-index: 1000;
        margin: 0;
        padding: 4px 0;
        list-style: none;
    }

    .dropdown-item {
        width: 100%;
        padding: 8px 12px;
        border: none;
        background: none;
        text-align: left;
        font-family: inherit;
        font-size: 0.95em;
        color: inherit;
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 8px;
    }

    .dropdown-item:hover {
        background-color: var(--neutral-light);
    }

    .dropdown-item.active {
        background-color: var(--theme-accent);
        color: white;
    }

    .new-item {
        color: var(--theme-accent);
        font-weight: 500;
    }

    .new-item:hover {
        background-color: var(--theme-accent);
        color: white;
    }
</style>

<AdminNavbar {siteName} />
<div class="page-wrapper">
    <div class="admin-layout">
        <div class="main-content">
            <div class="editor-container">
                <input
                    type="text"
                    class="title-input"
                    placeholder="輸入標題..."
                    bind:value={article.title}
                    required
                />
                <VditorEditor
                    bind:this={editor}
                    content={article.content}
                    on:change={handleEditorChange}
                />
            </div>
        </div>

        <div class="sidebar">
            <div class="sidebar-box publish-box">
                <h3>發布設定</h3>
                <div class="status-section">
                    <label for="status">文章狀態</label>
                    <div class="select-wrapper">
                        <select
                            id="status"
                            bind:value={article.status}
                            class="status-select"
                        >
                            {#each statusOptions as option}
                                <option value={option.value}>
                                    {option.label}
                                </option>
                            {/each}
                        </select>
                        <div class="select-icon">
                            <i class="fas fa-chevron-down"></i>
                        </div>
                    </div>
                </div>
                <div class="action-buttons">
                    <button class="button preview-button">
                        <i class="fas fa-eye"></i>
                        <span>預覽</span>
                    </button>
                    <button
                        class="button publish-button"
                        on:click|preventDefault={handleSubmit}
                    >
                        <i class="fas {submitButtonConfig.icon}"></i>
                        <span>{submitButtonConfig.text}</span>
                    </button>
                </div>
            </div>

            <div class="sidebar-box category-box">
                <h3>分類設定</h3>
                <div class="category-select">
                    <div class="select-container">
                        <input
                            type="text"
                            class="input search-input"
                            placeholder="搜尋或輸入新分類..."
                            bind:value={searchTerm}
                            on:focus={() => (showDropdown = true)}
                            on:input={() => {
                                showDropdown = true
                                if (!searchTerm) {
                                    article.category = ''
                                }
                            }}
                        />
                        <button
                            type="button"
                            class="toggle-button"
                            on:click|stopPropagation={() =>
                                (showDropdown = !showDropdown)}
                        >
                            <i class="fas fa-chevron-down"></i>
                        </button>
                    </div>

                    {#if showDropdown}
                        <ul class="dropdown-list">
                            {#each filteredCategories as category}
                                <li>
                                    <button
                                        class="dropdown-item"
                                        class:active={category ===
                                            article.category}
                                        on:click|stopPropagation={() => {
                                            article.category = category
                                            searchTerm = category
                                            showDropdown = false
                                        }}
                                    >
                                        <i class="fas fa-tag"></i>
                                        {category}
                                    </button>
                                </li>
                            {/each}

                            {#if searchTerm && !categories.includes(searchTerm)}
                                <li>
                                    <button
                                        class="dropdown-item new-item"
                                        on:click|stopPropagation={() => {
                                            const newCategory =
                                                searchTerm.trim()
                                            if (newCategory) {
                                                categories = [
                                                    ...categories,
                                                    newCategory,
                                                ]
                                                article.category = newCategory
                                                showDropdown = false
                                            }
                                        }}
                                    >
                                        <i class="fas fa-plus"></i>
                                        新增 "{searchTerm}"
                                    </button>
                                </li>
                            {/if}
                        </ul>
                    {/if}
                </div>
            </div>

            <div class="sidebar-box description-box">
                <h3>文章描述</h3>
                <textarea
                    class="textarea"
                    bind:value={article.description}
                    placeholder="輸入文章描述（建議 100-150 字）..."
                ></textarea>
            </div>
        </div>
    </div>
</div>

<Footer {siteName} />
