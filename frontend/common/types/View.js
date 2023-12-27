import { setView } from "../../lib/lib.js";

/**
 * Creates a new View element.
 * @param {{
 *   id?: string,
 *   className?: string,
 *   name?: string,
 *   children?: View[],
 *   style?: object,
 *   placeholder?: string
 * }} options - The options for creating the view.
 */
export class View {
  constructor(options = {}) {
    this.element;
    this.children = options.children || [];
    this.style = options.style || {};
    this.classList = options.className && options.className.split(" ");
    this.name = options.name;
    this.id = options.id;
    this.placeholder = options.placeholder;
    this.listeners = options.listeners
    this.onScroll = options.onScroll;
    this.constraints = options.constraints

    setView(this);
  }

  /**
   * Sets the attributes of the view element.
   */
  _setAttributes() {
    this.element.classList.add(this.classList);
    this.element.id = this.id && this.id;
    this.element.name = this.name && this.name;
    this.element.placeholder = this.placeholder && this.placeholder;
  }

  _setEventListeners() {
    if (this.listeners) {
      document.addEventListener("DOMContentLoaded", () => {
        for (const eventName in this.listeners) {
          if (Object.hasOwnProperty.call(this.listeners, eventName)) {
            const func = this.listeners[eventName];
            if (func) {
              this.element[eventName] = func
            }
          }
        }
      })
    }
  }

  _setConstraints() {
    if (this.constraints) {
      document.addEventListener("DOMContentLoaded", () => {

        if (this.constraints.top) {
          const top = this.constraints.top();
          this.element.style.top = top
        }
        if (this.constraints.bottom) {
          this.element.style.bottom = this.constraints.bottom();
        }
        if (this.constraints.left) {
          this.element.style.left = this.constraints.left();
        }
        if (this.constraints.right) {
          this.element.style.right = this.constraints.right();
        }
      });
    }
  }

  /**
   * Appends the child elements to the view element.
   */
  _appendChildren() {
    this.children.forEach((child) => this.element.appendChild(child.element));
  }

  /**
   * Applies the styles to the view element.
   */
  _applyStyles() {
    this.style &&
      Object.entries(this.style).forEach(
        ([key, value]) => (this.element.style[key] = value)
      );
  }

  /**
   * Adds a child element to the view.
   * @param {View} child - The child element to add.
   */
  addChild(child) {
    this.children.push(child);
    this.element.appendChild(child.element);
  }

  /**
   * Prepends a child element to the view.
   * @param {View} child - The child element to prepend.
   */
  prependChild(child) {
    this.children.unshift(child);
    this.element.insertBefore(child.element, this.element.firstChild);
  }

  hide() {
    this.element.style.display = "none";
  }

  show() {
    this.element.style.display = "block";
  }
}
