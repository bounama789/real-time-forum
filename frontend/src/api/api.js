import { getView } from "../lib/lib.js";

const baseUrl = "http://127.0.01:8000";

export async function get(path) {
  const url = baseUrl + path
  const token = localStorage.getItem("auth-token") || "none"

  return fetch(url, {
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'auth-token': JSON.stringify({ token })
    },
  })
    .then(response => {
      console.log(response);
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
    headers: {
      'auth-token': JSON.stringify({ token })
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
  return await post("/verifsess", { token }).catch(error => error)
}

export function setWSConnection() {
  const token = localStorage.getItem("auth-token")
  if (window.WebSocket) {
    const socket = new WebSocket(`ws://localhost:8000/ws?token=${JSON.stringify({ token })}}`)
    window.app.wsConnection = socket
  }

}

export async function getChatByUser(username) {
  const path = `/chat?username=${username}`
  return await get(path).catch(error => error)
}

export async function getUsersStatus() {
  const path = "/users-status"
  return await get(path).catch(error => error)
}

export const EventType = {
  WS_JOIN_EVENT: "join-event",
  WS_DISCONNECT_EVENT: "disconnect-event",
  WS_MESSAGE_EVENT: "msg-event"
}

export async function getMessages(queries) {
  const path = `/messages?${queries}`;
  return await get(path).catch(error => error)
}

export function handleWSEvent(wsEvent) {
  const event = JSON.parse(wsEvent)
  switch (event["Type"]) {
    case EventType.WS_JOIN_EVENT:
      setStatusOnline(event.From)
      break;
    case EventType.WS_DISCONNECT_EVENT:
      setStatusOffline(event.From)
      break;
    case EventType.WS_MESSAGE_EVENT:
      dispatchEvent(new CustomEvent("newMessage", { detail: event }))
    default:
      break;
  }
}

function setStatusOnline(username) {
  const dot = getView(`${username}status-dot`).element
  const text = getView(`${username}-status-text`).element
  const statusItmView = getView(`user-status${username}`)
  statusItmView.user.status = "online"
  text.innerText = "online"
  dot.style.backgroundColor = "green"
  if (getView(`chat${username}`)) {
    const dot = getView(`chat-${username}-status-dot`).element
    dot.style.backgroundColor = "green"
  }
}

function setStatusOffline(username) {
  const dot = getView(`${username}status-dot`).element
  const text = getView(`${username}-status-text`).element
  const statusItmView = getView(`user-status${username}`)
  statusItmView.user.status = "offline"
  text.innerText = "offline"
  dot.style.backgroundColor = "gray"
  if (getView(`chat${username}`)) {
    const dot = getView(`chat-${username}-status-dot`).element
    dot.style.backgroundColor = "gray"
  }
}