<script>
  import "../app.css"
  import { onMount } from 'svelte'
  import AdminNavbar from '../components/AdminNavbar.svelte'
  import Footer from '../components/Footer.svelte'
  import setTitle from '../js/setTitle'
  import VditorEditor from '../components/VditorEditor.svelte'

  export let title;
  export let siteName;
  setTitle(title, siteName)

  let newCategory = "";
  
  let articles = [
    {
      id: 1,
      title: "Rust繁中簡學！",
      created_at: "2024-03-20",
      category: "Rust"
    },
    {
      id: 2,
      title: "Web開發教學",
      created_at: "2024-03-21",
      category: "Web"
    }
  ];

  let categories = [
    { id: 1, name: "Rust", articleCount: 1 },
    { id: 2, name: "Web", articleCount: 1 },
    { id: 3, name: "Python", articleCount: 0 }
  ];

  let editingCategory = null;
  let editingCategoryName = "";
  
  const handleEditCategory = (category) => {
    editingCategory = category;
    editingCategoryName = category.name;
  };

  const handleSaveCategory = () => {
    if (!editingCategoryName.trim()) {
      alert('類別名稱不能為空！');
      return;
    }
    
    if (categories.some(c => c.name === editingCategoryName.trim() && c.id !== editingCategory.id)) {
      alert('類別名稱已存在！');
      return;
    }

    categories = categories.map(c => 
      c.id === editingCategory.id 
        ? { ...c, name: editingCategoryName.trim() }
        : c
    );
    
    editingCategory = null;
    editingCategoryName = "";
  };

  const handleDeleteCategory = async (id) => {
    const category = categories.find(c => c.id === id);
    
    if (category.articleCount > 0) {
      const confirmMessage = `警告：「${category.name}」類別下還有 ${category.articleCount} 篇文章！\n確定要刪除嗎？`;
      if (!confirm(confirmMessage)) return;
      
      // 二次確認
      const secondConfirm = `最後確認：刪除「${category.name}」類別將會影響 ${category.articleCount} 篇文章的分類。\n此操作無法復原，確定要繼續嗎？`;
      if (!confirm(secondConfirm)) return;
    } else {
      if (!confirm(`確定要刪除「${category.name}」類別嗎？`)) return;
    }

    categories = categories.filter(category => category.id !== id);
  };

  let selectedCategory = null;

  const viewCategoryArticles = (category) => {
    showArticles = true;
    selectedCategory = category;
  };

  let showArticles = true;

  // 新增日期篩選
  let dateFilter = {
    from: null,
    to: null
  };

  // 新增搜尋關鍵字
  let searchKeyword = "";

  // 添加文章排序相關變量
  let articleSort = "title"; // 'title' | 'date' | 'category'
  let articleSortDirection = "asc"; // 'asc' | 'desc'

  // 更新文章排序邏輯
  const toggleArticleSort = (field) => {
    if (articleSort === field) {
      articleSortDirection = articleSortDirection === "asc" ? "desc" : "asc";
    } else {
      articleSort = field;
      articleSortDirection = "asc";
    }
  };

  // 更新篩選和排序邏輯
  $: filteredArticles = articles
    .filter(article => {
      let matchCategory = selectedCategory 
        ? article.category === selectedCategory.name
        : true;

      let matchDate = true;
      if (dateFilter.from) {
        matchDate = matchDate && new Date(article.created_at) >= new Date(dateFilter.from);
      }
      if (dateFilter.to) {
        matchDate = matchDate && new Date(article.created_at) <= new Date(dateFilter.to);
      }

      let matchSearch = true;
      if (searchKeyword.trim()) {
        matchSearch = article.title.toLowerCase().includes(searchKeyword.toLowerCase());
      }

      return matchCategory && matchDate && matchSearch;
    })
    .sort((a, b) => {
      const direction = articleSortDirection === "asc" ? 1 : -1;
      
      switch (articleSort) {
        case "title":
          return direction * a.title.localeCompare(b.title);
        case "date":
          return direction * (new Date(a.created_at) - new Date(b.created_at));
        case "category":
          return direction * a.category.localeCompare(b.category);
        default:
          return 0;
      }
    });

  // 清除所有篩選
  const clearAllFilters = () => {
    selectedCategory = null;
    dateFilter = {
      from: null,
      to: null
    };
    searchKeyword = "";
  };

  // 更新類別文章數量
  const updateCategoryCount = () => {
    categories = categories.map(category => ({
      ...category,
      articleCount: articles.filter(article => article.category === category.name).length
    }));
  };

  const handleAddCategory = () => {
    if (!newCategory.trim()) return;
    if (categories.some(c => c.name === newCategory.trim())) {
      alert('類別名稱已存在！');
      return;
    }
    
    categories = [...categories, {
      id: categories.length + 1,
      name: newCategory.trim(),
      articleCount: 0
    }];
    newCategory = "";
  };

  // 新增日期驗證
  $: if (dateFilter.from && dateFilter.to) {
    if (new Date(dateFilter.to) < new Date(dateFilter.from)) {
      dateFilter.to = dateFilter.from;
    }
  }

  // 新增日期格式化函數
  const formatDate = (date) => {
    if (!date) return null;
    const d = new Date(date);
    if (isNaN(d.getTime())) return null;
    return d.toISOString().split('T')[0];
  };

  // 添加新的狀態變量
  let editingCategoryPage = null;
  let categoryPageContent = "";

  // 簡化編輯類別頁面的處理函數
  const handleEditCategoryPage = async (category) => {
    editingCategoryPage = category;
    // 這裡應該從後端獲取類別頁面內容
    categoryPageContent = "";  // 暫時為空，實際應該從後端獲取
  };

  // 簡化保存函數
  const handleSaveCategoryPage = async () => {
    if (!editingCategoryPage) return;
    
    try {
      // 這裡應該發送請求到後端保存類別頁面內容
      console.log('Saving content:', categoryPageContent);
      alert('類別頁面保存成功！');
      editingCategoryPage = null;
    } catch (error) {
      alert('保存失敗：' + error.message);
    }
  };

  // 簡化清理函數
  const cleanupEditor = () => {
    editingCategoryPage = null;
    categoryPageContent = "";
  };

  // 添加類別篩選相關變量
  let categorySearchKeyword = "";
  let categorySort = "name"; // 'name' | 'count'
  let categorySortDirection = "asc"; // 'asc' | 'desc'

  // 篩選和排序類別的計算屬性
  $: filteredCategories = categories
    .filter(category => {
      if (!categorySearchKeyword.trim()) return true;
      return category.name.toLowerCase().includes(categorySearchKeyword.toLowerCase());
    })
    .sort((a, b) => {
      const direction = categorySortDirection === "asc" ? 1 : -1;
      if (categorySort === "name") {
        return direction * a.name.localeCompare(b.name);
      } else {
        return direction * (a.articleCount - b.articleCount);
      }
    });

  // 切換排序方向
  const toggleSort = (field) => {
    if (categorySort === field) {
      categorySortDirection = categorySortDirection === "asc" ? "desc" : "asc";
    } else {
      categorySort = field;
      categorySortDirection = "asc";
    }
  };

  onMount(() => {
    updateCategoryCount();
  });
