<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
</script>

<header class="flex h-16 w-screen items-center justify-center bg-purple-700">
	<div class="flex flex-row items-center sm:min-w-0 md:min-w-2xl">
		<h1 class="text-3xl font-black">nosqd</h1>
		<nav class="ml-8">
			<a href={resolve('/')} aria-current={page.url.pathname === '/'} class="font-bold decoration-2 underline-offset-4 hover:underline aria-current:underline">about me</a>
		</nav>
	</div>
</header>
