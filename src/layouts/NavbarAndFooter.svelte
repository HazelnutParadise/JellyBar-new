<script lang="ts">
    import Navbar from '../components/Navbar.svelte'
    import Footer from '../components/Footer.svelte'
    export let siteName: string
</script>

<Navbar {siteName} />
<slot />
<Footer {siteName} />
