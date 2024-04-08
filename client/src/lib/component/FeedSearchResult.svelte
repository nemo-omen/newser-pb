<script lang="ts">
	import type { GoFeed } from '$lib/types';
	import SearchResultCard from './SearchResultCard.svelte';
	export let feeds: GoFeed[] = [];
</script>

<div class="stack">
	<h3>We found {feeds.length} feeds</h3>
	{#if feeds.length > 0}
		<ul class="stack card-list search-results">
			{#each feeds as feed}
				<SearchResultCard {feed} />
			{/each}
		</ul>
	{:else}
		<p>No feeds found</p>
	{/if}
</div>

<style>
	ul {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(30rem, 1fr));
		gap: 1rem;
		list-style-type: none;
		margin: 0;
		padding: 0;
	}
</style>
