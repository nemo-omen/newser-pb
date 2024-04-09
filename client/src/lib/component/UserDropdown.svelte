<script lang="ts">
	export let userName = '';
	export let email = 'hello@newser.app';
	export let userId = '';
	const avatarInitials = userName.length > 0 ? userName[0].toUpperCase() : email[0].toUpperCase();
</script>

<li class="dropdown">
	<button class="avatar" type="button" aria-expanded="false" aria-controls="user-dropdown">
		{avatarInitials}
	</button>
	<ul class="dropdown__menu" id="user-dropdown">
		<li>
			{#if userName.length > 0}
				<a href={`/app/user`}>{userName}</a>
			{:else}
				<a href={`/app/user`}>{email}</a>
			{/if}
		</li>
		<li>
			<form action="/auth/logout" method="post">
				<button type="submit" class="button-link">Log Out</button>
			</form>
		</li>
	</ul>
</li>

<style>
	.avatar {
		line-height: 2.5rem;
		text-align: center;
		text-anchor: middle;
	}
</style>
