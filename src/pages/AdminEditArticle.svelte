<script>
  import '../app.css'
  import Navbar from '../components/Navbar.svelte';
  import Footer from '../components/Footer.svelte';
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
  
  let vditor;
  
  // 定義繁體中文語言包
  const zhHant = {
    toolbar: {
      emoji: '表情',
      headings: '標題',
      bold: '粗體',
      italic: '斜體',
      strike: '刪除線',
      line: '分隔線',
      quote: '引用',
      list: '無序列表',
      'ordered-list': '有序列表',
      check: '待辦事項',
      code: '程式碼區塊',
      'inline-code': '行內程式碼',
      upload: '上傳',
      link: '連結',
      table: '表格',
      preview: '預覽',
      fullscreen: '全螢幕',
      outline: '大綱',
      help: '幫助',
    },
    hint: {
      emoji: '表情',
      placeholder: {
        emoji: '搜尋表情...',
        loading: '載入中...',
      }
    },
    preview: {
      mode: {
        editor: '編輯',
        preview: '預覽',
        both: '分欄'
      }
    },
    upload: {
      max: '文件大小不能超過',
      upload: '上傳',
      tip: '點擊或拖曳上傳',
      error: '上傳失敗',
      loading: '上傳中...'
    },
    dialog: {
      cancel: '取消',
      ok: '確定'
    },
    codeTheme: {
      github: 'GitHub',
      dark: '深色',
    }
  };
  
  onMount(() => {
    // 初始化 Vditor 編輯器
    vditor = new window.Vditor('vditor', {
      height: 500,
      mode: 'ir',
      lang: 'zh_TW', // 使用繁體中文語言包
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
        'help'
      ],
      placeholder: '請輸入文章內容...',
      theme: 'classic',
      cache: {
        enable: false
      },
      preview: {
        theme: {
          current: 'light'
        },
        hljs: {
          style: 'github'
        }
      },
      counter: {
        enable: true,
        type: 'text',
      },
      hint: {
        emoji: {
          '+1': '👍',
          '-1': '👎',
          'smile': '😄',
          'heart': '❤️',
          'star': '⭐',
        }
      },
      after: () => {
        if (article.content) {
          vditor.setValue(article.content);
        }
      }
    });
  });
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    article.content = vditor.getValue();
    console.log('儲存文章:', article);
  };
</script>

<svelte:head>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/vditor/dist/index.css" />
  <script src="https://cdn.jsdelivr.net/npm/vditor/dist/index.min.js"></script>
</svelte:head>

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
        <div id="vditor" class="vditor"></div>
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
  
  :global(.vditor) {
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