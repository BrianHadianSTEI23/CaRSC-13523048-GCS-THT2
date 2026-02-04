<script lang="ts">
	// imports from other files
	import welcomeFallback from '$lib/images/svelte-welcome.png';
	import welcome from '$lib/images/svelte-welcome.webp';

	import Counter from './Counter.svelte';
	import LeafletMap from './LeafletMap.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
	<link rel="stylesheet" href=
"https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
    integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
    crossorigin="" />
</svelte:head>

<section>
	<!-- <h1>
		<span class="welcome">
			<picture>
				<source srcset={welcome} type="image/webp"/>
				<img src={welcomeFallback} alt="Welcome" />
			</picture>
		</span>

		to your new<br />SvelteKit app
	</h1>

	<h2>
		try editing <strong>src/routes/+page.svelte</strong>
	</h2> -->

	<LeafletMap/>

	<!-- <Counter /> -->
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 1;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
