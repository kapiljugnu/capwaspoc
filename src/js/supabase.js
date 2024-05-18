import { createClient } from '@supabase/supabase-js'

const key = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImR4cWFndGFheGNxbXhhenJqa3ZrIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTUyNTY2MDgsImV4cCI6MjAzMDgzMjYwOH0.gIs3NI62Z2AmABwPe6IRJoGCKGFOZ2wnEXRPZ0X2Muw";
const url = "https://dxqagtaaxcqmxazrjkvk.supabase.co";

const supabase_client = createClient(url, key)

export function supabase_login() {
	const email = document.getElementById("email").value
	const password = document.getElementById("password").value

	supabase_client.auth.signInWithPassword({ email, password }).then(({ data, error }) => {
		if (error == undefined) {
			// localStorage.user = JSON.stringify(data.user)
			render_app('loggedin', JSON.stringify({ "email": data.user.email }))
		} else {
			render_app("loginfail")
		}
	})
	return false;
}

export function is_loggedin() {
	supabase_client.auth.getUser().then(({ data }) => {
		if (data.user) {
			render_app('loggedin', JSON.stringify({ "email": data.user.email }))
		} else {
			render_app('login')
		}

	})
}

export function supabase_logout() {
	supabase_client.auth.signOut().then(({ error }) => {
		if (!error) {
			render_app("login")
		}
	})
}

export function create_country() {
	supabase_client
		.from('countries')
		.insert({ id: 1, name: 'Denmark' })
		.then(({error}) => {
			if (!error) {
				alert('inserted')
			} else {
				console.log(error)
				alert('insert failed')
			}
		})
}

export function delete_country() {
	supabase_client
		.from('countries')
		.delete()
		.eq('id', 1)
		.then(({error}) => {
			if (!error) {
				alert('deleted')
			} else {
				alert('deletion failed')
			}
		})
}

export function update_country() {
	supabase_client
		.from('countries')
		.update({ name: 'Australia' })
		.eq('id', 1).then(({error}) => {
			if (!error) {
				alert('updated')
			} else {
				alert('update fail')
			}
		})
}

export function fetch_country() {
	supabase_client
		.from('countries')
		.select().then(({data}) => {
			if (data) {
				alert(data?.[0]?.name)
			} else {
				alert('fetch failed')
			}
		})
}