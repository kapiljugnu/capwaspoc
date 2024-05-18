import { is_loggedin, supabase_login, supabase_logout } from './services/app_auth';
import { insert_user_location } from './services/geolocation';
import { fetch_country, update_country, delete_country, create_country } from './services/countries';

window.is_loggedin = is_loggedin;
window.supabase_login = supabase_login;
window.supabase_logout = supabase_logout;
window.insert_user_location = insert_user_location;
window.db_op = {
    fetch_country,
    update_country,
    delete_country,
    create_country
}
