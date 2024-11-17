<script>
  import Navbar from '../components/Navbar.svelte';
  import Footer from '../components/Footer.svelte';
  
  export let id = null; // 如果是新增文章，id 為 null
  
  let article = {
    title: '',
    content: '',
    category: '',
    description: ''
  };
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    // TODO: 實作儲存邏輯
    console.log('儲存文章:', article);
  };
</script>

<Navbar siteName="後台管理" />

<div class="admin-container">
  <div class="container">
    <h1 class="title is-2 has-text-centered mb-6">
      {id ? '編輯文章' : '新增文章'}
    </h1>
    
    <form on:submit={handleSubmit} class="box">
      <div class="field">
        <label class="label">標題</label>
        <div class="control">
          <input 
            class="input" 
            type="text" 
            bind:value={article.title}
            required
          >
        </div>
      </div>
      
      <div class="field">
        <label class="label">分類</label>
        <div class="control">
          <input 
            class="input" 
            type="text" 
            bind:value={article.category}
            required
          >
        </div>
      </div>
      
      <div class="field">
        <label class="label">描述</label>
        <div class="control">
          <textarea 
            class="textarea" 
            bind:value={article.description}
            required
          ></textarea>
        </div>
      </div>
      
      <div class="field">
        <label class="label">內容</label>
        <div class="control">
          <textarea 
            class="textarea" 
            bind:value={article.content}
            rows="10"
            required
          ></textarea>
        </div>
      </div>
      
      <div class="field is-grouped is-grouped-right">
        <div class="control">
          <button type="submit" class="button is-primary">
            儲存
          </button>
        </div>
        <div class="control">
          <a href="/admin" class="button">
            取消
          </a>
        </div>
      </div>
    </form>
  </div>
</div>

<Footer />

<style>
  .admin-container {
    padding: 3rem 1.5rem;
    min-height: calc(100vh - 52px - 80px);
  }
  
  .box {
    max-width: 800px;
    margin: 0 auto;
  }
  
  @media screen and (max-width: 768px) {
    .admin-container {
      padding: 2rem 1rem;
    }
  }
</style> 