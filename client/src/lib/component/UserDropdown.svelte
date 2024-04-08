<script lang="ts">
	export let email = 'hello@newser.app';
	$: emailInitial = email[0].toUpperCase();
</script>

<li class="dropdown">
	<button class="avatar" type="button" aria-expanded="false" aria-controls="user-dropdown">
		{emailInitial}
	</button>
	<ul class="dropdown__menu" id="user-dropdown">
		<li>
			<a href="/app/user">{email}</a>
		</li>
		<li>
			<form action="/auth/logout" method="post">
				<button type="submit" class="button-link">Log Out</button>
			</form>
		</li>
	</ul>
</li>
