<script lang="ts">
	import { fade } from 'svelte/transition';
	import type { GoFeed } from '$lib/types';
	export let feed: GoFeed;
</script>

<li class="card">
	<div class="card-header flex-row flex-between">
		<div class="flex-row flex-align-center gap-1">
			{#if feed.image}
				<img src={feed.image.url} alt={feed.image.title} class="feed-logo" />
			{/if}
			<h3>
				{feed.title}
			</h3>
		</div>
		<!-- <span>{feed.feedType ?? ''}</span> -->
	</div>
	<div class="card-body">
		<p>
			{feed.description ?? ''}
		</p>
		<ul class="list">
			{#if feed.items.length > 2}
				{#each feed.items.slice(0, 2) as item}
					<li class="listitem">
						<h4>{item.title}</h4>
						{@html item.description ?? ''}
					</li>
				{/each}
			{/if}
		</ul>
	</div>
	<div class="card-footer">
		<form action="?/subscribe" method="POST">
			<!-- <input type="hidden" name="subscriptionurl" id="subscriptionurl" value={feed.feedLink} /> -->
			<input type="hidden" name="feed" value={JSON.stringify(feed)} />
			<button type="submit">Subscribe</button>
			<!-- <span>{feed.feedLink}</span> -->
		</form>
	</div>
</li>

<style>
	.list {
		list-style-type: none;
		margin: 0;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: var(--space-xs);
	}

	.listitem {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}
</style>
