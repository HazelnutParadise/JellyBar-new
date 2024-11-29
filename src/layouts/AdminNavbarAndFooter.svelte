<script lang="ts">
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    export let siteName: string
</script>

<AdminNavbar {siteName} />
<slot />
<Footer {siteName} />
