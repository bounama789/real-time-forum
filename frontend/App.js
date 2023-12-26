import { PostApi } from "./api/postApi.js";
import { navbarView } from "./views/components/navbarView.js";
import {
  rightSidebarView,
  leftSidebarView,
} from "./views/components/sidebarView.js ";
import { Post } from "./models/post/post.js";
import { PostView } from "./views/components/postView.js";

customElements.define("post-view", PostView);
customElements.define("navbar-view", navbarView);
customElements.define("right-sidebar-view", rightSidebarView);
customElements.define("left-sidebar-view", leftSidebarView);
class App {
  constructor() {
    this.mainWrapper = document.querySelector(".main-wrapper");
    this.postApi = new PostApi("/post");
  }
  async init() {
    const posts = await this.postApi.getPost();
    this.mainWrapper.innerHTML = `
    <navbar-view></navbar-view>
    <left-sidebar-view></left-sidebar-view>
    <section>
    ${posts
      .map((post) => new PostView(new Post(post)))
      .map((postView) => postView.render())
      .join("")}
      })}
    </section>
    <right-sidebar-view></right-sidebar-view>
    `;
  }
}

const app = new App();
app.init();
