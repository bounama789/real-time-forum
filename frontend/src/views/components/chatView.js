export class chatView {
  constructor() {
    this.chat = chat;
  }
  ChatComponent() {
    const chatCard = document.createElement("section");
    chatCard.classList.add("chat-card");
    const chatTemplate = `
    <div class="chat-card-sidebar">
    <div class="chat-card-sidebar-header">
    <div class="chat-card-sidebar-header-avatar">
    <img src="${this.chat._avatar}" alt="avatar"/>
    <h3>${this.chat._author}</h3>
    </div>
    <div class="chat-card-sidebar-header-search">
    <input type="text" placeholder="Search"/>
    </div>
    <div class="chat-card-sidebar-header-chatsinprogress">
    <h3>Chats in progress</h3>
    <div class="chat-card-sidebar-header-chatsinprogress-list">
    <ul>
    ${this.chat._chatsinprogress
      .map((chat) => {
        return `<li><a href="#" style="${chat.color}">${chat}</a></li>`;
      })
      .join("")}
    </ul>
    </div>
    </div>
    <div class="chat-card-content">
    <div class="chat-card-content-header">
<div class="chat-card-content-header-avatar">
<img src="${this.chat._avatar}" alt="avatar"/>
<h3>${this.chat._author}</h3>
</div>
<div class="chat-card-content-discussion">
<div class="chat-card-content-discussion-chat">
<ul>
${this.chat._discussion
  .map((chat) => {
    return `<li><pre style="${chat.color}">${chat.content}</pre></li>`;
  })
  .join("")}
</div>
<div class="chat-card-content-discussion-input">
<form>
<input type="text" placeholder="Type your message"/>
<button><span class="material-symbols-outlined">send</span>Send</button>
</form>
</div>
    </div>   `;
    chatCard.innerHTML = chatTemplate;
    return chatCard;
  }
}
