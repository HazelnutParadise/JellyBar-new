<script>
    export let siteName;
</script>

<footer class="footer">
    <div class="content has-text-centered">
        <p>
            &copy; {new Date().getFullYear()} {siteName} - 由 <a href="https://hazelnut-paradise.com">榛果繽紛樂</a> 版權所有
        </p>
    </div>
</footer>

<style>
    .footer {
        background-color: var(--theme-primary);
        color: white;
        padding: 1rem 0;
    }
</style> 