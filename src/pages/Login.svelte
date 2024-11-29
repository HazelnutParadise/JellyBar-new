<script>
  import { fade } from 'svelte/transition';
  
  let username = '';
  let password = '';
  let isLoading = false;

  async function handleLogin() {
    isLoading = true;
    // TODO: 實現登入邏輯
    await new Promise(resolve => setTimeout(resolve, 1000));
    isLoading = false;
  }
</script>

<div class="login-container" in:fade={{ duration: 300 }}>
  <div class="login-box">
    <h1>歡迎回來</h1>
    <form on:submit|preventDefault={handleLogin}>
      <div class="input-group">
        <input
          type="text"
          bind:value={username}
          placeholder="用戶名稱"
          required
        />
      </div>
      
      <div class="input-group">
        <input
          type="password"
          bind:value={password}
          placeholder="密碼"
          required
        />
      </div>

      <button type="submit" class:loading={isLoading}>
        {#if isLoading}
          登入中...
        {:else}
          登入
        {/if}
      </button>
    </form>
    
    <div class="links">
      <a href="/forgot-password">忘記密碼？</a>
      <a href="/register">註冊新帳號</a>
    </div>
  </div>
</div>

<style>
  .login-container {
    height: calc(100vh - 100px - 52px);
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  }

  .login-box {
    background: rgba(255, 255, 255, 0.95);
    padding: 2.5rem;
    border-radius: 1rem;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
    backdrop-filter: blur(10px);
  }

  h1 {
    color: #2d3748;
    text-align: center;
    margin-bottom: 2rem;
    font-size: 2rem;
  }

  .input-group {
    margin-bottom: 1.5rem;
  }

  input {
    width: 100%;
    padding: 0.75rem 1rem;
    border: 2px solid #e2e8f0;
    border-radius: 0.5rem;
    font-size: 1rem;
    transition: all 0.3s ease;
  }

  input:focus {
    border-color: #667eea;
    outline: none;
    box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
  }

  button {
    width: 100%;
    padding: 0.75rem;
    background: #667eea;
    color: white;
    border: none;
    border-radius: 0.5rem;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  button:hover {
    background: #764ba2;
    transform: translateY(-1px);
  }

  button.loading {
    opacity: 0.7;
    cursor: not-allowed;
  }

  .links {
    margin-top: 1.5rem;
    display: flex;
    justify-content: space-between;
    font-size: 0.875rem;
  }

  .links a {
    color: #667eea;
    text-decoration: none;
    transition: color 0.3s ease;
  }

  .links a:hover {
    color: #764ba2;
  }
</style>
