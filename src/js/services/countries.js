import { supabase_client } from '../supabase/instance';

export async function create_country() {
    const { error } = await supabase_client
        .from('countries')
        .insert({ id: 1, name: 'Denmark' });

    if (!error) {
        alert('inserted')
    } else {
        console.log(error)
        alert('insert failed')
    }
}

export async function delete_country() {
    const { error } = await supabase_client
        .from('countries')
        .delete()
        .eq('id', 1);

    if (!error) {
        alert('deleted')
    } else {
        alert('deletion failed')
    }
}

export async function update_country() {
    const { error } = await supabase_client
        .from('countries')
        .update({ name: 'Australia' })
        .eq('id', 1);

    if (!error) {
        alert('updated')
    } else {
        alert('update fail')
    }
}

export async function fetch_country() {
    const { data } = await supabase_client
        .from('countries')
        .select();
    if (data) {
        alert(data?.[0]?.name)
    } else {
        alert('fetch failed')
    }
}