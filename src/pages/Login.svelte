<script>
    import { fade } from 'svelte/transition'

    let username = ''
    let password = ''
    let isLoading = false

    async function handleLogin() {
        isLoading = true
        // TODO: 實現登入邏輯
        await new Promise((resolve) => setTimeout(resolve, 1000))
        isLoading = false
    }
</script>

<style>
    .login-container {
        min-height: calc(100vh - 100px - 52px);
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        position: relative;
        padding: 1rem;
    }

    .login-box {
        background: rgba(255, 255, 255, 0.98);
        padding: 1rem;
        margin: 0;
        border-radius: 24px;
        box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
        width: 100%;
        max-width: 600px;
        backdrop-filter: blur(10px);
        border-left: 5px solid #667eea;
        overflow: hidden;
    }

    /* 響應式調整 */
    @media screen and (max-width: 768px) {
        .login-box {
            padding: 1rem 0.5rem;
            margin: 1rem;
            border-radius: 16px;
        }

        .columns {
            margin: 0 !important;
        }

        .column {
            padding: 0.5rem !important;
        }

        .login-description {
            margin-bottom: 1.5rem;
        }

        h1 {
            font-size: 1.75rem !important;
            margin-bottom: 1.5rem !important;
        }

        .icon-wrapper {
            font-size: 2.5rem !important;
            margin-bottom: 1rem !important;
        }

        input {
            padding: 0.9rem !important;
        }

        button {
            padding: 0.9rem !important;
        }

        .links {
            flex-direction: column;
            align-items: center;
            gap: 1rem;
        }
    }

    @media screen and (max-width: 480px) {
        .login-container::before {
            display: none;
        }

        .login-box {
            padding: 1.25rem;
        }
    }

    /* 改進列布局 */
    .columns {
        margin: -0.5rem;
        display: flex;
        flex-wrap: wrap;
    }

    .column {
        padding: 10px;
        flex: 1;
        min-width: 250px;
    }

    /* 美化輸入框 */
    input {
        background: white;
        border: 2px solid #e2e8f0;
    }

    input:focus {
        border-color: #667eea;
        box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
    }

    /* 美化按鈕 */
    button {
        background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
    }

    button:hover {
        transform: translateY(-2px);
        box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
    }

    h1 {
        color: #2d3748;
        text-align: center;
        margin-bottom: 2rem;
        font-size: 2rem;
        font-weight: 700;
        line-height: 1.3;
    }

    .input-group {
        margin-bottom: 1.5rem;
        position: relative;
    }

    .input-group label {
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        background: white;
        padding: 0 0.5rem;
        color: #718096;
        transition: all 0.2s ease-out;
        pointer-events: none;
    }

    .input-group input:focus + label,
    .input-group input:not(:placeholder-shown) + label {
        top: 0;
        transform: translateY(-50%) scale(0.9);
        color: #667eea;
    }

    input {
        width: 100%;
        padding: 1.2rem;
        border: 2px solid var(--theme-subtle);
        border-radius: 8px;
        font-size: 1.1rem;
        transition: all 0.3s ease;
        background: rgba(255, 255, 255, 0.9);
    }

    input::placeholder {
        color: transparent;
    }

    button {
        width: 100%;
        padding: 1.2rem;
        background: #667eea;
        color: white;
        border: none;
        border-radius: 8px;
        font-size: 1.1rem;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.3s ease;
        text-transform: uppercase;
        letter-spacing: 0.5px;
    }

    button:hover {
        background: #764ba2;
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
    }

    button.loading {
        opacity: 0.7;
        cursor: not-allowed;
    }

    .links {
        margin-top: 1.5rem;
        display: flex;
        justify-content: space-between;
        font-size: 0.9rem;
    }

    .links a {
        color: #667eea;
        text-decoration: none;
        transition: all 0.3s ease;
        font-weight: 500;
    }

    .links a:hover {
        color: #764ba2;
        transform: translateY(-1px);
    }

    .login-description {
        max-width: 320px;
        margin: 0 auto;
        background: var(--theme-tertiary);
        padding: 1.2rem;
        margin-bottom: 30px;
        border-radius: 8px;
        border-left: 4px solid var(--decorative-pink-medium);
    }

    .login-description p {
        margin: 0;
        color: var(--theme-text);
        font-size: 1.1rem;
        line-height: 1.7;
    }

    .login-description a {
        color: #667eea;
        text-decoration: none;
        font-weight: 500;
        transition: all 0.3s ease;
    }

    .login-description a:hover {
        color: #764ba2;
    }

    .login-description strong {
        color: var(--decorative-pink-medium);
        font-weight: 600;
    }

    .login-description .tips {
        margin-top: 0.75rem;
        font-size: 0.9rem;
        color: var(--theme-text-light);
    }

    .icon-wrapper {
        font-size: 3rem;
        margin-bottom: 1.5rem;
        text-align: center;
        animation: bounceIn 0.6s ease-out;
    }

    @keyframes bounceIn {
        0% {
            transform: scale(0.3);
            opacity: 0;
        }
        50% {
            transform: scale(1.05);
        }
        70% {
            transform: scale(0.9);
        }
        100% {
            transform: scale(1);
            opacity: 1;
        }
    }

    .column.flex {
        display: flex;
        flex-direction: column;
        justify-content: center;
        min-height: 400px;
        padding: 2rem;
    }

    .w-full {
        max-width: 320px;
        margin: 0 auto;
    }
</style>

<div class="login-container" in:fade={{ duration: 300 }}>
    <div class="login-box columns">
        <div class="column">
            <div class="icon-wrapper">🔐</div>
            <h1>歡迎回來</h1>
            <div class="login-description">
                <p>使用 <strong>榛果繽紛樂</strong> 帳號登入</p>
                <p class="tips">
                    首次使用？請點擊下方「註冊新帳號」<br />
                    若遇登入問題，請<a
                        href="https://apps.hazelnut-paradise.com/mail"
                        target="_blank">聯繫我們</a
                    >
                </p>
            </div>
        </div>
        <div class="column flex flex-col items-center justify-center">
            <div class="w-full">
                <form on:submit|preventDefault={handleLogin}>
                    <div class="input-group">
                        <input
                            type="text"
                            bind:value={username}
                            id="username"
                            placeholder="用戶名稱"
                            required
                        />
                        <label for="username">用戶名稱</label>
                    </div>

                    <div class="input-group">
                        <input
                            type="password"
                            bind:value={password}
                            id="password"
                            placeholder="密碼"
                            required
                        />
                        <label for="password">密碼</label>
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
                    <a
                        href="https://hazelnut-paradise.com/account/forgot-password?redirect=https://jellybar.hazelnut-paradise.com"
                        >忘記密碼？</a
                    >
                    <a
                        href="https://support.hazelnut-paradise.com/index.php?category_id=2"
                        >支援中心</a
                    >
                    <a
                        href="https://hazelnut-paradise.com/account/sign-up?redirect=https://jellybar.hazelnut-paradise.com"
                        >註冊新帳號</a
                    >
                </div>
            </div>
        </div>
    </div>
</div>
