import Api from "./api";

export class CommentApi extends Api {
    constructor(url) {
      super(url);
    }
    async getComment() {
      return await this.get();
    }
  }
