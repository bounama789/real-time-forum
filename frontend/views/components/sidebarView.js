export class leftSidebarView extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
    <aside class="sidebar-left-container">
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
    </aside>
    `;
  }
}

const usersOnline = [
  {
    avatar: "https://i.pravatar.cc/300?img=3",
    username: "John Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=4",
    username: "Jane Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=5",
    username: "John Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=6",
    username: "Jane Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=7",
    username: "John Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=8",
    username: "Jane Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=9",
    username: "John Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=10",
    username: "Jane Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=11",
    username: "John Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=12",
    username: "Jane Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=13",
    username: "John Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=14",
    username: "Jane Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=15",
    username: "John Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=16",
    username: "Jane Doe",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=17",
    username: "John Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=18",
    username: "Jane Smith",
  },
  {
    avatar: "https://i.pravatar.cc/300?img=19",
    username: "John Doe",
  },
];
export class rightSidebarView extends HTMLElement {
  constructor(usersOnline) {
    super();
    this.usersOnline = usersOnline;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
      <aside class="sidebar-left__container">
      <ul>
      ${usersOnline
        .map((user) => {
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
        })
        .join("")}
      </ul>
      </aside>
      `;
  }
}
