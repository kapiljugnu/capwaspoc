import { Geolocation } from '@capacitor/geolocation';
import {insert} from '../supabase/user_location'

export async function insert_user_location() {
  const {coords} = await Geolocation.getCurrentPosition();
  const result = await insert({lat: coords.latitude, long: coords.longitude})
  console.log(result)
};