<script>
  import { onMount } from 'svelte';
  import Navbar from '../components/Navbar.svelte';
  import Footer from '../components/Footer.svelte';
  
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

  const handleDelete = async (id) => {
    if (!confirm('確定要刪除這篇文章嗎？')) return;
    articles = articles.filter(article => article.id !== id);
  };
</script>

<Navbar siteName="後台管理" />

<div class="admin-container">
  <div class="container">
    <h1 class="title is-2 has-text-centered mb-6">文章管理</h1>
    
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
</div>

<Footer />

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