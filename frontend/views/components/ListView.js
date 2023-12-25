import { View } from "../../common/types/View.js";
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
  constructor({ id, items = [], itemView, provider = async ()=>{} } = {}) {
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

    /**
     * @type {Div}
     */
    this.listContainer = new Div({
      className:`list-${this.id}-container`,
      style: {
        width: "100%",
        height: "100%",
        display: "flex",
        flexDirection: "column",
        alignItems: "start",
        padding: "1rem",
        borderRadius: "10px",
        gap: "1rem",
        backgroundColor:'transparent',
      },
    });

    // if (typeof this.provider == "Function") {
      this.provider(1).then((response)=>{
        // Add the items to the list
        response.forEach((item) => {
          this.listContainer.addChild(new this.itemView(item));
        });
      })
    // }

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

  fetchMoreItems() {
    this.provider(1).then((response)=>{
      // Add the items to the list
      response.forEach((item) => {
        this.listContainer.addChild(new this.itemView(item));
      });
    })
    try {
      this.addItem()
      this.page++;
      // this.provider(this.page+1).then(async(response)=>{
      //   this.page++;
      //   this.addItem(response)
      // })
    } catch (error) {
      console.log(error);
    }
  }
}
