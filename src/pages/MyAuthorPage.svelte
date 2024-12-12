<script>
  import { onMount } from 'svelte';
  import VditorEditor from '../components/VditorEditor.svelte';
  
  let thisUser;
  let authorName = '';
  let photoUrl = '';
  let photoFile;
  let isLoading = false;
  let vditorEditor;
  let authorDescription = '';
  
  onMount(async () => {
    // 載入當前作者資料
    if (thisUser) {
      authorName = thisUser.displayName || '';
      photoUrl = thisUser.photoURL || '';
      authorDescription = thisUser.description || '';
    }
  });

  async function handlePhotoChange(event) {
    const file = event.target.files[0];
    if (file) {
      photoFile = file;
      photoUrl = URL.createObjectURL(file);
    }
  }

  async function handleSubmit() {
    try {
      isLoading = true;
      const description = vditorEditor.getValue();
      // 這裡需要實作上傳照片到儲存空間的邏輯
      // 以及更新作者資料到資料庫的邏輯，包含 description
      alert('資料更新成功！');
    } catch (error) {
      console.error('更新失敗:', error);
      alert('更新失敗，請稍後再試');
    } finally {
      isLoading = false;
    }
  }
</script>

<style>
  /* 重置 Vditor 中的 Bulma 樣式影響 */
  :global(.vditor-reset) :is(h1, h2, h3, h4, h5, h6) {
    font-size: revert;
    margin: revert;
    font-weight: revert;
  }
  
  :global(.vditor-reset) p {
    margin: revert;
  }
  
  :global(.vditor-reset) :is(ul, ol) {
    margin: revert;
    padding: revert;
    list-style: revert;
  }
  
  :global(.vditor-reset) li {
    margin: revert;
    padding: revert;
  }
  
  :global(.vditor-reset) table {
    border-collapse: revert;
    width: revert;
  }
  
  :global(.vditor-reset) :is(th, td) {
    padding: revert;
    border: revert;
  }

  /* 確保編輯器容器樣式正確 */
  :global(.vditor-container) {
    border: 1px solid #dbdbdb;
    border-radius: 4px;
  }

  /* 修正工具列按鈕樣式 */
  :global(.vditor-toolbar button) {
    all: unset;
    background: none;
    border: none;
    padding: 6px !important;
    height: auto !important;
    line-height: normal !important;
  }

  :global(.vditor-toolbar button:hover) {
    background-color: #f5f5f5 !important;
  }

  /* 修正輸入區域樣式 */
  :global(.vditor-ir), :global(.vditor-wysiwyg) {
    padding: 10px !important;
  }

  /* 修正下拉選單樣式 */
  :global(.vditor-hint) {
    background-color: #fff !important;
    border: 1px solid #dbdbdb !important;
  }
</style>

<section class="section">
  <div class="container is-max-desktop">
    <h1 class="title is-2">編輯作者資料</h1>
    
    <form on:submit|preventDefault={handleSubmit}>
      <div class="field">
        <label class="label">作者名稱</label>
        <div class="control">
          <input
            class="input"
            type="text"
            bind:value={authorName}
            placeholder="請輸入作者名稱"
            required
          />
        </div>
      </div>
      
      <div class="field">
        <label class="label">個人照片</label>
        <div class="columns is-vcentered">
          {#if photoUrl}
            <div class="column is-narrow">
              <figure class="image is-96x96">
                <img
                  src={photoUrl}
                  alt="作者照片"
                  class="is-rounded"
                />
              </figure>
            </div>
          {/if}
          <div class="column">
            <div class="file">
              <label class="file-label">
                <input
                  class="file-input"
                  type="file"
                  accept="image/*"
                  on:change={handlePhotoChange}
                />
                <span class="file-cta">
                  <span class="file-label">
                    選擇照片
                  </span>
                </span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">作者介紹</label>
        <div class="control">
          <VditorEditor
            bind:this={vditorEditor}
            content={authorDescription}
            height={400}
            placeholder="請輸入作者介紹..."
          />
        </div>
      </div>
      
      <div class="field mt-5">
        <div class="control">
          <button
            type="submit"
            class="button is-primary"
            class:is-loading={isLoading}
            disabled={isLoading}
          >
            儲存變更
          </button>
        </div>
      </div>
    </form>
  </div>
</section>
