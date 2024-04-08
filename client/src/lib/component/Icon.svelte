<script>
	import { iconPaths } from '$lib/util/iconpaths';
	export let name = 'logo';
	$: paths = iconPaths[name] || [];
</script>

<svg
	class="icon"
	viewBox="0 0 24 24"
	fill-rule="evenodd"
	clip-rule="evenodd"
	xmlns="http://www.w3.org/2000/svg"
>
	{#if paths}
		{#each paths as path}
			<path d={path} />
		{/each}
	{/if}
</svg>
