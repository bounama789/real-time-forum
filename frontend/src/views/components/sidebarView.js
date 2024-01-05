export class sidebarView {
  constructor() {
    this.sidebar = sidebar;
  }

  SideBarLeftComponent() {
    const sideBarLeft = document.createElement("aside");
    sideBarLeft.classList.add("sidebar-left");
    const sideBarLeftContent = `
    <div class="sidebar-left__container">
    <ul>
    <li><a href="#"><span class="material-symbols-outlined">explore</span><b>Explorer</b></a></li>
    <li><a href="#"><span class="material-symbols-outlined">person</span><b>Profile</b></a></li>
    <li><a href="#"><span class="material-symbols-outlined">communities</span><b>Communities</b></a></li>
    <li><a href="#"><span class="material-symbols-outlined">
    category
    </span><b>Category</b></a></li>
    <li><a href="#"><span class="material-symbols-outlined">
    pages
    </span><b>Post</b></a></li>
    <li><a href="#"><span class="las la-home"></span><b></b></a></li>
    </ul>
    </div>
    `;
    sideBarLeft.innerHTML = sideBarLeftContent;
    return sideBarLeft;
  }

  SideBarRightComponent(usersOnline) {
    const usersListItems = usersOnline.map((user) => {
      return `<li>
                  <img
                    class="u-avatar"
                    src=${user.avatar}
                    alt="avatar"
                    style="border-radius: 50%"
                  />
                  <div class="name">
                    <div class="username">${user.username}</div>
                    <div><span class="blink"></span><h2 class="online-text">I'm Online</h2></div>
                  </div>
                </li>`;
    });
    const usersList = usersListItems.join("");
    const sideBarLeft = document.createElement("aside");
    sideBarLeft.classList.add("sidebar-left");
    const sideBarLeftContent = `
    <div class="sidebar-left__container">
    <ul>
    ${usersList}
    </ul>
    </div>
    `;
    sideBarLeft.innerHTML = sideBarLeftContent;
    return sideBarLeft;
  }
}
