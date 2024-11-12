<script>
    export let title = "繪果吉樂 BAR";
    import '../app.css';
    import ArticleGrid from '../components/ArticleGrid.svelte';

    const themes = [
        'default', 'forest', 'ocean', 'sunset', 'lavender', 'sakura', 
        'honey', 'mint', 'coral', 'twilight', 'sunshine', 'autumn', 
        'spring', 'winter', 'desert', 'galaxy', 'rose', 'emerald', 
        'dawn', 'dusk'
    ];

    const articles = themes.map(theme => ({
        title: `主題：${theme}`,
        description: `這是 ${theme} 主題的示例描述。`,
        theme
    }));
</script>

<nav class="navbar is-fixed-top">
    <div class="navbar-brand">
        <a class="navbar-item brand-logo" href="/">
            <img src="/logo.png" alt="Logo" height="28">
            <span class="ml-2 has-text-weight-bold">繪果吉樂 BAR</span>
        </a>
    </div>
    
    <div class="navbar-menu">
        <div class="navbar-end">
            <a class="navbar-item" href="/welcome">歡迎光臨 Jelly Bar！</a>
            <a class="navbar-item" href="/rust">Rust繁中簡學！</a>
            <a class="navbar-item" href="/articles">所有文章</a>
            <a class="navbar-item" href="/about">我的帳號</a>
            <a class="navbar-item" href="/about">關於</a>
            <a class="navbar-item" href="/fans">檸果織粉絲</a>
        </div>
    </div>
</nav>

<section class="hero is-fullheight main-content">
    <div class="hero-body">
        <div class="container">
            <h1 class="main-title has-text-centered mb-6">
                主題展示
            </h1>
            
            <ArticleGrid {articles} />
        </div>
    </div>
</section>

<style>
    /* 導航欄樣式 */
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

    /* 主要內容區域 */
    .main-content {
        background: linear-gradient(135deg, var(--color-green-light) 0%, var(--color-green-mint) 100%);
        padding-top: 4rem;
    }

    .main-title {
        font-size: 3rem;
        font-weight: 800;
        color: var(--color-green-dark);
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
        letter-spacing: 1px;
        animation: fadeInDown 1s ease-out;
    }

    /* 動畫 */
    @keyframes fadeInDown {
        from { opacity: 0; transform: translateY(-20px); }
        to { opacity: 1; transform: translateY(0); }
    }

    /* 響應式設計 */
    @media screen and (max-width: 768px) {
        .main-title {
            font-size: 2rem;
            padding: 0 1rem;
        }
    }
</style>