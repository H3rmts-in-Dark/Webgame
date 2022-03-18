<script>
	import {Svg} from '@smui/common/elements';
	import {Icon} from "@smui/icon-button";

	export var svg = ''
	export var cls = ''
</script>

<Icon class="margin0 {cls}" component={Svg} viewBox="0 0 24 24" color="secondary">
	<path fill="currentColor" d={svg} color="secondary"/>
</Icon>