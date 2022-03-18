<script lang="ts">
	import SvgIcon from "../SvgIconButton.svelte";
	import {mdiCogOutline} from "@mdi/js";

	let click = () => {
		console.log("test settings")
	}
</script>

<SvgIcon svg={mdiCogOutline} size="normal" click={click}/>

