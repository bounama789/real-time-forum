const user = {
  avatar: "https://i.pravatar.cc/300?img=9",
  username: "John Smith",
};
export class navbarView extends HTMLElement {
  constructor(user) {
    super();
    this.user = user;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
    <nav class="navbar">
    <div class="logo-hamburger">
    <span class="material-symbols-outlined">menu</span><img src="" alt="logo" class="logo"/>
    </div>
    <div class="search-bar">
    <input type="text" placeholder="Search here..." class="search-input"/>
    <span class="material-symbols-outlined">search</span>
    </div>
    <div class="icon-bar">
    <ul>
    <li><a href="#"><span class="material-symbols-outlined">home</span></a></li>
    <li><a href="#"><span class="material-symbols-outlined">notifications</span></a></li>
    <li><a href="#"><span class="material-symbols-outlined">message</span></a></li>
    <li><a href="#"><span class="material-symbols-outlined">person</span></a></li>
    <li><a href="#"><img src="${user.avatar}" alt="profile" class="profile-img"/><span>${user.username}</span></a></li>
    </ul>
    </div>
    </nav>`;
  }
}
