<script>
    import setTitle from '../js/setTitle.js'
    export let siteName
    export let title
    import '../app.css'
    import ArticleGrid from '../components/ArticleGrid.svelte'
    import Navbar from '../components/Navbar.svelte'
    import Footer from '../components/Footer.svelte'

    setTitle(title, siteName)

    const articles = [
        {
            title: 'Rust繁中簡學！',
            description:
                'Rust繁中簡學！系列文章以初學者的角度跟大家分享我自己翻譯成繁體中文的 Rust 自學資源以及學習過程。',
            theme: 'forest',
            icon: '🦀',
        },
        {
            title: '所有文章',
            description:
                '瀏覽本站所有技術文章，包含 Rust、Web 開發、系統設計等各種主題。',
            theme: 'ocean',
        },
        {
            title: '關於',
            description:
                '繪果吉樂 BAR 是 檸果織粉絲 的部落格。若您對本站有任何疑問，歡迎至 檸果織粉絲 網站信箱 留言！',
            theme: 'sunset',
            buttonText: '查看更多3',
        },
        {
            title: '檸果織粉絲',
            description: '檸果織粉絲 是 繪果吉樂 BAR 的粉絲專頁。',
            customTheme: {
                background: 'var(--accent-light)',
                borderColor: 'var(--accent-dark)',
            },
            icon: '🍋',
        },
    ]

    export let alert = undefined;
</script>

<style>
    .main-title {
        font-size: 3rem;
        font-weight: 800;
        color: var(--theme-primary);
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
        letter-spacing: 1px;
        animation: fadeInDown 1s ease-out;
    }

    /* 更新關於我們區塊的樣式 */
    .about-section {
        background-color: #f8f9fa;
        padding: 6rem 1.5rem;
    }

    .features-container {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
        gap: 2rem;
        margin-top: 2rem;
    }

    .feature-item {
        background: white;
        padding: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
        display: flex;
        align-items: center;
        gap: 1rem;
        transition: transform 0.2s ease;
    }

    .feature-item:hover {
        transform: translateY(-3px);
    }

    .feature-item .icon {
        font-size: 1.5rem;
    }

    @keyframes fadeInDown {
        from {
            opacity: 0;
            transform: translateY(-20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    @media screen and (max-width: 768px) {
        .main-title {
            font-size: 2rem;
            padding: 0 1rem;
        }

        .about-section {
            padding: 3rem 1rem;
        }

        .features-container {
            grid-template-columns: 1fr;
        }
    }
</style>

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

<section class="section about-section">
    <div class="container is-max-desktop">
        <div class="content has-text-centered">
            <h2 class="title is-3 mb-6">關於我們</h2>
            <p class="subtitle is-4 mb-6">
                繽果吉樂 BAR 是 <strong
                    ><a href="https://hazelnut-paradise.com">榛果繽紛樂</a
                    ></strong
                > 的部落格平台。
            </p>
        </div>

        <div class="columns is-centered">
            <div class="column is-8">
                <div class="content">
                    <p class="is-size-5 has-text-centered mb-5">
                        我們致力於提供高品質的技術文章，包括：
                    </p>
                    <div class="features-container">
                        <div class="feature-item">
                            <span class="icon">🦀</span>
                            <span>Rust 程式語言教學</span>
                        </div>
                        <div class="feature-item">
                            <span class="icon">💻</span>
                            <span>Web 開發技術分享</span>
                        </div>
                        <div class="feature-item">
                            <span class="icon">🔧</span>
                            <span>系統設計與架構</span>
                        </div>
                        <div class="feature-item">
                            <span class="icon">📚</span>
                            <span>學習資源整理</span>
                        </div>
                    </div>
                </div>
                <div class="has-text-centered mt-6">
                    <a
                        href="https://apps.hazelnut-paradise.com/mail"
                        class="button is-primary is-medium">聯絡我們</a
                    >
                </div>
            </div>
        </div>
    </div>
</section>

<Footer {siteName} />
