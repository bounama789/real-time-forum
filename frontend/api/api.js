    const baseUrl = "http://127.0.0.1:8000";

  export async function get(url) {
    return fetch(url, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',

      },
    })
    .then(response => response.json())
    .then(data => data)
    .catch(error => console.error('Error:', error));
  }

  export async function getPosts(){
    const url = baseUrl+"/posts/get"
    return await get(url)
  }

