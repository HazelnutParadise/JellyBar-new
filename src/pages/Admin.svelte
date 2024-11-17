<script>
  import "../app.css";
  import { onMount } from 'svelte';
  import Navbar from '../components/Navbar.svelte';
  import Footer from '../components/Footer.svelte';

  export let title;
  export let siteName;
  
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

  let newCategory = "";
  let isEditingCategory = false;
  let selectedCategory = null;
  let showArticles = true;

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

  const handleDeleteCategory = async (id) => {
    const category = categories.find(c => c.id === id);
    if (category.articleCount > 0) {
      alert(`無法刪除「${category.name}」類別，因為還有 ${category.articleCount} 篇文章使用此類別！`);
      return;
    }

    if (!confirm(`確定要刪除「${category.name}」類別嗎？此操作無法復原。`)) return;
    categories = categories.filter(category => category.id !== id);
  };

  onMount(() => {
    updateCategoryCount();
  });
</script>

<Navbar siteName={siteName} />

<div class="admin-container">
  <div class="container">
    <h1 class="title is-2 has-text-centered mb-6">後台管理</h1>
    
    <div class="tabs is-centered is-boxed">
      <ul>
        <li class:is-active={showArticles}>
          <a on:click={() => showArticles = true}>
            <span>文章管理</span>
          </a>
        </li>
        <li class:is-active={!showArticles}>
          <a on:click={() => showArticles = false}>
            <span>類別管理</span>
          </a>
        </li>
      </ul>
    </div>

    {#if showArticles}
      <!-- 文章管理區塊 -->
      <div class="box">
        <div class="admin-actions mb-5">
          <a href="/admin/new" class="button is-primary">
            <span>新增文章</span>
          </a>
        </div>

        <div class="table-container">
          <table class="table is-fullwidth is-striped">
            <thead>
              <tr>
                <th>標題</th>
                <th>分類</th>
                <th>發布日期</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {#each articles as article}
                <tr>
                  <td>{article.title}</td>
                  <td>{article.category}</td>
                  <td>{article.created_at}</td>
                  <td>
                    <div class="buttons are-small">
                      <a href={`/admin/edit/${article.id}`} class="button is-info">
                        編輯
                      </a>
                      <button class="button is-danger" on:click={() => handleDelete(article.id)}>
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
    {:else}
      <!-- 類別管理區塊 -->
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

        <div class="table-container mt-5">
          <table class="table is-fullwidth">
            <thead>
              <tr>
                <th>類別名稱</th>
                <th>文章數量</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {#each categories as category}
                <tr>
                  <td>{category.name}</td>
                  <td>{category.articleCount} 篇</td>
                  <td>
                    <button 
                      class="button is-danger is-small"
                      disabled={category.articleCount > 0}
                      on:click={() => handleDeleteCategory(category.id)}
                    >
                      刪除
                    </button>
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

<Footer siteName={siteName} />

<style>
  .admin-container {
    padding: 3rem 1.5rem;
    min-height: calc(100vh - 52px - 80px);
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
</style> 