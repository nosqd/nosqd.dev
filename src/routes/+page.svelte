<script lang="ts">
	import { memes } from '$lib/memes';
	let discordCopiedAlert = $state(false);

	async function copyDiscordToClipboard() {
		try {
			await navigator.clipboard.writeText('nosqd');
			discordCopiedAlert = true;
			setTimeout(() => {
				discordCopiedAlert = false;
			}, 2000);
		} catch {
			alert('Failed to copy text, copy by yourself: @nosqd');
		}
	}
</script>

<main class="flex w-screen flex-wrap items-center justify-center">
	<div class="flex max-w-4xl flex-col items-start justify-start p-4 sm:min-w-0 md:min-w-2xl">
		<span class="text-2xl">hi 👋, I am Dmitry Pshennikov, self-taught developer from Russia.</span>

		<section class="mt-4 mb-4 flex flex-col gap-2">
			<a class="text-2xl text-purple-200" href="#contacts" id="contacts"># contacts</a>
			<span>
				X: <a class="text-purple-200" href="https://x.com/nosqdgg" target="_blank">@nosqdgg</a>
			</span>
			<span>
				telegram: <a class="text-purple-200" href="https://t.me/pmnullsqd" target="_blank">@pmnullsqd</a>
			</span>
			<span>
				github: <a class="text-purple-200" href="https://github.com/nosqd" target="_blank">@nosqd</a>
			</span>
			<span>
				codeberg: <a class="text-purple-200" href="https://codeberg.org/nosqd" target="_blank">@nosqd</a>
			</span>
			<span>
				discord: <button class="cursor-pointer text-purple-200" onclick={copyDiscordToClipboard}>
					@nosqd
					{#if discordCopiedAlert}
						<span class="text-white">copied successfully!</span>
					{/if}
				</button>
			</span>
		</section>

		<section class="mt-4 mb-4 flex flex-col gap-2">
			<a class="text-2xl text-purple-200" href="#stack" id="stack"># my stack</a>
			<p>I'm familiar with:</p>
			<br />
			<p>Frontend: basics (HTML, CSS, JS), TypeScript, React, Vue, Svelte</p>
			<p>Backend: Moleculer (TS), Go</p>
			<p>DevOps: Docker, k8s (basics), Forgejo, Woodpecker CI</p>
			<p>Systems Programming (still learning): Zig, C/C++</p>
			<br />
			<p>Started as a webdev, now learning low-level gamedev (wgpu, vulkan, opengl) and systems programming in general.</p>
		</section>

		<section class="mt-4 mb-4 flex flex-col gap-2">
			<a class="text-2xl text-purple-200" href="#memes" id="memes"># memes</a>
			<div class="grid w-full gap-4 sm:grid-cols-2 md:grid-cols-4">
				{#each memes as meme (meme)}
					<img src={meme} alt="meme" width="250" />
				{/each}
			</div>
		</section>
	</div>
</main>
