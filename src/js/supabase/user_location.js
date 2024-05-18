import { supabase_client } from './instance';
import { getUser } from './user';

export async function insert(coordinates) {
    const user = await getUser()
    const result = await supabase_client
        .from('user_location')
        .insert({ user_id: user.id, ...coordinates })
    return result
}