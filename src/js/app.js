import { is_loggedin, supabase_login, supabase_logout } from './services/app_auth';
import { insert_user_location } from './services/geolocation';
import { fetch_country, update_country, delete_country, create_country } from './services/countries';
import { makeScriptsExecutable, render_app, render_test } from './services/render_app';

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
window.makeScriptsExecutable = makeScriptsExecutable;
window.render_app = render_app;
window.render_test = render_test;



const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    is_loggedin()
});
