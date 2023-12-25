export class Post {
  constructor(data) {
    this._author = data.username || data.email;
    this._avatar = data.avatar;
    this._duration = data.duration;
    this._title = data.title;
    this._content = data.content;
    this._categories = data.categories;
    this._comments = data.comments;
    this._nbrcomments = data.nbrcomments;
    this._likes = data.likes;
    this._dislikes = data.dislikes;
    this._bookmark = data.bookmark;
  }
  get author() {
    return this._author;
  }
  get avatar() {
    return this._avatar;
  }
  get duration() {
    return this._duration;
  }
  get title() {
    return this._title;
  }
  get content() {
    return this._content;
  }
  get categories() {
    return this._categories;
  }
  get comments() {
    return this._comments;
  }
  get nbrcomments() {
    return this._nbrcomments;
  }
  get likes() {
    return this._likes;
  }
  get dislikes() {
    return this._dislikes;
  }
  get bookmark() {
    return this._bookmark;
  }
}
