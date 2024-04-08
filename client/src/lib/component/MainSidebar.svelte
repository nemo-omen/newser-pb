<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import IconLink from '$lib/component/IconLink.svelte';
	import Icon from '$lib/component/Icon.svelte';

	$: current = $page.url.pathname ?? '';

	afterNavigate(() => {
		current = $page.url.pathname ?? '';
		console.log({ current });
	});
</script>

<aside class="sidebar" id="sidebar-main">
	<nav aria-label="Secondary Navigation">
		<ul>
			<li>
				<IconLink
					href="/app"
					iconName="list"
					text="All Posts"
					className="icon-link"
					current={current === '/app'}
				/>
			</li>
			<li>
				<IconLink
					href="/app/collection/saved"
					iconName="bookmark"
					text="Saved"
					className="icon-link"
					current={current === '/app/collection/saved'}
				/>
			</li>
			<li>
				<IconLink
					href="/app/note"
					iconName="note"
					text="Notes"
					className="icon-link"
					current={current === '/app/note'}
				/>
			</li>
			<li>
				<IconLink
					href="/app/search"
					iconName="add"
					text="Add Feed"
					className="icon-link"
					current={current === '/app/search'}
				/>
			</li>
		</ul>
	</nav>
	<!-- if len(getSidebarFeedInfo(ctx)) > 0 {
    <nav
      class="nav-vertical"
      aria-labelledby="#sidebar-subscriptions-heading"
      hx-get="/app/control/unreadcount"
      hx-trigger="updateUnreadCount from:body"
      hx-target="#main-feed-links"
      hx-swap="outerHTML"
      hx-push-url="false"
    >
      <a
        href="/app/subscriptions"
        class="sidebar-link"
        aria-current={ isCurrentString(ctx, "/app/subscriptions") }
        style="padding-block: var(--space-2xs); border-radius: 0.5rem;"
      >
        <h2 class="text-small" id="sidebar-subscriptions-heading" padding-block="var(--space-2xs)">
          Subscriptions
        </h2>
      </a>
      @MainFeedLinks()
    </nav>
  }-->
	<nav class="nav-vertical" aria-labelledby="#sidebar-collections-heading">
		<a
			href="/app/collection"
			class="sidebar-link"
			aria-current={current === '/app/collection' ? 'page' : undefined}
			style="padding-block: var(--space-2xs); border-radius: 0.5rem;"
		>
			<h2 class="text-small">Collections</h2>
		</a>
		<ul id="main-collection-links">
			<li>
				<a
					href="/app/collection/new"
					class="icon-link"
					aria-current={current === '/app/collection/new' ? 'page' : undefined}
				>
					<Icon name="folder_add" />
					<span class="link-text">New Collection</span>
				</a>
			</li>
		</ul>
	</nav>
</aside>
