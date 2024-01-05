import Api from "./api";

export class UserApi extends Api {
  constructor(url) {
    super(url);
  }
  async getUser() {
    return await this.get();
  }
}
