<script>
	import SvgIcon from "../SvgIcon.svelte";
	import {mdiLogout} from "@mdi/js";

	let click = () => {
		console.log("test logout")
	}
</script>

<SvgIcon svg={mdiLogout} size="normal" click={click}/>

