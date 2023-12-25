import { CommentView, CommentFormComponent } from "./commentView.js";

customElements.define("comment-view", CommentView);
customElements.define("comment-form-component", CommentFormComponent);
export class PostView extends HTMLElement {
  constructor(post) {
    super();
    this.post = post;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `

      <div class="post-card">
        <div class="post-container">
          <!-- Post header -->
          <div class="post-group">
            <div class="post-header">
              <!-- Post image and author info -->
              <div class="post-image">
                <img class="post-avatar" src="${
                  this.post._avatar
                }" alt="avatar"/>
                <div class="post-name">
                  <div class="post-username">${this.post._author}</div>
                  <div class="post-duration">${this.post._duration}</div>
                </div>
              </div>
              <button class="post-edit"><span class="material-symbols-outlined">edit</span></button>
            </div>

            <!-- Post content -->
            <div class="post-content">
              <div class="post-title">${this.post._title}</div>
              <pre class="post-text">${this.post._content}</pre>
              <nav class="post-categories">
                <ul>
                  ${this.post._categories
                    .map((category) => {
                      return `<li><a href="#" style="${category.color}">${category}</a></li>`;
                    })
                    .join("")}
                </ul>
              </nav>
            </div>

            <!-- Post reactions -->
            <div class="post-reactions">
              <div class="post-likedislike">
                <button><span class="material-symbols-outlined">thumb_up</span>${
                  this.post._likes
                } Likes</button>
                <button><span class="material-symbols-outlined">thumb_down</span>${
                  this.post._dislikes
                } Dislikes</button>
              </div>
              <div class="post-comshare">
                <button><span class="material-symbols-outlined">comment</span>${
                  this.post._nbrcomments
                } Comments</button>
                <button><span class="material-symbols-outlined">bookmark</span>${
                  this.post._bookmark
                } Bookmarks</button>
              </div>
            </div>
          </div>

          <!-- Post comments -->
          <div class="post-comment">
            ${this.post._comments
              .map((comment) => new CommentView(comment).render())
              .join("")}
          </div>
          <!-- Post comment form -->
          <div class="post-comment-form">
            <comment-form-component></comment-form-component>
          </div>
        </div>
      </div>
    `;
  }
}

class PostForm extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
      <section class="post-creation-card">
        <form method="post" class="post-created" id="post-created">
          <input type="text" class="post-created-title" placeholder="Enter Your Title" maxlength="30" required>
          <textarea class="post-created-text" placeholder="Enter Your Post" maxlength="280" required></textarea>
          <div class="post-created-categories"></div>
          <div class="post-created-toolbar">
            <button class="reset-button" type="reset"><span class="material-symbols-outlined">delete</span>RESET</button>
            <button class="confirm-button" type="submit"><span class="material-symbols-outlined">send</span>SEND</button>
          </div>
        </form>
      </section>
    `;
  }
}

customElements.define("post-form", PostForm);
