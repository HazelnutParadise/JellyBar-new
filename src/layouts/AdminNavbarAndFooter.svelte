<script lang="ts">
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    export let siteName: string
    export let logo: string
</script>

<AdminNavbar {siteName} {logo} />
<slot />
<Footer {siteName} />
