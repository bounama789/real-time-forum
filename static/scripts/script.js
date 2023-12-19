var profile = document.querySelector(".profile");
var dropdown = document.querySelector(".dropdown__wrapper");

if (profile) {
  profile.addEventListener("click", () => {
    dropdown.classList.remove("none");
    dropdown.classList.toggle("hide");
  });
}

if (dropdown) {
  document.addEventListener("click", (event) => {
    var isClickInsideDropdown = dropdown.contains(event.target);
    var isProfileClicked = profile.contains(event.target);

    if (!isClickInsideDropdown && !isProfileClicked) {
      dropdown.classList.add("hide");
      dropdown.classList.add("dropdown__wrapper--fade-in");
    }
    dropdown.childNodes.forEach(child => child.addEventListener("click", () => {
      dropdown.classList.add("hide");
    }))
  });
}

//dropdown.script ends here
//hamburger.script starts here
var hamburger = document.querySelector(".hamburger");
var navbar = document.querySelector(".accordion-container");

hamburger.addEventListener("click", (e) => {
  hamburger.textContent = hamburger.textContent === "close" ? "menu" : "close";
  hamburger.classList.toggle("active");
  navbar.classList.toggle("active");
})

document.addEventListener("click", (event) => {
  var isClickInsideNavbar = navbar.contains(event.target);
  var isHamburgerClicked = hamburger.contains(event.target);

  if (!isClickInsideNavbar && !isHamburgerClicked) {
    hamburger.textContent = "menu"
    hamburger.classList.remove("active");
    navbar.classList.remove("active");
  }
}
)

document.querySelectorAll(".Categories__Name").forEach(child => child.addEventListener("click", () => {
  hamburger.textContent = hamburger.textContent === "close" ? "menu" : "close";
  hamburger.classList.remove("active");
  navbar.classList.remove("active");
}
))


//hamburger.script ends here

//main-item.script starts here

var mainItems = document.querySelectorAll(".main-item");
mainItems.forEach((mainItem) => {
  mainItem.addEventListener("click", () => {
    mainItem.classList.toggle("main-item--open");
    document.addEventListener("click", (event) => {
      var isClickInsideMainItem = mainItem.contains(event.target);
      var isMainItemClicked = mainItem.contains(event.target);

      if (!isClickInsideMainItem && !isMainItemClicked) {
        mainItem.classList.remove("main-item--open");
      }
    }
    )
  });
});
//main-item.script ends here

//searchbar.script starts here
document.addEventListener('click', function (event) {
  // Check if the click target is not the search input
  if (event.target !== searchInput) {
    suggDiv.style.visibility = "hidden";
  }
});

