<script>
	let w;
	let h;
	let size = 20;
	let text = 'text to scale';
</script>

<div>
	
	<input type=range bind:value={size}>
	
	<p>size: {w}px x {h}px</p>
	
	<div bind:clientWidth={w} bind:clientHeight={h}>
		<span style="font-size: {size}px">{text}</span>
	</div>
</div>

<style>
    input {
        display: block;
    }

    div {
        display: inline-block;
    }

    span {
        word-break: break-all;
    }
</style>