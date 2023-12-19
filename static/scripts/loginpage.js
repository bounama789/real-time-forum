// password eye;
var inputIcon = document.querySelectorAll(".input__icon");
let isHidden = true;
inputIcon.forEach((icon) => {
  icon.addEventListener("click", async (e) => {
    e.preventDefault();
    isHidden = !isHidden;

    e.target.previousElementSibling.type = isHidden ? "password" : "text";
    e.target.src = isHidden
      ? "../static/assets/eye.svg"
      : "../static/assets/eye-off.svg";
  });
});

// check inputs
var form = document.getElementById("form");
var username = document.getElementById("username");
var email = document.getElementById("email");
var password = document.getElementById("password");
var password2 = document.getElementById("repeatpassword");

if (form) {
  var url = form.action
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    let f = e.currentTarget
    var url = f.action
    ok = await checkInputs()
    if (ok) {
      let formFields = new FormData(form)
      for (const [name, value] of formFields.entries()) {
        if (name != "password" && typeof value === 'string') {
          formFields.set(name, value.trim());
        }
    }
      var res = await postFormFieldsAsJson({ url, formFields })

      if (res.status == 422) {
        // Show error the user to the new URL
        let s = await res.text()
        let msg = document.createTextNode(JSON.parse(s).msg) 
        let p = document.createElement("p")
        p.style.color = "red"
        p.style.width = "100%"
        p.style.fontSize = "14px"
        p.style.display = "block"
        p.style.padding = "10px"
        p.style.textAlign = "center"
        p.classList.add("login__error__msg")
        p.wid
        if (form.querySelector(".login__error__msg")) {
          form.removeChild(form.querySelector(".login__error__msg"))
        }
        p.appendChild(msg)
        form.appendChild(p)
      } else if (res.redirected) {
        // Redirect the user to the new URL
        window.location.href = res.url;
      }

      
    }

  });

}

async function postFormFieldsAsJson({ url, formFields }) {
  //Create an object from the form data entries
  let formDataObject = Object.fromEntries(formFields.entries());
  // Format the plain form data as JSON
  let formDataJsonString = JSON.stringify(formDataObject);

  //Set the fetch options (headers, body)
  let fetchOptions = {
    //HTTP method set to POST.
    method: "POST",
    //Set the headers that specify you're sending a JSON body request and accepting JSON response
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    // POST request body as JSON string.
    body: formDataJsonString,
  };

  //Get the response body as JSON.
  //If the response was not OK, throw an error.
  let res = await fetch(url, fetchOptions);

  //If the response is not ok throw an error (for debugging)
//If the response was OK, return the response body.
  return res
}

async function checkInputs() {
  let ok = true;
  // trim to remove the whitespaces
  var usernameValue = username.value.trim();
  var emailValue = email.value.trim();
  var passwordValue = password.value;
  var password2Value = password2.value;
  emailCheckRes = await verifyExistingEmail({ email: emailValue })
  usernameCheckRes = await verifyExistingUsername({ username: usernameValue })

  if (usernameValue == "") {
    ok = false;
    setErrorFor(username, "invalid username");
  } else if (usernameCheckRes != null && usernameCheckRes != "valid") {
    ok = false
    setErrorFor(username, usernameCheckRes);
  } else {
    setSuccessFor(username);
  }
  if (emailValue === "") {
    ok = false;
    setErrorFor(email, "invalid email");
  } else if (!isEmail(emailValue)) {
    ok = false;
    setErrorFor(email, "invalid email");
  } else if (emailCheckRes != null && emailCheckRes !== "valid") {
    ok = false;
    setErrorFor(email, emailCheckRes);
  } else {
    setSuccessFor(email);
  }

  if (passwordValue === "") {
    ok = false;
    setErrorFor(password, "invalid password");
  } else {
    setSuccessFor(password);
  }

  if (password2Value === "") {
    ok = false;
    setErrorFor(password2, "password doesn't match");
  } else if (passwordValue !== password2Value) {
    ok = false;
    setErrorFor(password2, "password doesn't match");
  } else {
    setSuccessFor(password2);
  }

  return ok;
}

function setErrorFor(input, msg) {
  var msgNode = document.createTextNode(msg)
  var formControl = input.parentElement;
  let p = document.createElement("p")
  if (input.id !== "first" || input.id !== "last") {
    formControl = input.parentElement.parentElement;
  }
  formControl.classList.remove("success")
  formControl.classList.add("error");
  p.classList.add("error__msg")
  p.appendChild(msgNode)
  if (formControl.querySelector(".error__msg")) {
    formControl.removeChild(formControl.querySelector(".error__msg"))
  }
  formControl.appendChild(p)
}

