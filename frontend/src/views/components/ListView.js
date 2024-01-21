import { View } from "../../common/types/View.js";
import { setView } from "../../lib/lib.js";
import { Div } from "../elements/index.js";

/**
 * Creates a new ListView with the given options.
 * @param {{
 *  items: Array,
 *  itemView: View
 * }} options - The options for creating the ListView.
 * @returns {Div} The container element for the ListView.
 */
export class ListView {
  constructor({ id, items = [], itemView, provider = async ()=>{},providerParams = {}, style = {} } = {}) {
    this.page = 1;

    this.id=id

    /**
     * @type {Array}
     */
    this.items = items 

    /**
     * @type {View}
     */
    this.itemView = itemView;

    this.provider = provider;

    this.providerParams = providerParams

    /**
     * @type {Div}
     */
    this.listContainer = new Div({
      id:`list-${this.id}-container`,
      className: "list-container",
      style: style,
    });

    let query = ""

    for (const key in this.providerParams) {
      if (Object.hasOwnProperty.call(this.providerParams, key)) {

        if (query =! "") {
          query += "&"
        }
        const value = this.providerParams[key];
        query += `&${key}=${value}`;
      }
    }

    this.providerQueries = query

    query += `&page=${this.page}`

    // if (typeof this.provider == "Function") {
      this.provider(query).then((response)=>{
        response = response || []

        // Add the items to the list
        response.forEach((item) => {
          this.listContainer.addChild(new this.itemView(item));
        });
      })
    // }

    setView(this)

  }

  /**
   * Adds one or more items to the end of the list.
   * @param {...*} items - The items to add to the list.
   */
  addItem(...items) {
    this.items.push(items);
    items.forEach((item) => {
      this.listContainer.addChild(new this.itemView(item));
    });
  }

  /**
   * Adds one or more items to the beginning of the list.
   * @param {...*} items - The items to add to the list.
   */
  prependItem(...items) {
    this.items.unshift(items);
    items.forEach((item) => {
      this.listContainer.prependChild(new this.itemView(item));
    });
  }

  removeItemView(itemView) {
    this.listContainer.removeChild(itemView);
  }

  fetchMoreItems() {
    this.page++;
    this.providerQueries += `&page=${this.page}`

    this.provider(this.providerQueries).then((response)=>{
      // Add the items to the list
      response = response || []
      response.forEach((item) => {
        this.listContainer.addChild(new this.itemView(item));
      });
    })
    try {
      this.addItem()
      // this.provider(this.page+1).then(async(response)=>{
      //   this.page++;
      //   this.addItem(response)
      // })
    } catch (error) {
      console.log(error);
    }
  }
}
