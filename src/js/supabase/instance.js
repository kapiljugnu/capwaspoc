import { createClient } from '@supabase/supabase-js'

const key = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImR4cWFndGFheGNxbXhhenJqa3ZrIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTUyNTY2MDgsImV4cCI6MjAzMDgzMjYwOH0.gIs3NI62Z2AmABwPe6IRJoGCKGFOZ2wnEXRPZ0X2Muw";
const url = "https://dxqagtaaxcqmxazrjkvk.supabase.co";

export const supabase_client = createClient(url, key)

