<script lang="ts">
    import Navbar from '../components/Navbar.svelte'
    import Footer from '../components/Footer.svelte'
    export let siteName: string
    export let logo: string
</script>

<Navbar {siteName} {logo} />
<slot />
<Footer {siteName} />
