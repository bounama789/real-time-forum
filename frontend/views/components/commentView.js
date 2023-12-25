export class CommentView extends HTMLElement {
  constructor(comment) {
    super();
    this.commentComponent = comment;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
      <div class="comment-card">
        <div class="comment-container">
          <div class="comment-header">
            <div class="comment-image">
              <img class="comment-avatar" src="${this.commentComponent._avatar}" alt="avatar"/>
              <div class="comment-username">${this.commentComponent._author}</div>
            </div>
            <button class="comment-edit">
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div class="comment-content">
            <pre class="comment-text">${this.commentComponent._content}</pre>
          </div>
          <div class="comment-reactions">
            <button><span class="material-symbols-outlined">thumb_up</span>${this.commentComponent._likes} Likes</button>
            <button><span class="material-symbols-outlined">thumb_down</span>${this.commentComponent._dislikes} Dislikes</button>
            <div class="comment-duration">${this.commentComponent._duration} ago</div>
          </div>
        </div>
      </div>
    `;
  }
}

export class CommentFormComponent extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.innerHTML = `
      <section class="comment-creation-card">
        <form method="post" class="comment-created" id="comment-created">
          <input type="text" class="coomment-created-text" placeholder="Enter Your Comment" maxlength="280" required>
          <div class="comment-created-toolbar">
            <button class="reset-button" type="reset"><span class="material-symbols-outlined">delete</span></button>
            <button class="confirm-button" type="submit"><span class="material-symbols-outlined">send</span></button>
          </div>
        </form>
      </section>
    `;
  }
}