</script>

<AdminNavbar siteName={siteName} />

<div class="admin-container">
  <div class="container">
    <h1 class="title is-2 has-text-centered mb-6">{title}</h1>
    
    <div class="section-tabs mb-6">
      <button 
        class="section-tab" 
        class:active={showArticles}
        on:click={() => showArticles = true}
      >
        <span class="icon">
          <i class="fas fa-file-alt"></i>
        </span>
        <span>文章管理</span>
      </button>
      <button 
        class="section-tab" 
        class:active={!showArticles}
        on:click={() => showArticles = false}
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
                >
                <span class="icon is-left">
                  <i class="fas fa-search"></i>
                </span>
              </div>
            </div>

            <div class="filter-controls-right">
              <div class="field-group">
                <div class="select is-fullwidth">
                  <select 
                    bind:value={selectedCategory}
                    class="has-floating-label"
                  >
                    <option value={null} disabled selected={!selectedCategory}>
                      選擇類別
                    </option>
                    {#each categories as category}
                      <option value={category}>
                        {category.name} ({category.articleCount})
                      </option>
                    {/each}
                  </select>
                  <label class="floating-label" for="category">類別</label>
                </div>
              </div>

              <div class="field has-floating-label">
                開始日期
                <input 
                  type="date" 
                  class="input" 
                  bind:value={dateFilter.from}
                  max={dateFilter.to || undefined}
                >
              </div>
              
              <div class="field has-floating-label">
                結束日期
                <input 
                  type="date" 
                  class="input" 
                  bind:value={dateFilter.to}
                  min={dateFilter.from || undefined}
                >
              </div>

              <div class="field-group">
                <a href="/admin/article/new" class="button is-success">
                  <span class="icon">
                    <i class="fas fa-plus"></i>
                  </span>
                  <span>新增文章</span>
                </a>
              </div>
            </div>
          </div>

          <div class="filter-tags">
            {#if selectedCategory || dateFilter.from || dateFilter.to || searchKeyword}
              <div class="tags">
                {#if searchKeyword}
                  <span class="tag is-medium is-warning">
                    搜尋：{searchKeyword}
                    <button class="delete" on:click={() => searchKeyword = ""}></button>
                  </span>
                {/if}
                {#if selectedCategory}
                  <span class="tag is-medium is-info">
                    類別：{selectedCategory.name}
                    <button class="delete" on:click={() => selectedCategory = null}></button>
                  </span>
                {/if}
                {#if dateFilter.from}
                  <span class="tag is-medium is-success">
                    從：{dateFilter.from}
                    <button class="delete" on:click={() => dateFilter.from = null}></button>
                  </span>
                {/if}
                {#if dateFilter.to}
                  <span class="tag is-medium is-success">
                    至：{dateFilter.to}
                    <button class="delete" on:click={() => dateFilter.to = null}></button>
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
          <table class="table is-fullwidth is-hoverable">
            <thead>
              <tr>
                <th class="is-narrow">#</th>
                <th>
                  <a href="javascript:void(0)" 
                     class="sort-header"
                     on:click={() => toggleArticleSort('title')}
                  >
                    標題
                    {#if articleSort === 'title'}
                      <span class="icon">
                        <i class="fas fa-sort-{articleSortDirection === 'asc' ? 'up' : 'down'}"></i>
                      </span>
                    {/if}
                  </a>
                </th>
                <th class="is-narrow">
                  <a href="javascript:void(0)" 
                     class="sort-header"
                     on:click={() => toggleArticleSort('category')}
                  >
                    分類
                    {#if articleSort === 'category'}
                      <span class="icon">
                        <i class="fas fa-sort-{articleSortDirection === 'asc' ? 'up' : 'down'}"></i>
                      </span>
                    {/if}
                  </a>
                </th>
                <th class="is-narrow">
                  <a href="javascript:void(0)" 
                     class="sort-header"
                     on:click={() => toggleArticleSort('date')}
                  >
                    發布日期
                    {#if articleSort === 'date'}
                      <span class="icon">
                        <i class="fas fa-sort-{articleSortDirection === 'asc' ? 'up' : 'down'}"></i>
                      </span>
                    {/if}
                  </a>
                </th>
                <th class="is-narrow has-text-centered">操作</th>
              </tr>
            </thead>
            <tbody>
              {#each filteredArticles as article, index}
                <tr>
                  <td class="has-text-grey">{index + 1}</td>
                  <td>
                    <div class="article-title">
                      <a href={`/article/${article.id}`} target="_blank" class="has-text-dark">
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
                  <td>
                    <div class="buttons is-centered are-small">
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
                        on:click={() => handleDelete(article.id)}
                      >
                        <span class="icon">
                          <i class="fas fa-trash-alt"></i>
                        </span>
                      </button>
                    </div>
                  </td>
                </tr>
              {/each}
              {#if filteredArticles.length === 0}
                <tr>
                  <td colspan="5" class="has-text-centered py-6">
                    <p class="has-text-grey">
                      {selectedCategory 
                        ? `「${selectedCategory.name}」類別目前沒有文章`
                        : '目前沒有任何文章'}
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
              >
            </div>
            <div class="control">
              <button class="button is-primary" on:click={handleAddCategory}>
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
                >
                <span class="icon is-left">
                  <i class="fas fa-search"></i>
                </span>
              </div>
            </div>
          </div>
        </div>

        <div class="table-container">
          <table class="table is-fullwidth">
            <thead>
              <tr>
                <th>
                  <a href="javascript:void(0)" 
                     class="sort-header"
                     on:click={() => toggleSort('name')}
                  >
                    類別名稱
                    {#if categorySort === 'name'}
                      <span class="icon">
                        <i class="fas fa-sort-{categorySortDirection === 'asc' ? 'up' : 'down'}"></i>
                      </span>
                    {/if}
                  </a>
                </th>
                <th>
                  <a href="javascript:void(0)" 
                     class="sort-header"
                     on:click={() => toggleSort('count')}
                  >
                    文章數量
                    {#if categorySort === 'count'}
                      <span class="icon">
                        <i class="fas fa-sort-{categorySortDirection === 'asc' ? 'up' : 'down'}"></i>
                      </span>
                    {/if}
                  </a>
                </th>
                <th>操作</th>
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
                          >
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
                      <a 
                        href="javascript:void(0)" 
                        class="has-text-link"
                        on:click={() => viewCategoryArticles(category)}
                      >
                        {category.name}
                      </a>
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
                          on:click={() => handleEditCategory(category)}
                        >
                          編輯類別名稱
                        </button>
                      {/if}
                      <button 
                        class="button is-primary is-small"
                        on:click={() => handleEditCategoryPage(category)}
                      >
                        編輯類別頁面
                      </button>
                      <button 
                        class="button is-danger is-small"
                        class:is-light={category.articleCount > 0}
                        on:click={() => handleDeleteCategory(category.id)}
                      >
                        刪除
                      </button>
                    </div>
                  </td>
                </tr>
              {/each}
              {#if filteredCategories.length === 0}
                <tr>
                  <td colspan="3" class="has-text-centered py-6">
                    <p class="has-text-grey">
                      {categorySearchKeyword 
                        ? '沒有符合搜尋條件的類別'
                        : '目前沒有任何類別'}
                    </p>
                  </td>
                </tr>
              {/if}
            </tbody>
          </table>
        </div>
      </div>
    {/if}
  </div>
</div>
<Footer siteName={siteName} />

<!-- 修改模態框的內容部分，使用 VditorEditor 組件 -->
{#if editingCategoryPage}
  <div class="modal is-active">
    <div class="modal-background" on:click={cleanupEditor}></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">編輯「{editingCategoryPage.name}」類別頁面</p>
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
            on:change={(e) => categoryPageContent = e.detail}
          />
        </div>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" on:click={handleSaveCategoryPage}>
          儲存變更
        </button>
        <button class="button" on:click={cleanupEditor}>
          取消
        </button>
      </footer>
    </div>
  </div>
{/if}

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
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
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
    box-shadow: 0 2px 15px rgba(0,0,0,0.05);
  }

  .table.is-hoverable tr:hover {
    background-color: #f8faff !important;
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

  /* 更新模態框相關樣式 */
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
    overflow: hidden; /* 防止內容溢出 */
  }

  /* 添加排序標題樣式 */
  .sort-header {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    color: inherit;
    cursor: pointer;
  }

  .sort-header:hover {
    color: var(--theme-primary);
  }

  .sort-header .icon {
    font-size: 0.8em;
    opacity: 0.7;
  }

  /* 調整表格標題樣式 */
  .table th {
    white-space: nowrap;
  }

  .table th a {
    text-decoration: none;
  }

  /* 更新浮動標籤樣式 */
  .field.has-floating-label,
  .select,
  .control {
    position: relative;
    margin-top: 1.5em;
  }

  .floating-label,
  .select label,
  .control label,
  .field.has-floating-label > :first-child:not(input):not(select) {
    position: absolute;
    top: -1.5em;
    left: 0;
    font-size: 0.75em;
    color: #666;
    pointer-events: none;
    line-height: 1;
  }

  /* 統一所有輸入框高度 */
  .input,
  .select select,
  input[type="date"] {
    height: 2.5em;
  }

  /* 日期輸入框寬度 */
  input[type="date"] {
    width: 120px;
  }

  /* 確保所有控制項底部對齊 */
  .filter-controls,
  .filter-controls-right {
    display: flex;
    gap: 1rem;
    align-items: flex-end;
  }

  .filter-controls-right {
    margin-left: auto;
  }
</style> 