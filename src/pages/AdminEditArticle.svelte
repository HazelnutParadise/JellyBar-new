<script>
  import '../app.css'
  import Navbar from '../components/Navbar.svelte';
  import Footer from '../components/Footer.svelte';
  import VditorEditor from '../components/VditorEditor.svelte';
  import { onMount } from 'svelte';
  
  export let siteName;
  export let title;
  export let id = null;
  
  let article = {
    title: '',
    content: '',
    category: '',
    description: '',
    status: 'draft'
  };
  
  let editor;
  
  onMount(() => {
    // 禁用頁面上的 highlight.js
    if (window.hljs) {
      window.hljs.configure({ ignoreUnescapedHTML: true });
      window.hljs.highlightAll = () => {}; // 覆蓋 highlightAll 方法
    }
  });
  
  function handleEditorChange(event) {
    article.content = event.detail;
  }
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log('儲存文章:', article);
  };
</script>

<Navbar {siteName}/>
<div class="page-wrapper">
  <div class="admin-layout">
    <div class="main-content">
      <div class="editor-container">
        <input 
          type="text"
          class="title-input"
          placeholder="輸入標題"
          bind:value={article.title}
          required
        >
        <VditorEditor
          bind:this={editor}
          content={article.content}
          on:change={handleEditorChange}
        />
      </div>
    </div>
    
    <div class="sidebar">
      <div class="publish-box">
        <h3>發布</h3>
        <div class="status-section">
          <select bind:value={article.status} class="status-select">
            <option value="draft">草稿</option>
            <option value="publish">發布</option>
          </select>
        </div>
        <div class="action-buttons">
          <button class="button is-light">預覽</button>
          <button class="button is-primary" on:click|preventDefault={handleSubmit}>
            {article.status === 'draft' ? '儲存草稿' : '發布'}
          </button>
        </div>
      </div>
      
      <div class="category-box">
        <h3>分類</h3>
        <input 
          class="input" 
          type="text" 
          bind:value={article.category}
          placeholder="輸入分類"
        >
      </div>
      
      <div class="description-box">
        <h3>描述</h3>
        <textarea 
          class="textarea" 
          bind:value={article.description}
          placeholder="輸入文章描述"
        ></textarea>
      </div>
    </div>
  </div>
</div>

<Footer {siteName}/>

<style>
  .page-wrapper {
    padding-top: 80px;
    padding-bottom: 20px;
    min-height: 100vh;
    background-color: #f0f0f1;
  }

  .admin-layout {
    display: grid;
    grid-template-columns: minmax(0, 1fr) 300px;
    gap: 20px;
    padding: 20px;
    max-width: 1400px;
    margin: 0 auto;
    padding-top: 32px;
  }
  
  .main-content {
    background: #fff;
    border-radius: 4px;
    box-shadow: 0 1px 3px rgba(0,0,0,0.1);
    min-height: calc(100vh - 120px);
  }
  
  .editor-container {
    padding: 20px;
  }
  
  .title-input {
    width: 100%;
    font-size: 2em;
    border: none;
    padding: 10px;
    margin-bottom: 20px;
    border-bottom: 1px solid #eee;
    background: transparent;
  }
  
  .title-input:focus {
    outline: none;
    border-bottom-color: #4a4a4a;
  }
  
  .sidebar {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  
  .publish-box,
  .category-box,
  .description-box {
    background: #fff;
    padding: 15px;
    border-radius: 4px;
    box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  }
  
  .status-section {
    margin: 15px 0;
  }
  
  .status-select {
    width: 100%;
    padding: 8px;
    border-radius: 4px;
    border: 1px solid #ddd;
  }
  
  .action-buttons {
    display: flex;
    justify-content: space-between;
    gap: 10px;
  }
  
  .action-buttons button {
    flex: 1;
  }
  
  h3 {
    font-size: 1.1em;
    font-weight: 600;
    margin-bottom: 15px;
    color: #1d2327;
  }
  
  textarea.textarea {
    width: 100%;
    min-height: 100px;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
    resize: vertical;
  }
  
  .input {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }


  
  @media (max-width: 768px) {
    .admin-layout {
      grid-template-columns: 1fr;
    }
    
    .main-content {
      min-height: auto;
    }
  }
</style> 