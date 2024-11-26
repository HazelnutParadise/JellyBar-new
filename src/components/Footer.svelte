<script>
    export let siteName
</script>

<style>
    .footer {
        height: 100px;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--theme-primary);
        color: white;
        padding: 1rem 0;
    }
</style>

<footer class="footer">
    <div class="content has-text-centered">
        <p>
            &copy; {new Date().getFullYear()}
            {siteName} - 由
            <a href="https://hazelnut-paradise.com">榛果繽紛樂</a> 版權所有
        </p>
    </div>
</footer>
