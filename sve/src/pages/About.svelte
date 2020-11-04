<script>
	function do_login() {
		const ids = ['IdAccount', 'Password'];
		let data = {};
		for( let IdAccount of ids ) {
			data[ IdAccount ] = document.getElementById( IdAccount ).value;
		}
		fetch('', {
			method: 'post',
			body: JSON.stringify(data),
		} ).then( ( res ) => {
			res.json().then((json) => {
				document.getElementById( 'result' ).innerText = JSON.stringify(json);
			})
		} )
		return false
	}

</script>

<style>
	h1 {
		color: purple;
	}
</style>

<label for="IdAccount">Id Account</label>
<input name="IdAccount" id="IdAccount"/>
<p></p>
<label for="Password">Password</label>
<input name='Password' id="Password"/>
<p></p>
<button onClick="do_login()">Login</button>
