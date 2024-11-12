<script>
    export let title = "繪果吉樂 BAR";
    import Counter from '../components/Counter.svelte';
    import '../app.css';

    // 定義更豐富的卡片主題
    const cardThemes = {
        forest: {
            background: 'var(--color-green-pale)',
            borderColor: 'var(--color-green-dark)',
            hoverColor: 'var(--color-green-sage)',
            icon: '🌲'
        },
        ocean: {
            background: 'var(--color-blue-light)',
            borderColor: 'var(--color-blue-dark)',
            hoverColor: 'var(--color-blue-medium)',
            icon: '🌊'
        },
        sunset: {
            background: 'var(--color-orange-light)',
            borderColor: 'var(--color-orange-dark)',
            hoverColor: 'var(--color-orange-medium)',
            icon: '🌅'
        },
        honey: {
            background: 'var(--color-yellow-light)',
            borderColor: 'var(--color-yellow-dark)',
            hoverColor: 'var(--color-yellow-amber)',
            icon: '🍯'
        },
        lavender: {
            background: 'var(--color-purple-light)',
            borderColor: 'var(--color-purple-deep)',
            hoverColor: 'var(--color-purple-dusty)',
            icon: '💜'
        },
        cherry: {
            background: 'var(--color-red-light)',
            borderColor: 'var(--color-red-dark)',
            hoverColor: 'var(--color-red)',
            icon: '🍒'
        },
        mint: {
            background: 'var(--color-green-mint)',
            borderColor: 'var(--color-green-sage)',
            hoverColor: 'var(--color-green-dark)',
            icon: '🌿'
        },
        blossom: {
            background: 'var(--color-pink-light)',
            borderColor: 'var(--color-pink-dark)',
            hoverColor: 'var(--color-pink-medium)',
            icon: '🌸'
        },
        sunshine: {
            background: 'var(--color-yellow-pale)',
            borderColor: 'var(--color-yellow-dark)',
            hoverColor: 'var(--color-yellow-medium)',
            icon: '☀️'
        },
        twilight: {
            background: 'var(--color-purple-dusty)',
            borderColor: 'var(--color-purple-deep)',
            hoverColor: 'var(--color-purple)',
            icon: '🌙'
        }
    };

    const articles = [
        {
            title: "Rust繁中簡學！",
            description: "Rust繁中簡學！系列文章以初學者的角度跟大家分享我自己翻譯成繁體中文的 Rust 自學資源以及學習過程。",
            theme: 'forest',
            customIcon: '🦀'
        },
        {
            title: "所有文章",
            description: "瀏覽本站所有技術文章，包含 Rust、Web 開發、系統設計等各種主題。",
            theme: 'ocean'
        },
        {
            title: "關於",
            description: "繪果吉樂 BAR 是 檸果織粉絲 的部落格。若您對本站有任何疑問，歡迎至 檸果織粉絲 網站信箱 留言！",
            theme: 'sunset'
        },
        {
            title: "檸果織粉絲",
            description: "檸果織粉絲 是 繪果吉樂 BAR 的粉絲專頁。",
            theme: 'honey',
            customIcon: '🍋'
        },
        {
            title: "技術分享",
            description: "分享各種程式開發技術與心得。",
            theme: 'lavender'
        },
        {
            title: "最新消息",
            description: "網站最新消息與更新資訊。",
            theme: 'cherry'
        }
    ];
</script>

<nav class="navbar is-dark">
    <div class="navbar-brand">
        <a class="navbar-item" href="/">
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
            <h1 class="title has-text-centered is-size-2 mb-6">
                歡迎光臨 Jelly Bar！
            </h1>
            
            <div class="articles-grid">
                {#each articles as article}
                    <div 
                        class="article-card"
                        style="
                            background-color: {cardThemes[article.theme].background};
                            border-left-color: {cardThemes[article.theme].borderColor};
                        "
                    >
                        <div class="card-icon">
                            {article.customIcon || cardThemes[article.theme].icon}
                        </div>
                        <div class="card-content">
                            <h2 class="title is-4">{article.title}</h2>
                            <p class="description mb-4">
                                {article.description}
                            </p>
                            <div class="card-footer">
                                <button 
                                    class="button custom-button"
                                    style="background-color: {cardThemes[article.theme].borderColor}"
                                >
                                    查看更多
                                </button>
                            </div>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</section>

<style>
    .navbar {
        background-color: var(--color-green-dark) !important;
    }

    .navbar-item {
        color: var(--color-green-light) !important;
    }
    
    .navbar-item:hover {
        background-color: rgba(0, 0, 0, 0.1) !important;
    }
    
    .main-content {
        background-color: var(--color-green-light);
    }

    .box {
        transition: transform 0.3s ease;
    }
    
    .box:hover {
        transform: translateY(-5px);
    }

    .rust-box {
        background-color: var(--color-green-pale);
    }

    .articles-box {
        background-color: var(--color-yellow-pale);
    }

    .about-box {
        background-color: var(--color-purple);
    }
    
    .custom-button {
        width: 100%;
        padding: 1rem;
        font-weight: bold;
        transition: all 0.3s ease;
        border: none;
        color: var(--color-green-light);
    }
    
    .custom-button:hover {
        opacity: 0.9;
        color: var(--color-green-light);
    }

    .articles-grid {
        display: grid;
        gap: 2rem;
        grid-template-columns: repeat(auto-fit, minmax(min(100%, 400px), 1fr));
        margin: 0 auto;
        max-width: 1200px;
        padding: 0 1.5rem;
    }

    .article-card {
        border-radius: 12px;
        padding: 2rem;
        height: 100%;
        display: flex;
        flex-direction: column;
        position: relative;
        overflow: hidden;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        transition: all 0.3s ease;
        min-height: 300px;
        word-break: break-word;
        border-left: 5px solid;
    }

    .article-card:hover {
        transform: translateY(-5px);
        box-shadow: 0 8px 15px rgba(0, 0, 0, 0.2);
    }

    .card-icon {
        font-size: 2.5rem;
        margin-bottom: 1rem;
        text-align: center;
    }

    .card-content {
        flex: 1;
        display: flex;
        flex-direction: column;
    }

    .description {
        color: #4a4a4a;
        line-height: 1.6;
        margin-bottom: 1.5rem;
        flex: 1;
        overflow-wrap: break-word;
        hyphens: auto;
    }

    .card-footer {
        margin-top: auto;
    }

    .custom-button {
        width: 100%;
        padding: 1rem;
        font-weight: bold;
        transition: all 0.3s ease;
        border: none;
        color: var(--color-green-light);
    }

    .custom-button:hover {
        opacity: 0.9;
        color: var(--color-green-light);
    }

    @media screen and (max-width: 768px) {
        .articles-grid {
            grid-template-columns: 1fr;
            padding: 1rem;
            gap: 1.5rem;
        }

        .article-card {
            padding: 1.5rem;
            min-height: auto;
        }

        .hero-body {
            padding: 1rem;
        }

        .title.is-size-2 {
            font-size: 1.75rem !important;
        }
    }

    @media screen and (min-width: 769px) and (max-width: 1023px) {
        .articles-grid {
            grid-template-columns: repeat(2, 1fr);
        }
    }

    @media screen and (min-width: 1024px) {
        .articles-grid {
            grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
        }
    }
</style>
