    const baseUrl = "http://127.0.01:8000";

  export async function get(path) {
    const url = baseUrl+path
    return fetch(url, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',

      },
    })
    .then(response => response.json())
    .then(data => data).catch(error=>error)
  }

  export async function post(path,data) {
    const url = baseUrl+path
    return fetch(url, {
      method: 'POST',
      body: JSON.stringify(data)
    })
    .then(response =>{
      console.log(response.headers.getSetCookie());
       response.json()})
    .then(data => data).catch(error=>error)
  }

  export async function getPosts(){
    const path = "/posts/get"
    return await get(path).catch(error=>error)
  }

