import {supabase_client} from '../supabase/instance';
import {getUser} from '../supabase/user';

export async function supabase_login() {
	const email = document.getElementById("email").value
	const password = document.getElementById("password").value

	const {data, error } =  await supabase_client.auth.signInWithPassword({ email, password });
	if (error == undefined) {
		render_app('loggedin', JSON.stringify({ "email": data.user.email }))
	} else {
		render_app("loginfail")
	}
	return false;
}

export async function is_loggedin() {
	const user = await getUser()
	if (user) {
		render_app('loggedin', JSON.stringify({ "email": user.email }))
	} else {
		render_app('login')
	}
}

export async function supabase_logout() {
	const {error} = await supabase_client.auth.signOut()
	if (!error) {
		render_app("login")
	}
}

