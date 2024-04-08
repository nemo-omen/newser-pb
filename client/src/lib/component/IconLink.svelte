<script lang="ts">
	import Icon from './Icon.svelte';
	export let href: string;
	export let iconName: string;
	export let text: string;
	export let className: string;
	export let current = false;
</script>

<a {href} class={className ?? 'icon-link'} aria-current={current ? 'page' : undefined}>
	<Icon name={iconName} />
	{text}
</a>
