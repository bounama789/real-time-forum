export class CommentView {
  constructor(commentComponent) {
    this.commentComponent = commentComponent;
  }
  commentComponent() {
    const commentComponent = document.createElement("section");
    commentComponent.classList.add("comment-card");
    commentComponent.innerHTML = `
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
          <button><span class="material-symbols-outlined">thumb_up</span>${this.commentComponent._likes}Likes</button>
          <div class="comment-duration">${this.commentComponent._duration} ago</div>
        </div>
      </div>
    `;
  }
}
export class CommentformComponent {
  constructor() {
    this.commentform = commentform;
  }
  commentformComponent() {
    const commentformComponent = document.createElement("section");
    commentformComponent.classList.add("comment-creation-card");
    commentformComponent.innerHTML = `
        <form method="post" class="comment-created" id="comment-created">
          <input type="text" class="comment-created-text" placeholder="Enter Your Comment" maxlenght="280" required>
          <div class="comment-created-toolbar">
            <button class="reset-button" type="reset"><span class="material-symbols-outlined">delete</span></button>
            <button class="confirm-button" type="submit"><span class="material-symbols-outlined">send</span></button>
          </div>
        </form>
    `;
    return commentformComponent;
  }
}