function setSuccessFor(input) {
  var formControl = input.parentElement;
  if (input.id !== "first" || input.id !== "last") {
    formControl = input.parentElement.parentElement;
  }
  formControl.classList.remove("error")
  if (formControl.querySelector(".error__msg")) {
    formControl.removeChild(formControl.querySelector(".error__msg"))
  }
  formControl.classList.add("success");
}

function isEmail(email) {
  return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(
    email
  );
}

function isValidUSername(username) {
  r = /^[a-zA-Z][a-zA-Z0-9_]{5,15}$/.test(
    username
  );
  return r
}

//LOG IN Check Input

var formlogin = document.getElementById("formlogin")
var emaillogin = document.getElementById("email-login");
var passwordlogin = document.getElementById("password-login");


if (formlogin) {
  var url = formlogin.action
  formlogin.addEventListener("submit", async (e) => {
    e.preventDefault();
    let f = e.currentTarget
    var url = f.action

    if (checkLoginInputs()) {
      let formFields = new FormData(formlogin)
      var res = await postFormFieldsAsJson({ url, formFields })

      if (res.status == 401 || res.status == 422) {
        // Show error the user to the new URL
        let msg = document.createTextNode("wrong credentials")
        let p = document.createElement("p")
        p.style.color = "red"
        p.style.fontSize = "18px"
        p.style.display = "block"
        p.style.padding = "10px"
        p.style.textAlign = "center"
        p.classList.add("login__error__msg")
        p.wid
        if (formlogin.querySelector(".login__error__msg")) {
          formlogin.removeChild(formlogin.querySelector(".login__error__msg"))
        }
        p.appendChild(msg)
        formlogin.appendChild(p)
      } else if (res.status == 200) {
        if (formlogin.querySelector(".login__error__msg")) {
          formlogin.removeChild(formlogin.querySelector(".login__error__msg"))
        }
        window.location.href = res.url
      }
    }

  });
}



function checkLoginInputs() {
  let ok = true;
  // trim to remove the whitespaces
  var emailloginValue = emaillogin.value.trim();
  var passwordloginValue = passwordlogin.value.trim();
  var formControl = emaillogin.form
  if (emailloginValue === "") {
    ok = false;
    if (formControl.querySelector(".login__error__msg")) {
      formControl.removeChild(formControl.querySelector(".login__error__msg"))
    }
    setErrorFor(emaillogin,"this field should not be empty");
  } else if (emailloginValue.includes("@") && !isEmail(emailloginValue)) {
    ok = false;
    if (formControl.querySelector(".login__error__msg")) {
      formControl.removeChild(formControl.querySelector(".login__error__msg"))
    }
    setErrorFor(emaillogin,"wrong email");

  } else if (!isValidUSername(emailloginValue) && !emailloginValue.includes("@")) {
    ok = false;
    if (formControl.querySelector(".login__error__msg")) {
      formControl.removeChild(formControl.querySelector(".login__error__msg"))
    }
    setErrorFor(emaillogin,"wrong username");

  } else {
    if (formControl.querySelector(".login__error__msg")) {
      formControl.removeChild(formControl.querySelector(".login__error__msg"))
    }
    setSuccessFor(emaillogin);
  }

  // if (usernameValue === "") {
  //   ok = false;
  //   setErrorFor();
  // } else {
  //   setSuccessFor();
  // }
  return ok
}

// function setErrorFor() {
// var error = document.querySelectorAll(".Error_Message")
// error.classList.remove("messagesuccess")
// error.classList.add("messagerror")
// }

// function setSuccessFor() {
//   var error = document.querySelectorAll(".Error_Message")
//   error.classList.remove("messagerror")
//   error.classList.add("messagesuccess")
//   }
async function getRequest({ url }) {

  //Set the fetch options (headers, body)
  let fetchOptions = {
    method: "GET",
    //Set the headers that specify you're sending a JSON body request and accepting JSON response
    headers: {
      Accept: "application/json",
    },
  };

  //Get the response body as JSON.
  //If the response was not OK, throw an error.
  let res = await fetch(url, fetchOptions);

  //If the response is not ok throw an error (for debugging)
  // if (!res.ok) {
  //   let error = await res.text();
  //   throw new Error(error);
  // }
  //If the response was OK, return the response body.
  return res
}

async function verifyExistingEmail({ email }) {
  let url = `/verify/email?email=${email}`
  let res = await getRequest({ url })
  msg = await res.text()
  return JSON.parse(msg).msg
}

async function verifyExistingUsername({ username }) {
  let url = `/verify/username?username=${username}`
  let res = await getRequest({ url })
  msg = await res.text()
  return JSON.parse(msg).msg
}
