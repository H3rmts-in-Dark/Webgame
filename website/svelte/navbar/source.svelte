<script>
	import SvgIcon from "../SvgIcon.svelte";
	import {mdiTextSearch} from "@mdi/js";

	let click = () => {
		window.open("https://github.com/H3rmts-in-Dark/Webgame");
	}
</script>

<SvgIcon svg={mdiTextSearch} size="normal" click={click}/>

