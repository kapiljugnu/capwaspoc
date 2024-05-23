import {supabase_client} from './instance';

export async function getUser() {
	const { data }  = await supabase_client.auth.getUser()
    return data.user
}