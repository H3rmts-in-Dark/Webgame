<script>
	import {Svg} from '@smui/common/elements';
	import IconButton, {Icon} from "@smui/icon-button";

	export var svg = ''
	export var size = 'button'

	export var click = () => {
	}
</script>

<IconButton class="material-icons" size={size} color="secondary" on:click={click}>
	<Icon component={Svg} viewBox="0 0 24 24" color="secondary">
		<path fill="currentColor" d={svg} color="secondary"/>
	</Icon>
</IconButton>