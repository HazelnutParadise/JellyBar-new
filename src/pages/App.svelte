<script>
    export let siteName;
    export let title;
    import '../app.css';
    import ArticleGrid from '../components/ArticleGrid.svelte';
    import Navbar from '../components/Navbar.svelte';

    const articles = [
        {
            title: "Rust繁中簡學！",
            description: "Rust繁中簡學！系列文章以初學者的角度跟大家分享我自己翻譯成繁體中文的 Rust 自學資源以及學習過程。",
            theme: 'forest',
            icon: '🦀'
        },
        {
            title: "所有文章",
            description: "瀏覽本站所有技術文章，包含 Rust、Web 開發、系統設計等各種主題。",
            theme: 'ocean'
        },
        {
            title: "關於",
            description: "繪果吉樂 BAR 是 檸果織粉絲 的部落格。若您對本站有任何疑問，歡迎至 檸果織粉絲 網站信箱 留言！",
            theme: 'sunset',
            buttonText: '查看更多3'
        },
        {
            title: "檸果織粉絲",
            description: "檸果織粉絲 是 繪果吉樂 BAR 的粉絲專頁。",
            customTheme: {
                background: 'var(--accent-light)',
                borderColor: 'var(--accent-dark)',
            },
            icon: '🍋'
        }
    ];
</script>

<Navbar {siteName} />

<section class="hero is-fullheight main-content">
    <div class="hero-body">
        <div class="container">
            <h1 class="main-title has-text-centered mb-6">
                歡迎光臨 Jelly Bar！
            </h1>
            
            <ArticleGrid {articles} />
        </div>
    </div>
</section>

<style>
    /* 主要內容區域 */
    .main-content {
        padding-top: 4rem;
    }

    .main-title {
        font-size: 3rem;
        font-weight: 800;
        color: var(--theme-primary);
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
