import { CommentView, CommentformComponent } from "./commentView.js";

export class PostView {
  constructor(post) {
    this.post = post;
  }
  PostComponent() {
    const postComponent = document.createElement("div");
    postComponent.classList.add("post-card");
    const postTemplate = `
    <div class="post-container">
        <div class="post-group">
          <div class="post-header">
            <div class="post-image">
              <img class="post-avatar" src="${this.post._avatar}" alt="avatar"/>
              <div class="post-name">
                <div class="post-username">"${this.post._author}"</div>
                <div class="post-duration">"${this.post.duration}"</div>
              </div>
            </div>
                      <button class="post-edit"><span class="material-symbols-outlined">edit</span></button>
</div>
          <div class="post__content">
            <div class="post__title">"${this.post.title}"</div>
            <pre class="post-text">"${this.post.content}"</pre>
            <nav class="post-categories">
            <ul>
            ${this.post.categories
              .map((category) => {
                return `<li><a href="#" style="${category.color}">${category}</a></li>`;
              })
              .join("")}
            </ul>
          </div>
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
                this.post._comments
              } Comments</button>
              <button><span class="material-symbols-outlined">bookmark</span>${
                this.post._bookmark
              } Bookmarks</button>
            </div>
          </div>
        </div>
        <div class="post-comment">
        ${this.post._comments
          .map((comment) => new commentModel(comment))
          .map((comment) => {
            const commentComp = new CommentView(comment);
            return commentComp.commentComponent();
          })
          .join("")}
        </div>
        <div class="post-comment-form">
        ${CommentformComponent.commentformComponent()}
          </div>
    </div>
        `;
    postComponent.innerHTML = postTemplate;
    return postComponent;
  }
}

export class PostForm {
  constructor() {
    this.post = post;
  }
  postComponent() {
    const postComponent = document.createElement("section");
    postComponent.classList.add("post-creation-card");
    postComponent.innerHTML = `
    <form method="post" class="post-created" id="post-created">
      <input type="text" class="post-created-title" placeholder="Enter Your Title" maxlenght="30" required>
      <textarea class="post-created-text" placeholder="Enter Your Post" maxlenght="280" required></textarea>
      <div class="post-created-categories"></div>
      <div class="post-created-toolbar">
        <button class="reset-button" type="reset"><span class="material-symbols-outlined">delete</span>RESET</button>
        <button class="confirm-button" type="submit"><span class="material-symbols-outlined">send</span>SEND</button>
      </div>
    </form>
    `;
    return postComponent;
  }
}
