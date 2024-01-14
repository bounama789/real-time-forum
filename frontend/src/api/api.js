const baseUrl = "http://127.0.01:8000";

export async function get(path) {
  const url = baseUrl + path
  const token = localStorage.getItem("auth-token")  || "none"

  return fetch(url, {
    method: 'GET',
    headers: {
      // 'Accept': 'application/json',
      'auth-token': JSON.stringify({token})
    },
  })
    .then(response => {
      console.log(response);
      console.log(response.headers.getSetCookie());

      return response.json()
    })
    .then(data => data).catch(error => error)
}

export async function post(path, data) {
  const url = baseUrl + path
  const token = localStorage.getItem("auth-token") || "none"

  return fetch(url, {
    method: 'POST',
    body: JSON.stringify(data),
    headers:{
      'auth-token': JSON.stringify({token})
    }

  })
    .then(response => {
      return response.json()
    })
    .then(data => data).catch(error => error)
}

export async function getPosts() {
  const path = "/posts/get"
  return await get(path).catch(error => error)
}

export async function checkSession() {
  const token = localStorage.getItem("auth-token")
  return await post("/verifsess",{token}).catch(error => error)
}

export function setWSConnection(){
  const token = localStorage.getItem("auth-token")
  if (window.WebSocket) {
    const socket = new WebSocket(`ws://localhost:8000/ws?token=${JSON.stringify({token})}}`)
    window.app.wsConnection = socket
  }
  
}

