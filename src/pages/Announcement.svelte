<script lang="ts">
    import "../app.css"
    import Footer from "../components/Footer.svelte"
    import setTitle from "../js/setTitle"
    export let siteName: string
    export let siteLogo_base64: string
    export let announcementTitle: string
    export let announcementContent: string
    export let announcementType: string

    setTitle(announcementTitle, siteName)
</script>

<style>
    .announcement-page {
        min-height: calc(100vh - 100px);
        background-color: var(--neutral-medium);
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 40px 20px;
        gap: 40px;
    }

    .site-header {
        text-align: center;
    }

    .site-logo {
        width: 80px;
        height: 80px;
        margin-bottom: 10px;
    }

    .site-name {
        font-size: 1.5rem;
        color: var(--theme-primary);
        margin: 0;
    }

    .announcement-container {
        width: 100%;
        max-width: 800px;
        background-color: white;
        border-radius: 15px;
        box-shadow: 0 2px 15px rgba(0, 0, 0, 0.08);
        overflow: hidden;
        margin-bottom: 20px;
    }

    .announcement-header {
        padding: 25px 30px;
        position: relative;
    }

    .announcement-content {
        padding: 30px;
        background-color: white;
    }

    /* 根據不同類型的公告設計 */
    .error .announcement-header {
        background-color: #fff5f5;
        border-left: 4px solid var(--state-error);
        color: var(--state-error-dark);
    }

    .info .announcement-header {
        background-color: #ebf8ff;
        border-left: 4px solid var(--interactive-dark);
        color: var(--interactive-dark);
    }

    .success .announcement-header {
        background-color: #f0fff4;
        border-left: 4px solid var(--theme-accent);
        color: var(--theme-accent);
    }

    .warning .announcement-header {
        background-color: #fffbeb;
        border-left: 4px solid var(--accent-dark);
        color: var(--accent-dark);
    }

    h1 {
        margin: 0;
        font-size: 1.8rem;
        font-weight: 600;
    }

    p {
        margin: 0;
        line-height: 1.7;
        color: var(--theme-text);
        font-size: 1.1rem;
    }

    a {
        color: var(--interactive-dark);
        font-weight: 600;
    }

    a:hover {
        color: var(--interactive-dark-hover);
    }

    .announcement-icon {
        position: absolute;
        right: 20px;
        top: 50%;
        transform: translateY(-50%);
        width: 24px;
        height: 24px;
        opacity: 0.6;
    }
</style>

<div class="announcement-page">
    <div class="site-header">
        <img src="data:image/png;base64,{siteLogo_base64}" alt="網站 Logo" class="site-logo">
        <h2 class="site-name">{siteName}</h2>
    </div>

    <div class="announcement-container {announcementType}">
        <div class="announcement-header">
            <h1>{announcementTitle}</h1>
            <div class="announcement-icon">
                <!-- 這裡可以放置對應的 SVG 圖示 -->
            </div>
        </div>
        <div class="announcement-content">
            <p>{announcementContent}</p>
        </div>
    </div>
    {#if announcementType === "error"}
        <p>先來去 <a target="_blank" href="https://shopee.tw/timlai0102?utm_content=39EGJW1Nt4ij7gVhGfwj4eZp4qK5">榛划算二手商店</a> 逛逛吧！</p>
    {/if}
</div>
<Footer {siteName} />