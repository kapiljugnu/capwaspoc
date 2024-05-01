
const key = "key";
const url = "url";

function supabase_login() {
    const email = document.getElementById("email").value
    const password = document.getElementById("password").value

    const { createClient } = supabase
    const supabase_client = createClient(url, key)

	supabase_client.auth.signInWithPassword({ email,password }).then(({data, error})=>{
		 if (error == undefined) {
			// localStorage.user = JSON.stringify(data.user)
			render_app('loggedin', JSON.stringify({"email":data.user.email}))
		 } else {
			render_app("loginfail")
		 }
	})
    return false;
}

function is_loggedin() {
	const { createClient } = supabase
    const supabase_client = createClient(url, key)
	supabase_client.auth.getUser().then(({ data })=>{
		if (data.user) {
			render_app('loggedin', JSON.stringify({"email": data.user.email}))
		} else {
			render_app('login')
		}

	})
}