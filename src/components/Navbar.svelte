<script>
    export let siteName
</script>

<style>
    .navbar-wrapper {
        width: 100%;
    }

    .navbar-spacer {
        height: 52px;
        width: 100%;
    }

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
        color: var(--theme-secondary) !important;
        font-weight: 500;
        transition: all 0.3s ease;
    }

    .navbar-item:hover {
        background-color: rgba(255, 255, 255, 0.1) !important;
        transform: translateY(-2px);
        text-decoration: none;
    }
</style>

<div class="navbar-wrapper">
    <nav class="navbar is-fixed-top">
        <div class="navbar-brand">
            <a class="navbar-item brand-logo" href="/">
                <img
                    src="/assets/logo.png"
                    alt="Logo"
                    height="28"
                    style="margin-right: -0.5rem;"
                />
                <span class="ml-2 has-text-weight-bold">{siteName}</span>
            </a>
        </div>

        <div class="navbar-menu">
            <div class="navbar-end">
                <a class="navbar-item" href="/">歡迎光臨 Jelly Bar！</a>
                <a class="navbar-item" href="/article">所有文章</a>
                <a class="navbar-item" href="/about">我的帳號</a>
                <a class="navbar-item" href="/about">關於</a>
                <a class="navbar-item" href="/themes">主題展示</a>
                <a class="navbar-item" href="/author">作者帳號</a>
                <a class="navbar-item" href="/article/test">文章</a>
                <a class="navbar-item" href="/category">文章分類</a>
            </div>
        </div>
    </nav>
    <div class="navbar-spacer"></div>
</div>
