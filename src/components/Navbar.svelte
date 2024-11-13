<script>
    export let siteName;
</script>

<nav class="navbar is-fixed-top">
    <div class="navbar-brand">
        <a class="navbar-item brand-logo" href="/">
            <img src="/logo.png" alt="Logo" height="28">
            <span class="ml-2 has-text-weight-bold">{siteName}</span>
        </a>
    </div>
    
    <div class="navbar-menu">
        <div class="navbar-end">
            <a class="navbar-item" href="/">歡迎光臨 Jelly Bar！</a>
            <a class="navbar-item" href="/rust">Rust繁中簡學！</a>
            <a class="navbar-item" href="/articles">所有文章</a>
            <a class="navbar-item" href="/about">我的帳號</a>
            <a class="navbar-item" href="/about">關於</a>
            <a class="navbar-item" href="/fans">檸果織粉絲</a>
            <a class="navbar-item" href="/themes">主題展示</a>
        </div>
    </div>
</nav>

<style>
    .navbar {
        background: rgba(53, 70, 73, 0.95) !important;
        backdrop-filter: blur(10px);
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    }

    .brand-logo {
        font-size: 1.2rem;
        letter-spacing: 0.5px;
    }

    .navbar-item {
        color: var(--color-green-light) !important;
        font-weight: 500;
        transition: all 0.3s ease;
    }
    
    .navbar-item:hover {
        background-color: rgba(255, 255, 255, 0.1) !important;
        transform: translateY(-2px);
    }
</style> 