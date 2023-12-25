export default class Api {
  constructor(url) {
    this.url = url;
  }
  async get() {
    try {
      const response = await fetch(this.url);
      const data = await response.json();
      return data;
    } catch (error) {
      console.log(error);
    }
  }
}

/*async get() {
    return fetch(this.url)
      .then((res) => res.json())
      .then((res) => res.data)
      .catch((err) => console.log(err));
  } */