// COMMENT
document.addEventListener("DOMContentLoaded", function () {
  const colorSchemeQuery = window.matchMedia('(prefers-color-scheme: dark)');
  handleColorSchemeChange(colorSchemeQuery);

  // Initial check and handling

  ccs = localStorage.getItem('preferredColorScheme') == null ? window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light' : localStorage.getItem('preferredColorScheme');
  tmp = ccs != null ? ccs == "light" ? "dark" : "light" : null
  toggleColorScheme(tmp)
  themeToggleButton = document.querySelector(".theme__toggle")
  themeToggleButton.addEventListener("click", () => {
    currentColorScheme = localStorage.getItem('preferredColorScheme');
    toggleColorScheme(currentColorScheme)
  })

  colorSchemeQuery.addListener(handleColorSchemeChange);

  suggDiv = document.querySelector(".sugg__div")
  searchInput = document.querySelector(".searchTerm")
  searchButton = document.querySelector(".searchButton")
  sanitizedValue = searchInput.value
  searchButton.setAttribute("hx-vals", `{"q":"${sanitizedValue}"}`);

  searchInput.addEventListener("keyup", (e) => {
    sanitizedValue = searchInput.value
    searchButton.setAttribute("hx-vals", `{"q":"${sanitizedValue}"}`);
    document.addEventListener('htmx:afterRequest', function (event) {
      if (suggDiv.childElementCount > 0 && document.activeElement == searchInput) {
        suggDiv.style.visibility = "visible"
      } else if (document.activeElement == searchInput) {
        suggDiv.style.visibility = "hidden"
      } else {
        suggDiv.style.visibility = "hidden"

      }
      let suggs = document.querySelectorAll(".post__sugg__div")

      suggs.forEach((elem) => {
        elem.addEventListener("click", (e) => {
          document.addEventListener('htmx:afterRequest', function (event) {
            suggDiv.visibility = "hidden"
            searchInput.value = ""
          })
        })
      })
    });

  })

  var commentButtons = document.querySelectorAll(".post__nbrcomment");

  commentButtons.forEach(function (button) {
    button.addEventListener("click", (e) => {
      let postContainer = button.closest(".post__container");
      let commentDiv = postContainer.querySelector(".post__comment");

      // Check the initial state and toggle accordingly
      if (
        commentDiv.style.display === "" ||
        commentDiv.style.display === "none"
      ) {
        commentDiv.style.display = "block";
      } else {
        commentDiv.style.display = "none";
      }
    });
  });
  // RESET BUTTON
  var newCommentForm = document.getElementById("comment__created__content");

  // LIKES AND DISLIKES
  var like1 = document.querySelectorAll(".post__likes");
  var like2 = document.querySelectorAll(".post__dislikes");

  like1.forEach((button) => {
    button.addEventListener("click", function () {
      if (button.nextElementSibling.classList.contains("red")) {
        button.nextElementSibling.classList.remove("red");
      }
      button.style.borderRight = "1px solid green";
      button.nextElementSibling.style.borderLeft = "none"
      this.classList.toggle("green");
    });
  });
  like2.forEach((button) => {
    button.addEventListener("click", function () {
      if (button.previousElementSibling.classList.contains("green")) {
        button.previousElementSibling.classList.remove("green");
      }
      button.style.borderLeft = "1px solid red";
      button.previousElementSibling.style.borderRight = "none"
      this.classList.toggle("red");
    });
  });

});

function toggleColorScheme(currentColorScheme) {
  themeToggleButton = document.querySelector(".theme__toggle")
  themeToggleButton.innerHTML = currentColorScheme == "light" ? "light_mode" : "dark_mode"

  if (currentColorScheme === 'light') {
    switchToDarkMode();
    if (currentColorScheme != null) {

      setPreferredColorScheme("dark")
    }
  } else {
    switchToLightMode();
    if (currentColorScheme != null) {
      setPreferredColorScheme("light")

    }

  }
}

function setPreferredColorScheme(scheme) {
  localStorage.setItem('preferredColorScheme', scheme);
}


// Function to switch to light mode
function switchToLightMode() {
  document.documentElement.style.setProperty('--background', '#d6e9ff');
  document.documentElement.style.setProperty('--text', 'black');
  document.documentElement.style.setProperty('--secondary', '#00abe4');
  document.documentElement.style.setProperty('--primary', '#ffffff');
  document.documentElement.style.setProperty('--text-gray', '#2e4452');
  document.documentElement.style.setProperty('--textsize', 'clamp(1rem, 1.5vw, 2rem)');
  document.querySelectorAll(".affected").forEach((elem) => {
    elem.classList.remove('dark')
  })
}

// Function to switch to dark mode
function switchToDarkMode() {
  document.documentElement.style.setProperty('--background', 'rgb(61, 62, 63)');
  document.documentElement.style.setProperty('--text', 'aliceblue');
  document.documentElement.style.setProperty('--secondary', '#00abe4');
  document.documentElement.style.setProperty('--primary', 'rgb(53, 52, 52)');
  document.documentElement.style.setProperty('--text-gray', '#2e4452');
  document.documentElement.style.setProperty('--textsize', 'clamp(1rem, 1.5vw, 2rem)');
  document.querySelectorAll(".affected").forEach((elem) => {
    elem.classList.add('dark')
  })



}

function handleColorSchemeChange(event) {
  console.log(event)
  if (event.matches) {
    n = localStorage.getItem("preferredColorScheme") == "light" ? "dark" : "light"
    if (n != "dark") {
      toggleColorScheme(n)
      setPreferredColorScheme('dark')
    }

  } else {
    n = localStorage.getItem("preferredColorScheme") == "dark" ? "light" : "dark"
    if (n != "light") {
      setPreferredColorScheme('light')
      toggleColorScheme(n)
    }

  }
}

