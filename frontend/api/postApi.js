import Api from "./api.js";

export class PostApi extends Api {
  constructor(url) {
    super(url);
  }
  async getPost() {
    return await this.get();
  }
}
