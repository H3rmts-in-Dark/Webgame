<script>
	import {Svg} from '@smui/common/elements';
	import IconButton, {Icon} from "@smui/icon-button";

	export var svg = ''
	export var size = 'button'

</script>

<IconButton class="material-icons" size={size}>
	<Icon component={Svg} viewBox="0 0 24 24">
		<path fill="currentColor" d={svg}/>
	</Icon>
</IconButton>