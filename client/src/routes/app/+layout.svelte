<script lang="ts">
	import MainSidebar from '$lib/component/MainSidebar.svelte';
	import SecondarySidebar from '$lib/component/SecondarySidebar.svelte';
</script>

<div class="grid-container">
	<MainSidebar />
	<main>
		<slot />
	</main>
	<SecondarySidebar />
</div>
